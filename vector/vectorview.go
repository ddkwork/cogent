// Copyright (c) 2021, The Vector Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vector

//go:generate core generate

import (
	"errors"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"cogentcore.org/core/events"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/goosi"
	"cogentcore.org/core/ki"
	"cogentcore.org/core/mat32"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/svg"
	"cogentcore.org/core/units"
)

// VectorView is the Vector SVG vector drawing program: Go-rendered interactive drawing
//
//core:no-new
type VectorView struct {
	gi.Frame

	// full path to current drawing filename
	Filename gi.Filename `ext:".svg"`

	// current edit state
	EditState EditState
}

// AddNewVectorView adds a new editor to given parent node, with given name.
func AddNewVectorView(parent ki.Ki, name string) *VectorView {
	vv := parent.AddNewChild(KiT_VectorView, name).(*VectorView)
	vv.Defaults()
	return vv
}

func (g *VectorView) CopyFieldsFrom(frm any) {
	fr := frm.(*VectorView)
	g.Frame.CopyFieldsFrom(&fr.Frame)
	// todo: fill out
}

func (vv *VectorView) Defaults() {
	es := &vv.EditState
	es.ConfigDefaultGradient()
}

// OpenDrawingFile opens a new .svg drawing file -- just the basic opening
func (vv *VectorView) OpenDrawingFile(fnm gi.Filename) error {
	path, _ := filepath.Abs(string(fnm))
	vv.Filename = gi.Filename(path)
	sv := vv.SVG()
	err := sv.OpenXML(gi.Filename(path))
	if err != nil && err != io.EOF {
		log.Println(err)
		// return err
	}
	SavedPaths.AddPath(path, gi.Prefs.Params.SavedPathsMax)
	SavePaths()
	fdir, _ := filepath.Split(path)
	os.Chdir(fdir)
	vv.EditState.Init()
	vv.UpdateLayerView()

	vv.EditState.Gradients = sv.Gradients()
	sv.GatherIds() // also ensures uniqueness, key for json saving
	sv.ZoomToContents(false)
	sv.ReadMetaData()
	sv.SetTransform()
	return err
}

// OpenDrawing opens a new .svg drawing
func (vv *VectorView) OpenDrawing(fnm gi.Filename) error {
	wupdt := vv.TopUpdateStart()
	defer vv.TopUpdateEnd(wupdt)
	updt := vv.UpdateStart()
	vv.SetFullReRender()

	vv.Defaults()
	err := vv.OpenDrawingFile(fnm)

	sv := vv.SVG()
	vv.SetTitle()
	tv := vv.TreeView()
	tv.CloseAll()
	tv.ReSync()
	vv.SetStatus("Opened: " + string(vv.Filename))
	vv.UpdateEnd(updt)
	tv.CloseAll()
	sv.bgVectorEff = 0
	sv.UpdateView(true)
	return err
}

// NewDrawing opens a new drawing window
func (vv *VectorView) NewDrawing(sz PhysSize) *VectorView {
	ngr := NewDrawing(sz)
	return ngr
}

// PromptPhysSize prompts for physical size of drawing and sets it
func (vv *VectorView) PromptPhysSize() {
	sv := vv.SVG()
	sz := &PhysSize{}
	sz.SetFromSVG(sv)
	giv.StructViewDialog(vv.Viewport, sz, giv.DlgOpts{Title: "SVG Physical Size", Ok: true, Cancel: true}, vv.This(),
		func(recv, send ki.Ki, sig int64, d any) {
			if sig == int64(gi.DialogAccepted) {
				vv.SetPhysSize(sz)
				sv.bgVectorEff = -1
				sv.UpdateView(true)
			}
		})
}

// SetPhysSize sets physical size of drawing
func (vv *VectorView) SetPhysSize(sz *PhysSize) {
	if sz == nil {
		return
	}
	if sz.Size.IsNil() {
		sz.SetStdSize(Prefs.Size.StdSize)
	}
	sv := vv.SVG()
	sz.SetToSVG(sv)
	sv.SetMetaData()
	sv.ZoomToPage(false)
}

// SaveDrawing saves .svg drawing to current filename
func (vv *VectorView) SaveDrawing() error {
	if vv.Filename == "" {
		giv.CallMethod(vv, "SaveDrawingAs", vv.ViewportSafe())
		return nil
	}
	sv := vv.SVG()
	sv.RemoveOrphanedDefs()
	sv.SetMetaData()
	err := sv.SaveXML(vv.Filename)
	if err != nil && err != io.EOF {
		log.Println(err)
	} else {
		vv.AutoSaveDelete()
	}
	vv.SetStatus("Saved: " + string(vv.Filename))
	vv.EditState.Changed = false
	return err
}

// SaveDrawingAs saves .svg drawing to given filename
func (vv *VectorView) SaveDrawingAs(fname gi.Filename) error {
	if fname == "" {
		return errors.New("SaveDrawingAs: filename is empty")
	}
	path, _ := filepath.Abs(string(fname))
	vv.Filename = gi.Filename(path)
	SavedPaths.AddPath(path, gi.Prefs.Params.SavedPathsMax)
	SavePaths()
	sv := vv.SVG()
	sv.RemoveOrphanedDefs()
	sv.SetMetaData()
	err := sv.SaveXML(gi.Filename(path))
	if err != nil && err != io.EOF {
		log.Println(err)
	} else {
		vv.AutoSaveDelete()
	}
	vv.SetTitle()
	vv.SetStatus("Saved: " + path)
	vv.EditState.Changed = false
	return err
}

// ExportPNG exports drawing to a PNG image (auto-names to same name
// with .png suffix).  Calls inkscape -- needs to be on the PATH.
// specify either width or height of resulting image, or nothing for
// physical size as set.  Renders full current page -- do ResizeToContents
// to render just current contents.
func (vv *VectorView) ExportPNG(width, height float32) error {
	path, _ := filepath.Split(string(vv.Filename))
	fnm := filepath.Join(path, "export_png.svg")
	sv := vv.SVG()
	err := sv.SaveXML(gi.Filename(fnm))
	if err != nil && err != io.EOF {
		log.Println(err)
		return err
	}
	fext := filepath.Ext(string(vv.Filename))
	onm := strings.TrimSuffix(string(vv.Filename), fext) + ".png"
	cstr := "inkscape"
	args := []string{`--export-type=png`, "-o", onm}
	if width > 0 {
		args = append(args, fmt.Sprintf("--export-width=%g", width))
	}
	if height > 0 {
		args = append(args, fmt.Sprintf("--export-height=%g", height))
	}
	args = append(args, fnm)
	cmd := exec.Command(cstr, args...)
	fmt.Printf("executing command: %s %v\n", cstr, args)
	out, err := cmd.CombinedOutput()
	// if err != nil {
	fmt.Println(string(out))
	// }
	os.Remove(fnm)
	return err
}

// ExportPDF exports drawing to a PDF file (auto-names to same name
// with .pdf suffix).  Calls inkscape -- needs to be on the PATH.
// specify DPI of resulting image for effects rendering.
// Renders full current page -- do ResizeToContents
// to render just current contents.
func (vv *VectorView) ExportPDF(dpi float32) error {
	path, _ := filepath.Split(string(vv.Filename))
	fnm := filepath.Join(path, "export_pdf.svg")
	sv := vv.SVG()
	err := sv.SaveXML(gi.Filename(fnm))
	if err != nil && err != io.EOF {
		log.Println(err)
		return err
	}
	fext := filepath.Ext(string(vv.Filename))
	onm := strings.TrimSuffix(string(vv.Filename), fext) + ".pdf"
	cstr := "inkscape"
	args := []string{`--export-type=pdf`, "-o", onm}
	if dpi > 0 {
		args = append(args, fmt.Sprintf("--export-dpi=%g", dpi))
	}
	args = append(args, fnm)
	cmd := exec.Command(cstr, args...)
	fmt.Printf("executing command: %s %v\n", cstr, args)
	out, err := cmd.CombinedOutput()
	// if err != nil {
	fmt.Println(string(out))
	// }
	os.Remove(fnm)
	return err
}

// ResizeToContents resizes the drawing to just fit the current contents,
// including moving everything to start at upper-left corner,
// preserving the current grid offset, so grid snapping
// is preserved.
func (vv *VectorView) ResizeToContents() {
	sv := vv.SVG()
	sv.ResizeToContents(true)
	sv.UpdateView(true)
}

// AddImage adds a new image node set to given image
func (vv *VectorView) AddImage(fname gi.Filename, width, height float32) error {
	sv := vv.SVG()
	sv.UndoSave("AddImage", string(fname))
	ind := sv.NewEl(svg.KiT_Image).(*svg.Image)
	ind.Pos.X = 100 // todo: default pos
	ind.Pos.Y = 100 // todo: default pos
	err := ind.OpenImage(fname, width, height)
	sv.UpdateView(true)
	vv.ChangeMade()
	return err
}

//////////////////////////////////////////////////////////////////////////
//  GUI Config

func (vv *VectorView) MainToolbar() *gi.Toolbar {
	return vv.ChildByName("main-tb", 0).(*gi.Toolbar)
}

func (vv *VectorView) ModalToolbarStack() *gi.Layout {
	return vv.ChildByName("modal-tb", 1).(*gi.Layout)
}

// SetModalSelect sets the modal toolbar to be the select one
func (vv *VectorView) SetModalSelect() {
	tbs := vv.ModalToolbarStack()
	updt := tbs.UpdateStart()
	tbs.SetFullReRender()
	vv.UpdateSelectToolbar()
	idx, _ := tbs.Kids.IndexByName("select-tb", 0)
	tbs.StackTop = idx
	tbs.UpdateEnd(updt)
}

// SetModalNode sets the modal toolbar to be the node editing one
func (vv *VectorView) SetModalNode() {
	tbs := vv.ModalToolbarStack()
	updt := tbs.UpdateStart()
	tbs.SetFullReRender()
	vv.UpdateNodeToolbar()
	idx, _ := tbs.Kids.IndexByName("node-tb", 1)
	tbs.StackTop = idx
	tbs.UpdateEnd(updt)
}

// SetModalText sets the modal toolbar to be the text editing one
func (vv *VectorView) SetModalText() {
	tbs := vv.ModalToolbarStack()
	updt := tbs.UpdateStart()
	tbs.SetFullReRender()
	vv.UpdateTextToolbar()
	idx, _ := tbs.Kids.IndexByName("text-tb", 2)
	tbs.StackTop = idx
	tbs.UpdateEnd(updt)
}

func (vv *VectorView) HBox() *gi.Layout {
	return vv.ChildByName("hbox", 2).(*gi.Layout)
}

func (vv *VectorView) Tools() *gi.Toolbar {
	return vv.HBox().ChildByName("tools", 0).(*gi.Toolbar)
}

func (vv *VectorView) Splits() *gi.Splits {
	return vv.HBox().ChildByName("splits", 1).(*gi.Splits)
}

func (vv *VectorView) LayerTree() *gi.Layout {
	return vv.Splits().ChildByName("layer-tree", 0).(*gi.Layout)
}

func (vv *VectorView) LayerView() *giv.TableView {
	return vv.LayerTree().ChildByName("layers", 0).(*giv.TableView)
}

func (vv *VectorView) TreeView() *TreeView {
	return vv.LayerTree().ChildByName("tree-frame", 1).Child(0).(*TreeView)
}

func (vv *VectorView) SVG() *SVGView {
	return vv.Splits().Child(1).(*SVGView)
}

func (vv *VectorView) Tabs() *gi.Tabs {
	return vv.Splits().ChildByName("tabs", 2).(*gi.Tabs)
}

// StatusBar returns the statusbar widget
func (vv *VectorView) StatusBar() *gi.Frame {
	return vv.ChildByName("statusbar", 4).(*gi.Frame)
}

// StatusLabel returns the statusbar label widget
func (vv *VectorView) StatusLabel() *gi.Label {
	return vv.StatusBar().Child(0).Embed(gi.KiT_Label).(*gi.Label)
}

// Config configures entire view -- only runs if no children yet
func (vv *VectorView) Config() {
	if vv.HasChildren() {
		return
	}
	updt := vv.UpdateStart()
	vv.Lay = gi.LayoutVert
	// gv.SetProp("spacing", gi.StdDialogVSpaceUnits)
	gi.NewToolbar(vv, "main-tb")
	gi.NewLayout(vv, "modal-tb", gi.LayoutStacked)
	hb := gi.NewLayout(vv, "hbox", gi.LayoutHoriz)
	hb.SetStretchMax()
	gi.NewFrame(vv, "statusbar", gi.LayoutHoriz)

	tb := gi.NewToolbar(hb, "tools")
	tb.Lay = gi.LayoutVert
	spv := gi.NewSplits(hb, "splitview")
	spv.Dim = mat32.X

	tly := gi.NewLayout(spv, "layer-tree", gi.LayoutVert)
	tly.SetStretchMax()

	nly := gi.NewButton(tly, "add-layer")
	nly.SetText("Add Layer")
	nly.OnClicked(func() {
		vv.AddLayer()
	})

	lyv := giv.AddNewTableView(tly, "layers")
	lyv.SetMinPrefHeight(units.NewEm(6))
	lyv.SetStretchMax()

	tvfr := gi.NewFrame(tly, "tree-frame", gi.LayoutVert)
	tvfr.SetMinPrefHeight(units.NewEm(12))
	tvfr.SetStretchMax()
	tvfr.SetReRenderAnchor()
	tv := AddNewTreeView(tvfr, "treeview")
	tv.VectorView = vv
	tv.OpenDepth = 4

	sv := AddNewSVGView(spv, "svg", vv)

	tab := gi.NewTabView(spv, "tabs")
	tab.SetStretchMaxWidth()

	tv.SetRootNode(sv)

	tv.TreeViewSig.Connect(vv.This(), func(recv, send ki.Ki, sig int64, data any) {
		gvv := recv.Embed(KiT_VectorView).(*VectorView)
		if data == nil {
			return
		}
		if sig == int64(giv.TreeViewInserted) {
			sn, ok := data.(svg.Node)
			if ok {
				gvv.SVG().NodeEnsureUniqueId(sn)
				svg.CloneNodeGradientProp(sn, "fill")
				svg.CloneNodeGradientProp(sn, "stroke")
			}
			return
		}
		if sig == int64(giv.TreeViewDeleted) {
			sn, ok := data.(svg.Node)
			if ok {
				svg.DeleteNodeGradientProp(sn, "fill")
				svg.DeleteNodeGradientProp(sn, "stroke")
			}
			return
		}
		if sig != int64(giv.TreeViewOpened) {
			return
		}
		tvn, _ := data.(ki.Ki).Embed(KiT_TreeView).(*TreeView)
		_, issvg := tvn.SrcNode.(svg.Node)
		if !issvg {
			return
		}
		if tvn.SrcNode.HasChildren() {
			return
		}
		giv.StructViewDialog(gvv.Viewport, tvn.SrcNode, giv.DlgOpts{Title: "SVG Element View"}, nil, nil)
		// ggv, _ := recv.Embed(KiT_VectorView).(*VectorView)
		// 		stv := ggv.RecycleTab("Obj", giv.KiT_StructView, true).(*giv.StructView)
		// 		stv.SetStruct(tvn.SrcNode)
	})

	spv.SetSplits(0.15, 0.60, 0.25)

	vv.ConfigStatusBar()
	vv.ConfigMainToolbar()
	vv.ConfigModalToolbar()
	vv.ConfigTools()
	vv.ConfigTabs()

	vv.SetPhysSize(&Prefs.Size)

	vv.SyncLayers()
	lyv.SetSlice(&vv.EditState.Layers)
	vv.LayerViewSigs(lyv)

	sv.UpdateGradients(vv.EditState.Gradients)

	vv.UpdateEnd(updt)
}

// IsConfiged returns true if the view is fully configured
func (vv *VectorView) IsConfiged() bool {
	if !vv.HasChildren() {
		return false
	}
	return true
}

// UndoAvailFunc is an ActionUpdateFunc that inactivates action if no more undos
func (vv *VectorView) UndoAvailFunc(act *gi.Button) {
	es := &vv.EditState
	act.SetInactiveState(!es.UndoMgr.HasUndoAvail())
}

// RedoAvailFunc is an ActionUpdateFunc that inactivates action if no more redos
func (vv *VectorView) RedoAvailFunc(act *gi.Button) {
	es := &vv.EditState
	act.SetInactiveState(!es.UndoMgr.HasRedoAvail())
}

// PasteAvailFunc is an ActionUpdateFunc that inactivates action if no paste avail
func (vv *VectorView) PasteAvailFunc(act *gi.Button) {
	empty := goosi.TheApp.Clipboard(vv.ParentWindow().OSWin).IsEmpty()
	act.SetInactiveState(empty)
}

func (vv *VectorView) ConfigMainToolbar() {
	tb := vv.MainToolbar()
	tb.SetStretchMaxWidth()
	tb.AddAction(gi.ActOpts{Label: "Updt", Icon: "update", Tooltip: "update display -- should not be needed but sometimes, while still under development..."},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			grr.UpdateAll()
		})
	tb.AddAction(gi.ActOpts{Label: "New", Icon: "new", Tooltip: "create new drawing of specified size"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			ndr := grr.NewDrawing(Prefs.Size)
			ndr.PromptPhysSize()
		})
	szmen := tb.AddAction(gi.ActOpts{Label: "Size", Icon: "gear"}, nil, nil)
	szmen.Menu.AddAction(gi.ActOpts{Label: "Set Size...", Icon: "gear", Tooltip: "set size and grid spacing of drawing"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			grr.PromptPhysSize()
		})
	szmen.Menu.AddAction(gi.ActOpts{Label: "Resize To Contents", Icon: "gear", Tooltip: "resizes the drawing to fit the current contents, moving everything to upper-left corner while preserving grid alignment"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			grr.ResizeToContents()
		})
	tb.AddAction(gi.ActOpts{Label: "Open...", Icon: "file-open", Tooltip: "Open a drawing from .svg file"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			giv.CallMethod(grr, "OpenDrawing", grr.ViewportSafe())
		})
	tb.AddAction(gi.ActOpts{Label: "Save", Icon: "file-save", Tooltip: "Save drawing to .svg file, using current filename (if empty, prompts)"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			grr.SaveDrawing()
		})
	tb.AddAction(gi.ActOpts{Label: "Save As...", Icon: "file-save", Tooltip: "Save drawing to a new .svg file"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			giv.CallMethod(grr, "SaveDrawingAs", grr.ViewportSafe())
		})
	expmen := tb.AddAction(gi.ActOpts{Label: "Export", Icon: "file-save"}, nil, nil)
	expmen.Menu.AddAction(gi.ActOpts{Label: "Export PNG", Icon: "file-image", Tooltip: "Export drawing to a .png file -- requires cairosvg.org to be installed"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			giv.CallMethod(grr, "ExportPNG", grr.ViewportSafe())
		})
	expmen.Menu.AddAction(gi.ActOpts{Label: "Export PDF", Icon: "file-pdf", Tooltip: "Export drawing to a .pdf  file -- requires cairosvg.org to be installed"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			giv.CallMethod(grr, "ExportPDF", grr.ViewportSafe())
		})

	gi.NewSeparator(tb, "sep-undo")
	tb.AddAction(gi.ActOpts{Label: "Undo", Icon: "rotate-left", Tooltip: "Undo last action", UpdateFunc: vv.UndoAvailFunc},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			grr.Undo()
		})
	tb.AddAction(gi.ActOpts{Label: "Redo", Icon: "rotate-right", Tooltip: "Redo last undo action", UpdateFunc: vv.RedoAvailFunc},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			grr.Redo()
		})
	gi.NewSeparator(tb, "sep-edit")
	tb.AddAction(gi.ActOpts{Label: "Duplicate", Icon: "documents", Tooltip: "Duplicate current selection -- original items will remain selected", UpdateFunc: vv.SelectedEnableFunc},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			grr.DuplicateSelected()
		})
	tb.AddAction(gi.ActOpts{Label: "Copy", Icon: "copy", Tooltip: "Copy current selection to clipboard", UpdateFunc: vv.SelectedEnableFunc},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			grr.CopySelected()
		})
	tb.AddAction(gi.ActOpts{Label: "Cut", Icon: "cut", Tooltip: "Cut current selection -- delete and copy to clipboard", UpdateFunc: vv.SelectedEnableFunc},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			grr.CutSelected()
		})
	tb.AddAction(gi.ActOpts{Label: "Paste", Icon: "paste", Tooltip: "Paste clipboard contents", UpdateFunc: vv.PasteAvailFunc},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			grr.PasteClip()
		})
	gi.NewSeparator(tb, "sep-import")
	tb.AddAction(gi.ActOpts{Label: "Add Image...", Icon: "file-image", Tooltip: "add an image from a file"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			giv.CallMethod(grr, "AddImage", grr.ViewportSafe())
		})
	gi.NewSeparator(tb, "sep-view")
	tb.AddAction(gi.ActOpts{Label: "Zoom Page", Icon: "zoom-out", Tooltip: "zoom to see entire page size for drawing"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			svvv := grr.SVG()
			svvv.ZoomToPage(false)
			svvv.UpdateView(true)
		})
	tb.AddAction(gi.ActOpts{Label: "Zoom All", Icon: "zoom-out", Tooltip: "zoom to see entire contents"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			grr := recv.Embed(KiT_VectorView).(*VectorView)
			svvv := grr.SVG()
			svvv.ZoomToContents(false)
			svvv.UpdateView(true)
		})
}

func (vv *VectorView) ConfigModalToolbar() {
	tb := vv.ModalToolbarStack()
	if tb == nil || tb.HasChildren() {
		return
	}
	tb.SetStretchMaxWidth()
	gi.NewToolbar(tb, "select-tb")
	gi.NewToolbar(tb, "node-tb")
	gi.NewToolbar(tb, "text-tb")

	vv.ConfigSelectToolbar()
	vv.ConfigNodeToolbar()
	vv.ConfigTextToolbar()
}

// ConfigStatusBar configures statusbar with label
func (vv *VectorView) ConfigStatusBar() {
	sb := vv.StatusBar()
	if sb == nil || sb.HasChildren() {
		return
	}
	sb.SetStretchMaxWidth()
	sb.SetMinPrefHeight(units.NewValue(1.2, units.Em))
	sb.SetProp("overflow", "hidden") // no scrollbars!
	sb.SetProp("margin", 0)
	sb.SetProp("padding", 0)
	lbl := sb.AddNewChild(gi.KiT_Label, "sb-lbl").(*gi.Label)
	lbl.SetStretchMaxWidth()
	lbl.SetMinPrefHeight(units.NewValue(1, units.Em))
	lbl.SetProp("vertical-align", styles.AlignTop)
	lbl.SetProp("margin", 0)
	lbl.SetProp("padding", 0)
	lbl.SetProp("tab-size", 4)
}

// SetStatus updates the statusbar label with given message, along with other status info
func (vv *VectorView) SetStatus(msg string) {
	sb := vv.StatusBar()
	if sb == nil {
		return
	}
	updt := sb.UpdateStart()
	lbl := vv.StatusLabel()
	es := &vv.EditState
	str := "<b>" + strings.TrimSuffix(es.Tool.String(), "Tool") + "</b>\t"
	if es.CurLayer != "" {
		str += "Layer: " + es.CurLayer + "\t\t"
	}
	str += msg
	lbl.SetText(str)
	sb.UpdateEnd(updt)
}

// CloseWindowReq is called when user tries to close window -- we
// automatically save the project if it already exists (no harm), and prompt
// to save open files -- if this returns true, then it is OK to close --
// otherwise not
func (vv *VectorView) CloseWindowReq() bool {
	if !vv.EditState.Changed {
		return true
	}
	gi.ChoiceDialog(vv.Viewport, gi.DlgOpts{Title: "Close Drawing: There are Unsaved Changes",
		Prompt: fmt.Sprintf("In Drawing: %v There are <b>unsaved changes</b> -- do you want to save or cancel closing this drawing?", giv.DirAndFile(string(vv.Filename)))},
		[]string{"Cancel", "Save", "Close Without Saving"},
		vv.This(), func(recv, send ki.Ki, sig int64, data any) {
			switch sig {
			case 0:
				// do nothing, will have returned false already
			case 1:
				vv.SaveDrawing()
			case 2:
				vv.ParentWindow().OSWin.Close() // will not be prompted again!
			}
		})
	return false // not yet
}

// QuitReq is called when user tries to quit the app -- we go through all open
// main windows and look for grid windows and call their CloseWindowReq
// functions!
func QuitReq() bool {
	for _, win := range gi.MainWindows {
		if !strings.HasPrefix(win.Nm, "grid-") {
			continue
		}
		mfr, err := win.MainWidget()
		if err != nil {
			continue
		}
		gek := mfr.ChildByName("grid", 0)
		if gek == nil {
			continue
		}
		vv := gek.Embed(KiT_VectorView).(*VectorView)
		if !vv.CloseWindowReq() {
			return false
		}
	}
	return true
}

func (vv *VectorView) SetTitle() {
	if vv.Filename == "" {
		return
	}
	dfnm := giv.DirAndFile(string(vv.Filename))
	winm := "grid-" + dfnm
	wintitle := "grid: " + dfnm
	win := vv.ParentWindow()
	win.SetName(winm)
	win.SetTitle(wintitle)
}

// NewDrawing opens a new drawing window
func NewDrawing(sz PhysSize) *VectorView {
	_, ngr := NewVectorWindow("")
	ngr.SetPhysSize(&sz)
	return ngr
}

// NewVectorWindow returns a new VectorWindow loading given file if non-empty
func NewVectorWindow(fnm string) *VectorView {
	path := ""
	dfnm := ""
	if fnm != "" {
		path, _ = filepath.Abs(fnm)
		dfnm = giv.DirAndFile(path)
	}
	winm := "grid-" + dfnm
	wintitle := "grid: " + dfnm

	if win, found := gi.AllWindows.FindName(winm); found {
		mfr := win.SetMainFrame()
		vv := mfr.Child(0).Embed(KiT_VectorView).(*VectorView)
		if string(vv.Filename) == path {
			win.OSWin.Raise()
			return win, vv
		}
	}

	width := 1600
	height := 1280
	sc := goosi.TheApp.Screen(0)
	if sc != nil {
		scsz := sc.Geometry.Size()
		width = int(.9 * float64(scsz.X))
		height = int(.8 * float64(scsz.Y))
	}

	win := gi.NewMainWindow(winm, wintitle, width, height)

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	mfr := win.SetMainFrame()
	vv := AddNewVectorView(mfr, "vectorview")
	vv.Viewport = vp
	vv.Defaults()
	vv.Config()

	mmen := win.MainMenu
	giv.MainMenuView(vv, win, mmen)

	win.MainMenuUpdated()

	vp.UpdateEndNoSig(updt)

	win.GoStartEventLoop()

	if fnm != "" {
		vv.OpenDrawingFile(gi.Filename(path))
	}

	return win, vv
}

/////////////////////////////////////////////////////////////////////////
//   Controls

// RecycleTab returns a tab with given name, first by looking for an existing one,
// and if not found, making a new one with widget of given type.
// If sel, then select it.  returns widget for tab.
// func (gv *VectorView) RecycleTab(label string, typ reflect.Type, sel bool) gi.Node2D {
// 	tv := gv.Tabs()
// 	return tv.RecycleTab(label, typ, sel)
// }

// Tab returns tab with given label
// func (gv *VectorView) Tab(label string) gi.Node2D {
// 	tv := gv.Tabs()
// 	return tv.TabByName(label)
// }

func (vv *VectorView) ConfigTabs() {
	tv := vv.Tabs()
	tv.NoDeleteTabs = true
	pv := vv.RecycleTab("Paint", KiT_PaintView, false).(*PaintView)
	pv.Config(vv)
	av := vv.RecycleTab("Align", KiT_AlignView, false).(*AlignView)
	av.Config(vv)
	vv.EditState.Text.Defaults()
	txv := vv.RecycleTab("Text", giv.KiT_StructView, false).(*giv.StructView)
	txv.SetStruct(&vv.EditState.Text)
}

func (vv *VectorView) PaintView() *PaintView {
	return vv.Tab("Paint").(*PaintView)
}

func (vv *VectorView) UpdateAll() {
	vv.UpdateTabs()
	vv.UpdateTreeView()
	vv.UpdateDisp()
}

func (vv *VectorView) UpdateDisp() {
	sv := vv.SVG()
	sv.UpdateView(true)
}

func (vv *VectorView) UpdateTreeView() {
	tv := vv.TreeView()
	tv.ReSync()
}

func (vv *VectorView) SetDefaultStyle() {
	pv := vv.Tab("Paint").(*PaintView)
	es := &vv.EditState
	switch es.Tool {
	case TextTool:
		pv.Update(&Prefs.TextStyle, nil)
	case BezierTool:
		pv.Update(&Prefs.PathStyle, nil)
	default:
		pv.Update(&Prefs.ShapeStyle, nil)
	}
}

func (vv *VectorView) UpdateTabs() {
	// fmt.Printf("updt-tabs\n")
	es := &vv.EditState
	fsel := es.FirstSelectedNode()
	if fsel != nil {
		sel := fsel.AsNodeBase()
		pv := vv.Tab("Paint").(*PaintView)
		pv.Update(&sel.Pnt, sel.This())
		txt, istxt := fsel.(*svg.Text)
		if istxt {
			es.Text.SetFromNode(txt)
			txv := vv.Tab("Text").(*giv.StructView)
			txv.UpdateFields()
			// todo: only show text toolbar on double-click
			// gv.SetModalText()
			// gv.UpdateTextToolbar()
		} else {
			vv.SetModalToolbar()
		}
	}
}

// SelectNodeInSVG selects given svg node in SVG drawing
func (vv *VectorView) SelectNodeInSVG(kn ki.Ki, mode events.SelectModes) {
	sii, ok := kn.(svg.Node)
	if !ok {
		return
	}
	sv := vv.SVG()
	es := &vv.EditState
	es.SelectAction(sii, mode, image.ZP)
	sv.UpdateView(false)
}

// Undo undoes one step, returning name of action that was undone
func (vv *VectorView) Undo() string {
	sv := vv.SVG()
	act := sv.Undo()
	if act != "" {
		vv.SetStatus("Undid: " + act)
	} else {
		vv.SetStatus("Undo: no more to undo")
	}
	vv.UpdateAll()
	return act
}

// Redo redoes one step, returning name of action that was redone
func (vv *VectorView) Redo() string {
	sv := vv.SVG()
	act := sv.Redo()
	if act != "" {
		vv.SetStatus("Redid: " + act)
	} else {
		vv.SetStatus("Redo: no more to redo")
	}
	vv.UpdateAll()
	return act
}

// ChangeMade should be called after any change is completed on the drawing.
// Calls autosave.
func (vv *VectorView) ChangeMade() {
	go vv.AutoSave()
}

/////////////////////////////////////////////////////////////////////////
//   Basic infrastructure

func (vv *VectorView) ConnectEvents2D() {
	vv.OSFileEvent()
}

/*
func (gv *VectorView) OSFileEvent() {
	gv.ConnectEvent(oswin.OSOpenFilesEvent, gi.RegPri, func(recv, send ki.Ki, sig int64, d any) {
		ofe := d.(*osevent.OpenFilesEvent)
		for _, fn := range ofe.Files {
			NewVectorWindow(fn)
		}
	})
}
*/

// OpenRecent opens a recently-used file
func (vv *VectorView) OpenRecent(filename gi.Filename) {
	if string(filename) == VectorViewResetRecents {
		SavedPaths = nil
		gi.StringsAddExtras((*[]string)(&SavedPaths), SavedPathsExtras)
	} else if string(filename) == VectorViewEditRecents {
		vv.EditRecents()
	} else {
		vv.OpenDrawing(filename)
	}
}

// RecentsEdit opens a dialog editor for deleting from the recents project list
func (vv *VectorView) EditRecents() {
	tmp := make([]string, len(SavedPaths))
	copy(tmp, SavedPaths)
	gi.StringsRemoveExtras((*[]string)(&tmp), SavedPathsExtras)
	opts := giv.DlgOpts{Title: "Recent Project Paths", Prompt: "Delete paths you no longer use", Ok: true, Cancel: true, NoAdd: true}
	giv.SliceViewDialog(vv.Viewport, &tmp, opts,
		nil, vv, func(recv, send ki.Ki, sig int64, data any) {
			if sig == int64(gi.DialogAccepted) {
				SavedPaths = nil
				SavedPaths = append(SavedPaths, tmp...)
				gi.StringsAddExtras((*[]string)(&SavedPaths), SavedPathsExtras)
			}
		})
}

// SplitsSetView sets split view splitters to given named setting
func (vv *VectorView) SplitsSetView(split SplitName) {
	sv := vv.Splits()
	sp, _, ok := AvailSplits.SplitByName(split)
	if ok {
		sv.SetSplitsAction(sp.Splits...)
		Prefs.SplitName = split
	}
}

// SplitsSave saves current splitter settings to named splitter settings under
// existing name, and saves to prefs file
func (vv *VectorView) SplitsSave(split SplitName) {
	sv := vv.Splits()
	sp, _, ok := AvailSplits.SplitByName(split)
	if ok {
		sp.SaveSplits(sv.Splits)
		AvailSplits.SavePrefs()
	}
}

// SplitsSaveAs saves current splitter settings to new named splitter settings, and
// saves to prefs file
func (vv *VectorView) SplitsSaveAs(name, desc string) {
	spv := vv.Splits()
	AvailSplits.Add(name, desc, spv.Splits)
	AvailSplits.SavePrefs()
}

// SplitsEdit opens the SplitsView editor to customize saved splitter settings
func (vv *VectorView) SplitsEdit() {
	SplitsView(&AvailSplits)
}

// HelpWiki opens wiki page for grid on github
func (vv *VectorView) HelpWiki() {
	goosi.TheApp.OpenURL("https://goki.dev/grid/wiki")
}

////////////////////////////////////////////////////////////////////////////////////////
//		AutoSave

// AutoSaveFilename returns the autosave filename
func (vv *VectorView) AutoSaveFilename() string {
	path, fn := filepath.Split(string(vv.Filename))
	if fn == "" {
		fn = "new_file_" + vv.Nm + ".svg"
	}
	asfn := filepath.Join(path, "#"+fn+"#")
	return asfn
}

// AutoSave does the autosave -- safe to call in a separate goroutine
func (vv *VectorView) AutoSave() error {
	if vv.HasFlag(int(VectorViewAutoSaving)) {
		return nil
	}
	vv.SetFlag(int(VectorViewAutoSaving))
	asfn := vv.AutoSaveFilename()
	sv := vv.SVG()
	err := sv.SaveXML(gi.Filename(asfn))
	if err != nil && err != io.EOF {
		log.Println(err)
	}
	vv.ClearFlag(int(VectorViewAutoSaving))
	return err
}

// AutoSaveDelete deletes any existing autosave file
func (vv *VectorView) AutoSaveDelete() {
	asfn := vv.AutoSaveFilename()
	os.Remove(asfn)
}

// AutoSaveCheck checks if an autosave file exists -- logic for dealing with
// it is left to larger app -- call this before opening a file
func (vv *VectorView) AutoSaveCheck() bool {
	asfn := vv.AutoSaveFilename()
	if _, err := os.Stat(asfn); os.IsNotExist(err) {
		return false // does not exist
	}
	return true
}

// VectorViewFlags extend WidgetFlags to hold VectorView state
type VectorViewFlags gi.WidgetFlags //enums:bitflag -trim-prefix VectorViewFlag

const (
	// VectorViewAutoSaving means
	VectorViewAutoSaving VectorViewFlags = VectorViewFlags(gi.WidgetFlagsN) + iota
)

/////////////////////////////////////////////////////////////////////////
//   Props, MainMenu

/*
var VectorViewProps = ki.Props{
	"MainMenu": ki.PropSlice{
		{"AppMenu", ki.BlankProp{}},
		{"File", ki.PropSlice{
			{"OpenRecent", ki.Props{
				"submenu": &SavedPaths,
				"Args": ki.PropSlice{
					{"File Name", ki.Props{}},
				},
			}},
			{"OpenDrawing", ki.Props{
				"shortcut": keyfun.MenuOpen,
				"label":    "Open SVG...",
				"desc":     "open an SVG drawing",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".svg",
					}},
				},
			}},
			{"NewDrawing", ki.Props{
				"shortcut": keyfun.MenuNew,
				"label":    "New",
				"desc":     "Create a new drawing of given physical size (size units are used for ViewBox).",
				"Args": ki.PropSlice{
					{"Physical Size", ki.Props{
						"default": Prefs.Size,
					}},
				},
			}},
			{"SaveDrawing", ki.Props{
				"shortcut": keyfun.MenuSave,
				"label":    "Save Drawing",
			}},
			{"SaveDrawingAs", ki.Props{
				"shortcut": keyfun.MenuSaveAs,
				"label":    "Save As...",
				"desc":     "Save drawing to given svg file name",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".svg",
					}},
				},
			}},
			{"sep-file", ki.BlankProp{}},
			{"PromptPhysSize", ki.Props{
				"label": "Set Size",
				"desc":  "sets the physical size (size units are used for ViewBox)",
			}},
			{"ResizeToContents", ki.Props{
				"label": "Resize To Contents",
				"desc":  "resizes the drawing to fit the current contents, moving everything to upper-left corner while preserving grid alignment",
			}},
			{"sep-exp", ki.BlankProp{}},
			{"ExportPNG", ki.Props{
				"desc": "Export drawing as a PNG image file (uses cairosvg -- must install!) -- specify either width or height in pixels as non-zero, or both 0 to use physical size.  Renders full page -- do Resize To Contents to only render contents.",
				"Args": ki.PropSlice{
					{"Width", ki.Props{
						"default": 1280,
					}},
					{"Height", ki.Props{
						"default": 0,
					}},
				},
			}},
			{"ExportPDF", ki.Props{
				"desc": "Export drawing as a PDF file (uses cairosvg -- must install!), at given specified DPI (only relevant for rendered effects).  Renders full page -- do Resize To Contents to only render contents.",
				"Args": ki.PropSlice{
					{"DPI", ki.Props{
						"default": 300,
					}},
				},
			}},
			{"sep-imp", ki.BlankProp{}},
			{"AddImage", ki.Props{
				"label": "Add Image...",
				"desc":  "Add a new Image node with given image file for this image node, rescaling to given size -- use 0, 0 to use native image size.",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"default-field": "Filename",
						"ext":           ".png,.jpg,.jpeg",
					}},
					{"Width", ki.Props{}},
					{"Height", ki.Props{}},
				},
			}},
			{"sep-af", ki.BlankProp{}},
			{"Close Window", ki.BlankProp{}},
		}},
		{"Edit", ki.PropSlice{
			{"Duplicate", ki.Props{
				"keyfun": keyfun.Duplicate,
				// "updtfunc": VectorViewInactiveTextSelectionFunc,
			}},
			{"Copy", ki.Props{
				"keyfun": keyfun.Copy,
				// "updtfunc": VectorViewInactiveTextSelectionFunc,
			}},
			{"Cut", ki.Props{
				"keyfun": keyfun.Cut,
				// "updtfunc": VectorViewInactiveTextSelectionFunc,
			}},
			{"Paste", ki.Props{
				"keyfun": keyfun.Paste,
			}},
			{"sep-undo", ki.BlankProp{}},
			{"Undo", ki.Props{
				"keyfun": keyfun.Undo,
			}},
			{"Redo", ki.Props{
				"keyfun": keyfun.Redo,
			}},
		}},
		{"View", ki.PropSlice{
			{"Splits", ki.PropSlice{
				{"SplitsSetView", ki.Props{
					"label":   "Set View",
					"submenu": &AvailSplitNames,
					"Args": ki.PropSlice{
						{"Split Name", ki.Props{}},
					},
				}},
				{"SplitsSaveAs", ki.Props{
					"label": "Save As...",
					"desc":  "save current splitter values to a new named split configuration",
					"Args": ki.PropSlice{
						{"Name", ki.Props{
							"width": 60,
						}},
						{"Desc", ki.Props{
							"width": 60,
						}},
					},
				}},
				{"SplitsSave", ki.Props{
					"label":   "Save",
					"submenu": &AvailSplitNames,
					"Args": ki.PropSlice{
						{"Split Name", ki.Props{}},
					},
				}},
				{"SplitsEdit", ki.Props{
					"label": "Edit...",
				}},
			}},
		}},
		{"Window", "Windows"},
		{"Help", ki.PropSlice{
			{"HelpWiki", ki.Props{}},
		}},
	},
}

*/