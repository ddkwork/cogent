// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gear

//go:generate core generate -add-types

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/grr"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/keyfun"
	"cogentcore.org/core/strcase"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/texteditor"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/units"
	"cogentcore.org/core/views"
	"cogentcore.org/core/xe"
	"github.com/mattn/go-shellwords"
	"github.com/robert-nix/ansihtml"
)

// App is a GUI view of a gear command.
type App struct {
	core.Frame

	// Cmd is the root command associated with this app.
	Cmd *Cmd

	// CurCmd is the current root command being typed in.
	CurCmd string

	// Dir is the current directory of the app.
	Dir string
}

var _ tree.Node = (*App)(nil)

func (a *App) OnInit() {
	a.Frame.OnInit()
	a.Dir = grr.Log1(os.Getwd())
}

func (a *App) AppBar(tb *core.Toolbar) {
	for _, cmd := range a.Cmd.Cmds {
		cmd := cmd
		fields := strings.Fields(cmd.Cmd)
		text := strcase.ToSentence(strings.Join(fields[1:], " "))
		bt := core.NewButton(tb).SetText(text).SetTooltip(cmd.Doc)
		bt.OnClick(func(e events.Event) {
			d := core.NewBody().AddTitle(text).AddText(cmd.Doc)
			st := StructForFlags(cmd.Flags)
			views.NewStructView(d).SetStruct(st)
			d.AddBottomBar(func(parent core.Widget) {
				d.AddCancel(parent)
				d.AddOK(parent).SetText(text).OnClick(func(e events.Event) {
					grr.Log(xe.Verbose().Run(fields[0], fields[1:]...))
				})
			})
			d.NewFullDialog(bt).Run()
		})
	}
}

func (a *App) Config() {
	if a.HasChildren() {
		return
	}

	// st := StructForFlags(a.Cmd.Flags)
	// views.NewStructView(a).SetStruct(st)

	sp := core.NewSplits(a, "splits").SetSplits(0.8, 0.2)
	sp.Style(func(s *styles.Style) {
		s.Direction = styles.Column
	})

	cmds := core.NewFrame(sp, "commands")
	cmds.Style(func(s *styles.Style) {
		s.Wrap = true
		s.Align.Content = styles.End
	})

	ef := core.NewFrame(sp, "editor-frame").Style(func(s *styles.Style) {
		s.Direction = styles.Column
	})
	dir := core.NewLabel(ef, "dir").SetText(a.Dir)

	tb := texteditor.NewBuffer()
	tb.NewBuffer(0)
	tb.Hi.Lang = "Bash"
	tb.Opts.LineNos = false
	grr.Log(tb.Stat())
	te := texteditor.NewEditor(ef, "editor").SetBuffer(tb)
	te.Style(func(s *styles.Style) {
		s.Font.Family = string(core.AppearanceSettings.MonoFont)
	})
	te.OnKeyChord(func(e events.Event) {
		txt := string(tb.Text())
		txt = strings.TrimSuffix(txt, "\n")

		kf := keyfun.Of(e.KeyChord())
		if kf == keyfun.Enter && e.Modifiers() == 0 {
			e.SetHandled()
			tb.NewBuffer(0)

			grr.Log(a.RunCmd(txt, cmds, dir))
			return
		}

		envs, words := grr.Log2(shellwords.ParseWithEnvs(txt))
		if len(words) > 0 {
			a.CurCmd = words[0]
		} else {
			a.CurCmd = ""
		}
		_ = envs
	})
}

// RunCmd runs the given command in the context of the given commands frame
// and current directory label.
func (a *App) RunCmd(cmd string, cmds *core.Frame, dir *core.Label) error {
	ctx, cancel := context.WithCancel(context.Background())

	cfr := core.NewFrame(cmds).Style(func(s *styles.Style) {
		s.Grow.Set(1, 0)
		s.Direction = styles.Column
		s.Border.Radius = styles.BorderRadiusLarge
		s.Background = colors.C(colors.Scheme.SurfaceContainer)
	})
	tr := core.NewLayout(cfr, "tr").Style(func(s *styles.Style) {
		s.Align.Items = styles.Center
		s.Padding.Set(units.Dp(8)).SetBottom(units.Zero())
	})
	core.NewLabel(tr, "cmd").SetType(core.LabelTitleLarge).SetText(cmd).Style(func(s *styles.Style) {
		s.Font.Family = string(core.AppearanceSettings.MonoFont)
		s.Grow.Set(1, 0)
	})
	core.NewButton(tr, "kill").SetType(core.ButtonAction).SetIcon(icons.Close).OnClick(func(e events.Event) {
		cancel()
		fmt.Println("canceled")
	})

	// output and input readers and writers
	or, ow := io.Pipe()
	ir, iw := io.Pipe()
	var ib []byte

	buf := texteditor.NewBuffer()
	buf.NewBuffer(0)
	buf.Opts.LineNos = false

	te := texteditor.NewEditor(cfr).SetBuffer(buf)
	te.Style(func(s *styles.Style) {
		s.Font.Family = string(core.AppearanceSettings.MonoFont)
		s.Min.Set(units.Em(30), units.Em(10))
		s.Background = cfr.Styles.Background
	})
	te.OnKeyChord(func(e events.Event) {
		kc := e.KeyChord()
		kf := keyfun.Of(kc)

		fmt.Println(kc, kf)

		switch kf {
		case keyfun.Enter:
			iw.Write(ib)
			iw.Write([]byte{'\n'})
			ib = nil
		case keyfun.Backspace:
			if len(ib) > 0 {
				ib = slices.Delete(ib, len(ib)-1, len(ib))
			}
		default:
			ib = append(ib, kc...)
		}

	})

	ob := &texteditor.OutputBuffer{}
	ob.Init(or, buf, 0, func(line []byte) []byte {
		return ansihtml.ConvertToHTML(line)
	})
	go func() {
		ob.MonitorOutput()
	}()

	cmds.Update()

	words, err := shellwords.Parse(cmd)
	if err != nil {
		return err
	}
	if len(words) > 0 && words[0] == "cd" {
		d := ""
		if len(words) > 1 {
			d = filepath.Join(a.Dir, words[1])
			_, err := os.Stat(d)
			if err != nil {
				return err
			}
		} else {
			d, err = os.UserHomeDir()
			if err != nil {
				return err
			}
		}
		a.Dir = d
		dir.SetText(a.Dir).Update()
		return nil
	}

	c := exec.CommandContext(ctx, "bash", "-c", cmd)
	c.Stdout = ow
	c.Stderr = ow
	c.Stdin = ir
	c.Dir = a.Dir
	c.Cancel = func() error {
		fmt.Println("icf")
		return grr.Log(xe.Run("bash", "-c", "kill -2 "+strconv.Itoa(c.Process.Pid)))
	}
	go func() {
		grr.Log(c.Run())
	}()
	return nil
}

// StructForFlags returns a new struct object for the given flags.
func StructForFlags(flags []*Flag) any {
	sfs := make([]reflect.StructField, len(flags))

	used := map[string]bool{}
	for i, flag := range flags {
		sf := reflect.StructField{}
		sf.Name = strings.Trim(flag.Name, "-[]")
		sf.Name = strcase.ToCamel(sf.Name)

		// TODO(kai/gear): better type determination
		if flag.Type == "bool" {
			sf.Type = reflect.TypeOf(false)
		} else if flag.Type == "int" {
			sf.Type = reflect.TypeOf(0)
		} else if flag.Type == "float" || flag.Type == "float32" || flag.Type == "float64" || flag.Type == "number" {
			sf.Type = reflect.TypeOf(0.0)
		} else {
			sf.Type = reflect.TypeOf("")
		}

		sf.Tag = reflect.StructTag(`desc:"` + flag.Doc + `"`)

		if used[sf.Name] {
			// TODO(kai/gear): consider better approach to unique names
			nm := sf.Name + "1"
			for i := 2; used[nm]; i++ {
				nm = sf.Name + strconv.Itoa(i)
			}
			sf.Name = nm
		}
		used[sf.Name] = true
		sfs[i] = sf
	}
	stt := reflect.StructOf(sfs)
	st := reflect.New(stt)
	return st.Interface()
}
