// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"testing"
)

func init() {
	err := Init()
	if err != nil {
		panic(err)
	}
}

func TestWindow(t *testing.T) {
	mw := MainWindow()
	if mw.SetTitle("Hello").Title() != "Hello" {
		t.Error("SetTitle")
	}
	if mw.SetAlpha(0.9).Alpha() != 0.9 {
		t.Error("SetAlpha")
	}

	mw.SetVisible(false)
	if mw.IsVisible() {
		t.Error("SetVisible")
	}
	mw.SetVisible(true)
	if !mw.IsVisible() {
		t.Error("SetVisible")
	}

	mw.Iconify()
	if !mw.IsIconify() {
		t.Error("Iconify")
	}
	mw.ShowNormal()

	mw.SetTopmost(true)
	Update()
	if !mw.IsTopmost() {
		t.Error("SetTopmost")
	}
	mw.SetTopmost(false)

	mw.SetGeometry(100, 200, 300, 400)
	x, y, w, h := mw.Geometry()
	if x != 100 || y != 200 || w != 300 || h != 400 {
		t.Error("Geometry", x, y, w, h)
	}
	mw.SetPos(101, 202)
	x, y = mw.Pos()
	if x != 101 || y != 202 {
		t.Error("Pos", x, y)
	}
	mw.SetSize(301, 302)
	w, h = mw.Size()
	if w != 301 || h != 302 {
		t.Error("Size", w, h)
	}

	mw.ShowMaximized()
	if !mw.IsMaximized() {
		t.Error("IsMaximized")
	}
	mw.ShowNormal()

	mw.SetResizable(false, false)
	enableW, enableH := mw.IsResizable()
	if enableW != false || enableH != false {
		t.Error("Resizable")
	}
	mw.SetResizable(true, true)

	mw.SetWidth(311).SetHeight(312)
	if mw.Width() != 311 || mw.Height() != 312 {
		t.Error("Width/Height")
	}

	mw.SetFullScreen(true)
	//Update()
	if !mw.IsFullScreen() {
		t.Error("IsFullScreen")
	}
	mw.SetFullScreen(false)

	mw.SetMaximumSize(500, 600)
	w, h = mw.MaximumSize()
	if w != 500 || h != 600 {
		t.Error("MaximumSize")
	}

	mw.SetMinimumSize(200, 300)
	w, h = mw.MinimumSize()
	if w != 200 || h != 300 {
		t.Error("MinimumSize")
	}
}
