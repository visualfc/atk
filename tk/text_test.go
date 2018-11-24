// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("Text", testText)
}

func testText(t *testing.T) {
	w := NewText(nil, TextAttrBackground("blue"), TextAttrBorderWidth(20), TextAttrForeground("blue"), TextAttrHighlightBackground("blue"), TextAttrHighlightColor("blue"), TextAttrHighlightthickness(20), TextAttrInsertBackground("blue"), TextAttrInsertBorderWidth(20), TextAttrInsertOffTime(20), TextAttrInsertOnTime(20), TextAttrInsertWidth(20), TextAttrReliefStyle(1), TextAttrSelectBackground("blue"), TextAttrSelectborderwidth(20), TextAttrSelectforeground("blue"), TextAttrInactiveSelectBackground("blue"), TextAttrTakeFocus(true), TextAttrAutoSeparatorsOnUndo(true), TextAttrBlockCursor(true), TextAttrWidth(20), TextAttrHeight(20), TextAttrInsertUnfocussed(DisplyCursorHollow), TextAttrMaxUndo(20), TextAttrLineAboveSpace(20), TextAttrLineWrapSpace(20), TextAttrLineBelowSpace(20), TextAttrLineWrap(LineWrapNone), TextAttrEnableUndo(true))
	defer w.Destroy()

	w.SetBackground("blue")
	if v := w.Background(); v != "blue" {
		t.Fatal("Background", "blue", v)
	}

	w.SetBorderWidth(20)
	if v := w.BorderWidth(); v != 20 {
		t.Fatal("BorderWidth", 20, v)
	}

	w.SetForeground("blue")
	if v := w.Foreground(); v != "blue" {
		t.Fatal("Foreground", "blue", v)
	}

	w.SetHighlightBackground("blue")
	if v := w.HighlightBackground(); v != "blue" {
		t.Fatal("HighlightBackground", "blue", v)
	}

	w.SetHighlightColor("blue")
	if v := w.HighlightColor(); v != "blue" {
		t.Fatal("HighlightColor", "blue", v)
	}

	w.SetHighlightthickness(20)
	if v := w.Highlightthickness(); v != 20 {
		t.Fatal("Highlightthickness", 20, v)
	}

	w.SetInsertBackground("blue")
	if v := w.InsertBackground(); v != "blue" {
		t.Fatal("InsertBackground", "blue", v)
	}

	w.SetInsertBorderWidth(20)
	if v := w.InsertBorderWidth(); v != 20 {
		t.Fatal("InsertBorderWidth", 20, v)
	}

	w.SetInsertOffTime(20)
	if v := w.InsertOffTime(); v != 20 {
		t.Fatal("InsertOffTime", 20, v)
	}

	w.SetInsertOnTime(20)
	if v := w.InsertOnTime(); v != 20 {
		t.Fatal("InsertOnTime", 20, v)
	}

	w.SetInsertWidth(20)
	if v := w.InsertWidth(); v != 20 {
		t.Fatal("InsertWidth", 20, v)
	}

	w.SetReliefStyle(1)
	if v := w.ReliefStyle(); v != 1 {
		t.Fatal("ReliefStyle", 1, v)
	}

	w.SetSelectBackground("blue")
	if v := w.SelectBackground(); v != "blue" {
		t.Fatal("SelectBackground", "blue", v)
	}

	w.SetSelectborderwidth(20)
	if v := w.Selectborderwidth(); v != 20 {
		t.Fatal("Selectborderwidth", 20, v)
	}

	w.SetSelectforeground("blue")
	if v := w.Selectforeground(); v != "blue" {
		t.Fatal("Selectforeground", "blue", v)
	}

	w.SetInactiveSelectBackground("blue")
	if v := w.InactiveSelectBackground(); v != "blue" {
		t.Fatal("InactiveSelectBackground", "blue", v)
	}

	w.SetTakeFocus(true)
	if v := w.IsTakeFocus(); v != true {
		t.Fatal("IsTakeFocus", true, v)
	}

	w.SetAutoSeparatorsOnUndo(true)
	if v := w.IsAutoSeparatorsOnUndo(); v != true {
		t.Fatal("IsAutoSeparatorsOnUndo", true, v)
	}

	w.SetBlockCursor(true)
	if v := w.IsBlockCursor(); v != true {
		t.Fatal("IsBlockCursor", true, v)
	}

	w.SetWidth(20)
	if v := w.Width(); v != 20 {
		t.Fatal("Width", 20, v)
	}

	w.SetHeight(20)
	if v := w.Height(); v != 20 {
		t.Fatal("Height", 20, v)
	}

	w.SetInsertUnfocussed(DisplyCursorHollow)
	if v := w.InsertUnfocussed(); v != DisplyCursorHollow {
		t.Fatal("InsertUnfocussed", DisplyCursorHollow, v)
	}

	w.SetMaxUndo(20)
	if v := w.MaxUndo(); v != 20 {
		t.Fatal("MaxUndo", 20, v)
	}

	w.SetLineAboveSpace(20)
	if v := w.LineAboveSpace(); v != 20 {
		t.Fatal("LineAboveSpace", 20, v)
	}

	w.SetLineWrapSpace(20)
	if v := w.LineWrapSpace(); v != 20 {
		t.Fatal("LineWrapSpace", 20, v)
	}

	w.SetLineBelowSpace(20)
	if v := w.LineBelowSpace(); v != 20 {
		t.Fatal("LineBelowSpace", 20, v)
	}

	w.SetLineWrap(LineWrapNone)
	if v := w.LineWrap(); v != LineWrapNone {
		t.Fatal("LineWrap", LineWrapNone, v)
	}

	w.SetEnableUndo(true)
	if v := w.IsEnableUndo(); v != true {
		t.Fatal("IsEnableUndo", true, v)
	}
}
