// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"testing"
)

func init() {
	Init()
}

func TestWindow(t *testing.T) {
	root := RootWindow()
	if root.SetTitle("Hello").Title() != "Hello" {
		t.Errorf("SetTitle")
	}
	if root.SetAlpha(0.9).Alpha() != 0.9 {
		t.Errorf("SetAlpha")
	}

	root.SetVisible(false)
	if root.IsVisible() {
		t.Errorf("SetVisible")
	}
	root.SetVisible(true)
	if !root.IsVisible() {
		t.Errorf("SetVisible")
	}

	root.Iconify()
	if !root.IsIconify() {
		t.Errorf("Iconify")
	}
	root.ShowNormal()

	root.SetTopmost(true)
	Update()
	if !root.IsTopmost() {
		t.Errorf("SetTopmost")
	}
	root.SetTopmost(false)

	root.SetGeometry(100, 200, 300, 400)
	x, y, w, h := root.Geometry()
	if x != 100 || y != 200 || w != 300 || h != 400 {
		t.Errorf("Geometry")
	}
	root.SetPos(101, 202)
	x, y = root.Pos()
	if x != 101 || y != 202 {
		t.Errorf("Pos")
	}
	root.SetSize(301, 302)
	w, h = root.Size()
	if w != 301 || h != 302 {
		t.Errorf("Size")
	}

	root.ShowMaximized()
	if !root.IsMaximized() {
		t.Errorf("IsMaximized")
	}
	root.ShowNormal()

	root.SetResizable(false, false)
	enableW, enableH := root.IsResizable()
	if enableW != false || enableH != false {
		t.Errorf("Resizable")
	}
	root.SetResizable(true, true)

	root.SetWidth(311).SetHeight(312)
	if root.Width() != 311 || root.Height() != 312 {
		t.Errorf("Width/Height")
	}

	root.SetFullScreen(true)
	//Update()
	if !root.IsFullScreen() {
		t.Errorf("IsFullScreen")
	}
	root.SetFullScreen(false)

	root.SetMaximumSize(500, 600)
	w, h = root.MaximumSize()
	if w != 500 || h != 600 {
		t.Errorf("MaximumSize")
	}

	root.SetMinimumSize(200, 300)
	w, h = root.MinimumSize()
	if w != 200 || h != 300 {
		t.Errorf("MinimumSize")
	}
}
