// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("Menu", testMenu)
}

func testMenu(t *testing.T) {
	w := NewMenu(nil, MenuAttrActiveBackground("blue"), MenuAttrActiveForground("blue"), MenuAttrBackground("blue"), MenuAttrForground("blue"), MenuAttrSelectColor("blue"), MenuAttrDisabledForground("blue"), MenuAttrActiveBorderWidth(20), MenuAttrBorderWidth(20), MenuAttrReliefStyle(1), MenuAttrTearoffTitle("text"), MenuAttrTearoff(true), MenuAttrTakeFocus(true))
	defer w.Destroy()

	w.SetActiveBackground("blue")
	if v := w.ActiveBackground(); v != "blue" {
		t.Fatal("ActiveBackground", "blue", v)
	}

	w.SetActiveForground("blue")
	if v := w.ActiveForground(); v != "blue" {
		t.Fatal("ActiveForground", "blue", v)
	}

	w.SetBackground("blue")
	if v := w.Background(); v != "blue" {
		t.Fatal("Background", "blue", v)
	}

	w.SetForground("blue")
	if v := w.Forground(); v != "blue" {
		t.Fatal("Forground", "blue", v)
	}

	w.SetSelectColor("blue")
	if v := w.SelectColor(); v != "blue" {
		t.Fatal("SelectColor", "blue", v)
	}

	w.SetDisabledForground("blue")
	if v := w.DisabledForground(); v != "blue" {
		t.Fatal("DisabledForground", "blue", v)
	}

	w.SetActiveBorderWidth(20)
	if v := w.ActiveBorderWidth(); v != 20 {
		t.Fatal("ActiveBorderWidth", 20, v)
	}

	w.SetBorderWidth(20)
	if v := w.BorderWidth(); v != 20 {
		t.Fatal("BorderWidth", 20, v)
	}

	w.SetReliefStyle(1)
	if v := w.ReliefStyle(); v != 1 {
		t.Fatal("ReliefStyle", 1, v)
	}

	w.SetTearoffTitle("text")
	if v := w.TearoffTitle(); v != "text" {
		t.Fatal("TearoffTitle", "text", v)
	}

	w.SetTearoff(true)
	if v := w.IsTearoff(); v != true {
		t.Fatal("IsTearoff", true, v)
	}

	w.SetTakeFocus(true)
	if v := w.IsTakeFocus(); v != true {
		t.Fatal("IsTakeFocus", true, v)
	}
}
