// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("Label", testLabel)
}

func testLabel(t *testing.T) {
	w := NewLabel(nil, "test", LabelAttrBackground("blue"), LabelAttrBorderWidth(20), LabelAttrForground("blue"), LabelAttrReliefStyle(1), LabelAttrAnchor(AnchorCenter), LabelAttrJustify(1), LabelAttrWrapLength(20), LabelAttrCompound(CompoundNone), LabelAttrText("text"), LabelAttrWidth(20), LabelAttrState(StateNormal), LabelAttrTakeFocus(true))
	defer w.Destroy()

	w.SetBackground("blue")
	if v := w.Background(); v != "blue" {
		t.Fatal("Background", "blue", v)
	}

	w.SetBorderWidth(20)
	if v := w.BorderWidth(); v != 20 {
		t.Fatal("BorderWidth", 20, v)
	}

	w.SetForground("blue")
	if v := w.Forground(); v != "blue" {
		t.Fatal("Forground", "blue", v)
	}

	w.SetReliefStyle(1)
	if v := w.ReliefStyle(); v != 1 {
		t.Fatal("ReliefStyle", 1, v)
	}

	w.SetAnchor(AnchorCenter)
	if v := w.Anchor(); v != AnchorCenter {
		t.Fatal("Anchor", AnchorCenter, v)
	}

	w.SetJustify(1)
	if v := w.Justify(); v != 1 {
		t.Fatal("Justify", 1, v)
	}

	w.SetWrapLength(20)
	if v := w.WrapLength(); v != 20 {
		t.Fatal("WrapLength", 20, v)
	}

	w.SetCompound(CompoundNone)
	if v := w.Compound(); v != CompoundNone {
		t.Fatal("Compound", CompoundNone, v)
	}

	w.SetText("text")
	if v := w.Text(); v != "text" {
		t.Fatal("Text", "text", v)
	}

	w.SetWidth(20)
	if v := w.Width(); v != 20 {
		t.Fatal("Width", 20, v)
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
