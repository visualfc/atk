// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("Scale", testScale)
}

func testScale(t *testing.T) {
	w := NewScale(nil, Vertical)
	defer w.Destroy()

	w.SetOrient(Horizontal)
	if v := w.Orient(); v != Horizontal {
		t.Fatal("Orient", Horizontal, v)
	}

	w.SetTakeFocus(true)
	if v := w.IsTakeFocus(); v != true {
		t.Fatal("IsTakeFocus", true, v)
	}

	w.SetFrom(0)
	if v := w.From(); v != 0 {
		t.Fatal("From", 0, v)
	}

	w.SetTo(100)
	if v := w.To(); v != 100 {
		t.Fatal("To", 100, v)
	}

	w.SetValue(0.5)
	if v := w.Value(); v != 0.5 {
		t.Fatal("Value", 0.5, v)
	}

	w.SetLength(20)
	if v := w.Length(); v != 20 {
		t.Fatal("Length", 20, v)
	}
}
