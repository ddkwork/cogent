// Code generated by "goki generate"; DO NOT EDIT.

package glide

import (
	"goki.dev/gi/v2/gi"
	"goki.dev/gti"
	"goki.dev/ki/v2"
	"goki.dev/ordmap"
)

// PageType is the [gti.Type] for [Page]
var PageType = gti.AddType(&gti.Type{
	Name:       "goki.dev/glide/glide.Page",
	ShortName:  "glide.Page",
	IDName:     "page",
	Doc:        "Page represents one web browser page",
	Directives: gti.Directives{},
	Fields:     ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Frame", &gti.Field{Name: "Frame", Type: "goki.dev/gi/v2/gi.Frame", LocalType: "gi.Frame", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &Page{},
})

// NewPage adds a new [Page] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewPage(par ki.Ki, name ...string) *Page {
	return par.NewChild(PageType, name...).(*Page)
}

// KiType returns the [*gti.Type] of [Page]
func (t *Page) KiType() *gti.Type {
	return PageType
}

// New returns a new [*Page] value
func (t *Page) New() ki.Ki {
	return &Page{}
}

// SetTooltip sets the [Page.Tooltip]
func (t *Page) SetTooltip(v string) *Page {
	t.Tooltip = v
	return t
}

// SetClass sets the [Page.Class]
func (t *Page) SetClass(v string) *Page {
	t.Class = v
	return t
}

// SetCustomContextMenu sets the [Page.CustomContextMenu]
func (t *Page) SetCustomContextMenu(v func(m *gi.Scene)) *Page {
	t.CustomContextMenu = v
	return t
}

// SetLayout sets the [Page.Lay]
func (t *Page) SetLayout(v gi.Layouts) *Page {
	t.Lay = v
	return t
}

// SetStackTop sets the [Page.StackTop]
func (t *Page) SetStackTop(v int) *Page {
	t.StackTop = v
	return t
}

// SetStripes sets the [Page.Stripes]
func (t *Page) SetStripes(v gi.Stripes) *Page {
	t.Stripes = v
	return t
}
