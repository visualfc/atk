// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// NOTE: update must
func (w *Window) ShowMaximized() error {
	return eval(fmt.Sprintf("update\nwm state %v zoomed", w.id))
}

func (w *Window) IsMaximized() bool {
	r, _ := evalAsString(fmt.Sprintf("wm state %v", w.id))
	return r == "zoomed"
}
