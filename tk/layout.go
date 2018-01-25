// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

type LayoutAttr struct {
	key   string
	value interface{}
}

type LayoutItem struct {
	widget Widget
	attrs  []*LayoutAttr
}

type LayoutFrame struct {
	BaseWidget
}

func (w *LayoutFrame) SetPaddingN(padx int, pady int) *LayoutFrame {
	if w.info.IsTtk {
		evalAsString(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
	} else {
		evalAsString(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
	}
	return w
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
	eval("lower " + iid)
	return w
}

type Layout interface {
	Master() Widget
	AddWidget(widget Widget, attrs ...*LayoutAttr)
	AddLayout(layout Layout, attrs ...*LayoutAttr)
	RemoveWidget(widget Widget) bool
	RemoveLayout(layout Layout) bool
}

func IsValidLayout(layout Layout) bool {
	return layout != nil && IsValidWidget(layout.Master())
}

func AppendLayoutAttrs(org []*LayoutAttr, attributes ...*LayoutAttr) []*LayoutAttr {
	var remain []*LayoutAttr
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
