// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

type GridLayout struct {
	*LayoutFrame
	items []*LayoutItem
}

func (w *GridLayout) AddWidget(widget Widget, attrs ...*LayoutAttr) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	return Grid(widget, AppendLayoutAttrs(attrs, GridAttrInMaster(w))...)
}

func (w *GridLayout) AddWidgets(widgets ...Widget) error {
	return GridList(widgets, GridAttrInMaster(w))
}

func (w *GridLayout) AddWidgetList(widgets []Widget, attrs ...*LayoutAttr) error {
	return GridList(widgets, AppendLayoutAttrs(attrs, GridAttrInMaster(w))...)
}

func (w *GridLayout) AddWidgetEx(widget Widget, row int, column int, rowspan int, columnspan int, sticky Sticky) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	return Grid(widget, GridAttrRow(row), GridAttrColumn(column),
		GridAttrRowSpan(rowspan), GridAttrColumnSpan(columnspan),
		GridAttrSticky(sticky), GridAttrInMaster(w))
}

func (w *GridLayout) RemoveWidget(widget Widget) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	return GridRemove(widget)
}

func (w *GridLayout) Repack() error {
	return Pack(w, PackAttrFill(FillBoth), PackAttrExpand(true))
}

func (w *GridLayout) SetBorderWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -borderwidth {%v}", w.Id(), width))
}

func (w *GridLayout) BorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.Id()))
	return r
}

// row index from 0, -1=all
func (w *GridLayout) SetRowAttr(row int, pad int, weight int, group string) error {
	return GridRowIndex(w, row, GridIndexAttrPad(pad), GridIndexAttrWeight(weight), GridIndexAttrUniform(group))
}

// column index from 0, -1=all
func (w *GridLayout) SetColumnAttr(column int, pad int, weight int, group string) error {
	return GridColumnIndex(w, column, GridIndexAttrPad(pad), GridIndexAttrWeight(weight), GridIndexAttrUniform(group))
}

func NewGridLayout(parent Widget) *GridLayout {
	grid := &GridLayout{NewLayoutFrame(parent), nil}
	grid.Lower(nil)
	grid.Repack()
	return grid
}
