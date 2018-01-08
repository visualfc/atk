// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

func (w *Window) ShowMaximized() *Window {
	if !w.IsVisible() {
		w.ShowNormal()
	}
	eval(fmt.Sprintf("wm attributes %v -zoomed 1", w.id))
	return w
}

func (w *Window) IsMaximized() bool {
	r, _ := evalAsBool(fmt.Sprintf("wm attributes %v -zoomed", w.id))
	return r
}
