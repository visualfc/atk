// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("ListBox", testListBox)
}

func testListBox(t *testing.T) {
	w := NewListBox(nil, ListBoxAttrBackground("blue"), ListBoxAttrBorderWidth(20), ListBoxAttrForground("blue"), ListBoxAttrReliefStyle(1), ListBoxAttrJustify(1), ListBoxAttrWidth(20), ListBoxAttrHeight(20), ListBoxAttrState(StateNormal), ListBoxAttrSelectMode(ListSelectSingle), ListBoxAttrTakeFocus(true))
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

	w.SetState(StateNormal)
	if v := w.State(); v != StateNormal {
		t.Fatal("State", StateNormal, v)
	}

	w.SetSelectMode(ListSelectSingle)
	if v := w.SelectMode(); v != ListSelectSingle {
		t.Fatal("SelectMode", ListSelectSingle, v)
	}

	w.SetTakeFocus(true)
	if v := w.IsTakeFocus(); v != true {
		t.Fatal("IsTakeFocus", true, v)
	}

	w.SetState(StateNormal)
	w.SetItems([]string{"100", "$ok", "{300"})
	if v := w.Items(); v[0] != "100" || v[1] != "$ok" || v[2] != "{300" {
		t.Fatal("values", v)
	}
	if v := w.ItemCount(); v != 3 {
		t.Fatal("SetItems", v)
	}
	w.InsertItem(0, "first")
	if v := w.ItemCount(); v != 4 {
		t.Fatal("InsertItem", v)
	}
	if v := w.ItemText(0); v != "first" {
		t.Fatal("InsertItem", v)
	}
	w.SetItemText(1, "$mm")
	if v := w.ItemText(1); v != "$mm" {
		t.Fatal("SetItemText", v)
	}
}
