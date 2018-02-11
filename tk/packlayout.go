// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

type PackLayout struct {
	*LayoutFrame
	side  Side
	pad   *Pad
	items []*LayoutItem
}

func (w *PackLayout) SetSide(side Side) {
	w.side = side
	w.Repack()
}

func (w *PackLayout) SetPadding(pad Pad) {
	w.SetPaddingN(pad.X, pad.Y)
}

func (w *PackLayout) SetPaddingN(padx int, pady int) {
	w.pad = &Pad{padx, pady}
	w.Repack()
}

func (w *PackLayout) removeItem(widget Widget) bool {
	n := w.indexOfWidget(widget)
	if n == -1 {
		return false
	}
	PackRemove(widget)
	w.items = append(w.items[:n], w.items[n+1:]...)
	return true
}

func (w *PackLayout) indexOfWidget(widget Widget) int {
	for n, v := range w.items {
		if v.widget == widget {
			return n
		}
	}
	return -1
}

func (w *PackLayout) AddWidget(widget Widget, attributes ...*LayoutAttr) {
	if !IsValidWidget(widget) {
		return
	}
	n := w.indexOfWidget(widget)
	if n != -1 {
		w.items = append(w.items[:n], w.items[n+1:]...)
	}
	w.items = append(w.items, &LayoutItem{widget, attributes})
	w.Repack()
}

func (w *PackLayout) InsertWidget(index int, widget Widget, attributes ...*LayoutAttr) {
	if index < 0 {
		w.AddWidget(widget, attributes...)
		return
	}
	n := w.indexOfWidget(widget)
	if n != -1 {
		if n == index {
			return
		}
		w.items = append(w.items[:n], w.items[n+1:]...)
	}
	if index >= len(w.items) {
		w.AddWidget(widget, attributes...)
		return
	}
	w.items = append(w.items[:index], append([]*LayoutItem{&LayoutItem{widget, attributes}}, w.items[index:]...)...)
	w.Repack()
}

func (w *PackLayout) AddWidgetEx(widget Widget, fill Fill, expand bool, anchor Anchor) {
	w.AddWidget(widget,
		PackAttrFill(fill), PackAttrExpand(expand),
		PackAttrAnchor(anchor))
}

func (w *PackLayout) InsertWidgetEx(index int, widget Widget, fill Fill, expand bool, anchor Anchor) {
	w.InsertWidget(index, widget,
		PackAttrFill(fill), PackAttrExpand(expand),
		PackAttrAnchor(anchor))
}

func (w *PackLayout) AddWidgets(widgets ...Widget) {
	for _, widget := range widgets {
		n := w.indexOfWidget(widget)
		if n != -1 {
			w.items = append(w.items[:n], w.items[n+1:]...)
		}
		w.items = append(w.items, &LayoutItem{widget, nil})
	}
	w.Repack()
}

func (w *PackLayout) AddWidgetList(widgets []Widget, attributes ...*LayoutAttr) {
	for _, widget := range widgets {
		n := w.indexOfWidget(widget)
		if n != -1 {
			w.items = append(w.items[:n], w.items[n+1:]...)
		}
		w.items = append(w.items, &LayoutItem{widget, attributes})
	}
	w.Repack()
}

func (w *PackLayout) RemoveWidget(widget Widget) bool {
	if !IsValidWidget(widget) {
		return false
	}
	return w.removeItem(widget)
}

func (w *PackLayout) SetWidgetAttr(widget Widget, attributes ...*LayoutAttr) {
	if !IsValidWidget(widget) {
		return
	}
	n := w.indexOfWidget(widget)
	if n == -1 {
		return
	}
	w.items[n].attrs = attributes
	w.Repack()
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

func (w *PackLayout) Repack() {
	for _, item := range w.items {
		if item.widget == nil {
			continue
		}
		if s, ok := item.widget.(*LayoutSpacer); ok {
			w.resetSpacerAttr(item, s)
		}
		Pack(item.widget, AppendLayoutAttrs(item.attrs, w.itemAttr()...)...)
	}
	Pack(w, PackAttrFill(FillBoth), PackAttrExpand(true))
}

func (w *PackLayout) SetBorderWidth(width int) {
	evalAsInt(fmt.Sprintf("%v configure -borderwidth {%v}", w.Id(), width))
}

func (w *PackLayout) BorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.Id()))
	return r
}

func NewPackLayout(parent Widget, side Side) *PackLayout {
	pack := &PackLayout{NewLayoutFrame(parent), side, nil, nil}
	pack.Repack()
	return pack
}

func NewHPackLayout(parent Widget) *PackLayout {
	return NewPackLayout(parent, SideLeft)
}

func NewVPackLayout(parent Widget) *PackLayout {
	return NewPackLayout(parent, SideTop)
}
