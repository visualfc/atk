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

type Spacer struct {
	BaseWidget
	space  int
	expand bool
}

func (w *Spacer) Type() WidgetType {
	return WidgetTypeSpacer
}

func (w *Spacer) SetSpace(space int) *Spacer {
	w.space = space
	return w
}

func (w *Spacer) Space() int {
	return w.space
}

func (w *Spacer) SetExpand(expand bool) *Spacer {
	w.expand = expand
	return w
}

func (w *Spacer) IsExpand() bool {
	return w.expand
}

// width ignore for PackLayout
func (w *Spacer) SetWidth(width int) *Spacer {
	evalAsInt(fmt.Sprintf("%v configure -width {%v}", w.id, width))
	return w
}

func (w *Spacer) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

// height ignore for PackLayout
func (w *Spacer) SetHeight(height int) *Spacer {
	evalAsInt(fmt.Sprintf("%v configure -height {%v}", w.id, height))
	return w
}

func (w *Spacer) Height() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -height", w.id))
	return r
}

func NewSpacer(parent Widget, space int, expand bool) *Spacer {
	theme := checkInitUseTheme(nil)
	iid := makeNamedWidgetId(parent, "atk_spacer")
	info := CreateWidgetInfo(iid, WidgetTypeFrame, theme, nil)
	if info == nil {
		return nil
	}
	w := &Spacer{}
	w.id = iid
	w.info = info
	w.space = space
	w.expand = expand
	RegisterWidget(w)
	return w
}

func NewLayoutFrame(parent Widget, attributes ...*WidgetAttr) *Frame {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_layoutframe")
	info := CreateWidgetInfo(iid, WidgetTypeFrame, theme, attributes)
	if info == nil {
		return nil
	}
	w := &Frame{}
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
