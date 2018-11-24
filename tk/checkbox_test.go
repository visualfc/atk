// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("CheckButton", testCheckButton)
}

func testCheckButton(t *testing.T) {
	w := NewCheckButton(nil, "test", CheckButtonAttrText("text"), CheckButtonAttrWidth(20), CheckButtonAttrCompound(CompoundNone), CheckButtonAttrState(StateNormal), CheckButtonAttrTakeFocus(true))
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

	w.SetChecked(true)
	if v := w.IsChecked(); v != true {
		t.Fatal("IsChecked", true, v)
	}
}
