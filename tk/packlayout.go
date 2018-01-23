// Copyright 2018 visualfc. All rights reserved.

package tk

type LayoutItem struct {
	widget Widget
	attrs  []*PackAttr
}

type PackLayout struct {
	main  *Frame
	side  Side
	attrs []*PackAttr
	items []*LayoutItem
}

func (w *PackLayout) sideAttr() *PackAttr {
	return &PackAttr{"side", w.side}
}

func (w *PackLayout) UpdateSide(side Side) {
	w.side = side
	w.Repack()
}

func (w *PackLayout) UpdatePackAttrs(attributes ...*PackAttr) {
	w.attrs = AppendPackAttrs(w.attrs, attributes...)
	w.Repack()
}

func (w *PackLayout) AddWidget(widget Widget, attributes ...*PackAttr) {
	if !IsValidWidget(widget) {
		return
	}
	w.items = append(w.items, &LayoutItem{widget, AppendPackAttrs(attributes, PackAttrInMaster(w.main))})
	w.Repack()
}

func (w *PackLayout) RemoveWidget(widget Widget) bool {
	if !IsValidWidget(widget) {
		return false
	}
	for n, item := range w.items {
		if item.widget == widget {
			eval("pack forget " + widget.Id())
			w.items = append(w.items[:n], w.items[n+1:]...)
			return true
		}
	}
	return false
}

func (w *PackLayout) UpdateWidget(widget Widget, attributes ...*PackAttr) {
	if !IsValidWidget(widget) {
		return
	}
	for _, item := range w.items {
		if item.widget == widget {
			item.attrs = AppendPackAttrs(item.attrs, attributes...)
		}
	}
	w.Repack()
}

func (w *PackLayout) AddLayout(layout *PackLayout, attributes ...*PackAttr) {
	if layout == nil || !IsValidWidget(layout.main) {
		return
	}
	w.items = append(w.items, &LayoutItem{layout.main, AppendPackAttrs(attributes, PackAttrInMaster(w.main))})
	w.Repack()
}

func (w *PackLayout) UpdateLayout(layout *PackLayout, attributes ...*PackAttr) {
	if layout == nil || !IsValidWidget(layout.main) {
		return
	}
	for _, item := range w.items {
		if item.widget == layout.main {
			item.attrs = AppendPackAttrs(item.attrs, attributes...)
		}
	}
	w.Repack()
}

func (w *PackLayout) Repack() {
	for _, item := range w.items {
		Pack(item.widget, AppendPackAttrs(item.attrs, w.sideAttr())...)
	}
	Pack(w.main, w.attrs...)
}

func NewPackLayout(parent Widget, side Side, attributes ...*PackAttr) *PackLayout {
	attrs := []*PackAttr{PackAttrFill(FillBoth), PackAttrExpand(true)}
	return &PackLayout{NewFrame(parent), side, AppendPackAttrs(attrs, attributes...), nil}
}

func NewHPackLayout(parent Widget, attributes ...*PackAttr) *PackLayout {
	return NewPackLayout(parent, SideLeft, attributes...)
}

func NewVPackLayout(parent Widget, attributes ...*PackAttr) *PackLayout {
	return NewPackLayout(parent, SideTop, attributes...)
}

func AppendPackAttrs(org []*PackAttr, attributes ...*PackAttr) []*PackAttr {
	var remain []*PackAttr
	var find bool
	for _, attr := range attributes {
		find = false
		for _, old := range org {
			if old.key == attr.key {
				old.value = attr.value
				find = true
				break
			}
		}
		if !find {
			remain = append(remain, attr)
		}
	}
	return append(org, remain...)
}
