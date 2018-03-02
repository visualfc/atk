// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
)

type PackLayout struct {
	*LayoutFrame
	side  Side
	pad   *Pad
	items []*LayoutItem
}

func (w *PackLayout) SetSide(side Side) error {
	w.side = side
	return w.Repack()
}

func (w *PackLayout) SetPadding(pad Pad) error {
	return w.SetPaddingN(pad.X, pad.Y)
}

func (w *PackLayout) SetPaddingN(padx int, pady int) error {
	w.pad = &Pad{padx, pady}
	return w.Repack()
}

func (w *PackLayout) removeItem(widget Widget) error {
	n := w.indexOfWidget(widget)
	if n == -1 {
		return ErrInvalid
	}
	PackRemove(widget)
	w.items = append(w.items[:n], w.items[n+1:]...)
	return nil
}

func (w *PackLayout) indexOfWidget(widget Widget) int {
	for n, v := range w.items {
		if v.widget == widget {
			return n
		}
	}
	return -1
}

func (w *PackLayout) AddWidget(widget Widget, attributes ...*LayoutAttr) error {
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

func (w *PackLayout) InsertWidget(index int, widget Widget, attributes ...*LayoutAttr) error {
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

func (w *PackLayout) AddWidgetEx(widget Widget, fill Fill, expand bool, anchor Anchor) error {
	return w.AddWidget(widget,
		PackAttrFill(fill), PackAttrExpand(expand),
		PackAttrAnchor(anchor))
}

func (w *PackLayout) InsertWidgetEx(index int, widget Widget, fill Fill, expand bool, anchor Anchor) error {
	return w.InsertWidget(index, widget,
		PackAttrFill(fill), PackAttrExpand(expand),
		PackAttrAnchor(anchor))
}

func (w *PackLayout) AddWidgets(widgets ...Widget) error {
	for _, widget := range widgets {
		n := w.indexOfWidget(widget)
		if n != -1 {
			w.items = append(w.items[:n], w.items[n+1:]...)
		}
		w.items = append(w.items, &LayoutItem{widget, nil})
	}
	return w.Repack()
}

func (w *PackLayout) AddWidgetList(widgets []Widget, attributes ...*LayoutAttr) error {
	for _, widget := range widgets {
		n := w.indexOfWidget(widget)
		if n != -1 {
			w.items = append(w.items[:n], w.items[n+1:]...)
		}
		w.items = append(w.items, &LayoutItem{widget, attributes})
	}
	return w.Repack()
}

func (w *PackLayout) RemoveWidget(widget Widget) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	return w.removeItem(widget)
}

func (w *PackLayout) SetWidgetAttr(widget Widget, attributes ...*LayoutAttr) error {
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

func (w *PackLayout) itemAttr() []*LayoutAttr {
	itemsAttr := []*LayoutAttr{PackAttrSide(w.side), PackAttrInMaster(w)}
	if w.pad != nil {
		itemsAttr = append(itemsAttr, PackAttrPadx(w.pad.X), PackAttrPady(w.pad.Y))
	}
	return itemsAttr
}

func (w *PackLayout) resetSpacerAttr(item *LayoutItem, s *LayoutSpacer) {
	if s.IsExpand() {
		s.SetWidth(0)
		s.SetHeight(0)
		if w.side == SideTop || w.side == SideBottom {
			item.attrs = AppendLayoutAttrs(item.attrs, PackAttrFillY(), PackAttrExpand(true))
		} else {
			item.attrs = AppendLayoutAttrs(item.attrs, PackAttrFillX(), PackAttrExpand(true))
		}
	} else {
		item.attrs = AppendLayoutAttrs(item.attrs, PackAttrFillNone(), PackAttrExpand(false))
		if w.side == SideTop || w.side == SideBottom {
			s.SetHeight(s.space)
			s.SetWidth(0)
		} else {
			s.SetWidth(s.space)
			s.SetHeight(0)
		}
	}
}

func (w *PackLayout) Repack() error {
	for _, item := range w.items {
		if item.widget == nil {
			continue
		}
		if s, ok := item.widget.(*LayoutSpacer); ok {
			w.resetSpacerAttr(item, s)
		}
		Pack(item.widget, AppendLayoutAttrs(item.attrs, w.itemAttr()...)...)
	}
	return Pack(w, PackAttrFill(FillBoth), PackAttrExpand(true))
}

func (w *PackLayout) SetBorderWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -borderwidth {%v}", w.Id(), width))
}

func (w *PackLayout) BorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.Id()))
	return r
}

func NewPackLayout(parent Widget, side Side) *PackLayout {
	pack := &PackLayout{NewLayoutFrame(parent), side, nil, nil}
	pack.Lower(nil)
	pack.Repack()
	return pack
}

func NewHPackLayout(parent Widget) *PackLayout {
	return NewPackLayout(parent, SideLeft)
}

func NewVPackLayout(parent Widget) *PackLayout {
	return NewPackLayout(parent, SideTop)
}
