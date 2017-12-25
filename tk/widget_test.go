// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"testing"
)

type TestWidget struct {
	id string
}

func (w *TestWidget) Id() string {
	return w.id
}

func (w *TestWidget) Type() string {
	return "Test"
}

func (w *TestWidget) String() string {
	return fmt.Sprintf("TestWidget:%v", w.id)
}

func NewTestWidget(parent Widget, id string) *TestWidget {
	w := &TestWidget{MakeWidgetId(id, parent)}
	RegisterWidget(w)
	return w
}

func TestWidgetId(t *testing.T) {
	var id string
	parent := &TestWidget{".base"}
	id = MakeWidgetId("", nil)
	if id != ".gotk_id1025" {
		t.Fatal(id)
	}
	id = MakeWidgetId("", parent)
	if id != ".base.gotk_id1026" {
		t.Fatal(id)
	}
	id = MakeWidgetId(".idtest", parent)
	if id != ".idtest" {
		t.Fatal(id)
	}
	id = MakeWidgetId("idtest", nil)
	if id != ".idtest" {
		t.Fatal(id)
	}
	id = MakeWidgetId("idtest", parent)
	if id != ".base.idtest" {
		t.Fatal(id)
	}
}

func findOfList(w Widget, list []Widget) bool {
	for _, v := range list {
		if v == w {
			return true
		}
	}
	return false
}

func TestWidgetParent(t *testing.T) {
	a1 := NewTestWidget(nil, "a1")
	a2 := NewTestWidget(nil, "a2")
	b1 := NewTestWidget(a1, "b1")
	c1 := NewTestWidget(b1, "c1")
	c2 := NewTestWidget(b1, "c2")
	c3 := NewTestWidget(b1, "c3")
	if p := ParentOfWidget(a1); p != MainWindow() {
		t.Fatal("ParentWidget", p)
	}
	if p := ParentOfWidget(b1); p != a1 {
		t.Fatal("ParentWidget", p)
	}
	if p := ParentOfWidget(c1); p != b1 {
		t.Fatal("ParentWidget", p)
	}
	list := ChildrenOfWidget(mainWindow)
	if len(list) != 2 || !findOfList(a1, list) || !findOfList(a2, list) {
		t.Fatal("ChildrenOfWidget", list)
	}
	list = ChildrenOfWidget(b1)
	if len(list) != 3 || !findOfList(c1, list) || !findOfList(c2, list) || !findOfList(c3, list) {
		t.Fatal("ChildrenOfWidget", list)
	}
	DestroyWidget(c3)
	list = ChildrenOfWidget(b1)
	if len(list) != 2 {
		t.Fatal("DestroyWidget", list)
	}
	DestroyWidget(a1)
	list = ChildrenOfWidget(mainWindow)
	if len(list) != 1 {
		t.Fatal("DestroyWidget", list)
	}
	if IsValidWidget(c1) {
		t.Fatal("IsValidWidget", a2)
	}
}
