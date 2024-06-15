// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package code

import (
	"fmt"
	"strings"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/filetree"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/keymap"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/states"
	"cogentcore.org/core/tree"
)

func (cv *Code) MakeToolbar(p *core.Plan) { //types:add
	fmt.Printf("cv make toolbar: %p\n", cv)
	core.AddInit(p, "app-chooser", func(w *core.Chooser) {
		cv.AddChooserFiles(w)
		cv.AddChooserSymbols(w)
		if tb := core.ParentToolbar(w); tb != nil {
			tb.AddOverflowMenu(cv.OverflowMenu)
		}
		w.OnFirst(events.KeyChord, func(e events.Event) {
			kf := keymap.Of(e.KeyChord())
			if kf == keymap.Abort {
				w.ClearError()
				cv.FocusActiveTextEditor()
			}
		})
	})
	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.UpdateFiles).SetText("").SetIcon(icons.Refresh).SetShortcut("Command+U")
	})
	core.Add(p, func(w *core.Switch) {
		w.SetText("Go mod").SetTooltip("Toggles the use of go modules -- saved with project -- if off, uses old school GOPATH mode")
		w.Styler(func(s *styles.Style) {
			w.SetChecked(cv.Settings.GoMod) // todo: update
		})
		w.OnChange(func(e events.Event) {
			cv.Settings.GoMod = w.StateIs(states.Checked)
			SetGoMod(cv.Settings.GoMod)
		})
	})

	core.Add(p, func(w *core.Separator) {})

	core.Add(p, func(w *core.Button) {
		w.SetText("Open recent").SetMenu(func(m *core.Scene) {
			for _, rp := range RecentPaths {
				core.NewButton(m).SetText(rp).OnClick(func(e events.Event) {
					cv.OpenRecent(core.Filename(rp))
				})
			}
			core.NewSeparator(m)
			core.NewButton(m).SetText("Recent recent paths").OnClick(func(e events.Event) {
				RecentPaths = nil
			})
			core.NewButton(m).SetText("Edit recent paths").OnClick(func(e events.Event) {
				cv.EditRecentPaths()
			})
		})
	})
	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.OpenPath).
			SetText("Open").SetIcon(icons.Open).SetKey(keymap.Open)
		cv.ConfigActiveFilename(w)
	})
	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.SaveActiveView).SetText("Save").
			SetIcon(icons.Save).SetKey(keymap.Save)
	})
	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.SaveAll).SetIcon(icons.Save)
	})

	core.Add(p, func(w *core.Separator) {})

	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.CursorToHistPrev).SetText("").SetKey(keymap.HistPrev).
			SetIcon(icons.KeyboardArrowLeft)
	})
	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.CursorToHistNext).SetText("").SetKey(keymap.HistNext).
			SetIcon(icons.KeyboardArrowRight)
	})

	core.Add(p, func(w *core.Separator) {})

	// todo: this does not work to apply project defaults!
	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.Find).SetIcon(icons.FindReplace)
		cv.ConfigFindButton(w)
	})

	core.Add(p, func(w *core.Separator) {})

	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.Symbols).SetIcon(icons.List)
	})
	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.Spell).SetIcon(icons.Spellcheck)
	})

	core.Add(p, func(w *core.Separator) {})

	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.RunBuild).SetText("Build").SetIcon(icons.Build).
			SetShortcut(KeyBuildProject.Chord())
	})
	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.Run).SetIcon(icons.PlayArrow).
			SetShortcut(KeyRunProject.Chord())
	})
	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.Debug).SetIcon(icons.Debug)
	})
	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.DebugTest).SetIcon(icons.Debug)
	})

	core.Add(p, func(w *core.Separator) {})

	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.Commit).SetIcon(icons.Star)
	})
	core.Add(p, func(w *core.FuncButton) {
		fmt.Printf("add vcs log: %p\n", cv)
		w.SetFunc(cv.VCSLog).SetText("VCS Log").SetIcon(icons.List)
	})
	core.Add(p, func(w *core.FuncButton) {
		w.SetFunc(cv.EditProjectSettings).SetText("Settings").SetIcon(icons.Edit)
	})

	core.Add(p, func(w *core.Button) {
		w.SetText("Command").
			SetShortcut(KeyExecCmd.Chord()).
			SetMenu(func(m *core.Scene) {
				ec := ExecCmds(cv)
				for _, cc := range ec {
					cc := cc
					cat := cc[0]
					ic := icons.Icon(strings.ToLower(cat))
					core.NewButton(m).SetText(cat).SetIcon(ic).SetMenu(func(mm *core.Scene) {
						nc := len(cc)
						for i := 1; i < nc; i++ {
							cm := cc[i]
							core.NewButton(mm).SetText(cm).SetIcon(ic).OnClick(func(e events.Event) {
								e.SetHandled()
								cv.ExecCmdNameActive(CommandName(cat, cm))
							})
						}
					})
				}
			})
	})

	core.Add(p, func(w *core.Separator) {})

	core.Add(p, func(w *core.Button) {
		w.SetText("Splits").SetMenu(func(m *core.Scene) {
			core.NewButton(m).SetText("Set View").
				SetMenu(func(mm *core.Scene) {
					for _, sp := range AvailableSplitNames {
						sn := SplitName(sp)
						mb := core.NewButton(mm).SetText(sp).OnClick(func(e events.Event) {
							cv.SplitsSetView(sn)
						})
						if sn == cv.Settings.SplitName {
							mb.SetSelected(true)
						}
					}
				})
			core.NewFuncButton(m).SetFunc(cv.SplitsSaveAs).SetText("Save As")
			core.NewButton(m).SetText("Save").
				SetMenu(func(mm *core.Scene) {
					for _, sp := range AvailableSplitNames {
						sn := SplitName(sp)
						mb := core.NewButton(mm).SetText(sp).OnClick(func(e events.Event) {
							cv.SplitsSave(sn)
						})
						if sn == cv.Settings.SplitName {
							mb.SetSelected(true)
						}
					}
				})
			core.NewFuncButton(m).SetFunc(cv.SplitsEdit).SetText("Edit")
		})
	})
}

func (cv *Code) OverflowMenu(m *core.Scene) {
	core.NewButton(m).SetText("File").SetMenu(func(m *core.Scene) {
		core.NewFuncButton(m).SetFunc(cv.NewProject).SetIcon(icons.NewWindow).SetKey(keymap.New)
		core.NewFuncButton(m).SetFunc(cv.NewFile).SetText("New File").SetIcon(icons.NewWindow)

		core.NewSeparator(m)

		core.NewFuncButton(m).SetFunc(cv.OpenProject).SetIcon(icons.Open)

		core.NewSeparator(m)

		core.NewFuncButton(m).SetFunc(cv.SaveProject).SetIcon(icons.Save)

		cv.ConfigActiveFilename(core.NewFuncButton(m).SetFunc(cv.SaveProjectAs)).SetIcon(icons.SaveAs)

		core.NewSeparator(m)

		core.NewFuncButton(m).SetFunc(cv.RevertActiveView).SetText("Revert File").
			SetIcon(icons.Undo)

		cv.ConfigActiveFilename(core.NewFuncButton(m).SetFunc(cv.SaveActiveViewAs)).
			SetText("Save File As").SetIcon(icons.SaveAs).SetKey(keymap.SaveAs)

	})

	core.NewButton(m).SetText("Edit").SetMenu(func(m *core.Scene) {
		core.NewButton(m).SetText("Paste history").SetIcon(icons.Paste).
			SetKey(keymap.PasteHist)

		core.NewFuncButton(m).SetFunc(cv.RegisterPaste).SetIcon(icons.Paste).
			SetShortcut(KeyRegPaste.Chord())
		core.NewFuncButton(m).SetFunc(cv.RegisterCopy).SetIcon(icons.Copy).
			SetShortcut(KeyRegCopy.Chord())

		core.NewSeparator(m)

		core.NewFuncButton(m).SetFunc(cv.CopyRect).SetIcon(icons.Copy).
			SetShortcut(KeyRectCopy.Chord())
		core.NewFuncButton(m).SetFunc(cv.CutRect).SetIcon(icons.Cut).
			SetShortcut(KeyRectCut.Chord())
		core.NewFuncButton(m).SetFunc(cv.PasteRect).SetIcon(icons.Paste).
			SetShortcut(KeyRectPaste.Chord())

		core.NewSeparator(m)

		core.NewButton(m).SetText("Undo").SetIcon(icons.Undo).SetKey(keymap.Undo)
		core.NewButton(m).SetText("Redo").SetIcon(icons.Redo).SetKey(keymap.Redo)

		core.NewSeparator(m)

		core.NewFuncButton(m).SetFunc(cv.ReplaceInActive).SetText("Replace in File").
			SetIcon(icons.FindReplace)

		core.NewButton(m).SetText("Show completions").SetIcon(icons.CheckCircle).SetKey(keymap.Complete)
		core.NewButton(m).SetText("Lookup symbol").SetIcon(icons.Search).SetKey(keymap.Lookup)
		core.NewButton(m).SetText("Jump to line").SetIcon(icons.GoToLine).SetKey(keymap.Jump)

		core.NewSeparator(m)

		core.NewFuncButton(m).SetFunc(cv.CommentOut).SetText("Comment region").
			SetIcon(icons.Comment).SetShortcut(KeyCommentOut.Chord())
		core.NewFuncButton(m).SetFunc(cv.Indent).SetIcon(icons.FormatIndentIncrease).
			SetShortcut(KeyIndent.Chord())
		core.NewFuncButton(m).SetFunc(cv.ReCase).SetIcon(icons.MatchCase)
		core.NewFuncButton(m).SetFunc(cv.JoinParaLines).SetIcon(icons.Join)
		core.NewFuncButton(m).SetFunc(cv.TabsToSpaces).SetIcon(icons.TabMove)
		core.NewFuncButton(m).SetFunc(cv.SpacesToTabs).SetIcon(icons.TabMove)
	})

	core.NewButton(m).SetText("View").SetMenu(func(m *core.Scene) {
		core.NewFuncButton(m).SetFunc(cv.FocusPrevPanel).SetText("Focus prev").SetIcon(icons.KeyboardArrowLeft).
			SetShortcut(KeyPrevPanel.Chord())
		core.NewFuncButton(m).SetFunc(cv.FocusNextPanel).SetText("Focus next").SetIcon(icons.KeyboardArrowRight).
			SetShortcut(KeyNextPanel.Chord())
		core.NewFuncButton(m).SetFunc(cv.CloneActiveView).SetText("Clone active").SetIcon(icons.Copy).
			SetShortcut(KeyBufClone.Chord())

		core.NewSeparator(m)

		core.NewFuncButton(m).SetFunc(cv.CloseActiveView).SetText("Close file").SetIcon(icons.Close).
			SetShortcut(KeyBufClose.Chord())
		core.NewFuncButton(m).SetFunc(cv.OpenConsoleTab).SetText("Open console").SetIcon(icons.Terminal)
	})

	core.NewButton(m).SetText("Command").SetMenu(func(m *core.Scene) {
		core.NewFuncButton(m).SetFunc(cv.DebugAttach).SetText("Debug attach").SetIcon(icons.Debug)
		core.NewFuncButton(m).SetFunc(cv.VCSUpdateAll).SetText("VCS update all").SetIcon(icons.Update)

		core.NewSeparator(m)

		core.NewFuncButton(m).SetFunc(cv.CountWords).SetText("Count words all").SetIcon(icons.Counter5)
		core.NewFuncButton(m).SetFunc(cv.CountWordsRegion).SetText("Count words region").SetIcon(icons.Counter3)

		core.NewSeparator(m)

		core.NewFuncButton(m).SetFunc(cv.HelpWiki).SetText("Help").SetIcon(icons.Help)
	})

	core.NewSeparator(m)
}

// AddChooserFiles adds the files to the app chooser.
func (cv *Code) AddChooserFiles(ac *core.Chooser) {
	ac.AddItemsFunc(func() {
		if cv.Files == nil {
			return
		}
		cv.Files.WidgetWalkDown(func(wi core.Widget, wb *core.WidgetBase) bool {
			fn := filetree.AsNode(wi)
			if fn == nil || fn.IsIrregular() {
				return tree.Continue
			}
			rpath := fn.MyRelPath()
			nmpath := fn.Name + ":" + rpath
			switch {
			case fn.IsDir():
				ac.Items = append(ac.Items, core.ChooserItem{
					Text: nmpath,
					Icon: icons.Folder,
					Func: func() {
						if !fn.HasChildren() {
							fn.OpenEmptyDir()
						}
						fn.Open()
						fn.ScrollToMe()
						ac.CallItemsFuncs() // refresh avail files
					},
				})
			case fn.IsExec():
				ac.Items = append(ac.Items, core.ChooserItem{
					Text: nmpath,
					Icon: icons.FileExe,
					Func: func() {
						cv.FileNodeRunExe(fn)
					},
				})
			default:
				ac.Items = append(ac.Items, core.ChooserItem{
					Text: nmpath,
					Icon: fn.Info.Ic,
					Func: func() {
						cv.NextViewFileNode(fn)
						ac.CallItemsFuncs() // refresh avail files
					},
				})
			}
			return tree.Continue
		})
	})
}

// AddChooserSymbols adds the symbols to the app chooser.
func (cv *Code) AddChooserSymbols(ac *core.Chooser) {
	ac.AddItemsFunc(func() {
		tv := cv.ActiveTextEditor()
		if tv == nil || tv.Buffer == nil || !tv.Buffer.Hi.UsingParse() {
			return
		}
		pfs := tv.Buffer.ParseState.Done()
		if len(pfs.ParseState.Scopes) == 0 {
			return
		}
		pkg := pfs.ParseState.Scopes[0] // first scope of parse state is the full set of package symbols
		syms := NewSymNode()
		syms.SetName("syms")
		syms.OpenSyms(pkg, "", "")
		syms.WalkDown(func(k tree.Node) bool {
			sn := k.(*SymNode)
			ac.Items = append(ac.Items, core.ChooserItem{
				Text: sn.Symbol.Label(),
				Icon: sn.GetIcon(),
				Func: func() {
					SelectSymbol(cv, sn.Symbol)
				},
			})
			return tree.Continue
		})
	})
}
