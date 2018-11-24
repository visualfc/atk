// Copyright 2018 visualfc. All rights reserved.

package tk

import "testing"

func init() {
	registerTest("TreeView", testTreeView)
}

func testTreeView(t *testing.T) {
	w := NewTreeView(nil, TreeViewAttrTakeFocus(true), TreeViewAttrHeight(20), TreeViewAttrTreeSelectMode(TreeSelectBrowse))
	defer w.Destroy()

	w.SetTakeFocus(true)
	if v := w.IsTakeFocus(); v != true {
		t.Fatal("IsTakeFocus", true, v)
	}

	w.SetHeight(20)
	if v := w.Height(); v != 20 {
		t.Fatal("Height", 20, v)
	}

	w.SetTreeSelectMode(TreeSelectBrowse)
	if v := w.TreeSelectMode(); v != TreeSelectBrowse {
		t.Fatal("TreeSelectMode", TreeSelectBrowse, v)
	}
}
