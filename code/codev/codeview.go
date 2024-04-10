// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codev

//go:generate core generate

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"cogentcore.org/cogent/code/code"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/events/key"
	"cogentcore.org/core/fileinfo"
	"cogentcore.org/core/filetree"
	"cogentcore.org/core/grr"
	"cogentcore.org/core/spell"
	"cogentcore.org/core/texteditor"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/vci"
	"cogentcore.org/core/views"
)

// CodeView is the core editor and tab viewer framework for the Code system.  The
// default view has a tree browser of files on the left, editor panels in the
// middle, and a tabbed viewer on the right.
type CodeView struct {
	core.Frame

	// root directory for the project -- all projects must be organized within a top-level root directory, with all the files therein constituting the scope of the project -- by default it is the path for ProjFilename
	ProjRoot core.Filename

	// current project filename for saving / loading specific Code configuration information in a .code file (optional)
	ProjFilename core.Filename `ext:".code"`

	// filename of the currently-active texteditor
	ActiveFilename core.Filename `set:"-"`

	// language for current active filename
	ActiveLang fileinfo.Known

	// VCS repo for current active filename
	ActiveVCS vci.Repo `set:"-"`

	// VCS info for current active filename (typically branch or revision) -- for status
	ActiveVCSInfo string `set:"-"`

	// has the root changed?  we receive update signals from root for changes
	Changed bool `set:"-" json:"-"`

	// the last status update message
	StatusMessage string

	// timestamp for when a file was last saved -- provides dirty state for various updates including rebuilding in debugger
	LastSaveTStamp time.Time `set:"-" json:"-"`

	// all the files in the project directory and subdirectories
	Files *filetree.Tree `set:"-" json:"-"`

	// index of the currently-active texteditor -- new files will be viewed in other views if available
	ActiveTextEditorIndex int `set:"-" json:"-"`

	// list of open nodes, most recent first
	OpenNodes code.OpenNodes `json:"-"`

	// the command buffers for commands run in this project
	CmdBufs map[string]*texteditor.Buffer `set:"-" json:"-"`

	// history of commands executed in this session
	CmdHistory code.CmdNames `set:"-" json:"-"`

	// currently running commands in this project
	RunningCmds code.CmdRuns `set:"-" json:"-" xml:"-"`

	// current arg var vals
	ArgVals code.ArgVarVals `set:"-" json:"-" xml:"-"`

	// settings for this project -- this is what is saved in a .code project file
	Settings code.ProjSettings `set:"-"`

	// current debug view
	CurDbg *code.DebugView `set:"-"`

	// first key in sequence if needs2 key pressed
	KeySeq1 key.Chord `set:"-"`

	// mutex for protecting overall updates to CodeView
	UpdateMu sync.Mutex `set:"-"`
}

func init() {
	// core.URLHandler = URLHandler
	// paint.TextLinkHandler = TextLinkHandler
}

func (ge *CodeView) OnInit() {
	ge.Frame.OnInit()
	ge.HandleEvents()
}

func (ge *CodeView) OnAdd() {
	ge.Frame.OnAdd()
	ge.AddCloseDialog()
}

////////////////////////////////////////////////////////
// Code interface

func (ge *CodeView) ProjSettings() *code.ProjSettings {
	return &ge.Settings
}

func (ge *CodeView) FileTree() *filetree.Tree {
	return ge.Files
}

func (ge *CodeView) LastSaveTime() time.Time {
	return ge.LastSaveTStamp
}

// VersionControl returns the version control system in effect, using the file tree detected
// version or whatever is set in project settings
func (ge *CodeView) VersionControl() filetree.VersionControlName {
	vc := ge.Settings.VersionControl
	return vc
}

func (ge *CodeView) CmdRuns() *code.CmdRuns {
	return &ge.RunningCmds
}

func (ge *CodeView) CmdHist() *code.CmdNames {
	return &ge.CmdHistory
}

func (ge *CodeView) ArgVarVals() *code.ArgVarVals {
	return &ge.ArgVals
}

func (ge *CodeView) CurOpenNodes() *code.OpenNodes {
	return &ge.OpenNodes
}

func (ge *CodeView) FocusOnTabs() bool {
	return ge.FocusOnPanel(TabsIndex)
}

////////////////////////////////////////////////////////
//  Main project API

// UpdateFiles updates the list of files saved in project
func (ge *CodeView) UpdateFiles() { //gti:add
	if ge.Files != nil && ge.ProjRoot != "" {
		ge.Files.OpenPath(string(ge.ProjRoot))
		ge.Files.Open()
	}
}

func (ge *CodeView) IsEmpty() bool {
	return ge.ProjRoot == ""
}

// OpenRecent opens a recently-used file
func (ge *CodeView) OpenRecent(filename core.Filename) { //gti:add
	ext := strings.ToLower(filepath.Ext(string(filename)))
	if ext == ".code" {
		ge.OpenProj(filename)
	} else {
		ge.OpenPath(filename)
	}
}

// EditRecentPaths opens a dialog editor for editing the recent project paths list
func (ge *CodeView) EditRecentPaths() {
	d := core.NewBody().AddTitle("Recent project paths").
		AddText("You can delete paths you no longer use")
	views.NewSliceView(d).SetSlice(&code.RecentPaths)
	d.AddOKOnly().NewDialog(ge).Run()
}

// OpenFile opens file in an open project if it has the same path as the file
// or in a new window.
func (ge *CodeView) OpenFile(fnm string) { //gti:add
	abfn, _ := filepath.Abs(fnm)
	if strings.HasPrefix(abfn, string(ge.ProjRoot)) {
		ge.ViewFile(core.Filename(abfn))
		return
	}
	for _, win := range core.MainRenderWindows {
		msc := win.MainScene()
		geo := CodeInScene(msc)
		if strings.HasPrefix(abfn, string(geo.ProjRoot)) {
			geo.ViewFile(core.Filename(abfn))
			return
		}
	}
	// fmt.Printf("open path: %s\n", ge.ProjRoot)
	ge.OpenPath(core.Filename(abfn))
}

// SetWindowNameTitle sets the window name and title based on current project name
func (ge *CodeView) SetWindowNameTitle() {
	pnm := ge.Name()
	title := "Cogent Code • " + pnm
	ge.Scene.Body.SetTitle(title)
}

// OpenPath creates a new project by opening given path, which can either be a
// specific file or a folder containing multiple files of interest -- opens in
// current CodeView object if it is empty, or otherwise opens a new window.
func (ge *CodeView) OpenPath(path core.Filename) *CodeView { //gti:add
	if gproj, has := CheckForProjAtPath(string(path)); has {
		return ge.OpenProj(core.Filename(gproj))
	}
	if !ge.IsEmpty() {
		return NewCodeProjPath(string(path))
	}
	ge.Defaults()
	root, pnm, fnm, ok := ProjPathParse(string(path))
	if ok {
		os.Chdir(root)
		code.RecentPaths.AddPath(root, core.SystemSettings.SavedPathsMax)
		code.SavePaths()
		ge.ProjRoot = core.Filename(root)
		ge.SetName(pnm)
		ge.Scene.SetName(pnm)
		ge.Settings.ProjFilename = core.Filename(filepath.Join(root, pnm+".code"))
		ge.ProjFilename = ge.Settings.ProjFilename
		ge.Settings.ProjRoot = ge.ProjRoot
		ge.GuessMainLang()
		ge.LangDefaults()
		ge.SetWindowNameTitle()
		ge.UpdateFiles()
		ge.SplitsSetView(code.SplitName(code.AvailableSplitNames[0]))
		if fnm != "" {
			ge.NextViewFile(core.Filename(fnm))
		}
	}
	return ge
}

// OpenProj opens .code project file and its settings from given filename, in a standard
// toml-formatted file
func (ge *CodeView) OpenProj(filename core.Filename) *CodeView { //gti:add
	if !ge.IsEmpty() {
		return OpenCodeProj(string(filename))
	}
	ge.Defaults()
	if err := ge.Settings.Open(filename); err != nil {
		slog.Error("Project Settings had a loading error", "error", err)
		if ge.Settings.ProjRoot == "" {
			root, _, _, _ := ProjPathParse(string(filename))
			ge.Settings.ProjRoot = core.Filename(root)
			ge.GuessMainLang()
		}
	}
	ge.Settings.ProjFilename = filename // should already be set but..
	_, pnm, _, ok := ProjPathParse(string(ge.Settings.ProjRoot))
	if ok {
		code.SetGoMod(ge.Settings.GoMod)
		os.Chdir(string(ge.Settings.ProjRoot))
		ge.ProjRoot = core.Filename(ge.Settings.ProjRoot)
		code.RecentPaths.AddPath(string(filename), core.SystemSettings.SavedPathsMax)
		code.SavePaths()
		ge.SetName(pnm)
		ge.Scene.SetName(pnm)
		ge.ApplySettings()
		ge.UpdateFiles()
		ge.SetWindowNameTitle()
	}
	return ge
}

// NewProj creates a new project at given path, making a new folder in that
// path -- all CodeView projects are essentially defined by a path to a folder
// containing files.  If the folder already exists, then use OpenPath.
// Can also specify main language and version control type
func (ge *CodeView) NewProj(path core.Filename, folder string, mainLang fileinfo.Known, VersionControl filetree.VersionControlName) *CodeView { //gti:add
	np := filepath.Join(string(path), folder)
	err := os.MkdirAll(np, 0775)
	if err != nil {
		core.MessageDialog(ge, fmt.Sprintf("Could not make folder for project at: %v, err: %v", np, err), "Could not Make Folder")
		return nil
	}
	nge := ge.OpenPath(core.Filename(np))
	nge.Settings.MainLang = mainLang
	if VersionControl != "" {
		nge.Settings.VersionControl = VersionControl
	}
	return nge
}

// NewFile creates a new file in the project
func (ge *CodeView) NewFile(filename string, addToVcs bool) { //gti:add
	np := filepath.Join(string(ge.ProjRoot), filename)
	_, err := os.Create(np)
	if err != nil {
		core.MessageDialog(ge, fmt.Sprintf("Could not make new file at: %v, err: %v", np, err), "Could not Make File")
		return
	}
	ge.Files.UpdatePath(np)
	if addToVcs {
		nfn, ok := ge.Files.FindFile(np)
		if ok {
			nfn.AddToVCS()
		}
	}
}

// SaveProj saves project file containing custom project settings, in a
// standard toml-formatted file
func (ge *CodeView) SaveProj() { //gti:add
	if ge.Settings.ProjFilename == "" {
		return
	}
	ge.SaveProjAs(ge.Settings.ProjFilename)
	ge.SaveAllCheck(false, nil) // false = no cancel option
}

// SaveProjIfExists saves project file containing custom project settings, in a
// standard toml-formatted file, only if it already exists -- returns true if saved
// saveAllFiles indicates if user should be prompted for saving all files
func (ge *CodeView) SaveProjIfExists(saveAllFiles bool) bool {
	spell.SaveIfLearn()
	if ge.Settings.ProjFilename == "" {
		return false
	}
	if _, err := os.Stat(string(ge.Settings.ProjFilename)); os.IsNotExist(err) {
		return false // does not exist
	}
	ge.SaveProjAs(ge.Settings.ProjFilename)
	if saveAllFiles {
		ge.SaveAllCheck(false, nil)
	}
	return true
}

// SaveProjAs saves project custom settings to given filename, in a standard
// toml-formatted file
// saveAllFiles indicates if user should be prompted for saving all files
// returns true if the user was prompted, false otherwise
func (ge *CodeView) SaveProjAs(filename core.Filename) bool { //gti:add
	spell.SaveIfLearn()
	code.RecentPaths.AddPath(string(filename), core.SystemSettings.SavedPathsMax)
	code.SavePaths()
	ge.Settings.ProjFilename = filename
	ge.ProjFilename = ge.Settings.ProjFilename
	ge.GrabSettings()
	ge.Settings.Save(filename)
	ge.Files.UpdatePath(string(filename))
	ge.Changed = false
	return false
}

// SaveAllCheck -- check if any files have not been saved, and prompt to save them
// returns true if there were unsaved files, false otherwise.
// cancelOpt presents an option to cancel current command, in which case function is not called.
// if function is passed, then it is called in all cases except if the user selects cancel.
func (ge *CodeView) SaveAllCheck(cancelOpt bool, fun func()) bool {
	nch := ge.NChangedFiles()
	if nch == 0 {
		if fun != nil {
			fun()
		}
		return false
	}
	d := core.NewBody().AddTitle("There are Unsaved Files").
		AddText(fmt.Sprintf("In Project: %v There are <b>%v</b> opened files with <b>unsaved changes</b> -- do you want to save all?", ge.Nm, nch))
	d.AddBottomBar(func(parent core.Widget) {
		if cancelOpt {
			d.AddCancel(parent).SetText("Cancel Command")
		}
		core.NewButton(parent).SetText("Don't Save").OnClick(func(e events.Event) {
			d.Close()
			if fun != nil {
				fun()
			}
		})
		core.NewButton(parent).SetText("Save All").OnClick(func(e events.Event) {
			d.Close()
			ge.SaveAllOpenNodes()
			if fun != nil {
				fun()
			}
		})
	})
	d.NewDialog(ge).Run()
	return true
}

// ProjPathParse parses given project path into a root directory (which could
// be the path or just the directory portion of the path, depending in whether
// the path is a directory or not), and a bool if all is good (otherwise error
// message has been reported). projnm is always the last directory of the path.
func ProjPathParse(path string) (root, projnm, fnm string, ok bool) {
	if path == "" {
		return "", "blank", "", false
	}
	effpath := grr.Log1(filepath.EvalSymlinks(path))
	info, err := os.Lstat(effpath)
	if err != nil {
		emsg := fmt.Errorf("code.ProjPathParse: Cannot open at given path: %q: Error: %v", effpath, err)
		log.Println(emsg)
		return
	}
	path, _ = filepath.Abs(path)
	dir, fn := filepath.Split(path)
	pathIsDir := info.IsDir()
	if pathIsDir {
		root = path
		projnm = fn
	} else {
		root = filepath.Clean(dir)
		_, projnm = filepath.Split(root)
		fnm = fn
	}
	ok = true
	return
}

// CheckForProjAtPath checks if there is a .code project at the given path
// returns project path and true if found, otherwise false
func CheckForProjAtPath(path string) (string, bool) {
	root, pnm, _, ok := ProjPathParse(path)
	if !ok {
		return "", false
	}
	gproj := filepath.Join(root, pnm+".code")
	if _, err := os.Stat(gproj); os.IsNotExist(err) {
		return "", false // does not exist
	}
	return gproj, true
}

//////////////////////////////////////////////////////////////////////////////////////
//   Close / Quit Req

// NChangedFiles returns number of opened files with unsaved changes
func (ge *CodeView) NChangedFiles() int {
	return ge.OpenNodes.NChanged()
}

// AddCloseDialog adds the close dialog that automatically saves the project
// and prompts the user to save open files when they try to close the scene
// containing this code view.
func (ge *CodeView) AddCloseDialog() {
	ge.WidgetBase.AddCloseDialog(func(d *core.Body) bool {
		ge.SaveProjIfExists(false) // don't prompt here, as we will do it now..
		nch := ge.NChangedFiles()
		if nch == 0 {
			return false
		}
		d.AddTitle("Unsaved files").
			AddText(fmt.Sprintf("There are %d open files in %s with unsaved changes", nch, ge.Nm))
		d.AddBottomBar(func(parent core.Widget) {
			d.AddOK(parent, "cws").SetText("Close without saving").OnClick(func(e events.Event) {
				ge.Scene.Close()
			})
			core.NewButton(parent, "sa").SetText("Save and close").OnClick(func(e events.Event) {
				ge.SaveAllOpenNodes()
				ge.Scene.Close()
			})
		})
		return true
	})
}

//////////////////////////////////////////////////////////////////////////////////////
//   Project window

// NewCodeProjPath creates a new CodeView window with a new CodeView project for given
// path, returning the window and the path
func NewCodeProjPath(path string) *CodeView {
	root, projnm, _, _ := ProjPathParse(path)
	return NewCodeWindow(path, projnm, root, true)
}

// OpenCodeProj creates a new CodeView window opened to given CodeView project,
// returning the window and the path
func OpenCodeProj(projfile string) *CodeView {
	pp := &code.ProjSettings{}
	if err := pp.Open(core.Filename(projfile)); err != nil {
		slog.Debug("Project Settings had a loading error", "error", err)
	}
	path := string(pp.ProjRoot)
	root, projnm, _, _ := ProjPathParse(path)
	return NewCodeWindow(projfile, projnm, root, false)
}

func CodeInScene(sc *core.Scene) *CodeView {
	gv := sc.Body.ChildByType(CodeViewType, tree.NoEmbeds)
	if gv != nil {
		return gv.(*CodeView)
	}
	return nil
}

// NewCodeWindow is common code for Open CodeWindow from Proj or Path
func NewCodeWindow(path, projnm, root string, doPath bool) *CodeView {
	winm := "Cogent Code • " + projnm

	if win, found := core.AllRenderWindows.FindName(winm); found {
		sc := win.MainScene()
		ge := CodeInScene(sc)
		if ge != nil && string(ge.ProjRoot) == root {
			win.Raise()
			return ge
		}
	}

	b := core.NewBody(winm).SetTitle(winm)

	ge := NewCodeView(b)
	core.TheApp.AppBarConfig = ge.AppBarConfig
	b.AddAppBar(ge.ConfigToolbar)

	b.NewWindow().Run()

	if doPath {
		ge.OpenPath(core.Filename(path))
	} else {
		ge.OpenProj(core.Filename(path))
	}

	return ge
}
