// Code generated by "goki generate ./..."; DO NOT EDIT.

package gide

import (
	"regexp"
	"time"

	"goki.dev/colors"
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/giv"
	"goki.dev/girl/units"
	"goki.dev/gti"
	"goki.dev/icons"
	"goki.dev/ki/v2"
	"goki.dev/mat32/v2"
	"goki.dev/ordmap"
	"goki.dev/pi/v2/filecat"
	"goki.dev/pi/v2/lex"
	"goki.dev/pi/v2/syms"
)

// DebugViewType is the [gti.Type] for [DebugView]
var DebugViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.DebugView",
	ShortName:  "gide.DebugView",
	IDName:     "debug-view",
	Doc:        "DebugView is the debugger",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Sup", &gti.Field{Name: "Sup", Type: "goki.dev/pi/v2/filecat.Supported", LocalType: "filecat.Supported", Doc: "supported file type to determine debugger", Directives: gti.Directives{}, Tag: ""}},
		{"ExePath", &gti.Field{Name: "ExePath", Type: "string", LocalType: "string", Doc: "path to executable / dir to debug", Directives: gti.Directives{}, Tag: ""}},
		{"DbgTime", &gti.Field{Name: "DbgTime", Type: "time.Time", LocalType: "time.Time", Doc: "time when dbg was last restarted", Directives: gti.Directives{}, Tag: ""}},
		{"Dbg", &gti.Field{Name: "Dbg", Type: "goki.dev/gide/v2/gidebug.GiDebug", LocalType: "gidebug.GiDebug", Doc: "the debugger", Directives: gti.Directives{}, Tag: "set:\"-\" json:\"-\" xml:\"-\""}},
		{"State", &gti.Field{Name: "State", Type: "goki.dev/gide/v2/gidebug.AllState", LocalType: "gidebug.AllState", Doc: "all relevant debug state info", Directives: gti.Directives{}, Tag: "set:\"-\" json:\"-\" xml:\"-\""}},
		{"CurFileLoc", &gti.Field{Name: "CurFileLoc", Type: "goki.dev/gide/v2/gidebug.Location", LocalType: "gidebug.Location", Doc: "current ShowFile location -- cleared before next one or run", Directives: gti.Directives{}, Tag: "set:\"-\" json:\"-\" xml:\"-\""}},
		{"BBreaks", &gti.Field{Name: "BBreaks", Type: "[]*goki.dev/gide/v2/gidebug.Break", LocalType: "[]*gidebug.Break", Doc: "backup breakpoints list -- to track deletes", Directives: gti.Directives{}, Tag: "set:\"-\" json:\"-\" xml:\"-\""}},
		{"OutBuf", &gti.Field{Name: "OutBuf", Type: "*goki.dev/gi/v2/texteditor.Buf", LocalType: "*texteditor.Buf", Doc: "output from the debugger", Directives: gti.Directives{}, Tag: "set:\"-\" json:\"-\" xml:\"-\""}},
		{"Gide", &gti.Field{Name: "Gide", Type: "goki.dev/gide/v2/gide.Gide", LocalType: "Gide", Doc: "parent gide project", Directives: gti.Directives{}, Tag: "set:\"-\" json:\"-\" xml:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &DebugView{},
})

// NewDebugView adds a new [DebugView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewDebugView(par ki.Ki, name ...string) *DebugView {
	return par.NewChild(DebugViewType, name...).(*DebugView)
}

// KiType returns the [*gti.Type] of [DebugView]
func (t *DebugView) KiType() *gti.Type {
	return DebugViewType
}

// New returns a new [*DebugView] value
func (t *DebugView) New() ki.Ki {
	return &DebugView{}
}

// SetSup sets the [DebugView.Sup]:
// supported file type to determine debugger
func (t *DebugView) SetSup(v filecat.Supported) *DebugView {
	t.Sup = v
	return t
}

// SetExePath sets the [DebugView.ExePath]:
// path to executable / dir to debug
func (t *DebugView) SetExePath(v string) *DebugView {
	t.ExePath = v
	return t
}

// SetDbgTime sets the [DebugView.DbgTime]:
// time when dbg was last restarted
func (t *DebugView) SetDbgTime(v time.Time) *DebugView {
	t.DbgTime = v
	return t
}

// SetTooltip sets the [DebugView.Tooltip]
func (t *DebugView) SetTooltip(v string) *DebugView {
	t.Tooltip = v
	return t
}

// SetClass sets the [DebugView.Class]
func (t *DebugView) SetClass(v string) *DebugView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [DebugView.CustomContextMenu]
func (t *DebugView) SetCustomContextMenu(v func(m *gi.Scene)) *DebugView {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [DebugView.Lay]
func (t *DebugView) SetLayout(v gi.Layouts) *DebugView {
	t.Lay = v
	return t
}

// SetStackTop sets the [DebugView.StackTop]
func (t *DebugView) SetStackTop(v int) *DebugView {
	t.StackTop = v
	return t
}

// StackViewType is the [gti.Type] for [StackView]
var StackViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.StackView",
	ShortName:  "gide.StackView",
	IDName:     "stack-view",
	Doc:        "StackView is a view of the stack trace",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"FindFrames", &gti.Field{Name: "FindFrames", Type: "bool", LocalType: "bool", Doc: "if true, this is a find frames, not a regular stack", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &StackView{},
})

// NewStackView adds a new [StackView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewStackView(par ki.Ki, name ...string) *StackView {
	return par.NewChild(StackViewType, name...).(*StackView)
}

// KiType returns the [*gti.Type] of [StackView]
func (t *StackView) KiType() *gti.Type {
	return StackViewType
}

// New returns a new [*StackView] value
func (t *StackView) New() ki.Ki {
	return &StackView{}
}

// SetFindFrames sets the [StackView.FindFrames]:
// if true, this is a find frames, not a regular stack
func (t *StackView) SetFindFrames(v bool) *StackView {
	t.FindFrames = v
	return t
}

// SetTooltip sets the [StackView.Tooltip]
func (t *StackView) SetTooltip(v string) *StackView {
	t.Tooltip = v
	return t
}

// SetClass sets the [StackView.Class]
func (t *StackView) SetClass(v string) *StackView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [StackView.CustomContextMenu]
func (t *StackView) SetCustomContextMenu(v func(m *gi.Scene)) *StackView {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [StackView.Lay]
func (t *StackView) SetLayout(v gi.Layouts) *StackView {
	t.Lay = v
	return t
}

// SetStackTop sets the [StackView.StackTop]
func (t *StackView) SetStackTop(v int) *StackView {
	t.StackTop = v
	return t
}

// BreakViewType is the [gti.Type] for [BreakView]
var BreakViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.BreakView",
	ShortName:  "gide.BreakView",
	IDName:     "break-view",
	Doc:        "BreakView is a view of the breakpoints",
	Directives: gti.Directives{},
	Fields:     ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &BreakView{},
})

// NewBreakView adds a new [BreakView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewBreakView(par ki.Ki, name ...string) *BreakView {
	return par.NewChild(BreakViewType, name...).(*BreakView)
}

// KiType returns the [*gti.Type] of [BreakView]
func (t *BreakView) KiType() *gti.Type {
	return BreakViewType
}

// New returns a new [*BreakView] value
func (t *BreakView) New() ki.Ki {
	return &BreakView{}
}

// SetTooltip sets the [BreakView.Tooltip]
func (t *BreakView) SetTooltip(v string) *BreakView {
	t.Tooltip = v
	return t
}

// SetClass sets the [BreakView.Class]
func (t *BreakView) SetClass(v string) *BreakView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [BreakView.CustomContextMenu]
func (t *BreakView) SetCustomContextMenu(v func(m *gi.Scene)) *BreakView {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [BreakView.Lay]
func (t *BreakView) SetLayout(v gi.Layouts) *BreakView {
	t.Lay = v
	return t
}

// SetStackTop sets the [BreakView.StackTop]
func (t *BreakView) SetStackTop(v int) *BreakView {
	t.StackTop = v
	return t
}

// ThreadViewType is the [gti.Type] for [ThreadView]
var ThreadViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.ThreadView",
	ShortName:  "gide.ThreadView",
	IDName:     "thread-view",
	Doc:        "ThreadView is a view of the threads",
	Directives: gti.Directives{},
	Fields:     ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &ThreadView{},
})

// NewThreadView adds a new [ThreadView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewThreadView(par ki.Ki, name ...string) *ThreadView {
	return par.NewChild(ThreadViewType, name...).(*ThreadView)
}

// KiType returns the [*gti.Type] of [ThreadView]
func (t *ThreadView) KiType() *gti.Type {
	return ThreadViewType
}

// New returns a new [*ThreadView] value
func (t *ThreadView) New() ki.Ki {
	return &ThreadView{}
}

// SetTooltip sets the [ThreadView.Tooltip]
func (t *ThreadView) SetTooltip(v string) *ThreadView {
	t.Tooltip = v
	return t
}

// SetClass sets the [ThreadView.Class]
func (t *ThreadView) SetClass(v string) *ThreadView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [ThreadView.CustomContextMenu]
func (t *ThreadView) SetCustomContextMenu(v func(m *gi.Scene)) *ThreadView {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [ThreadView.Lay]
func (t *ThreadView) SetLayout(v gi.Layouts) *ThreadView {
	t.Lay = v
	return t
}

// SetStackTop sets the [ThreadView.StackTop]
func (t *ThreadView) SetStackTop(v int) *ThreadView {
	t.StackTop = v
	return t
}

// TaskViewType is the [gti.Type] for [TaskView]
var TaskViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.TaskView",
	ShortName:  "gide.TaskView",
	IDName:     "task-view",
	Doc:        "TaskView is a view of the threads",
	Directives: gti.Directives{},
	Fields:     ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &TaskView{},
})

// NewTaskView adds a new [TaskView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewTaskView(par ki.Ki, name ...string) *TaskView {
	return par.NewChild(TaskViewType, name...).(*TaskView)
}

// KiType returns the [*gti.Type] of [TaskView]
func (t *TaskView) KiType() *gti.Type {
	return TaskViewType
}

// New returns a new [*TaskView] value
func (t *TaskView) New() ki.Ki {
	return &TaskView{}
}

// SetTooltip sets the [TaskView.Tooltip]
func (t *TaskView) SetTooltip(v string) *TaskView {
	t.Tooltip = v
	return t
}

// SetClass sets the [TaskView.Class]
func (t *TaskView) SetClass(v string) *TaskView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [TaskView.CustomContextMenu]
func (t *TaskView) SetCustomContextMenu(v func(m *gi.Scene)) *TaskView {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [TaskView.Lay]
func (t *TaskView) SetLayout(v gi.Layouts) *TaskView {
	t.Lay = v
	return t
}

// SetStackTop sets the [TaskView.StackTop]
func (t *TaskView) SetStackTop(v int) *TaskView {
	t.StackTop = v
	return t
}

// VarsViewType is the [gti.Type] for [VarsView]
var VarsViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.VarsView",
	ShortName:  "gide.VarsView",
	IDName:     "vars-view",
	Doc:        "VarsView is a view of the variables",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"GlobalVars", &gti.Field{Name: "GlobalVars", Type: "bool", LocalType: "bool", Doc: "if true, this is global vars, not local ones", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &VarsView{},
})

// NewVarsView adds a new [VarsView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewVarsView(par ki.Ki, name ...string) *VarsView {
	return par.NewChild(VarsViewType, name...).(*VarsView)
}

// KiType returns the [*gti.Type] of [VarsView]
func (t *VarsView) KiType() *gti.Type {
	return VarsViewType
}

// New returns a new [*VarsView] value
func (t *VarsView) New() ki.Ki {
	return &VarsView{}
}

// SetGlobalVars sets the [VarsView.GlobalVars]:
// if true, this is global vars, not local ones
func (t *VarsView) SetGlobalVars(v bool) *VarsView {
	t.GlobalVars = v
	return t
}

// SetTooltip sets the [VarsView.Tooltip]
func (t *VarsView) SetTooltip(v string) *VarsView {
	t.Tooltip = v
	return t
}

// SetClass sets the [VarsView.Class]
func (t *VarsView) SetClass(v string) *VarsView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [VarsView.CustomContextMenu]
func (t *VarsView) SetCustomContextMenu(v func(m *gi.Scene)) *VarsView {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [VarsView.Lay]
func (t *VarsView) SetLayout(v gi.Layouts) *VarsView {
	t.Lay = v
	return t
}

// SetStackTop sets the [VarsView.StackTop]
func (t *VarsView) SetStackTop(v int) *VarsView {
	t.StackTop = v
	return t
}

// VarViewType is the [gti.Type] for [VarView]
var VarViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.VarView",
	ShortName:  "gide.VarView",
	IDName:     "var-view",
	Doc:        "VarView represents a struct, creating a property editor of the fields --\nconstructs Children widgets to show the field names and editor fields for\neach field, within an overall frame with an optional title, and a button\nbox at the bottom where methods can be invoked",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Var", &gti.Field{Name: "Var", Type: "*goki.dev/gide/v2/gidebug.Variable", LocalType: "*gidebug.Variable", Doc: "variable being edited", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"FrameInfo", &gti.Field{Name: "FrameInfo", Type: "string", LocalType: "string", Doc: "frame info", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"DbgView", &gti.Field{Name: "DbgView", Type: "*goki.dev/gide/v2/gide.DebugView", LocalType: "*DebugView", Doc: "parent DebugView", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Frame", &gti.Field{Name: "Frame", Type: "goki.dev/gi/v2/gi.Frame", LocalType: "gi.Frame", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &VarView{},
})

// NewVarView adds a new [VarView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewVarView(par ki.Ki, name ...string) *VarView {
	return par.NewChild(VarViewType, name...).(*VarView)
}

// KiType returns the [*gti.Type] of [VarView]
func (t *VarView) KiType() *gti.Type {
	return VarViewType
}

// New returns a new [*VarView] value
func (t *VarView) New() ki.Ki {
	return &VarView{}
}

// SetDbgView sets the [VarView.DbgView]:
// parent DebugView
func (t *VarView) SetDbgView(v *DebugView) *VarView {
	t.DbgView = v
	return t
}

// SetTooltip sets the [VarView.Tooltip]
func (t *VarView) SetTooltip(v string) *VarView {
	t.Tooltip = v
	return t
}

// SetClass sets the [VarView.Class]
func (t *VarView) SetClass(v string) *VarView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [VarView.CustomContextMenu]
func (t *VarView) SetCustomContextMenu(v func(m *gi.Scene)) *VarView {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [VarView.Lay]
func (t *VarView) SetLayout(v gi.Layouts) *VarView {
	t.Lay = v
	return t
}

// SetStackTop sets the [VarView.StackTop]
func (t *VarView) SetStackTop(v int) *VarView {
	t.StackTop = v
	return t
}

// SetStripes sets the [VarView.Stripes]
func (t *VarView) SetStripes(v gi.Stripes) *VarView {
	t.Stripes = v
	return t
}

// FileNodeType is the [gti.Type] for [FileNode]
var FileNodeType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.FileNode",
	ShortName:  "gide.FileNode",
	IDName:     "file-node",
	Doc:        "FileNode is Gide version of FileNode for FileTree view",
	Directives: gti.Directives{},
	Fields:     ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Node", &gti.Field{Name: "Node", Type: "goki.dev/gi/v2/filetree.Node", LocalType: "filetree.Node", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &FileNode{},
})

// NewFileNode adds a new [FileNode] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewFileNode(par ki.Ki, name ...string) *FileNode {
	return par.NewChild(FileNodeType, name...).(*FileNode)
}

// KiType returns the [*gti.Type] of [FileNode]
func (t *FileNode) KiType() *gti.Type {
	return FileNodeType
}

// New returns a new [*FileNode] value
func (t *FileNode) New() ki.Ki {
	return &FileNode{}
}

// SetTooltip sets the [FileNode.Tooltip]
func (t *FileNode) SetTooltip(v string) *FileNode {
	t.Tooltip = v
	return t
}

// SetClass sets the [FileNode.Class]
func (t *FileNode) SetClass(v string) *FileNode {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [FileNode.CustomContextMenu]
func (t *FileNode) SetCustomContextMenu(v func(m *gi.Scene)) *FileNode {
	t.CustomContextMenu = v
	return t
}

// SetIcon sets the [FileNode.Icon]
func (t *FileNode) SetIcon(v icons.Icon) *FileNode {
	t.Icon = v
	return t
}

// SetIndent sets the [FileNode.Indent]
func (t *FileNode) SetIndent(v units.Value) *FileNode {
	t.Indent = v
	return t
}

// SetOpenDepth sets the [FileNode.OpenDepth]
func (t *FileNode) SetOpenDepth(v int) *FileNode {
	t.OpenDepth = v
	return t
}

// SetViewIdx sets the [FileNode.ViewIdx]
func (t *FileNode) SetViewIdx(v int) *FileNode {
	t.ViewIdx = v
	return t
}

// SetWidgetSize sets the [FileNode.WidgetSize]
func (t *FileNode) SetWidgetSize(v mat32.Vec2) *FileNode {
	t.WidgetSize = v
	return t
}

// SetRootView sets the [FileNode.RootView]
func (t *FileNode) SetRootView(v *giv.TreeView) *FileNode {
	t.RootView = v
	return t
}

// SetSelectedNodes sets the [FileNode.SelectedNodes]
func (t *FileNode) SetSelectedNodes(v []*giv.TreeView) *FileNode {
	t.SelectedNodes = v
	return t
}

// FindViewType is the [gti.Type] for [FindView]
var FindViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.FindView",
	ShortName:  "gide.FindView",
	IDName:     "find-view",
	Doc:        "FindView is a find / replace widget that displays results in a TextView\nand has a toolbar for controlling find / replace process.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Gide", &gti.Field{Name: "Gide", Type: "goki.dev/gide/v2/gide.Gide", LocalType: "Gide", Doc: "parent gide project", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\""}},
		{"LangVV", &gti.Field{Name: "LangVV", Type: "goki.dev/gi/v2/giv.Value", LocalType: "giv.Value", Doc: "langs value view", Directives: gti.Directives{}, Tag: ""}},
		{"Time", &gti.Field{Name: "Time", Type: "time.Time", LocalType: "time.Time", Doc: "time of last find", Directives: gti.Directives{}, Tag: ""}},
		{"Re", &gti.Field{Name: "Re", Type: "*regexp.Regexp", LocalType: "*regexp.Regexp", Doc: "compiled regexp", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &FindView{},
})

// NewFindView adds a new [FindView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewFindView(par ki.Ki, name ...string) *FindView {
	return par.NewChild(FindViewType, name...).(*FindView)
}

// KiType returns the [*gti.Type] of [FindView]
func (t *FindView) KiType() *gti.Type {
	return FindViewType
}

// New returns a new [*FindView] value
func (t *FindView) New() ki.Ki {
	return &FindView{}
}

// SetGide sets the [FindView.Gide]:
// parent gide project
func (t *FindView) SetGide(v Gide) *FindView {
	t.Gide = v
	return t
}

// SetLangVv sets the [FindView.LangVV]:
// langs value view
func (t *FindView) SetLangVv(v giv.Value) *FindView {
	t.LangVV = v
	return t
}

// SetTime sets the [FindView.Time]:
// time of last find
func (t *FindView) SetTime(v time.Time) *FindView {
	t.Time = v
	return t
}

// SetRe sets the [FindView.Re]:
// compiled regexp
func (t *FindView) SetRe(v *regexp.Regexp) *FindView {
	t.Re = v
	return t
}

// SetTooltip sets the [FindView.Tooltip]
func (t *FindView) SetTooltip(v string) *FindView {
	t.Tooltip = v
	return t
}

// SetClass sets the [FindView.Class]
func (t *FindView) SetClass(v string) *FindView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [FindView.CustomContextMenu]
func (t *FindView) SetCustomContextMenu(v func(m *gi.Scene)) *FindView {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [FindView.Lay]
func (t *FindView) SetLayout(v gi.Layouts) *FindView {
	t.Lay = v
	return t
}

// SetStackTop sets the [FindView.StackTop]
func (t *FindView) SetStackTop(v int) *FindView {
	t.StackTop = v
	return t
}

// SpellViewType is the [gti.Type] for [SpellView]
var SpellViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.SpellView",
	ShortName:  "gide.SpellView",
	IDName:     "spell-view",
	Doc:        "SpellView is a widget that displays results of spell check",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Gide", &gti.Field{Name: "Gide", Type: "goki.dev/gide/v2/gide.Gide", LocalType: "Gide", Doc: "parent gide project", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\" copy:\"-\""}},
		{"Text", &gti.Field{Name: "Text", Type: "*goki.dev/gide/v2/gide.TextView", LocalType: "*TextView", Doc: "textview that we're spell-checking", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\" copy:\"-\""}},
		{"Errs", &gti.Field{Name: "Errs", Type: "goki.dev/pi/v2/lex.Line", LocalType: "lex.Line", Doc: "current spelling errors", Directives: gti.Directives{}, Tag: ""}},
		{"CurLn", &gti.Field{Name: "CurLn", Type: "int", LocalType: "int", Doc: "current line in text we're on", Directives: gti.Directives{}, Tag: ""}},
		{"CurIdx", &gti.Field{Name: "CurIdx", Type: "int", LocalType: "int", Doc: "current index in Errs we're on", Directives: gti.Directives{}, Tag: ""}},
		{"UnkLex", &gti.Field{Name: "UnkLex", Type: "goki.dev/pi/v2/lex.Lex", LocalType: "lex.Lex", Doc: "current unknown lex token", Directives: gti.Directives{}, Tag: ""}},
		{"UnkWord", &gti.Field{Name: "UnkWord", Type: "string", LocalType: "string", Doc: "current unknown word", Directives: gti.Directives{}, Tag: ""}},
		{"Suggest", &gti.Field{Name: "Suggest", Type: "[]string", LocalType: "[]string", Doc: "a list of suggestions from spell checker", Directives: gti.Directives{}, Tag: ""}},
		{"LastAction", &gti.Field{Name: "LastAction", Type: "*goki.dev/gi/v2/gi.Button", LocalType: "*gi.Button", Doc: "last user action (ignore, change, learn)", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &SpellView{},
})

// NewSpellView adds a new [SpellView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewSpellView(par ki.Ki, name ...string) *SpellView {
	return par.NewChild(SpellViewType, name...).(*SpellView)
}

// KiType returns the [*gti.Type] of [SpellView]
func (t *SpellView) KiType() *gti.Type {
	return SpellViewType
}

// New returns a new [*SpellView] value
func (t *SpellView) New() ki.Ki {
	return &SpellView{}
}

// SetGide sets the [SpellView.Gide]:
// parent gide project
func (t *SpellView) SetGide(v Gide) *SpellView {
	t.Gide = v
	return t
}

// SetText sets the [SpellView.Text]:
// textview that we're spell-checking
func (t *SpellView) SetText(v *TextView) *SpellView {
	t.Text = v
	return t
}

// SetErrs sets the [SpellView.Errs]:
// current spelling errors
func (t *SpellView) SetErrs(v lex.Line) *SpellView {
	t.Errs = v
	return t
}

// SetCurLn sets the [SpellView.CurLn]:
// current line in text we're on
func (t *SpellView) SetCurLn(v int) *SpellView {
	t.CurLn = v
	return t
}

// SetCurIdx sets the [SpellView.CurIdx]:
// current index in Errs we're on
func (t *SpellView) SetCurIdx(v int) *SpellView {
	t.CurIdx = v
	return t
}

// SetUnkLex sets the [SpellView.UnkLex]:
// current unknown lex token
func (t *SpellView) SetUnkLex(v lex.Lex) *SpellView {
	t.UnkLex = v
	return t
}

// SetUnkWord sets the [SpellView.UnkWord]:
// current unknown word
func (t *SpellView) SetUnkWord(v string) *SpellView {
	t.UnkWord = v
	return t
}

// SetSuggest sets the [SpellView.Suggest]:
// a list of suggestions from spell checker
func (t *SpellView) SetSuggest(v []string) *SpellView {
	t.Suggest = v
	return t
}

// SetLastAction sets the [SpellView.LastAction]:
// last user action (ignore, change, learn)
func (t *SpellView) SetLastAction(v *gi.Button) *SpellView {
	t.LastAction = v
	return t
}

// SetTooltip sets the [SpellView.Tooltip]
func (t *SpellView) SetTooltip(v string) *SpellView {
	t.Tooltip = v
	return t
}

// SetClass sets the [SpellView.Class]
func (t *SpellView) SetClass(v string) *SpellView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [SpellView.CustomContextMenu]
func (t *SpellView) SetCustomContextMenu(v func(m *gi.Scene)) *SpellView {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [SpellView.Lay]
func (t *SpellView) SetLayout(v gi.Layouts) *SpellView {
	t.Lay = v
	return t
}

// SetStackTop sets the [SpellView.StackTop]
func (t *SpellView) SetStackTop(v int) *SpellView {
	t.StackTop = v
	return t
}

// SymbolsViewType is the [gti.Type] for [SymbolsView]
var SymbolsViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.SymbolsView",
	ShortName:  "gide.SymbolsView",
	IDName:     "symbols-view",
	Doc:        "SymbolsView is a widget that displays results of a file or package parse",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Gide", &gti.Field{Name: "Gide", Type: "goki.dev/gide/v2/gide.Gide", LocalType: "Gide", Doc: "parent gide project", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\""}},
		{"SymParams", &gti.Field{Name: "SymParams", Type: "goki.dev/gide/v2/gide.SymbolsParams", LocalType: "SymbolsParams", Doc: "params for structure display", Directives: gti.Directives{}, Tag: ""}},
		{"Syms", &gti.Field{Name: "Syms", Type: "*goki.dev/gide/v2/gide.SymNode", LocalType: "*SymNode", Doc: "all the symbols for the file or package in a tree", Directives: gti.Directives{}, Tag: ""}},
		{"Match", &gti.Field{Name: "Match", Type: "string", LocalType: "string", Doc: "only show symbols that match this string", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &SymbolsView{},
})

// NewSymbolsView adds a new [SymbolsView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewSymbolsView(par ki.Ki, name ...string) *SymbolsView {
	return par.NewChild(SymbolsViewType, name...).(*SymbolsView)
}

// KiType returns the [*gti.Type] of [SymbolsView]
func (t *SymbolsView) KiType() *gti.Type {
	return SymbolsViewType
}

// New returns a new [*SymbolsView] value
func (t *SymbolsView) New() ki.Ki {
	return &SymbolsView{}
}

// SetGide sets the [SymbolsView.Gide]:
// parent gide project
func (t *SymbolsView) SetGide(v Gide) *SymbolsView {
	t.Gide = v
	return t
}

// SetSymParams sets the [SymbolsView.SymParams]:
// params for structure display
func (t *SymbolsView) SetSymParams(v SymbolsParams) *SymbolsView {
	t.SymParams = v
	return t
}

// SetSyms sets the [SymbolsView.Syms]:
// all the symbols for the file or package in a tree
func (t *SymbolsView) SetSyms(v *SymNode) *SymbolsView {
	t.Syms = v
	return t
}

// SetMatch sets the [SymbolsView.Match]:
// only show symbols that match this string
func (t *SymbolsView) SetMatch(v string) *SymbolsView {
	t.Match = v
	return t
}

// SetTooltip sets the [SymbolsView.Tooltip]
func (t *SymbolsView) SetTooltip(v string) *SymbolsView {
	t.Tooltip = v
	return t
}

// SetClass sets the [SymbolsView.Class]
func (t *SymbolsView) SetClass(v string) *SymbolsView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [SymbolsView.CustomContextMenu]
func (t *SymbolsView) SetCustomContextMenu(v func(m *gi.Scene)) *SymbolsView {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [SymbolsView.Lay]
func (t *SymbolsView) SetLayout(v gi.Layouts) *SymbolsView {
	t.Lay = v
	return t
}

// SetStackTop sets the [SymbolsView.StackTop]
func (t *SymbolsView) SetStackTop(v int) *SymbolsView {
	t.StackTop = v
	return t
}

// SymNodeType is the [gti.Type] for [SymNode]
var SymNodeType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.SymNode",
	ShortName:  "gide.SymNode",
	IDName:     "sym-node",
	Doc:        "SymNode represents a language symbol -- the name of the node is\nthe name of the symbol. Some symbols, e.g. type have children",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Symbol", &gti.Field{Name: "Symbol", Type: "goki.dev/pi/v2/syms.Symbol", LocalType: "syms.Symbol", Doc: "the symbol", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Node", &gti.Field{Name: "Node", Type: "goki.dev/ki/v2.Node", LocalType: "ki.Node", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &SymNode{},
})

// NewSymNode adds a new [SymNode] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewSymNode(par ki.Ki, name ...string) *SymNode {
	return par.NewChild(SymNodeType, name...).(*SymNode)
}

// KiType returns the [*gti.Type] of [SymNode]
func (t *SymNode) KiType() *gti.Type {
	return SymNodeType
}

// New returns a new [*SymNode] value
func (t *SymNode) New() ki.Ki {
	return &SymNode{}
}

// SetSymbol sets the [SymNode.Symbol]:
// the symbol
func (t *SymNode) SetSymbol(v syms.Symbol) *SymNode {
	t.Symbol = v
	return t
}

// SymTreeViewType is the [gti.Type] for [SymTreeView]
var SymTreeViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.SymTreeView",
	ShortName:  "gide.SymTreeView",
	IDName:     "sym-tree-view",
	Doc:        "SymTreeView is a TreeView that knows how to operate on FileNode nodes",
	Directives: gti.Directives{},
	Fields:     ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"TreeView", &gti.Field{Name: "TreeView", Type: "goki.dev/gi/v2/giv.TreeView", LocalType: "giv.TreeView", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &SymTreeView{},
})

// NewSymTreeView adds a new [SymTreeView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewSymTreeView(par ki.Ki, name ...string) *SymTreeView {
	return par.NewChild(SymTreeViewType, name...).(*SymTreeView)
}

// KiType returns the [*gti.Type] of [SymTreeView]
func (t *SymTreeView) KiType() *gti.Type {
	return SymTreeViewType
}

// New returns a new [*SymTreeView] value
func (t *SymTreeView) New() ki.Ki {
	return &SymTreeView{}
}

// SetTooltip sets the [SymTreeView.Tooltip]
func (t *SymTreeView) SetTooltip(v string) *SymTreeView {
	t.Tooltip = v
	return t
}

// SetClass sets the [SymTreeView.Class]
func (t *SymTreeView) SetClass(v string) *SymTreeView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [SymTreeView.CustomContextMenu]
func (t *SymTreeView) SetCustomContextMenu(v func(m *gi.Scene)) *SymTreeView {
	t.CustomContextMenu = v
	return t
}

// SetIcon sets the [SymTreeView.Icon]
func (t *SymTreeView) SetIcon(v icons.Icon) *SymTreeView {
	t.Icon = v
	return t
}

// SetIndent sets the [SymTreeView.Indent]
func (t *SymTreeView) SetIndent(v units.Value) *SymTreeView {
	t.Indent = v
	return t
}

// SetOpenDepth sets the [SymTreeView.OpenDepth]
func (t *SymTreeView) SetOpenDepth(v int) *SymTreeView {
	t.OpenDepth = v
	return t
}

// SetViewIdx sets the [SymTreeView.ViewIdx]
func (t *SymTreeView) SetViewIdx(v int) *SymTreeView {
	t.ViewIdx = v
	return t
}

// SetWidgetSize sets the [SymTreeView.WidgetSize]
func (t *SymTreeView) SetWidgetSize(v mat32.Vec2) *SymTreeView {
	t.WidgetSize = v
	return t
}

// SetRootView sets the [SymTreeView.RootView]
func (t *SymTreeView) SetRootView(v *giv.TreeView) *SymTreeView {
	t.RootView = v
	return t
}

// SetSelectedNodes sets the [SymTreeView.SelectedNodes]
func (t *SymTreeView) SetSelectedNodes(v []*giv.TreeView) *SymTreeView {
	t.SelectedNodes = v
	return t
}

// TextViewType is the [gti.Type] for [TextView]
var TextViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/gide/v2/gide.TextView",
	ShortName:  "gide.TextView",
	IDName:     "text-view",
	Doc:        "TextView is the Gide-specific version of the TextView, with support for\nsetting / clearing breakpoints, etc",
	Directives: gti.Directives{},
	Fields:     ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Editor", &gti.Field{Name: "Editor", Type: "goki.dev/gi/v2/texteditor.Editor", LocalType: "texteditor.Editor", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &TextView{},
})

// NewTextView adds a new [TextView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewTextView(par ki.Ki, name ...string) *TextView {
	return par.NewChild(TextViewType, name...).(*TextView)
}

// KiType returns the [*gti.Type] of [TextView]
func (t *TextView) KiType() *gti.Type {
	return TextViewType
}

// New returns a new [*TextView] value
func (t *TextView) New() ki.Ki {
	return &TextView{}
}

// SetTooltip sets the [TextView.Tooltip]
func (t *TextView) SetTooltip(v string) *TextView {
	t.Tooltip = v
	return t
}

// SetClass sets the [TextView.Class]
func (t *TextView) SetClass(v string) *TextView {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [TextView.CustomContextMenu]
func (t *TextView) SetCustomContextMenu(v func(m *gi.Scene)) *TextView {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [TextView.Lay]
func (t *TextView) SetLayout(v gi.Layouts) *TextView {
	t.Lay = v
	return t
}

// SetStackTop sets the [TextView.StackTop]
func (t *TextView) SetStackTop(v int) *TextView {
	t.StackTop = v
	return t
}

// SetPlaceholder sets the [TextView.Placeholder]
func (t *TextView) SetPlaceholder(v string) *TextView {
	t.Placeholder = v
	return t
}

// SetCursorWidth sets the [TextView.CursorWidth]
func (t *TextView) SetCursorWidth(v units.Value) *TextView {
	t.CursorWidth = v
	return t
}

// SetLineNumberColor sets the [TextView.LineNumberColor]
func (t *TextView) SetLineNumberColor(v colors.Full) *TextView {
	t.LineNumberColor = v
	return t
}

// SetSelectColor sets the [TextView.SelectColor]
func (t *TextView) SetSelectColor(v colors.Full) *TextView {
	t.SelectColor = v
	return t
}

// SetHighlightColor sets the [TextView.HighlightColor]
func (t *TextView) SetHighlightColor(v colors.Full) *TextView {
	t.HighlightColor = v
	return t
}

// SetCursorColor sets the [TextView.CursorColor]
func (t *TextView) SetCursorColor(v colors.Full) *TextView {
	t.CursorColor = v
	return t
}
