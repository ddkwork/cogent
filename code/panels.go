// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package code

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles/states"
	"cogentcore.org/core/texteditor"
)

// PanelIsOpen returns true if the given panel has not been collapsed and is avail
// and visible for displaying something
func (cv *Code) PanelIsOpen(panel int) bool {
	sv := cv.Splits()
	if panel < 0 || panel >= len(sv.Children) {
		return false
	}
	if sv.Splits[panel] <= 0.01 {
		return false
	}
	return true
}

// CurPanel returns the splitter panel that currently has keyboard focus
func (cv *Code) CurPanel() int {
	sv := cv.Splits()
	for i, ski := range sv.Children {
		_, sk := core.AsWidget(ski)
		if sk.HasStateWithin(states.Focused) {
			return i
		}
	}
	return -1 // nobody
}

// FocusOnPanel moves keyboard focus to given panel -- returns false if nothing at that tab
func (cv *Code) FocusOnPanel(panel int) bool {
	sv := cv.Splits()
	switch panel {
	case TextEditor1Index:
		cv.SetActiveTextEditorIndex(0)
	case TextEditor2Index:
		cv.SetActiveTextEditorIndex(1)
	case TabsIndex:
		tv := cv.Tabs()
		ct, _ := tv.CurrentTab()
		if ct != nil {
			cv.Scene.Events.FocusNextFrom(ct)
		} else {
			return false
		}
	default:
		ski, _ := core.AsWidget(sv.Children[panel])
		cv.Scene.Events.FocusNextFrom(ski)
	}
	cv.NeedsRender()
	return true
}

// FocusNextPanel moves the keyboard focus to the next panel to the right
func (cv *Code) FocusNextPanel() { //types:add
	sv := cv.Splits()
	cp := cv.CurPanel()
	cp++
	np := len(sv.Children)
	if cp >= np {
		cp = 0
	}
	for sv.Splits[cp] <= 0.01 {
		cp++
		if cp >= np {
			cp = 0
		}
	}
	cv.FocusOnPanel(cp)
}

// FocusPrevPanel moves the keyboard focus to the previous panel to the left
func (cv *Code) FocusPrevPanel() { //types:add
	sv := cv.Splits()
	cp := cv.CurPanel()
	cp--
	np := len(sv.Children)
	if cp < 0 {
		cp = np - 1
	}
	for sv.Splits[cp] <= 0.01 {
		cp--
		if cp < 0 {
			cp = np - 1
		}
	}
	cv.FocusOnPanel(cp)
}

// TabByName returns a tab with given name, nil if not found.
func (cv *Code) TabByName(name string) core.Widget {
	tv := cv.Tabs()
	return tv.TabByName(name)
}

// SelectTabByName Selects given main tab, and returns all of its contents as well.
func (cv *Code) SelectTabByName(name string) core.Widget {
	tv := cv.Tabs()
	if tv == nil {
		return nil
	}
	return tv.SelectTabByName(name)
}

// RecycleTabTextEditor returns a text editor in a tab with the given name,
// first by looking for an existing one, and if not found, making a new
// one with a text editor in it.
func (cv *Code) RecycleTabTextEditor(name string, buf *texteditor.Buffer) *texteditor.Editor {
	tv := cv.Tabs()
	if tv == nil {
		return nil
	}

	fr := tv.RecycleTab(name)
	if fr.HasChildren() {
		return fr.Child(0).(*texteditor.Editor)
	}
	txv := texteditor.NewEditor(fr)
	txv.SetName(fr.Name)
	if buf != nil {
		txv.SetBuffer(buf)
	}
	ConfigOutputTextEditor(txv)
	tv.Update()
	return txv
}
