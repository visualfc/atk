// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"testing"
)

type TestWidget struct {
	id string
}

func (w *TestWidget) Id() string {
	return w.id
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
	a1 := &TestWidget{".a1"}
	a2 := &TestWidget{".a2"}
	b1 := &TestWidget{".a1.b1"}
	c1 := &TestWidget{".a1.b1.c1"}
	c2 := &TestWidget{".a1.b1.c2"}
	RegisterWidget(a1)
	RegisterWidget(a2)
	RegisterWidget(b1)
	RegisterWidget(c1)
	RegisterWidget(c2)
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
	if len(list) != 2 || !findOfList(c1, list) || !findOfList(c2, list) {
		t.Fatal("ChildrenOfWidget", list)
	}
}
