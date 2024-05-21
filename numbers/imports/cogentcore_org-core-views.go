// Code generated by 'yaegi extract cogentcore.org/core/views'. DO NOT EDIT.

package imports

import (
	"cogentcore.org/core/base/fileinfo/mimedata"
	"cogentcore.org/core/core"
	"cogentcore.org/core/enums"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/abilities"
	"cogentcore.org/core/styles/states"
	"cogentcore.org/core/system"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
	"cogentcore.org/core/views"
	"go/constant"
	"go/token"
	"image"
	"reflect"
)

func init() {
	Symbols["cogentcore.org/core/views/views"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AddValue":                     reflect.ValueOf(views.AddValue),
		"ArgViewType":                  reflect.ValueOf(&views.ArgViewType).Elem(),
		"AsTreeView":                   reflect.ValueOf(views.AsTreeView),
		"CallFunc":                     reflect.ValueOf(views.CallFunc),
		"ColorButtonType":              reflect.ValueOf(&views.ColorButtonType).Elem(),
		"ColorViewType":                reflect.ValueOf(&views.ColorViewType).Elem(),
		"Config":                       reflect.ValueOf(views.Config),
		"ConfigBase":                   reflect.ValueOf(views.ConfigBase),
		"ConfigDialogValue":            reflect.ValueOf(views.ConfigDialogValue),
		"ConfigDialogWidget":           reflect.ValueOf(views.ConfigDialogWidget),
		"ConfigImageToolbar":           reflect.ValueOf(views.ConfigImageToolbar),
		"ConfigSVGToolbar":             reflect.ValueOf(views.ConfigSVGToolbar),
		"DateViewType":                 reflect.ValueOf(&views.DateViewType).Elem(),
		"FieldToValue":                 reflect.ValueOf(views.FieldToValue),
		"FileViewDialog":               reflect.ValueOf(views.FileViewDialog),
		"FileViewDirOnlyFilter":        reflect.ValueOf(views.FileViewDirOnlyFilter),
		"FileViewExtOnlyFilter":        reflect.ValueOf(views.FileViewExtOnlyFilter),
		"FileViewKindColorMap":         reflect.ValueOf(&views.FileViewKindColorMap).Elem(),
		"FileViewType":                 reflect.ValueOf(&views.FileViewType).Elem(),
		"FuncButtonType":               reflect.ValueOf(&views.FuncButtonType).Elem(),
		"InspectorType":                reflect.ValueOf(&views.InspectorType).Elem(),
		"InspectorView":                reflect.ValueOf(views.InspectorView),
		"InspectorWindow":              reflect.ValueOf(views.InspectorWindow),
		"JoinViewPath":                 reflect.ValueOf(views.JoinViewPath),
		"MapButtonType":                reflect.ValueOf(&views.MapButtonType).Elem(),
		"MapViewType":                  reflect.ValueOf(&views.MapViewType).Elem(),
		"NewArgView":                   reflect.ValueOf(views.NewArgView),
		"NewColorButton":               reflect.ValueOf(views.NewColorButton),
		"NewColorView":                 reflect.ValueOf(views.NewColorView),
		"NewDateView":                  reflect.ValueOf(views.NewDateView),
		"NewFileView":                  reflect.ValueOf(views.NewFileView),
		"NewFuncButton":                reflect.ValueOf(views.NewFuncButton),
		"NewInspector":                 reflect.ValueOf(views.NewInspector),
		"NewMapButton":                 reflect.ValueOf(views.NewMapButton),
		"NewMapView":                   reflect.ValueOf(views.NewMapView),
		"NewSliceButton":               reflect.ValueOf(views.NewSliceButton),
		"NewSliceView":                 reflect.ValueOf(views.NewSliceView),
		"NewSliceViewBase":             reflect.ValueOf(views.NewSliceViewBase),
		"NewSliceViewGrid":             reflect.ValueOf(views.NewSliceViewGrid),
		"NewSliceViewInline":           reflect.ValueOf(views.NewSliceViewInline),
		"NewSoloFuncButton":            reflect.ValueOf(views.NewSoloFuncButton),
		"NewStructButton":              reflect.ValueOf(views.NewStructButton),
		"NewStructView":                reflect.ValueOf(views.NewStructView),
		"NewTableView":                 reflect.ValueOf(views.NewTableView),
		"NewTimeView":                  reflect.ValueOf(views.NewTimeView),
		"NewTreeView":                  reflect.ValueOf(views.NewTreeView),
		"NewTreeViewFrame":             reflect.ValueOf(views.NewTreeViewFrame),
		"NewValue":                     reflect.ValueOf(views.NewValue),
		"NoSentenceCaseFor":            reflect.ValueOf(&views.NoSentenceCaseFor).Elem(),
		"NoSentenceCaseForType":        reflect.ValueOf(views.NoSentenceCaseForType),
		"OpenDialog":                   reflect.ValueOf(views.OpenDialog),
		"OpenDialogBase":               reflect.ValueOf(views.OpenDialogBase),
		"OpenDialogValue":              reflect.ValueOf(views.OpenDialogValue),
		"OpenDialogValueBase":          reflect.ValueOf(views.OpenDialogValueBase),
		"SettingsView":                 reflect.ValueOf(views.SettingsView),
		"SettingsViewToolbarBase":      reflect.ValueOf(views.SettingsViewToolbarBase),
		"SettingsWindow":               reflect.ValueOf(views.SettingsWindow),
		"SliceButtonType":              reflect.ValueOf(&views.SliceButtonType).Elem(),
		"SliceIndexByValue":            reflect.ValueOf(views.SliceIndexByValue),
		"SliceViewBaseType":            reflect.ValueOf(&views.SliceViewBaseType).Elem(),
		"SliceViewColProperty":         reflect.ValueOf(constant.MakeFromLiteral("\"sv-col\"", token.STRING, 0)),
		"SliceViewConfigured":          reflect.ValueOf(views.SliceViewConfigured),
		"SliceViewFlagsN":              reflect.ValueOf(views.SliceViewFlagsN),
		"SliceViewFlagsValues":         reflect.ValueOf(views.SliceViewFlagsValues),
		"SliceViewGridType":            reflect.ValueOf(&views.SliceViewGridType).Elem(),
		"SliceViewInFocusGrab":         reflect.ValueOf(views.SliceViewInFocusGrab),
		"SliceViewInFullRebuild":       reflect.ValueOf(views.SliceViewInFullRebuild),
		"SliceViewInlineType":          reflect.ValueOf(&views.SliceViewInlineType).Elem(),
		"SliceViewIsArray":             reflect.ValueOf(views.SliceViewIsArray),
		"SliceViewReadOnlyKeyNav":      reflect.ValueOf(views.SliceViewReadOnlyKeyNav),
		"SliceViewReadOnlyMultiSelect": reflect.ValueOf(views.SliceViewReadOnlyMultiSelect),
		"SliceViewRowProperty":         reflect.ValueOf(constant.MakeFromLiteral("\"sv-row\"", token.STRING, 0)),
		"SliceViewSelectMode":          reflect.ValueOf(views.SliceViewSelectMode),
		"SliceViewShowIndex":           reflect.ValueOf(views.SliceViewShowIndex),
		"SliceViewType":                reflect.ValueOf(&views.SliceViewType).Elem(),
		"StructButtonType":             reflect.ValueOf(&views.StructButtonType).Elem(),
		"StructSliceIndexByValue":      reflect.ValueOf(views.StructSliceIndexByValue),
		"StructViewType":               reflect.ValueOf(&views.StructViewType).Elem(),
		"TableViewType":                reflect.ValueOf(&views.TableViewType).Elem(),
		"TimeViewType":                 reflect.ValueOf(&views.TimeViewType).Elem(),
		"ToValue":                      reflect.ValueOf(views.ToValue),
		"TreeViewFlagClosed":           reflect.ValueOf(views.TreeViewFlagClosed),
		"TreeViewFlagSelectMode":       reflect.ValueOf(views.TreeViewFlagSelectMode),
		"TreeViewFlagsN":               reflect.ValueOf(views.TreeViewFlagsN),
		"TreeViewFlagsValues":          reflect.ValueOf(views.TreeViewFlagsValues),
		"TreeViewInOpen":               reflect.ValueOf(views.TreeViewInOpen),
		"TreeViewPageSteps":            reflect.ValueOf(&views.TreeViewPageSteps).Elem(),
		"TreeViewTempMovedTag":         reflect.ValueOf(constant.MakeFromLiteral("\"_\\\\&MOVED\\\\&\"", token.STRING, 0)),
		"TreeViewType":                 reflect.ValueOf(&views.TreeViewType).Elem(),
		"ValueDialogNewWindow":         reflect.ValueOf(views.ValueDialogNewWindow),
		"ValueFlagsN":                  reflect.ValueOf(views.ValueFlagsN),
		"ValueFlagsValues":             reflect.ValueOf(views.ValueFlagsValues),
		"ValueHasSavedDoc":             reflect.ValueOf(views.ValueHasSavedDoc),
		"ValueHasSavedLabel":           reflect.ValueOf(views.ValueHasSavedLabel),
		"ValueMap":                     reflect.ValueOf(&views.ValueMap).Elem(),
		"ValueMapKey":                  reflect.ValueOf(views.ValueMapKey),
		"ValueReadOnly":                reflect.ValueOf(views.ValueReadOnly),

		// type definitions
		"ArgView":            reflect.ValueOf((*views.ArgView)(nil)),
		"BitFlagValue":       reflect.ValueOf((*views.BitFlagValue)(nil)),
		"BoolValue":          reflect.ValueOf((*views.BoolValue)(nil)),
		"ByteSliceValue":     reflect.ValueOf((*views.ByteSliceValue)(nil)),
		"ColorButton":        reflect.ValueOf((*views.ColorButton)(nil)),
		"ColorMapName":       reflect.ValueOf((*views.ColorMapName)(nil)),
		"ColorMapValue":      reflect.ValueOf((*views.ColorMapValue)(nil)),
		"ColorValue":         reflect.ValueOf((*views.ColorValue)(nil)),
		"ColorView":          reflect.ValueOf((*views.ColorView)(nil)),
		"ConfigDialoger":     reflect.ValueOf((*views.ConfigDialoger)(nil)),
		"DateView":           reflect.ValueOf((*views.DateView)(nil)),
		"DurationValue":      reflect.ValueOf((*views.DurationValue)(nil)),
		"EnumValue":          reflect.ValueOf((*views.EnumValue)(nil)),
		"FieldValuer":        reflect.ValueOf((*views.FieldValuer)(nil)),
		"FileValue":          reflect.ValueOf((*views.FileValue)(nil)),
		"FileView":           reflect.ValueOf((*views.FileView)(nil)),
		"FileViewFilterFunc": reflect.ValueOf((*views.FileViewFilterFunc)(nil)),
		"FontValue":          reflect.ValueOf((*views.FontValue)(nil)),
		"FuncButton":         reflect.ValueOf((*views.FuncButton)(nil)),
		"FuncValue":          reflect.ValueOf((*views.FuncValue)(nil)),
		"IconValue":          reflect.ValueOf((*views.IconValue)(nil)),
		"Inspector":          reflect.ValueOf((*views.Inspector)(nil)),
		"KeyChordValue":      reflect.ValueOf((*views.KeyChordValue)(nil)),
		"KeyMapValue":        reflect.ValueOf((*views.KeyMapValue)(nil)),
		"MapButton":          reflect.ValueOf((*views.MapButton)(nil)),
		"MapInlineValue":     reflect.ValueOf((*views.MapInlineValue)(nil)),
		"MapValue":           reflect.ValueOf((*views.MapValue)(nil)),
		"MapView":            reflect.ValueOf((*views.MapView)(nil)),
		"NilValue":           reflect.ValueOf((*views.NilValue)(nil)),
		"NumberValue":        reflect.ValueOf((*views.NumberValue)(nil)),
		"OpenDialoger":       reflect.ValueOf((*views.OpenDialoger)(nil)),
		"RuneSliceValue":     reflect.ValueOf((*views.RuneSliceValue)(nil)),
		"SliceButton":        reflect.ValueOf((*views.SliceButton)(nil)),
		"SliceInlineValue":   reflect.ValueOf((*views.SliceInlineValue)(nil)),
		"SliceValue":         reflect.ValueOf((*views.SliceValue)(nil)),
		"SliceView":          reflect.ValueOf((*views.SliceView)(nil)),
		"SliceViewBase":      reflect.ValueOf((*views.SliceViewBase)(nil)),
		"SliceViewFlags":     reflect.ValueOf((*views.SliceViewFlags)(nil)),
		"SliceViewGrid":      reflect.ValueOf((*views.SliceViewGrid)(nil)),
		"SliceViewInline":    reflect.ValueOf((*views.SliceViewInline)(nil)),
		"SliceViewStyleFunc": reflect.ValueOf((*views.SliceViewStyleFunc)(nil)),
		"SliceViewer":        reflect.ValueOf((*views.SliceViewer)(nil)),
		"SliderValue":        reflect.ValueOf((*views.SliderValue)(nil)),
		"StringValue":        reflect.ValueOf((*views.StringValue)(nil)),
		"StructButton":       reflect.ValueOf((*views.StructButton)(nil)),
		"StructInlineValue":  reflect.ValueOf((*views.StructInlineValue)(nil)),
		"StructValue":        reflect.ValueOf((*views.StructValue)(nil)),
		"StructView":         reflect.ValueOf((*views.StructView)(nil)),
		"TableView":          reflect.ValueOf((*views.TableView)(nil)),
		"TableViewStyleFunc": reflect.ValueOf((*views.TableViewStyleFunc)(nil)),
		"TimeValue":          reflect.ValueOf((*views.TimeValue)(nil)),
		"TimeView":           reflect.ValueOf((*views.TimeView)(nil)),
		"TreeValue":          reflect.ValueOf((*views.TreeValue)(nil)),
		"TreeView":           reflect.ValueOf((*views.TreeView)(nil)),
		"TreeViewFlags":      reflect.ValueOf((*views.TreeViewFlags)(nil)),
		"TreeViewer":         reflect.ValueOf((*views.TreeViewer)(nil)),
		"TypeValue":          reflect.ValueOf((*views.TypeValue)(nil)),
		"Value":              reflect.ValueOf((*views.Value)(nil)),
		"ValueData":          reflect.ValueOf((*views.ValueData)(nil)),
		"ValueFlags":         reflect.ValueOf((*views.ValueFlags)(nil)),
		"Valuer":             reflect.ValueOf((*views.Valuer)(nil)),

		// interface wrapper definitions
		"_ConfigDialoger": reflect.ValueOf((*_cogentcore_org_core_views_ConfigDialoger)(nil)),
		"_FieldValuer":    reflect.ValueOf((*_cogentcore_org_core_views_FieldValuer)(nil)),
		"_OpenDialoger":   reflect.ValueOf((*_cogentcore_org_core_views_OpenDialoger)(nil)),
		"_SliceViewer":    reflect.ValueOf((*_cogentcore_org_core_views_SliceViewer)(nil)),
		"_TreeViewer":     reflect.ValueOf((*_cogentcore_org_core_views_TreeViewer)(nil)),
		"_Value":          reflect.ValueOf((*_cogentcore_org_core_views_Value)(nil)),
		"_Valuer":         reflect.ValueOf((*_cogentcore_org_core_views_Valuer)(nil)),
	}
}

// _cogentcore_org_core_views_ConfigDialoger is an interface wrapper for ConfigDialoger type
type _cogentcore_org_core_views_ConfigDialoger struct {
	IValue        interface{}
	WConfigDialog func(d *core.Body) (bool, func())
}

func (W _cogentcore_org_core_views_ConfigDialoger) ConfigDialog(d *core.Body) (bool, func()) {
	return W.WConfigDialog(d)
}

// _cogentcore_org_core_views_FieldValuer is an interface wrapper for FieldValuer type
type _cogentcore_org_core_views_FieldValuer struct {
	IValue      interface{}
	WFieldValue func(field string, fval any) views.Value
}

func (W _cogentcore_org_core_views_FieldValuer) FieldValue(field string, fval any) views.Value {
	return W.WFieldValue(field, fval)
}

// _cogentcore_org_core_views_OpenDialoger is an interface wrapper for OpenDialoger type
type _cogentcore_org_core_views_OpenDialoger struct {
	IValue      interface{}
	WOpenDialog func(ctx core.Widget, fun func())
}

func (W _cogentcore_org_core_views_OpenDialoger) OpenDialog(ctx core.Widget, fun func()) {
	W.WOpenDialog(ctx, fun)
}

// _cogentcore_org_core_views_SliceViewer is an interface wrapper for SliceViewer type
type _cogentcore_org_core_views_SliceViewer struct {
	IValue            interface{}
	WAsSliceViewBase  func() *views.SliceViewBase
	WConfigRow        func(c *core.Plan, i int, si int)
	WCopySelectToMime func() mimedata.Mimes
	WDragDrop         func(e events.Event)
	WDragStart        func(e events.Event)
	WDropDeleteSource func(e events.Event)
	WDropFinalize     func(de *events.DragDrop)
	WHasStyleFunc     func() bool
	WMakePasteMenu    func(m *core.Scene, md mimedata.Mimes, idx int, mod events.DropMods, fun func())
	WMimeDataType     func() string
	WPasteAssign      func(md mimedata.Mimes, idx int)
	WPasteAtIndex     func(md mimedata.Mimes, idx int)
	WRowFirstWidget   func(row int) (*core.WidgetBase, bool)
	WRowGrabFocus     func(row int) *core.WidgetBase
	WRowWidgetNs      func() (nWidgPerRow int, idxOff int)
	WSliceDeleteAt    func(idx int)
	WSliceGrid        func() *views.SliceViewGrid
	WSliceNewAt       func(idx int)
	WStyleRow         func(w core.Widget, idx int, fidx int)
	WStyleValue       func(w core.Widget, s *styles.Style, row int, col int)
	WUpdateMaxWidths  func()
	WUpdateSliceSize  func() int
}

func (W _cogentcore_org_core_views_SliceViewer) AsSliceViewBase() *views.SliceViewBase {
	return W.WAsSliceViewBase()
}
func (W _cogentcore_org_core_views_SliceViewer) ConfigRow(c *core.Plan, i int, si int) {
	W.WConfigRow(c, i, si)
}
func (W _cogentcore_org_core_views_SliceViewer) CopySelectToMime() mimedata.Mimes {
	return W.WCopySelectToMime()
}
func (W _cogentcore_org_core_views_SliceViewer) DragDrop(e events.Event) {
	W.WDragDrop(e)
}
func (W _cogentcore_org_core_views_SliceViewer) DragStart(e events.Event) {
	W.WDragStart(e)
}
func (W _cogentcore_org_core_views_SliceViewer) DropDeleteSource(e events.Event) {
	W.WDropDeleteSource(e)
}
func (W _cogentcore_org_core_views_SliceViewer) DropFinalize(de *events.DragDrop) {
	W.WDropFinalize(de)
}
func (W _cogentcore_org_core_views_SliceViewer) HasStyleFunc() bool {
	return W.WHasStyleFunc()
}
func (W _cogentcore_org_core_views_SliceViewer) MakePasteMenu(m *core.Scene, md mimedata.Mimes, idx int, mod events.DropMods, fun func()) {
	W.WMakePasteMenu(m, md, idx, mod, fun)
}
func (W _cogentcore_org_core_views_SliceViewer) MimeDataType() string {
	return W.WMimeDataType()
}
func (W _cogentcore_org_core_views_SliceViewer) PasteAssign(md mimedata.Mimes, idx int) {
	W.WPasteAssign(md, idx)
}
func (W _cogentcore_org_core_views_SliceViewer) PasteAtIndex(md mimedata.Mimes, idx int) {
	W.WPasteAtIndex(md, idx)
}
func (W _cogentcore_org_core_views_SliceViewer) RowFirstWidget(row int) (*core.WidgetBase, bool) {
	return W.WRowFirstWidget(row)
}
func (W _cogentcore_org_core_views_SliceViewer) RowGrabFocus(row int) *core.WidgetBase {
	return W.WRowGrabFocus(row)
}
func (W _cogentcore_org_core_views_SliceViewer) RowWidgetNs() (nWidgPerRow int, idxOff int) {
	return W.WRowWidgetNs()
}
func (W _cogentcore_org_core_views_SliceViewer) SliceDeleteAt(idx int) {
	W.WSliceDeleteAt(idx)
}
func (W _cogentcore_org_core_views_SliceViewer) SliceGrid() *views.SliceViewGrid {
	return W.WSliceGrid()
}
func (W _cogentcore_org_core_views_SliceViewer) SliceNewAt(idx int) {
	W.WSliceNewAt(idx)
}
func (W _cogentcore_org_core_views_SliceViewer) StyleRow(w core.Widget, idx int, fidx int) {
	W.WStyleRow(w, idx, fidx)
}
func (W _cogentcore_org_core_views_SliceViewer) StyleValue(w core.Widget, s *styles.Style, row int, col int) {
	W.WStyleValue(w, s, row, col)
}
func (W _cogentcore_org_core_views_SliceViewer) UpdateMaxWidths() {
	W.WUpdateMaxWidths()
}
func (W _cogentcore_org_core_views_SliceViewer) UpdateSliceSize() int {
	return W.WUpdateSliceSize()
}

// _cogentcore_org_core_views_TreeViewer is an interface wrapper for TreeViewer type
type _cogentcore_org_core_views_TreeViewer struct {
	IValue               interface{}
	WAbilityIs           func(flag abilities.Abilities) bool
	WAddChild            func(kid tree.Node) error
	WAddChildNode        func()
	WAddContextMenu      func(menu func(m *core.Scene)) *core.WidgetBase
	WApplyContextMenus   func(m *core.Scene)
	WApplyStyle          func()
	WAsTreeNode          func() *tree.NodeBase
	WAsTreeView          func() *views.TreeView
	WAsWidget            func() *core.WidgetBase
	WBaseType            func() *types.Type
	WCanOpen             func() bool
	WChild               func(i int) tree.Node
	WChildBackground     func(child core.Widget) image.Image
	WChildByName         func(name string, startIndex ...int) tree.Node
	WChildByType         func(t *types.Type, embeds bool, startIndex ...int) tree.Node
	WChildren            func() *tree.Slice
	WClone               func() tree.Node
	WConfig              func(c *core.Plan)
	WConfigChildren      func(config tree.Config) bool
	WConfigWidget        func()
	WContextMenuPos      func(e events.Event) image.Point
	WCopy                func(reset bool)
	WCopyFieldsFrom      func(from tree.Node)
	WCopyFrom            func(src tree.Node)
	WCut                 func()
	WDelete              func()
	WDeleteChild         func(child tree.Node) bool
	WDeleteChildAtIndex  func(idx int) bool
	WDeleteChildByName   func(name string) bool
	WDeleteChildren      func()
	WDeleteNode          func()
	WDeleteProperty      func(key string)
	WDestroy             func()
	WDirectRenderDraw    func(drw system.Drawer, idx int, flipY bool)
	WDirectRenderImage   func(drw system.Drawer, idx int)
	WDragDrop            func(e events.Event)
	WDragStart           func(e events.Event)
	WDropDeleteSource    func(e events.Event)
	WDropFinalize        func(de *events.DragDrop)
	WDuplicate           func()
	WFieldByName         func(field string) (tree.Node, error)
	WFindPath            func(path string) tree.Node
	WFlagType            func() enums.BitFlagSetter
	WHandleEvent         func(e events.Event)
	WHasChildren         func() bool
	WIndexInParent       func() int
	WInsertAfter         func()
	WInsertBefore        func()
	WInsertChild         func(kid tree.Node, at int) error
	WInsertNewChild      func(typ *types.Type, at int) tree.Node
	WIs                  func(f enums.BitFlag) bool
	WIsVisible           func() bool
	WMakePasteMenu       func(m *core.Scene, md mimedata.Mimes, fun func())
	WMimeData            func(md *mimedata.Mimes)
	WName                func() string
	WNew                 func() tree.Node
	WNewChild            func(typ *types.Type) tree.Node
	WNodeType            func() *types.Type
	WNodeWalkDown        func(fun func(n tree.Node) bool)
	WNumChildren         func() int
	WNumLifetimeChildren func() uint64
	WOn                  func(etype events.Types, fun func(e events.Event)) *core.WidgetBase
	WOnAdd               func()
	WOnChildAdded        func(child tree.Node)
	WOnClick             func(fun func(e events.Event)) *core.WidgetBase
	WOnClose             func()
	WOnDoubleClick       func(e events.Event)
	WOnInit              func()
	WOnOpen              func()
	WOnWidgetAdded       func(f func(w core.Widget)) *core.WidgetBase
	WParent              func() tree.Node
	WParentByName        func(name string) tree.Node
	WParentByType        func(t *types.Type, embeds bool) tree.Node
	WParentLevel         func(parent tree.Node) int
	WPaste               func()
	WPath                func() string
	WPathFrom            func(parent tree.Node) string
	WPosition            func()
	WProperties          func() map[string]any
	WProperty            func(key string) any
	WRender              func()
	WRenderWidget        func()
	WScenePos            func()
	WSend                func(e events.Types, orig ...events.Event)
	WSetAbilities        func(on bool, able ...abilities.Abilities) *core.WidgetBase
	WSetChild            func(kid tree.Node, idx int) error
	WSetFlag             func(on bool, f ...enums.BitFlag)
	WSetNChildren        func(n int, typ *types.Type, nameStub ...string) bool
	WSetName             func(name string)
	WSetProperty         func(key string, value any)
	WSetState            func(on bool, state ...states.States) *core.WidgetBase
	WShowContextMenu     func(e events.Event)
	WSizeDown            func(iter int) bool
	WSizeFinal           func()
	WSizeUp              func()
	WStateIs             func(flag states.States) bool
	WStyle               func(s func(s *styles.Style)) *core.WidgetBase
	WThis                func() tree.Node
	WUpdate              func()
	WUpdateBranchIcons   func()
	WWalkDown            func(fun func(n tree.Node) bool)
	WWalkDownBreadth     func(fun func(n tree.Node) bool)
	WWalkDownPost        func(doChildTest func(n tree.Node) bool, fun func(n tree.Node) bool)
	WWalkUp              func(fun func(n tree.Node) bool) bool
	WWalkUpParent        func(fun func(n tree.Node) bool) bool
	WWidgetTooltip       func(pos image.Point) (string, image.Point)
}

func (W _cogentcore_org_core_views_TreeViewer) AbilityIs(flag abilities.Abilities) bool {
	return W.WAbilityIs(flag)
}
func (W _cogentcore_org_core_views_TreeViewer) AddChild(kid tree.Node) error {
	return W.WAddChild(kid)
}
func (W _cogentcore_org_core_views_TreeViewer) AddChildNode() {
	W.WAddChildNode()
}
func (W _cogentcore_org_core_views_TreeViewer) AddContextMenu(menu func(m *core.Scene)) *core.WidgetBase {
	return W.WAddContextMenu(menu)
}
func (W _cogentcore_org_core_views_TreeViewer) ApplyContextMenus(m *core.Scene) {
	W.WApplyContextMenus(m)
}
func (W _cogentcore_org_core_views_TreeViewer) ApplyStyle() {
	W.WApplyStyle()
}
func (W _cogentcore_org_core_views_TreeViewer) AsTreeNode() *tree.NodeBase {
	return W.WAsTreeNode()
}
func (W _cogentcore_org_core_views_TreeViewer) AsTreeView() *views.TreeView {
	return W.WAsTreeView()
}
func (W _cogentcore_org_core_views_TreeViewer) AsWidget() *core.WidgetBase {
	return W.WAsWidget()
}
func (W _cogentcore_org_core_views_TreeViewer) BaseType() *types.Type {
	return W.WBaseType()
}
func (W _cogentcore_org_core_views_TreeViewer) CanOpen() bool {
	return W.WCanOpen()
}
func (W _cogentcore_org_core_views_TreeViewer) Child(i int) tree.Node {
	return W.WChild(i)
}
func (W _cogentcore_org_core_views_TreeViewer) ChildBackground(child core.Widget) image.Image {
	return W.WChildBackground(child)
}
func (W _cogentcore_org_core_views_TreeViewer) ChildByName(name string, startIndex ...int) tree.Node {
	return W.WChildByName(name, startIndex...)
}
func (W _cogentcore_org_core_views_TreeViewer) ChildByType(t *types.Type, embeds bool, startIndex ...int) tree.Node {
	return W.WChildByType(t, embeds, startIndex...)
}
func (W _cogentcore_org_core_views_TreeViewer) Children() *tree.Slice {
	return W.WChildren()
}
func (W _cogentcore_org_core_views_TreeViewer) Clone() tree.Node {
	return W.WClone()
}
func (W _cogentcore_org_core_views_TreeViewer) Make(c *core.Plan) {
	W.WConfig(c)
}
func (W _cogentcore_org_core_views_TreeViewer) ConfigChildren(config tree.Config) bool {
	return W.WConfigChildren(config)
}
func (W _cogentcore_org_core_views_TreeViewer) ConfigWidget() {
	W.WConfigWidget()
}
func (W _cogentcore_org_core_views_TreeViewer) ContextMenuPos(e events.Event) image.Point {
	return W.WContextMenuPos(e)
}
func (W _cogentcore_org_core_views_TreeViewer) Copy(reset bool) {
	W.WCopy(reset)
}
func (W _cogentcore_org_core_views_TreeViewer) CopyFieldsFrom(from tree.Node) {
	W.WCopyFieldsFrom(from)
}
func (W _cogentcore_org_core_views_TreeViewer) CopyFrom(src tree.Node) {
	W.WCopyFrom(src)
}
func (W _cogentcore_org_core_views_TreeViewer) Cut() {
	W.WCut()
}
func (W _cogentcore_org_core_views_TreeViewer) Delete() {
	W.WDelete()
}
func (W _cogentcore_org_core_views_TreeViewer) DeleteChild(child tree.Node) bool {
	return W.WDeleteChild(child)
}
func (W _cogentcore_org_core_views_TreeViewer) DeleteChildAtIndex(idx int) bool {
	return W.WDeleteChildAtIndex(idx)
}
func (W _cogentcore_org_core_views_TreeViewer) DeleteChildByName(name string) bool {
	return W.WDeleteChildByName(name)
}
func (W _cogentcore_org_core_views_TreeViewer) DeleteChildren() {
	W.WDeleteChildren()
}
func (W _cogentcore_org_core_views_TreeViewer) DeleteNode() {
	W.WDeleteNode()
}
func (W _cogentcore_org_core_views_TreeViewer) DeleteProperty(key string) {
	W.WDeleteProperty(key)
}
func (W _cogentcore_org_core_views_TreeViewer) Destroy() {
	W.WDestroy()
}
func (W _cogentcore_org_core_views_TreeViewer) DirectRenderDraw(drw system.Drawer, idx int, flipY bool) {
	W.WDirectRenderDraw(drw, idx, flipY)
}
func (W _cogentcore_org_core_views_TreeViewer) DirectRenderImage(drw system.Drawer, idx int) {
	W.WDirectRenderImage(drw, idx)
}
func (W _cogentcore_org_core_views_TreeViewer) DragDrop(e events.Event) {
	W.WDragDrop(e)
}
func (W _cogentcore_org_core_views_TreeViewer) DragStart(e events.Event) {
	W.WDragStart(e)
}
func (W _cogentcore_org_core_views_TreeViewer) DropDeleteSource(e events.Event) {
	W.WDropDeleteSource(e)
}
func (W _cogentcore_org_core_views_TreeViewer) DropFinalize(de *events.DragDrop) {
	W.WDropFinalize(de)
}
func (W _cogentcore_org_core_views_TreeViewer) Duplicate() {
	W.WDuplicate()
}
func (W _cogentcore_org_core_views_TreeViewer) FieldByName(field string) (tree.Node, error) {
	return W.WFieldByName(field)
}
func (W _cogentcore_org_core_views_TreeViewer) FindPath(path string) tree.Node {
	return W.WFindPath(path)
}
func (W _cogentcore_org_core_views_TreeViewer) FlagType() enums.BitFlagSetter {
	return W.WFlagType()
}
func (W _cogentcore_org_core_views_TreeViewer) HandleEvent(e events.Event) {
	W.WHandleEvent(e)
}
func (W _cogentcore_org_core_views_TreeViewer) HasChildren() bool {
	return W.WHasChildren()
}
func (W _cogentcore_org_core_views_TreeViewer) IndexInParent() int {
	return W.WIndexInParent()
}
func (W _cogentcore_org_core_views_TreeViewer) InsertAfter() {
	W.WInsertAfter()
}
func (W _cogentcore_org_core_views_TreeViewer) InsertBefore() {
	W.WInsertBefore()
}
func (W _cogentcore_org_core_views_TreeViewer) InsertChild(kid tree.Node, at int) error {
	return W.WInsertChild(kid, at)
}
func (W _cogentcore_org_core_views_TreeViewer) InsertNewChild(typ *types.Type, at int) tree.Node {
	return W.WInsertNewChild(typ, at)
}
func (W _cogentcore_org_core_views_TreeViewer) Is(f enums.BitFlag) bool {
	return W.WIs(f)
}
func (W _cogentcore_org_core_views_TreeViewer) IsVisible() bool {
	return W.WIsVisible()
}
func (W _cogentcore_org_core_views_TreeViewer) MakePasteMenu(m *core.Scene, md mimedata.Mimes, fun func()) {
	W.WMakePasteMenu(m, md, fun)
}
func (W _cogentcore_org_core_views_TreeViewer) MimeData(md *mimedata.Mimes) {
	W.WMimeData(md)
}
func (W _cogentcore_org_core_views_TreeViewer) Name() string {
	return W.WName()
}
func (W _cogentcore_org_core_views_TreeViewer) New() tree.Node {
	return W.WNew()
}
func (W _cogentcore_org_core_views_TreeViewer) NewChild(typ *types.Type) tree.Node {
	return W.WNewChild(typ)
}
func (W _cogentcore_org_core_views_TreeViewer) NodeType() *types.Type {
	return W.WNodeType()
}
func (W _cogentcore_org_core_views_TreeViewer) NodeWalkDown(fun func(n tree.Node) bool) {
	W.WNodeWalkDown(fun)
}
func (W _cogentcore_org_core_views_TreeViewer) NumChildren() int {
	return W.WNumChildren()
}
func (W _cogentcore_org_core_views_TreeViewer) NumLifetimeChildren() uint64 {
	return W.WNumLifetimeChildren()
}
func (W _cogentcore_org_core_views_TreeViewer) On(etype events.Types, fun func(e events.Event)) *core.WidgetBase {
	return W.WOn(etype, fun)
}
func (W _cogentcore_org_core_views_TreeViewer) OnAdd() {
	W.WOnAdd()
}
func (W _cogentcore_org_core_views_TreeViewer) OnChildAdded(child tree.Node) {
	W.WOnChildAdded(child)
}
func (W _cogentcore_org_core_views_TreeViewer) OnClick(fun func(e events.Event)) *core.WidgetBase {
	return W.WOnClick(fun)
}
func (W _cogentcore_org_core_views_TreeViewer) OnClose() {
	W.WOnClose()
}
func (W _cogentcore_org_core_views_TreeViewer) OnDoubleClick(e events.Event) {
	W.WOnDoubleClick(e)
}
func (W _cogentcore_org_core_views_TreeViewer) OnInit() {
	W.WOnInit()
}
func (W _cogentcore_org_core_views_TreeViewer) OnOpen() {
	W.WOnOpen()
}
func (W _cogentcore_org_core_views_TreeViewer) OnWidgetAdded(f func(w core.Widget)) *core.WidgetBase {
	return W.WOnWidgetAdded(f)
}
func (W _cogentcore_org_core_views_TreeViewer) Parent() tree.Node {
	return W.WParent()
}
func (W _cogentcore_org_core_views_TreeViewer) ParentByName(name string) tree.Node {
	return W.WParentByName(name)
}
func (W _cogentcore_org_core_views_TreeViewer) ParentByType(t *types.Type, embeds bool) tree.Node {
	return W.WParentByType(t, embeds)
}
func (W _cogentcore_org_core_views_TreeViewer) ParentLevel(parent tree.Node) int {
	return W.WParentLevel(parent)
}
func (W _cogentcore_org_core_views_TreeViewer) Paste() {
	W.WPaste()
}
func (W _cogentcore_org_core_views_TreeViewer) Path() string {
	return W.WPath()
}
func (W _cogentcore_org_core_views_TreeViewer) PathFrom(parent tree.Node) string {
	return W.WPathFrom(parent)
}
func (W _cogentcore_org_core_views_TreeViewer) Position() {
	W.WPosition()
}
func (W _cogentcore_org_core_views_TreeViewer) Properties() map[string]any {
	return W.WProperties()
}
func (W _cogentcore_org_core_views_TreeViewer) Property(key string) any {
	return W.WProperty(key)
}
func (W _cogentcore_org_core_views_TreeViewer) Render() {
	W.WRender()
}
func (W _cogentcore_org_core_views_TreeViewer) RenderWidget() {
	W.WRenderWidget()
}
func (W _cogentcore_org_core_views_TreeViewer) ScenePos() {
	W.WScenePos()
}
func (W _cogentcore_org_core_views_TreeViewer) Send(e events.Types, orig ...events.Event) {
	W.WSend(e, orig...)
}
func (W _cogentcore_org_core_views_TreeViewer) SetAbilities(on bool, able ...abilities.Abilities) *core.WidgetBase {
	return W.WSetAbilities(on, able...)
}
func (W _cogentcore_org_core_views_TreeViewer) SetChild(kid tree.Node, idx int) error {
	return W.WSetChild(kid, idx)
}
func (W _cogentcore_org_core_views_TreeViewer) SetFlag(on bool, f ...enums.BitFlag) {
	W.WSetFlag(on, f...)
}
func (W _cogentcore_org_core_views_TreeViewer) SetNChildren(n int, typ *types.Type, nameStub ...string) bool {
	return W.WSetNChildren(n, typ, nameStub...)
}
func (W _cogentcore_org_core_views_TreeViewer) SetName(name string) {
	W.WSetName(name)
}
func (W _cogentcore_org_core_views_TreeViewer) SetProperty(key string, value any) {
	W.WSetProperty(key, value)
}
func (W _cogentcore_org_core_views_TreeViewer) SetState(on bool, state ...states.States) *core.WidgetBase {
	return W.WSetState(on, state...)
}
func (W _cogentcore_org_core_views_TreeViewer) ShowContextMenu(e events.Event) {
	W.WShowContextMenu(e)
}
func (W _cogentcore_org_core_views_TreeViewer) SizeDown(iter int) bool {
	return W.WSizeDown(iter)
}
func (W _cogentcore_org_core_views_TreeViewer) SizeFinal() {
	W.WSizeFinal()
}
func (W _cogentcore_org_core_views_TreeViewer) SizeUp() {
	W.WSizeUp()
}
func (W _cogentcore_org_core_views_TreeViewer) StateIs(flag states.States) bool {
	return W.WStateIs(flag)
}
func (W _cogentcore_org_core_views_TreeViewer) Style(s func(s *styles.Style)) *core.WidgetBase {
	return W.WStyle(s)
}
func (W _cogentcore_org_core_views_TreeViewer) This() tree.Node {
	return W.WThis()
}
func (W _cogentcore_org_core_views_TreeViewer) Update() {
	W.WUpdate()
}
func (W _cogentcore_org_core_views_TreeViewer) UpdateBranchIcons() {
	W.WUpdateBranchIcons()
}
func (W _cogentcore_org_core_views_TreeViewer) WalkDown(fun func(n tree.Node) bool) {
	W.WWalkDown(fun)
}
func (W _cogentcore_org_core_views_TreeViewer) WalkDownBreadth(fun func(n tree.Node) bool) {
	W.WWalkDownBreadth(fun)
}
func (W _cogentcore_org_core_views_TreeViewer) WalkDownPost(doChildTest func(n tree.Node) bool, fun func(n tree.Node) bool) {
	W.WWalkDownPost(doChildTest, fun)
}
func (W _cogentcore_org_core_views_TreeViewer) WalkUp(fun func(n tree.Node) bool) bool {
	return W.WWalkUp(fun)
}
func (W _cogentcore_org_core_views_TreeViewer) WalkUpParent(fun func(n tree.Node) bool) bool {
	return W.WWalkUpParent(fun)
}
func (W _cogentcore_org_core_views_TreeViewer) WidgetTooltip(pos image.Point) (string, image.Point) {
	return W.WWidgetTooltip(pos)
}

// _cogentcore_org_core_views_Value is an interface wrapper for Value type
type _cogentcore_org_core_views_Value struct {
	IValue          interface{}
	WAllTags        func() map[string]string
	WAsValueData    func() *views.ValueData
	WAsWidget       func() core.Widget
	WAsWidgetBase   func() *core.WidgetBase
	WConfig         func()
	WDoc            func() string
	WIs             func(f enums.BitFlag) bool
	WIsReadOnly     func() bool
	WLabel          func() string
	WName           func() string
	WOnChange       func(fun func(e events.Event))
	WOwnerKind      func() reflect.Kind
	WSendChange     func(orig ...events.Event)
	WSetDoc         func(doc string) *views.ValueData
	WSetFlag        func(on bool, f ...enums.BitFlag)
	WSetLabel       func(label string) *views.ValueData
	WSetMapKey      func(val reflect.Value, owner any)
	WSetMapValue    func(val reflect.Value, owner any, key any, keyView views.Value, viewPath string)
	WSetName        func(name string)
	WSetReadOnly    func(ro bool)
	WSetSliceValue  func(val reflect.Value, owner any, idx int, viewPath string)
	WSetSoloValue   func(val reflect.Value)
	WSetStructValue func(val reflect.Value, owner any, field *reflect.StructField, viewPath string)
	WSetTag         func(tag string, value string)
	WSetTags        func(tags map[string]string)
	WSetValue       func(val any) bool
	WSetWidget      func(w core.Widget)
	WString         func() string
	WTag            func(tag string) (string, bool)
	WUpdate         func()
	WVal            func() reflect.Value
	WWidgetType     func() *types.Type
}

func (W _cogentcore_org_core_views_Value) AllTags() map[string]string {
	return W.WAllTags()
}
func (W _cogentcore_org_core_views_Value) AsValueData() *views.ValueData {
	return W.WAsValueData()
}
func (W _cogentcore_org_core_views_Value) AsWidget() core.Widget {
	return W.WAsWidget()
}
func (W _cogentcore_org_core_views_Value) AsWidgetBase() *core.WidgetBase {
	return W.WAsWidgetBase()
}
func (W _cogentcore_org_core_views_Value) Config() {
	W.WConfig()
}
func (W _cogentcore_org_core_views_Value) Doc() string {
	return W.WDoc()
}
func (W _cogentcore_org_core_views_Value) Is(f enums.BitFlag) bool {
	return W.WIs(f)
}
func (W _cogentcore_org_core_views_Value) IsReadOnly() bool {
	return W.WIsReadOnly()
}
func (W _cogentcore_org_core_views_Value) Label() string {
	return W.WLabel()
}
func (W _cogentcore_org_core_views_Value) Name() string {
	return W.WName()
}
func (W _cogentcore_org_core_views_Value) OnChange(fun func(e events.Event)) {
	W.WOnChange(fun)
}
func (W _cogentcore_org_core_views_Value) OwnerKind() reflect.Kind {
	return W.WOwnerKind()
}
func (W _cogentcore_org_core_views_Value) SendChange(orig ...events.Event) {
	W.WSendChange(orig...)
}
func (W _cogentcore_org_core_views_Value) SetDoc(doc string) *views.ValueData {
	return W.WSetDoc(doc)
}
func (W _cogentcore_org_core_views_Value) SetFlag(on bool, f ...enums.BitFlag) {
	W.WSetFlag(on, f...)
}
func (W _cogentcore_org_core_views_Value) SetLabel(label string) *views.ValueData {
	return W.WSetLabel(label)
}
func (W _cogentcore_org_core_views_Value) SetMapKey(val reflect.Value, owner any) {
	W.WSetMapKey(val, owner)
}
func (W _cogentcore_org_core_views_Value) SetMapValue(val reflect.Value, owner any, key any, keyView views.Value, viewPath string) {
	W.WSetMapValue(val, owner, key, keyView, viewPath)
}
func (W _cogentcore_org_core_views_Value) SetName(name string) {
	W.WSetName(name)
}
func (W _cogentcore_org_core_views_Value) SetReadOnly(ro bool) {
	W.WSetReadOnly(ro)
}
func (W _cogentcore_org_core_views_Value) SetSliceValue(val reflect.Value, owner any, idx int, viewPath string) {
	W.WSetSliceValue(val, owner, idx, viewPath)
}
func (W _cogentcore_org_core_views_Value) SetSoloValue(val reflect.Value) {
	W.WSetSoloValue(val)
}
func (W _cogentcore_org_core_views_Value) SetStructValue(val reflect.Value, owner any, field *reflect.StructField, viewPath string) {
	W.WSetStructValue(val, owner, field, viewPath)
}
func (W _cogentcore_org_core_views_Value) SetTag(tag string, value string) {
	W.WSetTag(tag, value)
}
func (W _cogentcore_org_core_views_Value) SetTags(tags map[string]string) {
	W.WSetTags(tags)
}
func (W _cogentcore_org_core_views_Value) SetValue(val any) bool {
	return W.WSetValue(val)
}
func (W _cogentcore_org_core_views_Value) SetWidget(w core.Widget) {
	W.WSetWidget(w)
}
func (W _cogentcore_org_core_views_Value) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
func (W _cogentcore_org_core_views_Value) Tag(tag string) (string, bool) {
	return W.WTag(tag)
}
func (W _cogentcore_org_core_views_Value) Update() {
	W.WUpdate()
}
func (W _cogentcore_org_core_views_Value) Val() reflect.Value {
	return W.WVal()
}
func (W _cogentcore_org_core_views_Value) WidgetType() *types.Type {
	return W.WWidgetType()
}

// _cogentcore_org_core_views_Valuer is an interface wrapper for Valuer type
type _cogentcore_org_core_views_Valuer struct {
	IValue interface{}
	WValue func() views.Value
}

func (W _cogentcore_org_core_views_Valuer) Value() views.Value {
	return W.WValue()
}
