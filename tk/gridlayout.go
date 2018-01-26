// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

type GridLayout struct {
	master *Frame
	items  []*LayoutItem
}

func (w *GridLayout) Master() Widget {
	return w.master
}

func (w *GridLayout) AddWidget(widget Widget, attrs ...*LayoutAttr) {
	if !IsValidWidget(widget) {
		return
	}
	Grid(widget, AppendLayoutAttrs(attrs, GridAttrInMaster(w.master))...)
}

func (w *GridLayout) AddWidgets(widgets []Widget, attrs ...*LayoutAttr) {
	GridList(widgets, AppendLayoutAttrs(attrs, GridAttrInMaster(w.master))...)
}

func (w *GridLayout) AddWidgetEx(widget Widget, row int, column int, rowspan int, columnspan int, sticky Sticky) {
	if !IsValidWidget(widget) {
		return
	}
	Grid(widget, GridAttrRow(row), GridAttrColumn(column),
		GridAttrRowSpan(rowspan), GridAttrColumnSpan(columnspan),
		GridAttrSticky(sticky), GridAttrInMaster(w.master))
}

func (w *GridLayout) AddLayout(layout Layout, attrs ...*LayoutAttr) {
	if !IsValidLayout(layout) {
		return
	}
	Grid(layout.Master(), AppendLayoutAttrs(attrs, GridAttrInMaster(w.master))...)
}

func (w *GridLayout) AddLayoutEx(layout Layout, row int, column int, rowspan int, columnspan int, sticky Sticky) {
	if !IsValidLayout(layout) {
		return
	}
	Grid(layout.Master(), GridAttrRow(row), GridAttrColumn(column),
		GridAttrRowSpan(rowspan), GridAttrColumnSpan(columnspan),
		GridAttrSticky(sticky), GridAttrInMaster(w.master))
}

func (w *GridLayout) RemoveWidget(widget Widget) bool {
	if !IsValidWidget(widget) {
		return false
	}
	err := GridRemove(widget)
	return err == nil
}

func (w *GridLayout) RemoveLayout(layout Layout) bool {
	if !IsValidLayout(layout) {
		return false
	}
	err := GridRemove(layout.Master())
	return err == nil
}

func (w *GridLayout) Repack() {
	Pack(w.master, PackAttrFill(FillBoth), PackAttrExpand(true))
}

func (w *GridLayout) SetBorderWidth(width int) {
	evalAsInt(fmt.Sprintf("%v configure -borderwidth {%v}", w.master.Id(), width))
}

func (w *GridLayout) BorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.master.Id()))
	return r
}

// row index from 0, -1=all
func (w *GridLayout) SetRow(row int, pad int, weight int, group string) {
	GridRowIndex(w.master, row, GridIndexAttrPad(pad), GridIndexAttrWeight(weight), GridIndexAttrUniform(group))
}

// column index from 0, -1=all
func (w *GridLayout) SetColumn(column int, pad int, weight int, group string) {
	GridColumnIndex(w.master, column, GridIndexAttrPad(pad), GridIndexAttrWeight(weight), GridIndexAttrUniform(group))
}

func NewGridLayout(parent Widget) *GridLayout {
	grid := &GridLayout{NewLayoutFrame(parent), nil}
	grid.Repack()
	return grid
}
