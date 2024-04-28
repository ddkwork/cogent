// Code generated by "core generate"; DO NOT EDIT.

package vector

import (
	"cogentcore.org/core/icons"
	"cogentcore.org/core/math32"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/views"
)

// AlignViewType is the [types.Type] for [AlignView]
var AlignViewType = types.AddType(&types.Type{Name: "cogentcore.org/cogent/vector.AlignView", IDName: "align-view", Doc: "AlignView provides a range of alignment actions on selected objects.", Embeds: []types.Field{{Name: "Layout"}}, Fields: []types.Field{{Name: "AlignAnchorView"}, {Name: "VectorView", Doc: "the parent vectorview"}}, Instance: &AlignView{}})

// NewAlignView adds a new [AlignView] with the given name to the given parent:
// AlignView provides a range of alignment actions on selected objects.
func NewAlignView(parent tree.Node, name ...string) *AlignView {
	return parent.NewChild(AlignViewType, name...).(*AlignView)
}

// NodeType returns the [*types.Type] of [AlignView]
func (t *AlignView) NodeType() *types.Type { return AlignViewType }

// New returns a new [*AlignView] value
func (t *AlignView) New() tree.Node { return &AlignView{} }

// SetAlignAnchorView sets the [AlignView.AlignAnchorView]
func (t *AlignView) SetAlignAnchorView(v views.EnumValue) *AlignView { t.AlignAnchorView = v; return t }

// SetVectorView sets the [AlignView.VectorView]:
// the parent vectorview
func (t *AlignView) SetVectorView(v *VectorView) *AlignView { t.VectorView = v; return t }

// SetTooltip sets the [AlignView.Tooltip]
func (t *AlignView) SetTooltip(v string) *AlignView { t.Tooltip = v; return t }

// PaintViewType is the [types.Type] for [PaintView]
var PaintViewType = types.AddType(&types.Type{Name: "cogentcore.org/cogent/vector.PaintView", IDName: "paint-view", Doc: "PaintView provides editing of basic Stroke and Fill painting parameters\nfor selected items", Embeds: []types.Field{{Name: "Layout"}}, Fields: []types.Field{{Name: "StrokeType", Doc: "paint type for stroke"}, {Name: "StrokeStops", Doc: "name of gradient with stops"}, {Name: "FillType", Doc: "paint type for fill"}, {Name: "FillStops", Doc: "name of gradient with stops"}, {Name: "VectorView", Doc: "the parent vectorview"}}, Instance: &PaintView{}})

// NewPaintView adds a new [PaintView] with the given name to the given parent:
// PaintView provides editing of basic Stroke and Fill painting parameters
// for selected items
func NewPaintView(parent tree.Node, name ...string) *PaintView {
	return parent.NewChild(PaintViewType, name...).(*PaintView)
}

// NodeType returns the [*types.Type] of [PaintView]
func (t *PaintView) NodeType() *types.Type { return PaintViewType }

// New returns a new [*PaintView] value
func (t *PaintView) New() tree.Node { return &PaintView{} }

// SetStrokeType sets the [PaintView.StrokeType]:
// paint type for stroke
func (t *PaintView) SetStrokeType(v PaintTypes) *PaintView { t.StrokeType = v; return t }

// SetStrokeStops sets the [PaintView.StrokeStops]:
// name of gradient with stops
func (t *PaintView) SetStrokeStops(v string) *PaintView { t.StrokeStops = v; return t }

// SetFillType sets the [PaintView.FillType]:
// paint type for fill
func (t *PaintView) SetFillType(v PaintTypes) *PaintView { t.FillType = v; return t }

// SetFillStops sets the [PaintView.FillStops]:
// name of gradient with stops
func (t *PaintView) SetFillStops(v string) *PaintView { t.FillStops = v; return t }

// SetVectorView sets the [PaintView.VectorView]:
// the parent vectorview
func (t *PaintView) SetVectorView(v *VectorView) *PaintView { t.VectorView = v; return t }

// SetTooltip sets the [PaintView.Tooltip]
func (t *PaintView) SetTooltip(v string) *PaintView { t.Tooltip = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/cogent/vector.PhysSize", IDName: "phys-size", Doc: "PhysSize specifies the physical size of the drawing, when making a new one", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "StandardSize", Doc: "select a standard size -- this will set units and size"}, {Name: "Portrait", Doc: "for standard size, use first number as width, second as height"}, {Name: "Units", Doc: "default units to use, e.g., in line widths etc"}, {Name: "Size", Doc: "drawing size, in Units"}, {Name: "Grid", Doc: "grid spacing, in units of ViewBox size"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/cogent/vector.Preferences", IDName: "preferences", Doc: "Preferences is the overall Vector preferences", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Size", Doc: "default physical size, when app is started without opening a file"}, {Name: "Colors", Doc: "active color preferences"}, {Name: "ColorSchemes", Doc: "named color schemes -- has Light and Dark schemes by default"}, {Name: "ShapeStyle", Doc: "default shape styles"}, {Name: "TextStyle", Doc: "default text styles"}, {Name: "PathStyle", Doc: "default line styles"}, {Name: "LineStyle", Doc: "default line styles"}, {Name: "VectorDisp", Doc: "turns on the grid display"}, {Name: "SnapVector", Doc: "snap positions and sizes to underlying grid"}, {Name: "SnapGuide", Doc: "snap positions and sizes to line up with other elements"}, {Name: "SnapNodes", Doc: "snap node movements to align with guides"}, {Name: "SnapTol", Doc: "number of screen pixels around target point (in either direction) to snap"}, {Name: "SplitName", Doc: "named-split config in use for configuring the splitters"}, {Name: "EnvVars", Doc: "environment variables to set for this app -- if run from the command line, standard shell environment variables are inherited, but on some OS's (Mac), they are not set when run as a gui app"}, {Name: "Changed", Doc: "flag that is set by StructView by virtue of changeflag tag, whenever an edit is made.  Used to drive save menus etc."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/cogent/vector.ColorSettings", IDName: "color-settings", Doc: "ColorSettings for", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Background", Doc: "drawing background color"}, {Name: "Border", Doc: "border color of the drawing"}, {Name: "Vector", Doc: "grid line color"}}})

// SVGViewType is the [types.Type] for [SVGView]
var SVGViewType = types.AddType(&types.Type{Name: "cogentcore.org/cogent/vector.SVGView", IDName: "svg-view", Doc: "SVGView is the element for viewing, interacting with the SVG", Embeds: []types.Field{{Name: "SVG"}}, Fields: []types.Field{{Name: "VectorView", Doc: "the parent vectorview"}, {Name: "Trans", Doc: "view translation offset (from dragging)"}, {Name: "Scale", Doc: "view scaling (from zooming)"}, {Name: "Grid", Doc: "grid spacing, in native ViewBox units"}, {Name: "VectorEff", Doc: "effective grid spacing given Scale level"}, {Name: "BgPixels", Doc: "background pixels, includes page outline and grid"}, {Name: "bgTrans", Doc: "bg rendered translation"}, {Name: "bgScale", Doc: "bg rendered scale"}, {Name: "bgVectorEff", Doc: "bg rendered grid"}}, Instance: &SVGView{}})

// NewSVGView adds a new [SVGView] with the given name to the given parent:
// SVGView is the element for viewing, interacting with the SVG
func NewSVGView(parent tree.Node, name ...string) *SVGView {
	return parent.NewChild(SVGViewType, name...).(*SVGView)
}

// NodeType returns the [*types.Type] of [SVGView]
func (t *SVGView) NodeType() *types.Type { return SVGViewType }

// New returns a new [*SVGView] value
func (t *SVGView) New() tree.Node { return &SVGView{} }

// SetTooltip sets the [SVGView.Tooltip]
func (t *SVGView) SetTooltip(v string) *SVGView { t.Tooltip = v; return t }

// TreeViewType is the [types.Type] for [TreeView]
var TreeViewType = types.AddType(&types.Type{Name: "cogentcore.org/cogent/vector.TreeView", IDName: "tree-view", Doc: "TreeView is a TreeView that knows how to operate on FileNode nodes", Embeds: []types.Field{{Name: "TreeView"}}, Fields: []types.Field{{Name: "VectorView", Doc: "the parent vectorview"}}, Instance: &TreeView{}})

// NewTreeView adds a new [TreeView] with the given name to the given parent:
// TreeView is a TreeView that knows how to operate on FileNode nodes
func NewTreeView(parent tree.Node, name ...string) *TreeView {
	return parent.NewChild(TreeViewType, name...).(*TreeView)
}

// NodeType returns the [*types.Type] of [TreeView]
func (t *TreeView) NodeType() *types.Type { return TreeViewType }

// New returns a new [*TreeView] value
func (t *TreeView) New() tree.Node { return &TreeView{} }

// SetVectorView sets the [TreeView.VectorView]:
// the parent vectorview
func (t *TreeView) SetVectorView(v *VectorView) *TreeView { t.VectorView = v; return t }

// SetTooltip sets the [TreeView.Tooltip]
func (t *TreeView) SetTooltip(v string) *TreeView { t.Tooltip = v; return t }

// SetText sets the [TreeView.Text]
func (t *TreeView) SetText(v string) *TreeView { t.Text = v; return t }

// SetIcon sets the [TreeView.Icon]
func (t *TreeView) SetIcon(v icons.Icon) *TreeView { t.Icon = v; return t }

// SetIconOpen sets the [TreeView.IconOpen]
func (t *TreeView) SetIconOpen(v icons.Icon) *TreeView { t.IconOpen = v; return t }

// SetIconClosed sets the [TreeView.IconClosed]
func (t *TreeView) SetIconClosed(v icons.Icon) *TreeView { t.IconClosed = v; return t }

// SetIconLeaf sets the [TreeView.IconLeaf]
func (t *TreeView) SetIconLeaf(v icons.Icon) *TreeView { t.IconLeaf = v; return t }

// SetIndent sets the [TreeView.Indent]
func (t *TreeView) SetIndent(v units.Value) *TreeView { t.Indent = v; return t }

// SetOpenDepth sets the [TreeView.OpenDepth]
func (t *TreeView) SetOpenDepth(v int) *TreeView { t.OpenDepth = v; return t }

// SetViewIndex sets the [TreeView.ViewIndex]
func (t *TreeView) SetViewIndex(v int) *TreeView { t.ViewIndex = v; return t }

// SetWidgetSize sets the [TreeView.WidgetSize]
func (t *TreeView) SetWidgetSize(v math32.Vector2) *TreeView { t.WidgetSize = v; return t }

// SetRootView sets the [TreeView.RootView]
func (t *TreeView) SetRootView(v *views.TreeView) *TreeView { t.RootView = v; return t }

// SetSelectedNodes sets the [TreeView.SelectedNodes]
func (t *TreeView) SetSelectedNodes(v ...views.TreeViewer) *TreeView { t.SelectedNodes = v; return t }

// VectorViewType is the [types.Type] for [VectorView]
var VectorViewType = types.AddType(&types.Type{Name: "cogentcore.org/cogent/vector.VectorView", IDName: "vector-view", Doc: "VectorView is the Vector SVG vector drawing program", Methods: []types.Method{{Name: "AddLayer", Doc: "AddLayer adds a new layer", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SelectGroup", Doc: "SelectGroup groups items together", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SelectUnGroup", Doc: "SelectUnGroup ungroups items from each other", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SelectRotateLeft", Doc: "SelectRotateLeft rotates the selection 90 degrees counter-clockwise", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SelectRotateRight", Doc: "SelectRotateRight rotates the selection 90 degrees clockwise", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SelectFlipHorizontal", Doc: "SelectFlipHorizontal flips the selection horizontally", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SelectFlipVertical", Doc: "SelectFlipVertical flips the selection vertically", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SelectRaiseTop", Doc: "SelectRaiseTop raises the selection to the top of the layer", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SelectRaise", Doc: "SelectRaise raises the selection by one level in the layer", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SelectLowerBottom", Doc: "SelectLowerBottom lowers the selection to the bottom of the layer", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SelectLower", Doc: "SelectLower lowers the selection by one level in the layer", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "DuplicateSelected", Doc: "DuplicateSelected duplicates selected items in SVG view, using TreeView methods", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "CopySelected", Doc: "CopySelected copies selected items in SVG view, using TreeView methods", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "CutSelected", Doc: "CutSelected cuts selected items in SVG view, using TreeView methods", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "PasteClip", Doc: "PasteClip pastes clipboard, using cur layer etc", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "OpenDrawing", Doc: "OpenDrawing opens a new .svg drawing", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fnm"}, Returns: []string{"error"}}, {Name: "PromptPhysSize", Doc: "PromptPhysSize prompts for the physical size of the drawing and sets it", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SaveDrawing", Doc: "SaveDrawing saves .svg drawing to current filename", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Returns: []string{"error"}}, {Name: "SaveDrawingAs", Doc: "SaveDrawingAs saves .svg drawing to given filename", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname"}, Returns: []string{"error"}}, {Name: "ExportPNG", Doc: "ExportPNG exports drawing to a PNG image (auto-names to same name\nwith .png suffix).  Calls inkscape -- needs to be on the PATH.\nspecify either width or height of resulting image, or nothing for\nphysical size as set.  Renders full current page -- do ResizeToContents\nto render just current contents.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"width", "height"}, Returns: []string{"error"}}, {Name: "ExportPDF", Doc: "ExportPDF exports drawing to a PDF file (auto-names to same name\nwith .pdf suffix).  Calls inkscape -- needs to be on the PATH.\nspecify DPI of resulting image for effects rendering.\nRenders full current page -- do ResizeToContents\nto render just current contents.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"dpi"}, Returns: []string{"error"}}, {Name: "ResizeToContents", Doc: "ResizeToContents resizes the drawing to just fit the current contents,\nincluding moving everything to start at upper-left corner,\npreserving the current grid offset, so grid snapping\nis preserved.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "AddImage", Doc: "AddImage adds a new image node set to the given image", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname", "width", "height"}, Returns: []string{"error"}}, {Name: "UpdateAll", Doc: "UpdateAll updates the display", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "Undo", Doc: "Undo undoes the last action", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Returns: []string{"string"}}, {Name: "Redo", Doc: "Redo redoes the previously undone action", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Returns: []string{"string"}}}, Embeds: []types.Field{{Name: "Frame"}}, Fields: []types.Field{{Name: "Filename", Doc: "full path to current drawing filename"}, {Name: "EditState", Doc: "current edit state"}}, Instance: &VectorView{}})

// NewVectorView adds a new [VectorView] with the given name to the given parent:
// VectorView is the Vector SVG vector drawing program
func NewVectorView(parent tree.Node, name ...string) *VectorView {
	return parent.NewChild(VectorViewType, name...).(*VectorView)
}

// NodeType returns the [*types.Type] of [VectorView]
func (t *VectorView) NodeType() *types.Type { return VectorViewType }

// New returns a new [*VectorView] value
func (t *VectorView) New() tree.Node { return &VectorView{} }

// SetTooltip sets the [VectorView.Tooltip]
func (t *VectorView) SetTooltip(v string) *VectorView { t.Tooltip = v; return t }