// Copyright 2018 visualfc. All rights reserved.

package tk

type PlaceFrame struct {
	*Frame
	items []*LayoutItem
}

func (w *PlaceFrame) removeItem(widget Widget) error {
	n := w.indexOfWidget(widget)
	if n == -1 {
		return ErrInvalid
	}
	PlaceRemove(widget)
	w.items = append(w.items[:n], w.items[n+1:]...)
	return nil
}

func (w *PlaceFrame) indexOfWidget(widget Widget) int {
	for n, v := range w.items {
		if v.widget == widget {
			return n
		}
	}
	return -1
}

func (w *PlaceFrame) AddWidget(widget Widget, attributes ...*LayoutAttr) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	n := w.indexOfWidget(widget)
	if n != -1 {
		w.items = append(w.items[:n], w.items[n+1:]...)
	}
	w.items = append(w.items, &LayoutItem{widget, attributes})
	return w.Repack()
}

func (w *PlaceFrame) InsertWidget(index int, widget Widget, attributes ...*LayoutAttr) error {
	if index < 0 {
		return w.AddWidget(widget, attributes...)
	}
	n := w.indexOfWidget(widget)
	if n != -1 {
		if n == index {
			return ErrExist
		}
		w.items = append(w.items[:n], w.items[n+1:]...)
	}
	if index >= len(w.items) {
		return w.AddWidget(widget, attributes...)
	}
	w.items = append(w.items[:index], append([]*LayoutItem{&LayoutItem{widget, attributes}}, w.items[index:]...)...)
	return w.Repack()
}

func (w *PlaceFrame) RemoveWidget(widget Widget) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	return w.removeItem(widget)
}

func (w *PlaceFrame) SetWidgetAttr(widget Widget, attributes ...*LayoutAttr) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	n := w.indexOfWidget(widget)
	if n == -1 {
		return ErrInvalid
	}
	w.items[n].attrs = attributes
	return w.Repack()
}

func (w *PlaceFrame) Repack() error {
	for _, item := range w.items {
		if item.widget == nil {
			continue
		}
		Place(item.widget, AppendLayoutAttrs(item.attrs, PlaceAttrInMaster(w))...)
	}
	return Pack(w, PackAttrFill(FillBoth), PackAttrExpand(true))
}

func NewPlaceFrame(parent Widget) *PlaceFrame {
	place := &PlaceFrame{NewFrame(parent), nil}
	place.Lower(nil)
	place.Repack()
	return place
}
