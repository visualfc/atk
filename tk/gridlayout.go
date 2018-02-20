// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

type GridLayout struct {
	*LayoutFrame
	items []*LayoutItem
}

func (w *GridLayout) AddWidget(widget Widget, attrs ...*LayoutAttr) {
	if !IsValidWidget(widget) {
		return
	}
	Grid(widget, AppendLayoutAttrs(attrs, GridAttrInMaster(w))...)
}

func (w *GridLayout) AddWidgets(widgets ...Widget) {
	GridList(widgets, GridAttrInMaster(w))
}

func (w *GridLayout) AddWidgetEx(widget Widget, row int, column int, rowspan int, columnspan int, sticky Sticky) {
	if !IsValidWidget(widget) {
		return
	}
	Grid(widget, GridAttrRow(row), GridAttrColumn(column),
		GridAttrRowSpan(rowspan), GridAttrColumnSpan(columnspan),
		GridAttrSticky(sticky), GridAttrInMaster(w))
}

func (w *GridLayout) RemoveWidget(widget Widget) bool {
	if !IsValidWidget(widget) {
		return false
	}
	err := GridRemove(widget)
	return err == nil
}

func (w *GridLayout) Repack() {
	Pack(w, PackAttrFill(FillBoth), PackAttrExpand(true))
}

func (w *GridLayout) SetBorderWidth(width int) {
	evalAsInt(fmt.Sprintf("%v configure -borderwidth {%v}", w.Id(), width))
}

func (w *GridLayout) BorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.Id()))
	return r
}

// row index from 0, -1=all
func (w *GridLayout) SetRowAttr(row int, pad int, weight int, group string) {
	GridRowIndex(w, row, GridIndexAttrPad(pad), GridIndexAttrWeight(weight), GridIndexAttrUniform(group))
}

// column index from 0, -1=all
func (w *GridLayout) SetColumnAttr(column int, pad int, weight int, group string) {
	GridColumnIndex(w, column, GridIndexAttrPad(pad), GridIndexAttrWeight(weight), GridIndexAttrUniform(group))
}

func NewGridLayout(parent Widget) *GridLayout {
	grid := &GridLayout{NewLayoutFrame(parent), nil}
	grid.Lower(nil)
	grid.Repack()
	return grid
}
