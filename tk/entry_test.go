// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("Entry", testEntry)
}

func testEntry(t *testing.T) {
	w := NewEntry(nil, EntryAttrForeground("blue"), EntryAttrBackground("blue"), EntryAttrWidth(20), EntryAttrJustify(1), EntryAttrShow("text"), EntryAttrState(StateNormal), EntryAttrTakeFocus(true), EntryAttrExportSelection(true))
	defer w.Destroy()

	w.SetForeground("blue")
	if v := w.Foreground(); v != "blue" {
		t.Fatal("Foreground", "blue", v)
	}

	w.SetBackground("blue")
	if v := w.Background(); v != "blue" {
		t.Fatal("Background", "blue", v)
	}

	w.SetWidth(20)
	if v := w.Width(); v != 20 {
		t.Fatal("Width", 20, v)
	}

	w.SetJustify(1)
	if v := w.Justify(); v != 1 {
		t.Fatal("Justify", 1, v)
	}

	w.SetShow("text")
	if v := w.Show(); v != "text" {
		t.Fatal("Show", "text", v)
	}

	w.SetState(StateNormal)
	if v := w.State(); v != StateNormal {
		t.Fatal("State", StateNormal, v)
	}

	w.SetTakeFocus(true)
	if v := w.IsTakeFocus(); v != true {
		t.Fatal("IsTakeFocus", true, v)
	}

	w.SetExportSelection(true)
	if v := w.IsExportSelection(); v != true {
		t.Fatal("IsExportSelection", true, v)
	}

	w.SetText("text")
	if v := w.Text(); v != "text" {
		t.Fatal("Text", "text", v)
	}

	w.SetState(StateNormal)
	w.SetShow("")

	w.SetText("abc中文")
	if w.TextLength() != 5 {
		t.Fatal("TextLength", w.TextLength())
	}
	if w.SetCursorPosition(1).CursorPosition() != 1 {
		t.Fatal("CursorPostion")
	}
	if w.HasSelectedText() {
		t.Fatal("HasSelectedText")
	}
	w.SelectAll()
	if !w.HasSelectedText() {
		t.Fatal("HasSelectedText")
	}
	if w.SelectedText() != w.Text() {
		t.Fatal("SelectAll", w.SelectionStart(), w.SelectionEnd())
	}
	w.ClearSelection()
	if w.HasSelectedText() {
		t.Fatal("ClearSelection")
	}
	w.SetSelection(2, 4)
	if w.SelectionStart() != 2 || w.SelectionEnd() != 4 {
		t.Fatal("SetSelection")
	}
	Update()
	if w.SelectedText() != "c中" {
		t.Fatal("SelectedText", w.SelectedText())
	}
	w.DeleteRange(2, 4)
	if w.HasSelectedText() {
		t.Fatal("DeleteRange", w.Text(), w.SelectedText())
	}
	w.Insert(2, "中")
	w.Insert(3, "c-")
	w.Delete(3)
	if w.Text() != "ab中-文" {
		t.Fatal("Text", w.Text())
	}
}
