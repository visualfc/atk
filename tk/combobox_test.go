// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("ComboBox", testComboBox)
}

func testComboBox(t *testing.T) {
	w := NewComboBox(nil, ComboBoxAttrBackground("blue"), ComboBoxAttrForground("blue"), ComboBoxAttrJustify(1), ComboBoxAttrWidth(20), ComboBoxAttrHeight(20), ComboBoxAttrEcho("*"), ComboBoxAttrState(StateNormal), ComboBoxAttrTakeFocus(true))
	defer w.Destroy()

	w.SetBackground("blue")
	if v := w.Background(); v != "blue" {
		t.Fatal("Background", "blue", v)
	}

	w.SetForground("blue")
	if v := w.Forground(); v != "blue" {
		t.Fatal("Forground", "blue", v)
	}

	w.SetJustify(1)
	if v := w.Justify(); v != 1 {
		t.Fatal("Justify", 1, v)
	}

	w.SetWidth(20)
	if v := w.Width(); v != 20 {
		t.Fatal("Width", 20, v)
	}

	w.SetHeight(20)
	if v := w.Height(); v != 20 {
		t.Fatal("Height", 20, v)
	}

	w.SetEcho("*")
	if v := w.Echo(); v != "*" {
		t.Fatal("Echo", "*", v)
	}

	w.SetState(StateNormal)
	if v := w.State(); v != StateNormal {
		t.Fatal("State", StateNormal, v)
	}

	w.SetTakeFocus(true)
	if v := w.IsTakeFocus(); v != true {
		t.Fatal("IsTakeFocus", true, v)
	}

	w.SetValues([]string{"100", "$ok", "{300"})
	if v := w.Values(); v[0] != "100" || v[1] != "$ok" || v[2] != "{300" {
		t.Fatal("values", v)
	}

	w.SetCurrentText("$ok hello}")
	if v := w.CurrentText(); v != "$ok hello}" {
		t.Fatal("CurrentText", v)
	}

	w.SetCurrentIndex(2)
	if v := w.CurrentIndex(); v != 2 {
		t.Fatal("CurrentIndex", v)
	}
}
