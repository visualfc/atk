// Copyright 2018 visualfc. All rights reserved.

package tk

type PackLayout struct {
	master *LayoutFrame
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

func (w *PackLayout) removeItem(id string) bool {
	for n, item := range w.items {
		if item.widget.Id() == id {
			eval("pack forget " + id)
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

func (w *PackLayout) RemoveWidget(widget Widget) bool {
	if widget == nil {
		return false
	}
	return w.removeItem(widget.Id())
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

func (w *PackLayout) RemoveLayout(layout Layout) bool {
	if layout == nil || layout.Master() == nil {
		return false
	}
	return w.removeItem(layout.Master().Id())
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

func (w *PackLayout) Repack() {
	itemsAttr := []*LayoutAttr{PackAttrSide(w.side), PackAttrInMaster(w.master)}
	if w.pad != nil {
		itemsAttr = append(itemsAttr, PackAttrPadx(w.pad.X), PackAttrPady(w.pad.Y))
	}
	for _, item := range w.items {
		Pack(item.widget, AppendLayoutAttrs(item.attrs, itemsAttr...)...)
	}
	Pack(w.master, PackAttrFill(FillBoth), PackAttrExpand(true))
}

func NewPackLayout(parent Widget, side Side) *PackLayout {
	return &PackLayout{NewLayoutFrame(parent), side, nil, nil}
}

func NewHPackLayout(parent Widget) *PackLayout {
	return NewPackLayout(parent, SideLeft)
}

func NewVPackLayout(parent Widget) *PackLayout {
	return NewPackLayout(parent, SideTop)
}
