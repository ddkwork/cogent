// Copyright (c) 2018, The Gide Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gide

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/goki/gi/gi"
	"github.com/goki/gi/giv"
	"github.com/goki/gi/oswin"
	"github.com/goki/ki/ki"
	"github.com/goki/ki/kit"
	"github.com/goki/pi/complete"
	"github.com/goki/pi/filecat"
	"github.com/goki/pi/lex"
	"github.com/goki/vci"
	"github.com/mattn/go-shellwords"
)

// CmdAndArgs contains the name of an external program to execute and args to
// pass to that program
type CmdAndArgs struct {
	Cmd            string  `width:"25" desc:"external program to execute -- must be on path or have full path specified -- use {RunExec} for the project RunExec executable."`
	Args           CmdArgs `complete:"arg" width:"25" desc:"args to pass to the program, one string per arg -- use {FileName} etc to refer to special variables -- just start typing { and you'll get a completion menu of options, and use backslash-quoted bracket to insert a literal curly bracket.  Use unix-standard path separators (/) -- they will be replaced with proper os-specific path separator (e.g., on Windows)."`
	Default        string  `width:"25" desc:"default value for prompt string, for first use -- thereafter it uses last value provided for given command"`
	PromptIsString bool    `desc:"if true, then do not split any prompted string into separate space-separated fields -- otherwise do so, except for values within quotes"`
}

// Label satisfies the Labeler interface
func (cm CmdAndArgs) Label() string {
	return cm.Cmd
}

// CmdArgs is a slice of arguments for a command
type CmdArgs []string

// SetCompleter specifies the functions that do completion and post selection
// editing when inserting the chosen completion
func (cm *CmdArgs) SetCompleter(tf *gi.TextField, id string) {
	if id == "arg" {
		tf.SetCompleter(cm, CompleteArg, CompleteArgEdit)
		return
	}
	fmt.Printf("no match for SetCompleter id argument")
}

// HasPrompts returns true if any prompts are required before running command,
// and the set of such args
func (cm *CmdAndArgs) HasPrompts() (map[string]struct{}, bool) {
	var ps map[string]struct{}
	if aps, has := ArgVarPrompts(cm.Cmd); has {
		ps = aps
	}
	for _, av := range cm.Args {
		if aps, has := ArgVarPrompts(av); has {
			if ps == nil {
				ps = aps
			} else {
				for key := range aps {
					ps[key] = struct{}{}
				}
			}
		}
	}
	if len(ps) > 0 {
		return ps, true
	}
	return nil, false
}

// BindArgs replaces any variables in the args with their values, and returns resulting args
func (cm *CmdAndArgs) BindArgs(avp *ArgVarVals) []string {
	sz := len(cm.Args)
	if sz == 0 {
		return nil
	}
	args := []string{}
	for i := range cm.Args {
		argNm := cm.Args[i]
		av := avp.Bind(argNm)
		if len(av) == 0 {
			continue
		}
		switch {
		case !cm.PromptIsString && argNm == "{PromptString1}":
			fallthrough
		case !cm.PromptIsString && argNm == "{PromptString2}":
			flds, err := shellwords.Parse(av)
			if err != nil {
				fmt.Println(err)
			}
			args = append(args, flds...)
			continue
		case av[0] == '*': // only allow at *start* of args -- for *.ext exprs
			glob, err := filepath.Glob(av)
			if err == nil && len(glob) > 0 {
				args = append(args, glob...)
			}
			continue
		}
		args = append(args, av)
	}
	return args
}

// PrepCmd prepares to run command, returning *exec.Cmd and a string of the full command
func (cm *CmdAndArgs) PrepCmd(avp *ArgVarVals) (*exec.Cmd, string) {
	cstr := avp.Bind(cm.Cmd)
	switch cm.Cmd {
	case "{PromptString1}": // special case -- expand args
		cmdstr := cstr
		args, err := shellwords.Parse(cmdstr)
		if err != nil {
			fmt.Println(err)
		}
		if len(args) > 1 {
			cstr = args[0]
			args = args[1:]
		} else {
			cstr = args[0]
			args = nil
		}
		cmd := exec.Command(cstr, args...)
		return cmd, cmdstr
	case "open":
		cstr = giv.OSOpenCommand()
		cmdstr := cstr
		args := cm.BindArgs(avp)
		if args != nil {
			astr := strings.Join(args, " ")
			cmdstr += " " + astr
		}
		cmd := exec.Command(cstr, args...)
		return cmd, cmdstr
	default:
		cmdstr := cstr
		args := cm.BindArgs(avp)
		if args != nil {
			astr := strings.Join(args, " ")
			cmdstr += " " + astr
		}
		cmd := exec.Command(cstr, args...)
		return cmd, cmdstr
	}
}

///////////////////////////////////////////////////////////////////////////
//  CmdRun, RunningCmds

// CmdRun tracks running commands
type CmdRun struct {
	Name    string      `desc:"Name of command being run -- same as Command.Name"`
	CmdStr  string      `desc:"command string"`
	CmdArgs *CmdAndArgs `desc:"Details of the command and args"`
	Exec    *exec.Cmd   `desc:"exec.Cmd for the command"`
}

// Kill kills the process
func (cm *CmdRun) Kill() {
	if cm.Exec.Process != nil {
		cm.Exec.Process.Kill()
	}
}

// CmdRuns is a slice list of running commands
type CmdRuns []*CmdRun

// Add adds a new running command
func (rc *CmdRuns) Add(cm *CmdRun) {
	if *rc == nil {
		*rc = make(CmdRuns, 0, 100)
	}
	*rc = append(*rc, cm)
}

// AddCmd adds a new running command, creating CmdRun via args
func (rc *CmdRuns) AddCmd(name, cmdstr string, cmdargs *CmdAndArgs, ex *exec.Cmd) {
	cm := &CmdRun{name, cmdstr, cmdargs, ex}
	rc.Add(cm)
}

// DeleteIdx delete command at given index
func (rc *CmdRuns) DeleteIdx(idx int) {
	*rc = append((*rc)[:idx], (*rc)[idx+1:]...)
}

// ByName returns command with given name
func (rc *CmdRuns) ByName(name string) (*CmdRun, int) {
	for i, cm := range *rc {
		if cm.Name == name {
			return cm, i
		}
	}
	return nil, -1
}

// DeleteByName deletes command by name
func (rc *CmdRuns) DeleteByName(name string) bool {
	_, idx := rc.ByName(name)
	if idx >= 0 {
		rc.DeleteIdx(idx)
		return true
	}
	return false
}

// KillByName kills a running process by name, and removes it from the list of
// running commands
func (rc *CmdRuns) KillByName(name string) bool {
	cm, idx := rc.ByName(name)
	if idx >= 0 {
		cm.Kill()
		rc.DeleteIdx(idx)
		return true
	}
	return false
}

///////////////////////////////////////////////////////////////////////////
//  Command

// Command defines different types of commands that can be run in the project.
// The output of the commands shows up in an associated tab.
type Command struct {
	Cat     string            `desc:"category for the command -- commands are organized in to hierarchical menus according to category"`
	Name    string            `width:"20" desc:"name of this command (must be unique in list of commands)"`
	Desc    string            `width:"40" desc:"brief description of this command"`
	Lang    filecat.Supported `desc:"supported language / file type that this command applies to -- choose Any or e.g., AnyCode for subtypes -- filters the list of commands shown based on file language type"`
	Cmds    []CmdAndArgs      `tableview-select:"-" desc:"sequence of commands to run for this overall command."`
	Dir     string            `width:"20" complete:"arg" desc:"if specified, will change to this directory before executing the command -- e.g., use {FileDirPath} for current file's directory -- only use directory values here -- if not specified, directory will be project root directory."`
	Wait    bool              `desc:"if true, we wait for the command to run before displaying output -- mainly for post-save commands and those with subsequent steps: if multiple commands are present, then it uses Wait mode regardless."`
	Focus   bool              `desc:"if true, keyboard focus is directed to the command output tab panel after the command runs."`
	Confirm bool              `desc:"if true, command requires Ok / Cancel confirmation dialog -- only needed for non-prompt commands"`
}

// CommandName returns a qualified command name as cat: cmd
func CommandName(cat, cmd string) string {
	return cat + ": " + cmd
}

// Label satisfies the Labeler interface
func (cm Command) Label() string {
	return CommandName(cm.Cat, cm.Name)
}

// HasPrompts returns true if any prompts are required before running command,
// and the set of such args
func (cm *Command) HasPrompts() (map[string]struct{}, bool) {
	var ps map[string]struct{}
	for i := range cm.Cmds {
		cma := &cm.Cmds[i]
		if aps, has := cma.HasPrompts(); has {
			if ps == nil {
				ps = aps
			} else {
				for key := range aps {
					ps[key] = struct{}{}
				}
			}
		}
	}
	if len(ps) > 0 {
		return ps, true
	}
	return nil, false
}

// CmdNoUserPrompt can be set to true to prevent user from being prompted for strings
// this is useful when a custom outer-loop has already set the string values.
// this will be reset automatically after command is run.
var CmdNoUserPrompt bool

// CmdWaitOverride will cause the next commands that are run to be in wait mode
// (sequentially, waiting for completion after each), instead of running each in
// a separate process as is typical.  Don't forget to reset it after commands.
// This is important when running multiple of the same command, to prevent collisions
// in the output buffer.
var CmdWaitOverride bool

// CmdPrompt1Vals holds last values  for PromptString1 per command, so that
// each such command has its own appropriate history
var CmdPrompt1Vals = map[string]string{}

// CmdPrompt2Vals holds last values  for PromptString2 per command, so that
// each such command has its own appropriate history
var CmdPrompt2Vals = map[string]string{}

// RepoCurBranches returns the current branch and a list of all branches
// ensuring that the current also appears on the list of all.
// In git, a new branch may not so appear.
func RepoCurBranches(repo vci.Repo) (string, []string, error) {
	cur, err := repo.Current()
	if err != nil {
		return "", nil, err
	}
	br, err := repo.Branches()
	if err != nil {
		return cur, nil, err
	}
	hasCur := false
	for _, b := range br {
		if b == cur {
			hasCur = true
			break
		}
	}
	if !hasCur {
		br = append(br, cur)
	}
	return cur, br, nil
}

// PromptUser prompts for values that need prompting for, and then runs
// RunAfterPrompts if not otherwise cancelled by user
func (cm *Command) PromptUser(ge Gide, buf *giv.TextBuf, pvals map[string]struct{}) {
	sz := len(pvals)
	avp := ge.ArgVarVals()
	cnt := 0
	var cmvals map[string]string
	for pv := range pvals {
		switch pv {
		case "{PromptString1}":
			cmvals = CmdPrompt1Vals
			fallthrough
		case "{PromptString2}":
			if cmvals == nil {
				cmvals = CmdPrompt2Vals
			}
			curval, _ := cmvals[cm.Label()] // (*avp)[pv]
			if curval == "" && cm.Cmds[0].Default != "" {
				curval = cm.Cmds[0].Default
			}
			gi.StringPromptDialog(ge.VPort(), curval, "Enter string value here..",
				gi.DlgOpts{Title: "Gide Command Prompt", Prompt: fmt.Sprintf("Command: %v: %v", cm.Name, cm.Desc)},
				ge.This(), func(recv, send ki.Ki, sig int64, data any) {
					dlg := send.(*gi.Dialog)
					if sig == int64(gi.DialogAccepted) {
						val := gi.StringPromptDialogValue(dlg)
						cmvals[cm.Label()] = val
						(*avp)[pv] = val
						cnt++
						if cnt == sz {
							cm.RunAfterPrompts(ge, buf)
						}
					}
				})
		// todo: looks like all the file prompts are not supported?
		case "{PromptBranch}":
			fn := ge.ActiveFileNode()
			if fn != nil {
				repo, _ := fn.Repo()
				if repo != nil {
					cur, br, err := RepoCurBranches(repo)
					if err == nil {
						gi.StringsChooserPopup(br, cur, ge.VPort(), func(recv, send ki.Ki, sig int64, data any) {
							ac := send.(*gi.Action)
							brnm := ac.Text
							(*avp)[pv] = brnm
							cnt++
							if cnt == sz {
								cm.RunAfterPrompts(ge, buf)
							}
						})
					} else {
						fmt.Println(err)
					}
				}
			}
		}
	}
}

// Run runs the command and saves the output in the Buf if it is non-nil,
// which can be displayed -- if !wait, then Buf is updated online as output
// occurs.  Status is updated with status of command exec.  User is prompted
// for any values that might be needed for command.
func (cm *Command) Run(ge Gide, buf *giv.TextBuf) {
	if cm.Confirm {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Confirm Command", Prompt: fmt.Sprintf("Command: %v: %v", cm.Label(), cm.Desc)}, gi.AddOk, gi.AddCancel, ge.This(), func(recv, send ki.Ki, sig int64, data any) {
			if sig == int64(gi.DialogAccepted) {
				cm.RunAfterPrompts(ge, buf)
			}
		})
		return
	}
	pvals, hasp := cm.HasPrompts()
	if !hasp || CmdNoUserPrompt {
		cm.RunAfterPrompts(ge, buf)
		return
	}
	cm.PromptUser(ge, buf, pvals)
}

// RunAfterPrompts runs after any prompts have been set, if needed
func (cm *Command) RunAfterPrompts(ge Gide, buf *giv.TextBuf) {
	// ge.CmdRuns().KillByName(cm.Label()) // make sure nothing still running for us..
	CmdNoUserPrompt = false
	cdir := "{ProjPath}"
	if cm.Dir != "" {
		cdir = cm.Dir
	}
	cds := ge.ArgVarVals().Bind(cdir)
	err := os.Chdir(cds)
	cm.AppendCmdOut(ge, buf, []byte(fmt.Sprintf("cd %v (from: %v)\n", cds, cdir)))
	if err != nil {
		cm.AppendCmdOut(ge, buf, []byte(fmt.Sprintf("Could not change to directory %v -- error: %v\n", cds, err)))
	}

	if CmdWaitOverride || cm.Wait || len(cm.Cmds) > 1 {
		for i := range cm.Cmds {
			cma := &cm.Cmds[i]
			if buf == nil {
				if !cm.RunNoBuf(ge, cma) {
					break
				}
			} else {
				if !cm.RunBufWait(ge, buf, cma) {
					break
				}
			}
		}
	} else if len(cm.Cmds) > 0 {
		cma := &cm.Cmds[0]
		if buf == nil {
			go cm.RunNoBuf(ge, cma)
		} else {
			go cm.RunBuf(ge, buf, cma)
		}
	}
}

// RunBufWait runs a command with output to the buffer, using CombinedOutput
// so it waits for completion -- returns overall command success, and logs one
// line of the command output to gide statusbar
func (cm *Command) RunBufWait(ge Gide, buf *giv.TextBuf, cma *CmdAndArgs) bool {
	cmd, cmdstr := cma.PrepCmd(ge.ArgVarVals())
	ge.CmdRuns().AddCmd(cm.Label(), cmdstr, cma, cmd)
	out, err := cmd.CombinedOutput()
	cm.AppendCmdOut(ge, buf, out)
	return cm.RunStatus(ge, buf, cmdstr, err, out)
}

// RunBuf runs a command with output to the buffer, incrementally updating the
// buffer with new results line-by-line as they come in
func (cm *Command) RunBuf(ge Gide, buf *giv.TextBuf, cma *CmdAndArgs) bool {
	cmd, cmdstr := cma.PrepCmd(ge.ArgVarVals())
	ge.CmdRuns().AddCmd(cm.Label(), cmdstr, cma, cmd)
	stdout, err := cmd.StdoutPipe()
	if err == nil {
		cmd.Stderr = cmd.Stdout
		err = cmd.Start()
		if err == nil {
			obuf := giv.OutBuf{}
			obuf.Init(stdout, buf, 0, MarkupCmdOutput)
			obuf.MonOut()
		}
		err = cmd.Wait()
	}
	return cm.RunStatus(ge, buf, cmdstr, err, nil)
}

// RunNoBuf runs a command without any output to the buffer -- can call using
// go as a goroutine for no-wait case -- returns overall command success, and
// logs one line of the command output to gide statusbar
func (cm *Command) RunNoBuf(ge Gide, cma *CmdAndArgs) bool {
	cmd, cmdstr := cma.PrepCmd(ge.ArgVarVals())
	ge.CmdRuns().AddCmd(cm.Label(), cmdstr, cma, cmd)
	out, err := cmd.CombinedOutput()
	return cm.RunStatus(ge, nil, cmdstr, err, out)
}

// AppendCmdOut appends command output to buffer, applying markup for links
func (cm *Command) AppendCmdOut(ge Gide, buf *giv.TextBuf, out []byte) {
	if buf == nil {
		return
	}

	wupdt := ge.VPort().TopUpdateStart()
	defer ge.VPort().TopUpdateEnd(wupdt)

	buf.SetInactive(true)

	lns := bytes.Split(out, []byte("\n"))
	sz := len(lns)
	outmus := make([][]byte, sz)
	for i, txt := range lns {
		outmus[i] = MarkupCmdOutput(txt)
	}
	lfb := []byte("\n")
	mlns := bytes.Join(outmus, lfb)
	mlns = append(mlns, lfb...)

	buf.AppendTextMarkup(out, mlns, giv.EditSignal)
	buf.AutoScrollViews()
}

// CmdOutStatusLen is amount of command output to include in the status update
var CmdOutStatusLen = 80

// RunStatus reports the status of the command run (given in cmdstr) to
// ge.StatusBar -- returns true if there are no errors, and false if there
// were errors
func (cm *Command) RunStatus(ge Gide, buf *giv.TextBuf, cmdstr string, err error, out []byte) bool {
	ge.CmdRuns().DeleteByName(cm.Label())
	var rval bool
	outstr := ""
	if out != nil {
		outstr = string(out[:CmdOutStatusLen])
	}
	finstat := ""
	tstr := time.Now().Format("Mon Jan  2 15:04:05 MST 2006")
	if err == nil {
		finstat = fmt.Sprintf("%v <b>successful</b> at: %v", cmdstr, tstr)
		rval = true
	} else if ee, ok := err.(*exec.ExitError); ok {
		finstat = fmt.Sprintf("%v <b>failed</b> at: %v with error: %v", cmdstr, tstr, ee.Error())
		rval = false
	} else {
		finstat = fmt.Sprintf("%v <b>exec error</b> at: %v error: %v", cmdstr, tstr, err.Error())
		rval = false
	}
	if buf != nil {
		buf.SetInactive(true)
		if err != nil {
			ge.SelectTabByName(cm.Label()) // sometimes it isn't
		}
		fsb := []byte(finstat)
		buf.AppendTextLineMarkup([]byte(""), []byte(""), giv.EditSignal)
		buf.AppendTextLineMarkup(fsb, MarkupCmdOutput(fsb), giv.EditSignal)
		buf.RefreshViews()
		buf.AutoScrollViews()
		if cm.Focus {
			ge.FocusOnTabs()
		}
	}
	ge.SetStatus(cmdstr + " " + outstr)
	return rval
}

// LangMatch returns true if the given language matches the command Lang constraints
func (cm *Command) LangMatch(lang filecat.Supported) bool {
	return filecat.IsMatch(cm.Lang, lang)
}

// MarkupCmdOutput applies links to the first element in command output line
// if it looks like a file name / position
func MarkupCmdOutput(out []byte) []byte {
	flds := strings.Fields(string(out))
	if len(flds) == 0 {
		return out
	}
	orig, link := lex.MarkupPathsAsLinks(flds, 2) // only first 2 fields
	if len(link) > 0 {
		nt := bytes.Replace(out, orig, link, -1)
		return nt
	}
	return out
}

////////////////////////////////////////////////////////////////////////////////
//  Commands

// Commands is a list of different commands
type Commands []*Command

var KiT_Commands = kit.Types.AddType(&Commands{}, CommandsProps)

// CmdName has an associated ValueView for selecting from the list of
// available command names, for use in preferences etc.
// Formatted as Cat: Name as in Command.Label()
type CmdName string

// IsValid checks if command name exists on AvailCmds list
func (cn CmdName) IsValid() bool {
	_, _, ok := AvailCmds.CmdByName(cn, false)
	return ok
}

// Command returns command associated with command name in AvailCmds, and
// false if it doesn't exist
func (cn CmdName) Command() (*Command, bool) {
	cmd, _, ok := AvailCmds.CmdByName(cn, true)
	return cmd, ok
}

// CmdNames is a slice of command names
type CmdNames []CmdName

// Add adds a name to the list
func (cn *CmdNames) Add(cmd CmdName) {
	*cn = append(*cn, cmd)
}

// AvailCmds is the current list of ALL available commands for use -- it
// combines StdCmds and CustomCmds.  Custom overrides Std items with
// the same names.
var AvailCmds Commands

// CustomCmds is user-specific list of commands saved in preferences available
// for all Gide projects.  These will override StdCmds with the same names.
var CustomCmds = Commands{}

// FilterCmdNames returns a slice of commands organized by category
// that are compatible with given language and version control system.
func (cm *Commands) FilterCmdNames(lang filecat.Supported, vcnm giv.VersCtrlName) [][]string {
	vnm := strings.ToLower(string(vcnm))
	var cmds [][]string
	cat := ""
	var csub []string
	for _, cmd := range *cm {
		if cmd.LangMatch(lang) {
			if cmd.Cat != cat {
				lcat := strings.ToLower(cmd.Cat)
				if giv.IsVersCtrlSystem(lcat) && lcat != vnm {
					continue
				}
				cat = cmd.Cat
				csub = []string{cat}
				cmds = append(cmds, csub)
			}
			csub = append(csub, cmd.Name)
			cmds[len(cmds)-1] = csub // in case updated
		}
	}
	return cmds
}

func init() {
	AvailCmds.CopyFrom(StdCmds)
}

// CmdByName returns a command and index by name -- returns false and emits a
// message to log if not found if msg is true
func (cm *Commands) CmdByName(name CmdName, msg bool) (*Command, int, bool) {
	for i, cmd := range *cm {
		if cmd.Label() == string(name) {
			return cmd, i, true
		}
	}
	if msg {
		log.Printf("gi.Commands.CmdByName: command named: %v not found\n", name)
	}
	return nil, -1, false
}

// PrefsCmdsFileName is the name of the preferences file in App prefs
// directory for saving / loading your CustomCmds commands list
var PrefsCmdsFileName = "command_prefs.json"

// OpenJSON opens commands from a JSON-formatted file.
func (cm *Commands) OpenJSON(filename gi.FileName) error {
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		// gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, gi.AddOk, gi.AddCancel, nil, nil)
		// log.Println(err)
		return err
	}
	*cm = make(Commands, 0, 10) // reset
	return json.Unmarshal(b, cm)
}

// SaveJSON saves commands to a JSON-formatted file.
func (cm *Commands) SaveJSON(filename gi.FileName) error {
	b, err := json.MarshalIndent(cm, "", "  ")
	if err != nil {
		log.Println(err) // unlikely
		return err
	}
	err = ioutil.WriteFile(string(filename), b, 0644)
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, gi.AddOk, gi.NoCancel, nil, nil)
		log.Println(err)
	}
	return err
}

// OpenPrefs opens custom Commands from App standard prefs directory, using
// PrefsCmdsFileName
func (cm *Commands) OpenPrefs() error {
	pdir := oswin.TheApp.AppPrefsDir()
	pnm := filepath.Join(pdir, PrefsCmdsFileName)
	CustomCmdsChanged = false
	err := cm.OpenJSON(gi.FileName(pnm))
	if err == nil {
		MergeAvailCmds()
	} else {
		cm = &Commands{} // restore a blank
	}
	return err
}

// SavePrefs saves custom Commands to App standard prefs directory, using
// PrefsCmdsFileName
func (cm *Commands) SavePrefs() error {
	pdir := oswin.TheApp.AppPrefsDir()
	pnm := filepath.Join(pdir, PrefsCmdsFileName)
	CustomCmdsChanged = false
	err := cm.SaveJSON(gi.FileName(pnm))
	if err == nil {
		MergeAvailCmds()
	}
	return err
}

// CopyFrom copies commands from given other map
func (cm *Commands) CopyFrom(cp Commands) {
	*cm = make(Commands, 0, len(cp)) // reset
	b, err := json.Marshal(cp)
	if err != nil {
		fmt.Printf("json err: %v\n", err.Error())
	}
	json.Unmarshal(b, cm)
}

// MergeAvailCmds updates the AvailCmds list from CustomCmds and StdCmds
func MergeAvailCmds() {
	AvailCmds.CopyFrom(StdCmds)
	for _, cmd := range CustomCmds {
		_, idx, has := AvailCmds.CmdByName(CmdName(cmd.Label()), false)
		if has {
			AvailCmds[idx] = cmd // replace
		} else {
			AvailCmds = append(AvailCmds, cmd)
		}
	}
}

// ViewStd shows the standard types that are compiled into the program and have
// all the lastest standards.  Useful for comparing against custom lists.
func (cm *Commands) ViewStd() {
	CmdsView(&StdCmds)
}

// CustomCmdsChanged is used to update giv.CmdsView toolbars via following
// menu, toolbar props update methods.
var CustomCmdsChanged = false

// CommandsProps define the ToolBar and MenuBar for TableView of Commands, e.g., CmdsView
var CommandsProps = ki.Props{
	"MainMenu": ki.PropSlice{
		{"AppMenu", ki.BlankProp{}},
		{"File", ki.PropSlice{
			{"OpenPrefs", ki.Props{}},
			{"SavePrefs", ki.Props{
				"shortcut": "Command+S",
				"updtfunc": giv.ActionUpdateFunc(func(cmi any, act *gi.Action) {
					act.SetActiveState(CustomCmdsChanged && cmi.(*Commands) == &CustomCmds)
				}),
			}},
			{"sep-file", ki.BlankProp{}},
			{"OpenJSON", ki.Props{
				"label":    "Open from file",
				"desc":     "You can save and open commands to / from files to share, experiment, transfer, etc",
				"shortcut": "Command+O",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
			{"SaveJSON", ki.Props{
				"label": "Save to file",
				"desc":  "You can save and open commands to / from files to share, experiment, transfer, etc",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
		}},
		{"Edit", "Copy Cut Paste Dupe"},
		{"Window", "Windows"},
	},
	"ToolBar": ki.PropSlice{
		{"SavePrefs", ki.Props{
			"desc": "saves Commands to App standard prefs directory, in file proj_types_prefs.json, which will be loaded automatically at startup if prefs SaveCommands is checked (should be if you're using custom commands)",
			"icon": "file-save",
			"updtfunc": giv.ActionUpdateFunc(func(cmi any, act *gi.Action) {
				act.SetActiveState(CustomCmdsChanged && cmi.(*Commands) == &CustomCmds)
			}),
		}},
		{"sep-file", ki.BlankProp{}},
		{"OpenJSON", ki.Props{
			"label": "Open from file",
			"icon":  "file-open",
			"desc":  "You can save and open commands to / from files to share, experiment, transfer, etc",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".json",
				}},
			},
		}},
		{"SaveJSON", ki.Props{
			"label": "Save to file",
			"icon":  "file-save",
			"desc":  "You can save and open commands to / from files to share, experiment, transfer, etc",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".json",
				}},
			},
		}},
		{"sep-std", ki.BlankProp{}},
		{"ViewStd", ki.Props{
			"desc": "Shows the standard commands that are compiled into the program (edits will not be saved -- even though the viewer is editable).  Custom commands override standard ones of the same name, so that is the way to change any existing commands.",
			"updtfunc": giv.ActionUpdateFunc(func(cmi any, act *gi.Action) {
				act.SetActiveState(cmi.(*Commands) != &StdCmds)
			}),
		}},
	},
}

// SetCompleter adds a completer to the textfield - each field
// can have its own match and edit functions
// For this to be called add a "complete" tag to the struct field
func (cmd *Command) SetCompleter(tf *gi.TextField, id string) {
	if id == "arg" {
		tf.SetCompleter(cmd, CompleteArg, CompleteArgEdit)
		return
	}
	fmt.Printf("no match for SetCompleter id argument")
}

// CompleteArg supplies directory variables to the completer
func CompleteArg(data any, text string, posLn, posCh int) (md complete.Matches) {
	md.Seed = complete.SeedWhiteSpace(text)
	possibles := complete.MatchSeedString(ArgVarKeys(), md.Seed)
	for _, p := range possibles {
		m := complete.Completion{Text: p, Icon: ""}
		md.Matches = append(md.Matches, m)
	}
	return md
}

// CompleteArgEdit edits completer text field after the user chooses from the candidate completions
func CompleteArgEdit(data any, text string, cursorPos int, c complete.Completion, seed string) (ed complete.Edit) {
	ed = complete.EditWord(text, cursorPos, c.Text, seed)
	return ed
}
