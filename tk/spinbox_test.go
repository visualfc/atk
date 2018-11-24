// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("SpinBox", testSpinBox)
}

func testSpinBox(t *testing.T) {
	w := NewSpinBox(nil, SpinBoxAttrTakeFocus(true), SpinBoxAttrFrom(0), SpinBoxAttrTo(100), SpinBoxAttrIncrement(0.5), SpinBoxAttrWrap(true))
	defer w.Destroy()

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

	w.SetIncrement(0.5)
	if v := w.Increment(); v != 0.5 {
		t.Fatal("Increment", 0.5, v)
	}

	w.SetWrap(true)
	if v := w.IsWrap(); v != true {
		t.Fatal("IsWrap", true, v)
	}
}
