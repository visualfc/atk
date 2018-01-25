// Copyright 2018 visualfc. All rights reserved.

package tk

type PackLayout struct {
	master *Frame
	side   Side
	pad    *Pad
	items  []*LayoutItem
}

func (w *PackLayout) Master() Widget {
	return w.master
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
	for n, item := range w.items {
		if item.widget == widget {
			PackRemove(widget)
			w.items = append(w.items[:n], w.items[n+1:]...)
			return true
		}
	}
	return false
}

func (w *PackLayout) AddWidget(widget Widget, attributes ...*LayoutAttr) {
	if !IsValidWidget(widget) {
		return
	}
	w.items = append(w.items, &LayoutItem{widget, attributes})
	w.Repack()
}

func (w *PackLayout) AddWidgetEx(widget Widget, fill Fill, expand bool, anchor Anchor, before Widget, after Widget) {
	w.AddWidget(widget,
		PackAttrFill(fill), PackAttrExpand(expand),
		PackAttrAnchor(anchor),
		PackAttrBefore(before), PackAttrAfter(after))
}

func (w *PackLayout) AddWidgets(widgets []Widget, attributes ...*LayoutAttr) {
	for _, widget := range widgets {
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
	for _, item := range w.items {
		if item.widget == widget {
			item.attrs = attributes
		}
	}
	w.Repack()
}

func (w *PackLayout) AddLayout(layout Layout, attributes ...*LayoutAttr) {
	if !IsValidLayout(layout) {
		return
	}
	w.items = append(w.items, &LayoutItem{layout.Master(), attributes})
	w.Repack()
}

func (w *PackLayout) AddLayoutEx(layout Layout, fill Fill, expand bool, anchor Anchor, before Widget, after Widget) {
	w.AddLayout(layout,
		PackAttrFill(fill), PackAttrExpand(expand),
		PackAttrAnchor(anchor),
		PackAttrBefore(before), PackAttrAfter(after))
}

func (w *PackLayout) RemoveLayout(layout Layout) bool {
	if !IsValidLayout(layout) {
		return false
	}
	return w.removeItem(layout.Master())
}

func (w *PackLayout) SetLayoutAttr(layout *PackLayout, attributes ...*LayoutAttr) {
	if layout == nil || !IsValidWidget(layout.master) {
		return
	}
	for _, item := range w.items {
		if item.widget == layout.master {
			item.attrs = attributes
		}
	}
	w.Repack()
}

func (w *PackLayout) itemAttr() []*LayoutAttr {
	itemsAttr := []*LayoutAttr{PackAttrSide(w.side), PackAttrInMaster(w.master)}
	if w.pad != nil {
		itemsAttr = append(itemsAttr, PackAttrPadx(w.pad.X), PackAttrPady(w.pad.Y))
	}
	return itemsAttr
}

func (w *PackLayout) Repack() {
	for _, item := range w.items {
		Pack(item.widget, AppendLayoutAttrs(item.attrs, w.itemAttr()...)...)
	}
	Pack(w.master, PackAttrFill(FillBoth), PackAttrExpand(true))
}

func (w *PackLayout) SetBorderWidth(width int) {
	w.master.SetBorderWidth(width)
}

func (w *PackLayout) BorderWidth() int {
	return w.master.BorderWidth()
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
