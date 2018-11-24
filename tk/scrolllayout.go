// Copyright 2018 visualfc. All rights reserved.

package tk

type ScrollLayout struct {
	*GridLayout
	XScrollBar *ScrollBar
	YScrollBar *ScrollBar
}

func NewScrollLayout(parent Widget) *ScrollLayout {
	grid := NewGridLayout(parent)
	xscroll := NewScrollBar(parent, Horizontal)
	yscroll := NewScrollBar(parent, Vertical)
	grid.AddWidget(xscroll, GridAttrRow(1), GridAttrColumn(0), GridAttrSticky(StickyEW))
	grid.AddWidget(yscroll, GridAttrRow(0), GridAttrColumn(1), GridAttrSticky(StickyNS))
	return &ScrollLayout{grid, xscroll, yscroll}
}

func (w *ScrollLayout) SetWidget(widget Widget) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	w.AddWidget(widget, GridAttrRow(0), GridAttrColumn(0), GridAttrSticky(StickyAll))
	w.SetRowAttr(0, 0, 1, "")
	w.SetColumnAttr(0, 0, 1, "")
	return nil
}

//export embedded id
func (w *ScrollLayout) Id() string {
	return w.id
}

func (w *ScrollLayout) ShowXScrollBar(b bool) (err error) {
	if b {
		err = w.AddWidget(w.XScrollBar, GridAttrRow(1), GridAttrColumn(0), GridAttrSticky(StickyEW))
	} else {
		err = w.RemoveWidget(w.XScrollBar)
	}
	return
}

func (w *ScrollLayout) ShowYScrollBar(b bool) (err error) {
	if b {
		err = w.AddWidget(w.YScrollBar, GridAttrRow(1), GridAttrColumn(0), GridAttrSticky(StickyEW))
	} else {
		err = w.RemoveWidget(w.YScrollBar)
	}
	return
}
