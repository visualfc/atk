// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("Canvas", testCanvas)
}

func testCanvas(t *testing.T) {
	w := NewCanvas(nil, CanvasAttrBackground("blue"), CanvasAttrBorderWidth(20), CanvasAttrHighlightBackground("blue"), CanvasAttrHighlightColor("blue"), CanvasAttrHighlightthickness(20), CanvasAttrInsertBackground("blue"), CanvasAttrInsertBorderWidth(20), CanvasAttrInsertOffTime(20), CanvasAttrInsertOnTime(20), CanvasAttrInsertWidth(20), CanvasAttrReliefStyle(1), CanvasAttrSelectBackground("blue"), CanvasAttrSelectborderwidth(20), CanvasAttrSelectforeground("blue"), CanvasAttrTakeFocus(true), CanvasAttrCloseEnough(0.5), CanvasAttrConfine(true), CanvasAttrWidth(20), CanvasAttrHeight(20), CanvasAttrState(StateNormal), CanvasAttrXScrollIncrement(20), CanvasAttrYScrollIncrement(20))
	defer w.Destroy()

	w.SetBackground("blue")
	if v := w.Background(); v != "blue" {
		t.Fatal("Background", "blue", v)
	}

	w.SetBorderWidth(20)
	if v := w.BorderWidth(); v != 20 {
		t.Fatal("BorderWidth", 20, v)
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

	w.SetTakeFocus(true)
	if v := w.IsTakeFocus(); v != true {
		t.Fatal("IsTakeFocus", true, v)
	}

	w.SetCloseEnough(0.5)
	if v := w.CloseEnough(); v != 0.5 {
		t.Fatal("CloseEnough", 0.5, v)
	}

	w.SetConfine(true)
	if v := w.IsConfine(); v != true {
		t.Fatal("IsConfine", true, v)
	}

	w.SetWidth(20)
	if v := w.Width(); v != 20 {
		t.Fatal("Width", 20, v)
	}

	w.SetHeight(20)
	if v := w.Height(); v != 20 {
		t.Fatal("Height", 20, v)
	}

	w.SetState(StateNormal)
	if v := w.State(); v != StateNormal {
		t.Fatal("State", StateNormal, v)
	}

	w.SetXScrollIncrement(20)
	if v := w.XScrollIncrement(); v != 20 {
		t.Fatal("XScrollIncrement", 20, v)
	}

	w.SetYScrollIncrement(20)
	if v := w.YScrollIncrement(); v != 20 {
		t.Fatal("YScrollIncrement", 20, v)
	}
}
