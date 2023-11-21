// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gidom

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"goki.dev/colors"
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/giv"
	"goki.dev/gi/v2/texteditor"
	"goki.dev/girl/paint"
	"goki.dev/girl/styles"
	"goki.dev/grows/images"
	"goki.dev/grr"
	"goki.dev/icons"
	"goki.dev/ki/v2"
	"golang.org/x/net/html"
)

// Handler is a function that can be used to describe the behavior
// of gidom parsing for a specific type of element. It takes the
// parent [gi.Widget] to add widgets to and the [*html.Node] to read from.
// It returns the widget, if any, that has been constructed for this node.
// If it returns false, that indicates that the children of this node have
// already been handled (or will not be handled), and thus should not
// be handled further. If it returns a true, then the children
// will be handled, with the returned widget as their parent.
type Handler func(ctx Context)

// ElementHandlers is a map of [Handler] functions for each HTML element
// type (eg: "button", "input", "p"). It is empty by default, but can be
// used by anyone in need of behavior different than the default behavior
// defined in [HandleElement].
var ElementHandlers = map[string]Handler{}

// New adds a new widget of the given type to the context parent.
func New[T gi.Widget](ctx Context) T {
	rules := ctx.Style()
	display := ""
	for _, rule := range rules {
		for _, decl := range rule.Declarations {
			if decl.Property == "display" {
				display = decl.Value
			}
		}
	}
	var par gi.Widget
	switch display {
	case "inline", "inline-block", "":
		par = ctx.InlineParent()
	default:
		par = ctx.BlockParent()
		ctx.SetInlineParent(nil)
	}
	w := ki.New[T](par)
	wb := w.AsWidget()

	for _, attr := range ctx.Node().Attr {
		switch attr.Key {
		case "id":
			wb.SetName(attr.Val)
		case "class":
			wb.SetClass(attr.Val)
		default:
			wb.SetProp(attr.Key, attr.Val)
		}
	}
	wb.SetProp("tag", ctx.Node().Data)
	w.Style(func(s *styles.Style) {
		for _, rule := range rules {
			for _, decl := range rule.Declarations {
				// TODO(kai/styprops): parent style and context
				s.StyleFromProp(s, decl.Property, decl.Value, colors.BaseContext(s.Color))
			}
		}
	})
	return w
}

// HandleELement calls the [Handler] associated with the given element [*html.Node]
// in [ElementHandlers] and returns the result, using the given context. If there
// is no handler associated with it, it uses default hardcoded configuration code.
func HandleElement(ctx Context) {
	tag := ctx.Node().Data
	h, ok := ElementHandlers[tag]
	if ok {
		h(ctx)
		return
	}

	if slices.Contains(TextTags, tag) {
		HandleLabelTag(ctx)
		return
	}

	switch tag {
	case "script", "title", "meta":
		// we don't render anything
	case "link":
		rel := GetAttr(ctx.Node(), "rel")
		// TODO(kai/gidom): maybe handle preload
		if rel == "preload" {
			return
		}
		// TODO(kai/gidom): support links other than stylesheets
		if rel != "stylesheet" {
			return
		}
		resp, err := Get(ctx, GetAttr(ctx.Node(), "href"))
		if grr.Log0(err) != nil {
			return
		}
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if grr.Log0(err) != nil {
			return
		}
		ctx.AddStyle(string(b))
	case "style":
		ctx.AddStyle(ExtractText(ctx))
	case "body", "main", "div", "section", "nav", "footer", "header":
		f := New[*gi.Frame](ctx)
		f.Style(func(s *styles.Style) {
			s.Direction = styles.Column
		})
		ctx.SetNewParent(f)
	case "button":
		New[*gi.Button](ctx).SetText(ExtractText(ctx))
	case "h1":
		HandleLabel(ctx).SetType(gi.LabelHeadlineLarge)
	case "h2":
		HandleLabel(ctx).SetType(gi.LabelHeadlineSmall)
	case "h3":
		HandleLabel(ctx).SetType(gi.LabelTitleLarge)
	case "h4":
		HandleLabel(ctx).SetType(gi.LabelTitleMedium)
	case "h5":
		HandleLabel(ctx).SetType(gi.LabelTitleSmall)
	case "h6":
		HandleLabel(ctx).SetType(gi.LabelLabelSmall)
	case "p":
		HandleLabel(ctx)
	case "pre":
		HandleLabel(ctx).Style(func(s *styles.Style) {
			s.Text.WhiteSpace = styles.WhiteSpacePre
		})
	case "ol", "ul":
		// if we are already in a treeview, we just return in the last item in it
		// (which is the list item we are contained in), which fixes the associativity
		// of nested list items and prevents the created of duplicated tree view items.
		if ptv, ok := ctx.Parent().(*giv.TreeView); ok {
			w := ki.LastChild(ptv).(gi.Widget)
			ctx.SetNewParent(w)
			return
		}
		tv := New[*giv.TreeView](ctx).SetText("").SetIcon(icons.None)
		tv.RootView = tv
		ctx.SetNewParent(tv)
		return
	case "li":
		ntv := New[*giv.TreeView](ctx)
		ftxt := ""
		ptv, ok := ctx.Parent().(*giv.TreeView)
		if ok {
			ntv.RootView = ptv.RootView
			if ptv.Prop("tag") == "ol" {
				ip, _ := ntv.IndexInParent()
				ftxt = strconv.Itoa(ip+1) + ". " // start at 1
			} else {
				// TODO(kai/gidom): have different bullets for different depths
				ftxt = "• "
			}
		} else {
			ntv.RootView = ntv
		}

		etxt := ExtractText(ctx)
		ntv.SetName(etxt)
		ntv.SetText(ftxt + etxt)
		ntv.OnWidgetAdded(func(w gi.Widget) {
			switch w := w.(type) {
			case *gi.Label:
				w.HandleLabelClick(func(tl *paint.TextLink) {
					grr.Log0(ctx.OpenURL(tl.URL))
				})
			}
		})
	case "img":
		img := New[*gi.Image](ctx)
		n := ctx.Node()
		go func() {
			src := GetAttr(n, "src")
			resp, err := Get(ctx, src)
			if grr.Log0(err) != nil {
				return
			}
			defer resp.Body.Close()
			if strings.Contains(resp.Header.Get("Content-Type"), "svg") {
				// TODO(kai/gidom): support svg
			} else {
				im, _, err := images.Read(resp.Body)
				if err != nil {
					slog.Error("error loading image", "url", src, "err", err)
					return
				}
				img.Filename = gi.FileName(src)
				img.SetImage(im, 0, 0)
				img.Update()
			}
		}()
	case "input":
		ityp := GetAttr(ctx.Node(), "type")
		switch ityp {
		case "number":
			New[*gi.Spinner](ctx)
		case "color":
			// TODO(kai/gidom): handle giv values with New structure correctly
			giv.NewValue(ctx.Parent(), colors.Black).AsWidget()
		case "datetime":
			giv.NewValue(ctx.Parent(), time.Now()).AsWidget()
		default:
			New[*gi.TextField](ctx)
		}
	case "textarea":
		buf := texteditor.NewBuf()
		buf.SetText([]byte(ExtractText(ctx)))
		New[*texteditor.Editor](ctx).SetBuf(buf)
	default:
		ctx.SetNewParent(ctx.Parent())
	}
}

// ConfigWidget sets the properties of the given widget based on the properties
// of the given node. It should be called on all widgets in [HandleElement] and
// [Handler] functions.
func ConfigWidget[T gi.Widget](ctx Context, w T, n *html.Node) T {
	wb := w.AsWidget()
	// if we already have the tag prop, we have already been configured
	if _, err := wb.PropTry("tag"); err == nil {
		return w
	}
	for _, attr := range n.Attr {
		switch attr.Key {
		case "id":
			wb.SetName(attr.Val)
		case "class":
			wb.SetClass(attr.Val)
		default:
			wb.SetProp(attr.Key, attr.Val)
		}
	}
	wb.SetProp("tag", n.Data)
	ctx.SetWidgetForNode(w, n)
	rules := ctx.Style()
	w.Style(func(s *styles.Style) {
		for _, rule := range rules {
			for _, decl := range rule.Declarations {
				// TODO(kai/styprops): parent style and context
				s.StyleFromProp(s, decl.Property, decl.Value, colors.BaseContext(s.Color))
			}
		}
	})
	return w
}

// HandleLabel creates a new label from the given information, setting the text and
// the label click function so that URLs are opened according to [Context.OpenURL].
func HandleLabel(ctx Context) *gi.Label {
	lb := New[*gi.Label](ctx).SetText(ExtractText(ctx))
	lb.HandleLabelClick(func(tl *paint.TextLink) {
		grr.Log0(ctx.OpenURL(tl.URL))
	})
	return lb
}

// HandleLabelTag creates a new label from the given information, setting the text and
// the label click function so that URLs are opened according to [Context.OpenURL]. Also,
// it wraps the label text with the [NodeString] of the given node, meaning that it
// should be used for standalone elements that are meant to only exist in labels
// (eg: a, span, b, code, etc).
func HandleLabelTag(ctx Context) *gi.Label {
	start, end := NodeString(ctx.Node())
	str := start + ExtractText(ctx) + end
	lb := New[*gi.Label](ctx).SetText(str)
	lb.HandleLabelClick(func(tl *paint.TextLink) {
		grr.Log0(ctx.OpenURL(tl.URL))
	})
	return lb
}

// GetAttr gets the given attribute from the given node, returning ""
// if the attribute is not found.
func GetAttr(n *html.Node, attr string) string {
	res := ""
	for _, a := range n.Attr {
		if a.Key == attr {
			res = a.Val
		}
	}
	return res
}

// HasAttr returns whether the given node has the given attribute defined.
func HasAttr(n *html.Node, attr string) bool {
	return slices.ContainsFunc(n.Attr, func(a html.Attribute) bool {
		return a.Key == attr
	})
}

// Get is a helper function that calls [http.Get] with the given URL, parsed
// relative to the page URL of the given context. It also checks the status
// code of the response and closes the response body and returns an error if
// it is not [http.StatusOK]. If the error is nil, then the response body is
// not closed and must be closed by the caller.
func Get(ctx Context, url string) (*http.Response, error) {
	u, err := ParseRelativeURL(url, ctx.PageURL())
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return resp, fmt.Errorf("got error status %q (code %d)", resp.Status, resp.StatusCode)
	}
	return resp, nil
}
