// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("RadioButton", testRadioButton)
}

func testRadioButton(t *testing.T) {
	w := NewRadioButton(nil, "test", RadioButtonAttrText("text"), RadioButtonAttrWidth(20), RadioButtonAttrCompound(CompoundNone), RadioButtonAttrState(StateNormal), RadioButtonAttrTakeFocus(true))
	defer w.Destroy()

	w.SetText("text")
	if v := w.Text(); v != "text" {
		t.Fatal("Text", "text", v)
	}

	w.SetWidth(20)
	if v := w.Width(); v != 20 {
		t.Fatal("Width", 20, v)
	}

	w.SetCompound(CompoundNone)
	if v := w.Compound(); v != CompoundNone {
		t.Fatal("Compound", CompoundNone, v)
	}

	w.SetState(StateNormal)
	if v := w.State(); v != StateNormal {
		t.Fatal("State", StateNormal, v)
	}

	w.SetTakeFocus(true)
	if v := w.IsTakeFocus(); v != true {
		t.Fatal("IsTakeFocus", true, v)
	}
}
