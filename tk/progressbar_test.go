// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("ProgressBar", testProgressBar)
}

func testProgressBar(t *testing.T) {
	w := NewProgressBar(nil, Vertical)
	defer w.Destroy()

	w.SetOrient(Horizontal)
	if v := w.Orient(); v != Horizontal {
		t.Fatal("Orient", Horizontal, v)
	}

	w.SetTakeFocus(true)
	if v := w.IsTakeFocus(); v != true {
		t.Fatal("IsTakeFocus", true, v)
	}

	w.SetLength(20)
	if v := w.Length(); v != 20 {
		t.Fatal("Length", 20, v)
	}

	w.SetMaximum(0.5)
	if v := w.Maximum(); v != 0.5 {
		t.Fatal("Maximum", 0.5, v)
	}

	w.SetValue(0.5)
	if v := w.Value(); v != 0.5 {
		t.Fatal("Value", 0.5, v)
	}
}
