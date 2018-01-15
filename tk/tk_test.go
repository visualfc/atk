// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"testing"
	"time"
)

func init() {
	err := Init()
	if err != nil {
		panic(err)
	}
	fnErrorHandle = func(err error) {
		panic(err)
	}
}

var (
	allTestMap = make(map[string]func(*testing.T))
)

func registerTest(name string, fn func(*testing.T)) {
	allTestMap[name] = fn
}

func TestMain(t *testing.T) {
	MainLoop(func() {
		t.Log("sub test", "MainWindow")
		testMainWindow(t)

		for name, fn := range allTestMap {
			t.Log("sub test", name)
			fn(t)
		}
		t.Log("sub test", "Async")
		go func() {
			<-time.After(1)
			Async(func() {
				Quit()
			})
		}()
	})
}

func testMainWindow(t *testing.T) {
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

	mw.SetGeometryN(100, 200, 300, 400)
	x, y, w, h := mw.GeometryN()
	if x != 100 || y != 200 || w != 300 || h != 400 {
		t.Error("Geometry", x, y, w, h)
	}
	mw.SetPosN(101, 202)
	x, y = mw.PosN()
	if x != 101 || y != 202 {
		t.Error("Pos", x, y)
	}
	mw.SetSizeN(301, 302)
	w, h = mw.SizeN()
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

	mw.SetMaximumSizeN(500, 600)
	w, h = mw.MaximumSizeN()
	if w != 500 || h != 600 {
		t.Error("MaximumSize")
	}

	mw.SetMinimumSizeN(200, 300)
	w, h = mw.MinimumSizeN()
	if w != 200 || h != 300 {
		t.Error("MinimumSize")
	}
}
