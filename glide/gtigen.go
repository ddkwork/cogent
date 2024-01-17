// Code generated by "core generate -add-types"; DO NOT EDIT.

package glide

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/gti"
	"cogentcore.org/core/ki"
)

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/cogent/glide.Context", IDName: "context", Doc: "Context implements [gidom.Context]", Embeds: []gti.Field{{Name: "ContextBase"}}, Fields: []gti.Field{{Name: "Page", Doc: "Page is the page associated with the context"}}})

// PageType is the [gti.Type] for [Page]
var PageType = gti.AddType(&gti.Type{Name: "cogentcore.org/cogent/glide.Page", IDName: "page", Doc: "Page represents one web browser page", Embeds: []gti.Field{{Name: "Frame"}}, Fields: []gti.Field{{Name: "History", Doc: "The history of URLs that have been visited. The oldest page is first."}, {Name: "PageURL", Doc: "PageURL is the current page URL"}}, Instance: &Page{}})

// NewPage adds a new [Page] with the given name to the given parent:
// Page represents one web browser page
func NewPage(par ki.Ki, name ...string) *Page {
	return par.NewChild(PageType, name...).(*Page)
}

// KiType returns the [*gti.Type] of [Page]
func (t *Page) KiType() *gti.Type { return PageType }

// New returns a new [*Page] value
func (t *Page) New() ki.Ki { return &Page{} }

// SetHistory sets the [Page.History]:
// The history of URLs that have been visited. The oldest page is first.
func (t *Page) SetHistory(v ...string) *Page { t.History = v; return t }

// SetPageUrl sets the [Page.PageURL]:
// PageURL is the current page URL
func (t *Page) SetPageUrl(v string) *Page { t.PageURL = v; return t }

// SetTooltip sets the [Page.Tooltip]
func (t *Page) SetTooltip(v string) *Page { t.Tooltip = v; return t }

// SetStackTop sets the [Page.StackTop]
func (t *Page) SetStackTop(v int) *Page { t.StackTop = v; return t }

// SetStripes sets the [Page.Stripes]
func (t *Page) SetStripes(v gi.Stripes) *Page { t.Stripes = v; return t }
