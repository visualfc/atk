// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

func (w *Window) ShowMaximized() error {
	if !w.IsVisible() {
		w.ShowNormal()
	}
	return eval(fmt.Sprintf("wm attributes %v -zoomed 1", w.id))
}

func (w *Window) IsMaximized() bool {
	r, _ := evalAsBool(fmt.Sprintf("wm attributes %v -zoomed", w.id))
	return r
}
