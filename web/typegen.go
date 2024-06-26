// Code generated by "core generate -add-types"; DO NOT EDIT.

package web

import (
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/cogent/web.Page", IDName: "page", Doc: "Page represents one web browser page.", Embeds: []types.Field{{Name: "Frame"}}, Fields: []types.Field{{Name: "History", Doc: "The history of URLs that have been visited. The oldest page is first."}, {Name: "Context", Doc: "Context is the page's [htmlcore.Context]."}}})

// NewPage returns a new [Page] with the given optional parent:
// Page represents one web browser page.
func NewPage(parent ...tree.Node) *Page { return tree.New[Page](parent...) }
