// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("Notebook", testNotebook)
}

func testNotebook(t *testing.T) {
	w := NewNotebook(nil, NotebookAttrWidth(20), NotebookAttrHeight(20), NotebookAttrTakeFocus(true))
	defer w.Destroy()

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
