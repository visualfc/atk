// Copyright 2018 visualfc. All rights reserved.

package tk

type LayoutItem struct {
	widget Widget
	attrs  []*PackAttr
}

type LayoutFrame struct {
	BaseWidget
}

func (w *LayoutFrame) Type() string {
	return "LayoutFrame"
}

func NewLayoutFrame(parent Widget, attributes ...*WidgetAttr) *LayoutFrame {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_layoutframe")
	info := CreateWidgetInfo(iid, WidgetTypeFrame, theme, attributes)
	if info == nil {
		return nil
	}
	w := &LayoutFrame{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

type PackLayout struct {
	main  *LayoutFrame
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

func (w *PackLayout) AddWidget(widget Widget, attributes ...*PackAttr) {
	if !IsValidWidget(widget) {
		return
	}
	w.items = append(w.items, &LayoutItem{widget, attributes})
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

func (w *PackLayout) SetWidgetAttr(widget Widget, attributes ...*PackAttr) {
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

func (w *PackLayout) AddLayout(layout *PackLayout, attributes ...*PackAttr) {
	if layout == nil || !IsValidWidget(layout.main) {
		return
	}
	w.items = append(w.items, &LayoutItem{layout.main, attributes})
	w.Repack()
}

func (w *PackLayout) SetLayoutAttr(layout *PackLayout, attributes ...*PackAttr) {
	if layout == nil || !IsValidWidget(layout.main) {
		return
	}
	for _, item := range w.items {
		if item.widget == layout.main {
			item.attrs = attributes
		}
	}
	w.Repack()
}

func (w *PackLayout) Repack() {
	itemsAttr := []*PackAttr{PackAttrSide(w.side), PackAttrInMaster(w.main)}
	if w.pad != nil {
		itemsAttr = append(itemsAttr, PackAttrPadx(w.pad.X), PackAttrPady(w.pad.Y))
	}
	for _, item := range w.items {
		Pack(item.widget, AppendPackAttrs(item.attrs, itemsAttr...)...)
	}
	Pack(w.main, PackAttrFill(FillBoth), PackAttrExpand(true))
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
