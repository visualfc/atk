// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("Frame", testFrame)
}

func testFrame(t *testing.T) {
	w := NewFrame(nil, FrameAttrBorderWidth(20), FrameAttrReliefStyle(1), FrameAttrWidth(20), FrameAttrHeight(20), FrameAttrTakeFocus(true))
	defer w.Destroy()

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
