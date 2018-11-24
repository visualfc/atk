// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("LabelFrame", testLabelFrame)
}

func testLabelFrame(t *testing.T) {
	w := NewLabelFrame(nil, LabelFrameAttrLabelText("text"), LabelFrameAttrLabelAnchor(AnchorNorthWest), LabelFrameAttrBorderWidth(20), LabelFrameAttrReliefStyle(1), LabelFrameAttrWidth(20), LabelFrameAttrHeight(20), LabelFrameAttrTakeFocus(true))
	defer w.Destroy()

	w.SetLabelText("text")
	if v := w.LabelText(); v != "text" {
		t.Fatal("LabelText", "text", v)
	}

	w.SetLabelAnchor(AnchorNorthWest)
	if v := w.LabelAnchor(); v != AnchorNorthWest {
		t.Fatal("LabelAnchor", AnchorNorthWest, v)
	}

	w.SetBorderWidth(20)
	if v := w.BorderWidth(); v != 20 {
		t.Fatal("BorderWidth", 20, v)
	}

	w.SetReliefStyle(1)
	if v := w.ReliefStyle(); v != 1 {
		t.Fatal("ReliefStyle", 1, v)
	}

	w.SetWidth(20)
	if v := w.Width(); v != 20 {
		t.Fatal("Width", 20, v)
	}

	w.SetHeight(20)
	if v := w.Height(); v != 20 {
		t.Fatal("Height", 20, v)
	}

	w.SetTakeFocus(true)
	if v := w.IsTakeFocus(); v != true {
		t.Fatal("IsTakeFocus", true, v)
	}
}
