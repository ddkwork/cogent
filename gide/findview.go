// Copyright (c) 2018, The Gide Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gide

import (
	"bytes"
	"fmt"
	"html"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/goki/gi/gi"
	"github.com/goki/gi/giv"
	"github.com/goki/gi/giv/textbuf"
	"github.com/goki/ki/ki"
	"github.com/goki/ki/kit"
	"github.com/goki/pi/filecat"
	"github.com/goki/pi/lex"
)

// FindLoc corresponds to the search scope
type FindLoc int

const (
	// FindLocAll finds in all open folders in the left file browser
	FindLocAll FindLoc = iota

	// FindLocFile only finds in the current active file
	FindLocFile

	// FindLocDir only finds in the directory of the current active file
	FindLocDir

	// FindLocNotTop finds in all open folders *except* the top-level folder
	FindLocNotTop

	// FindLocN is the number of find locations (scopes)
	FindLocN
)

//go:generate stringer -type=FindLoc

var KiT_FindLoc = kit.Enums.AddEnumAltLower(FindLocN, kit.NotBitFlag, nil, "FindLoc")

func (ev FindLoc) MarshalJSON() ([]byte, error)  { return kit.EnumMarshalJSON(ev) }
func (ev *FindLoc) UnmarshalJSON(b []byte) error { return kit.EnumUnmarshalJSON(ev, b) }

// FindParams are parameters for find / replace
type FindParams struct {
	Find       string              `desc:"find string"`
	Replace    string              `desc:"replace string"`
	IgnoreCase bool                `desc:"ignore case"`
	Regexp     bool                `desc:"use regexp regular expression search and replace"`
	Langs      []filecat.Supported `desc:"languages for files to search"`
	Loc        FindLoc             `desc:"locations to search in"`
	FindHist   []string            `desc:"history of finds"`
	ReplHist   []string            `desc:"history of replaces"`
}

// FindView is a find / replace widget that displays results in a TextView
// and has a toolbar for controlling find / replace process.
type FindView struct {
	gi.Layout
	Gide   Gide           `json:"-" xml:"-" desc:"parent gide project"`
	LangVV giv.ValueView  `desc:"langs value view"`
	Time   time.Time      `desc:"time of last find"`
	Re     *regexp.Regexp `desc:"compiled regexp"`
}

var KiT_FindView = kit.Types.AddType(&FindView{}, FindViewProps)

// Params returns the find params
func (fv *FindView) Params() *FindParams {
	return &fv.Gide.ProjPrefs().Find
}

// ShowResults shows the results in the buffer
func (fv *FindView) ShowResults(res []FileSearchResults) {
	ftv := fv.TextView()
	fbuf := ftv.Buf
	outlns := make([][]byte, 0, 100)
	outmus := make([][]byte, 0, 100) // markups
	for _, fs := range res {
		fp := fs.Node.Info.Path
		fn := fs.Node.MyRelPath()
		fbStLn := len(outlns) // find buf start ln
		lstr := fmt.Sprintf(`%v: %v`, fn, fs.Count)
		outlns = append(outlns, []byte(lstr))
		mstr := fmt.Sprintf(`<b>%v</b>`, lstr)
		outmus = append(outmus, []byte(mstr))
		for _, mt := range fs.Matches {
			txt := bytes.TrimSpace(mt.Text)
			txt = append([]byte{'\t'}, txt...)
			ln := mt.Reg.Start.Ln + 1
			ch := mt.Reg.Start.Ch + 1
			ech := mt.Reg.End.Ch + 1
			fnstr := fmt.Sprintf("%v:%d:%d", fn, ln, ch)
			nomu := bytes.Replace(txt, []byte("<mark>"), nil, -1)
			nomu = bytes.Replace(nomu, []byte("</mark>"), nil, -1)
			nomus := html.EscapeString(string(nomu))
			lstr = fmt.Sprintf(`%v: %s`, fnstr, nomus) // note: has tab embedded at start of lstr

			outlns = append(outlns, []byte(lstr))
			mstr = fmt.Sprintf(`	<a href="find:///%v#R%vN%vL%vC%v-L%vC%v">%v</a>: %s`, fp, fbStLn, fs.Count, ln, ch, ln, ech, fnstr, txt)
			outmus = append(outmus, []byte(mstr))
		}
		outlns = append(outlns, []byte(""))
		outmus = append(outmus, []byte(""))
	}
	ltxt := bytes.Join(outlns, []byte("\n"))
	mtxt := bytes.Join(outmus, []byte("\n"))
	fbuf.SetInactive(true)
	fbuf.AppendTextMarkup(ltxt, mtxt, giv.EditSignal)
	ftv.CursorStartDoc()
	ok := ftv.CursorNextLink(false) // no wrap
	if ok {
		ftv.OpenLinkAt(ftv.CursorPos)
	}
}

// SaveFindString saves the given find string to the find params history and current str
func (fv *FindView) SaveFindString(find string) {
	fv.Params().Find = find
	gi.StringsInsertFirstUnique(&fv.Params().FindHist, find, gi.Prefs.Params.SavedPathsMax)
	ftc := fv.FindText()
	if ftc != nil {
		ftc.ItemsFromStringList(fv.Params().FindHist, true, 0)
	}
}

// SaveReplString saves the given replace string to the find params history and current str
func (fv *FindView) SaveReplString(repl string) {
	fv.Params().Replace = repl
	gi.StringsInsertFirstUnique(&fv.Params().ReplHist, repl, gi.Prefs.Params.SavedPathsMax)
	rtc := fv.ReplText()
	if rtc != nil {
		rtc.ItemsFromStringList(fv.Params().ReplHist, true, 0)
	}
}

// FindAction runs a new find with current params
func (fv *FindView) FindAction() {
	fp := fv.Params()
	fv.SaveFindString(fp.Find)
	if !fv.CompileRegexp() {
		return
	}
	fv.Gide.Find(fp.Find, fp.Replace, fp.IgnoreCase, fp.Regexp, fp.Loc, fp.Langs)
}

// CheckValidRegexp returns false if using regexp and it is not valid
func (fv *FindView) CheckValidRegexp() bool {
	fp := fv.Params()
	if !fp.Regexp {
		return true
	}
	if fv.Re == nil {
		return false
	}
	return true
}

// ReplaceAction performs the replace -- if using regexp mode, regexp must be compiled in advance
func (fv *FindView) ReplaceAction() bool {
	if !fv.CheckValidRegexp() {
		return false
	}
	wupdt := fv.TopUpdateStart()
	defer fv.TopUpdateEnd(wupdt)

	fp := fv.Params()
	fv.SaveReplString(fp.Replace)
	gi.StringsInsertFirstUnique(&fp.ReplHist, fp.Replace, gi.Prefs.Params.SavedPathsMax)

	ftv := fv.TextView()
	tl, ok := ftv.OpenLinkAt(ftv.CursorPos)
	if !ok {
		ok = ftv.CursorNextLink(false) // no wrap
		if !ok {
			return false
		}
		tl, ok = ftv.OpenLinkAt(ftv.CursorPos)
		if !ok {
			return false
		}
	}
	ge := fv.Gide
	tv, reg, _, _, ok := ge.ParseOpenFindURL(tl.URL, ftv)
	if !ok {
		return false
	}
	if reg.IsNil() {
		ok = ftv.CursorNextLink(false) // no wrap
		if !ok {
			return false
		}
		tl, ok = ftv.OpenLinkAt(ftv.CursorPos)
		if !ok {
			return false
		}
		tv, reg, _, _, ok = ge.ParseOpenFindURL(tl.URL, ftv)
		if !ok || reg.IsNil() {
			return false
		}
	}
	reg.Time.SetTime(fv.Time)
	reg = tv.Buf.AdjustReg(reg)
	if !reg.IsNil() {
		if fp.Regexp {
			rg := tv.Buf.Region(reg.Start, reg.End)
			b := rg.ToBytes()
			rb := fv.Re.ReplaceAll(b, []byte(fp.Replace))
			tv.Buf.ReplaceText(reg.Start, reg.End, reg.Start, string(rb), giv.EditSignal, false)
		} else {
			// MatchCase only if doing IgnoreCase
			tv.Buf.ReplaceText(reg.Start, reg.End, reg.Start, fp.Replace, giv.EditSignal, fp.IgnoreCase)
		}

		// delete the link for the just done replace
		ftvln := ftv.CursorPos.Ln
		st := lex.Pos{Ln: ftvln, Ch: 0}
		len := len(ftv.Buf.Lines[ftvln])
		en := lex.Pos{Ln: ftvln, Ch: len}
		ftv.Buf.DeleteText(st, en, giv.EditSignal)
	}

	tv.ClearHighlights()

	ok = ftv.CursorNextLink(false) // no wrap
	if ok {
		ftv.OpenLinkAt(ftv.CursorPos) // move to next
	}
	return ok
}

// ReplaceAllAction performs replace all, prompting before proceeding
func (fv *FindView) ReplaceAllAction() {
	gi.PromptDialog(nil, gi.DlgOpts{Title: "Confirm Replace All", Prompt: "Are you sure you want to Replace All?"}, gi.AddOk, gi.AddCancel, fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.DialogAccepted) {
			fv.ReplaceAll()
		}
	})
}

// CompileRegexp compiles the regexp if necessary -- returns false if it is invalid
func (fv *FindView) CompileRegexp() bool {
	fp := fv.Params()
	if !fp.Regexp {
		fv.Re = nil
		return true
	}
	var err error
	fv.Re, err = regexp.Compile(fp.Find)
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Regexp is Invalid", Prompt: fmt.Sprintf("The regular expression was invalid: %v", err)}, gi.AddOk, gi.NoCancel, nil, nil)
		return false
	}
	return true
}

// ReplaceAll performs replace all
func (fv *FindView) ReplaceAll() {
	if !fv.CheckValidRegexp() {
		return
	}
	for {
		ok := fv.ReplaceAction()
		if !ok {
			break
		}
	}
}

// NextFind shows next find result
func (fv *FindView) NextFind() {
	ftv := fv.TextView()
	ok := ftv.CursorNextLink(true) // wrap
	if ok {
		ftv.OpenLinkAt(ftv.CursorPos)
	}
}

// PrevFind shows previous find result
func (fv *FindView) PrevFind() {
	ftv := fv.TextView()
	ok := ftv.CursorPrevLink(true) // wrap
	if ok {
		ftv.OpenLinkAt(ftv.CursorPos)
	}
}

// OpenFindURL opens given find:/// url from Find
func (fv *FindView) OpenFindURL(ur string, ftv *giv.TextView) bool {
	ge := fv.Gide
	tv, reg, fbBufStLn, fCount, ok := ge.ParseOpenFindURL(ur, ftv)
	if !ok {
		return false
	}
	reg.Time.SetTime(fv.Time)
	reg = tv.Buf.AdjustReg(reg)
	find := fv.Params().Find
	giv.PrevISearchString = find
	tve := tv.Embed(giv.KiT_TextView).(*giv.TextView)
	fv.HighlightFinds(tve, ftv, fbBufStLn, fCount, find)
	tv.SetNeedsRefresh()
	tv.RefreshIfNeeded()
	tv.SetCursorShow(reg.Start)
	return true
}

// HighlightFinds highlights all the find results in ftv buffer
func (fv *FindView) HighlightFinds(tv, ftv *giv.TextView, fbStLn, fCount int, find string) {
	lnka := []byte(`<a href="`)
	lnkasz := len(lnka)

	fb := ftv.Buf

	if len(tv.Highlights) != fCount { // highlight
		hi := make([]textbuf.Region, fCount)
		for i := 0; i < fCount; i++ {
			fln := fbStLn + 1 + i
			if fln >= len(fb.Markup) {
				continue
			}
			ltxt := fb.Markup[fln]
			fpi := bytes.Index(ltxt, lnka)
			if fpi < 0 {
				continue
			}
			fpi += lnkasz
			epi := fpi + bytes.Index(ltxt[fpi:], []byte(`"`))
			lnk := string(ltxt[fpi:epi])
			iup, err := url.Parse(lnk)
			if err != nil {
				continue
			}
			ireg := textbuf.Region{}
			lidx := strings.Index(iup.Fragment, "L")
			ireg.FromString(iup.Fragment[lidx:])
			ireg.Time.SetTime(fv.Time)
			hi[i] = ireg
		}
		tv.Highlights = hi
	}
}

//////////////////////////////////////////////////////////////////////////////////////
//    GUI config

// Config configures the view
func (fv *FindView) Config(ge Gide) {
	fv.Gide = ge
	fv.Lay = gi.LayoutVert
	fv.SetProp("spacing", gi.StdDialogVSpaceUnits)
	config := kit.TypeAndNameList{}
	config.Add(gi.KiT_ToolBar, "findbar")
	config.Add(gi.KiT_ToolBar, "replbar")
	config.Add(gi.KiT_Layout, "findtext")
	mods, updt := fv.ConfigChildren(config)
	if !mods {
		updt = fv.UpdateStart()
	}
	fp := fv.Params()
	fv.ConfigToolbar()
	ft := fv.FindText()
	ft.ItemsFromStringList(fp.FindHist, true, 0)
	ft.SetText(fp.Find)
	rt := fv.ReplText()
	rt.ItemsFromStringList(fp.ReplHist, true, 0)
	rt.SetText(fp.Replace)
	ib := fv.IgnoreBox()
	ib.SetChecked(fp.IgnoreCase)
	rb := fv.RegexpBox()
	rb.SetChecked(fp.Regexp)
	cf := fv.LocCombo()
	cf.SetCurIndex(int(fp.Loc))
	tvly := fv.TextViewLay()
	ConfigOutputTextView(tvly)
	if mods {
		na := fv.FindNextAct()
		na.GrabFocus()
	}
	fv.UpdateEnd(updt)
}

// FindBar returns the find toolbar
func (fv *FindView) FindBar() *gi.ToolBar {
	return fv.ChildByName("findbar", 0).(*gi.ToolBar)
}

// ReplBar returns the replace toolbar
func (fv *FindView) ReplBar() *gi.ToolBar {
	return fv.ChildByName("replbar", 1).(*gi.ToolBar)
}

// FindText returns the find textfield in toolbar
func (fv *FindView) FindText() *gi.ComboBox {
	return fv.FindBar().ChildByName("find-str", 1).(*gi.ComboBox)
}

// ReplText returns the replace textfield in toolbar
func (fv *FindView) ReplText() *gi.ComboBox {
	return fv.ReplBar().ChildByName("repl-str", 1).(*gi.ComboBox)
}

// IgnoreBox returns the ignore case checkbox in toolbar
func (fv *FindView) IgnoreBox() *gi.CheckBox {
	return fv.FindBar().ChildByName("ignore-case", 2).(*gi.CheckBox)
}

// RegexpBox returns the regexp checkbox in toolbar
func (fv *FindView) RegexpBox() *gi.CheckBox {
	return fv.FindBar().ChildByName("regexp", 3).(*gi.CheckBox)
}

// LocCombo returns the loc combobox
func (fv *FindView) LocCombo() *gi.ComboBox {
	return fv.FindBar().ChildByName("loc", 5).(*gi.ComboBox)
}

// FindNextAct returns the find next action in toolbar -- selected first
func (fv *FindView) FindNextAct() *gi.Action {
	return fv.ReplBar().ChildByName("next", 3).(*gi.Action)
}

// TextViewLay returns the find results TextView layout
func (fv *FindView) TextViewLay() *gi.Layout {
	return fv.ChildByName("findtext", 1).(*gi.Layout)
}

// TextView returns the find results TextView
func (fv *FindView) TextView() *giv.TextView {
	return fv.TextViewLay().Child(0).Embed(giv.KiT_TextView).(*giv.TextView)
}

// ConfigToolbar adds toolbar.
func (fv *FindView) ConfigToolbar() {
	fb := fv.FindBar()
	if fb.HasChildren() {
		return
	}
	fb.SetStretchMaxWidth()

	rb := fv.ReplBar()
	rb.SetStretchMaxWidth()

	fb.AddAction(gi.ActOpts{Label: "Find:", Tooltip: "Find given string in project files. Only open folders in file browser will be searched -- adjust those to scope the search"},
		fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
			fvv, _ := recv.Embed(KiT_FindView).(*FindView)
			fvv.FindAction()
		})

	finds := fb.AddNewChild(gi.KiT_ComboBox, "find-str").(*gi.ComboBox)
	finds.Editable = true
	finds.SetStretchMaxWidth()
	finds.Tooltip = "String to find -- hit enter or tab to update search -- click for history"
	finds.ConfigParts()
	finds.ItemsFromStringList(fv.Params().FindHist, true, 0)
	ftf, _ := finds.TextField()
	finds.ComboSig.Connect(fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		fvv, _ := recv.Embed(KiT_FindView).(*FindView)
		cb := send.(*gi.ComboBox)
		fvv.Params().Find = cb.CurVal.(string)
		fvv.FindAction()
	})
	ftf.TextFieldSig.Connect(fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.TextFieldDone) || sig == int64(gi.TextFieldDeFocused) {
			fvv, _ := recv.Embed(KiT_FindView).(*FindView)
			tf := send.(*gi.TextField)
			fvv.Params().Find = tf.Text()
			fvv.FindAction()
		} else if sig == int64(gi.TextFieldCleared) {
			tv := fv.Gide.ActiveTextView()
			if tv != nil {
				tv.ClearHighlights()
			}
			fvtv := fv.TextView()
			if fvtv != nil {
				fvtv.Buf.New(0)
			}
		}
	})

	ic := fb.AddNewChild(gi.KiT_CheckBox, "ignore-case").(*gi.CheckBox)
	ic.SetText("Ignore Case")
	ic.ButtonSig.Connect(fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonToggled) {
			fvv, _ := recv.Embed(KiT_FindView).(*FindView)
			cb := send.(*gi.CheckBox)
			fvv.Params().IgnoreCase = cb.IsChecked()
		}
	})

	rx := fb.AddNewChild(gi.KiT_CheckBox, "regexp").(*gi.CheckBox)
	rx.SetText("Regexp")
	rx.Tooltip = "use regular expression for search and replace -- see https://github.com/google/re2/wiki/Syntax"
	rx.ButtonSig.Connect(fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonToggled) {
			fvv, _ := recv.Embed(KiT_FindView).(*FindView)
			cb := send.(*gi.CheckBox)
			fvv.Params().Regexp = cb.IsChecked()
		}
	})

	locl := fb.AddNewChild(gi.KiT_Label, "loc-lbl").(*gi.Label)
	locl.SetText("Loc:")
	locl.Tooltip = "location to find in: all = all open folders in browser; file = current active file; dir = directory of current active file; nottop = all except the top-level in browser"
	// locl.SetProp("vertical-align", gi.AlignMiddle)

	cf := fb.AddNewChild(gi.KiT_ComboBox, "loc").(*gi.ComboBox)
	cf.SetText("Loc")
	cf.Tooltip = locl.Tooltip
	cf.ItemsFromEnum(KiT_FindLoc, false, 0)
	cf.ComboSig.Connect(fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		fvv, _ := recv.Embed(KiT_FindView).(*FindView)
		cb := send.(*gi.ComboBox)
		eval := cb.CurVal.(kit.EnumValue)
		fvv.Params().Loc = FindLoc(eval.Value)
	})

	//////////////// ReplBar

	rb.AddAction(gi.ActOpts{Name: "prev", Icon: "wedge-up", Tooltip: "go to previous result"},
		fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
			fvv, _ := recv.Embed(KiT_FindView).(*FindView)
			fvv.PrevFind()
		})

	rb.AddAction(gi.ActOpts{Name: "next", Icon: "wedge-down", Tooltip: "go to next result"},
		fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
			fvv, _ := recv.Embed(KiT_FindView).(*FindView)
			fvv.NextFind()
		})

	rb.AddAction(gi.ActOpts{Label: "Replace:", Tooltip: "Replace find string with replace string for currently-selected find result"}, fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		fvv, _ := recv.Embed(KiT_FindView).(*FindView)
		fvv.CompileRegexp()
		fvv.ReplaceAction()
	})

	repls := rb.AddNewChild(gi.KiT_ComboBox, "repl-str").(*gi.ComboBox)
	repls.Editable = true
	repls.SetStretchMaxWidth()
	repls.Tooltip = "String to replace find string -- click for history -- use ${n} for regexp submatch where n = 1 for first submatch, etc"
	repls.ConfigParts()
	repls.ItemsFromStringList(fv.Params().ReplHist, true, 0)
	rtf, _ := repls.TextField()
	repls.ComboSig.Connect(fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		fvv, _ := recv.Embed(KiT_FindView).(*FindView)
		cb := send.(*gi.ComboBox)
		fvv.Params().Replace = cb.CurVal.(string)
	})
	rtf.TextFieldSig.Connect(fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.TextFieldDone) {
			fvv, _ := recv.Embed(KiT_FindView).(*FindView)
			tf := send.(*gi.TextField)
			fvv.Params().Replace = tf.Text()
		}
	})
	rb.AddAction(gi.ActOpts{Label: "All", Tooltip: "replace all find strings with replace string"},
		fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
			fvv, _ := recv.Embed(KiT_FindView).(*FindView)
			fvv.ReplaceAllAction()
		})

	langl := rb.AddNewChild(gi.KiT_Label, "lang-lbl").(*gi.Label)
	langl.SetText("Lang:")
	langl.Tooltip = "Language(s) to restrict search / replace to"
	// langl.SetProp("vertical-align", gi.AlignMiddle)

	fv.LangVV = giv.ToValueView(&fv.Params().Langs, "")
	fv.LangVV.SetSoloValue(reflect.ValueOf(&fv.Params().Langs))
	vtyp := fv.LangVV.WidgetType()

	langw := rb.AddNewChild(vtyp, "langs").(gi.Node2D)
	fv.LangVV.ConfigWidget(langw)
	langw.AsWidget().Tooltip = langl.Tooltip
	//	vvb := vv.AsValueViewBase()
	//	vvb.ViewSig.ConnectOnly(fv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
	//		fvv, _ := recv.Embed(KiT_FindView).(*FindView)
	// hmm, langs updated..
	//	})

}

// FindViewProps are style properties for FindView
var FindViewProps = ki.Props{
	"EnumType:Flag":    gi.KiT_NodeFlags,
	"background-color": &gi.Prefs.Colors.Background,
	"color":            &gi.Prefs.Colors.Font,
	"max-width":        -1,
	"max-height":       -1,
}
