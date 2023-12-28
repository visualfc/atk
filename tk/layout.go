// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

type LayoutWidget interface {
	Widget
	LayoutWidget() Widget
}

func checkLayoutWidget(widget Widget) Widget {
	if w, ok := (widget).(LayoutWidget); ok {
		return w.LayoutWidget()
	}
	return widget
}

type Layout interface {
	Widget
	AddWidget(widget Widget, attrs ...*LayoutAttr) error
	AddLayout(layout Layout, attrs ...*LayoutAttr) error
	RemoveWidget(widget Widget) error
	RemoveLayout(layout Layout) error
}

type LayoutAttr struct {
	Key   string
	Value interface{}
}

type LayoutItem struct {
	widget Widget
	attrs  []*LayoutAttr
}

type LayoutSpacer struct {
	BaseWidget
	space  int
	expand bool
}

func (w *LayoutSpacer) Type() WidgetType {
	return WidgetTypeLayoutSpacer
}

func (w *LayoutSpacer) TypeName() string {
	return "LayoutSpacer"
}

func (w *LayoutSpacer) SetSpace(space int) error {
	w.space = space
	return nil
}

func (w *LayoutSpacer) Space() int {
	return w.space
}

func (w *LayoutSpacer) SetExpand(expand bool) error {
	w.expand = expand
	return nil
}

func (w *LayoutSpacer) IsExpand() bool {
	return w.expand
}

// width ignore for PackLayout
func (w *LayoutSpacer) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *LayoutSpacer) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

// height ignore for PackLayout
func (w *LayoutSpacer) SetHeight(height int) error {
	return eval(fmt.Sprintf("%v configure -height {%v}", w.id, height))
}

func (w *LayoutSpacer) Height() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -height", w.id))
	return r
}

func NewLayoutSpacer(parent Widget, space int, expand bool) *LayoutSpacer {
	theme := checkInitUseTheme(nil)
	iid := makeNamedWidgetId(parent, "atk_layoutspacer")
	info := CreateWidgetInfo(iid, WidgetTypeFrame, theme, nil)
	if info == nil {
		return nil
	}
	w := &LayoutSpacer{}
	w.id = iid
	w.info = info
	w.space = space
	w.expand = expand
	RegisterWidget(w)
	return w
}

type LayoutFrame struct {
	BaseWidget
}

func (w *LayoutFrame) Type() WidgetType {
	return WidgetTypeLayoutFrame
}

func (w *LayoutFrame) TypeName() string {
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

func AppendLayoutAttrs(org []*LayoutAttr, attributes ...*LayoutAttr) []*LayoutAttr {
	var remain []*LayoutAttr
	var find bool
	for _, attr := range attributes {
		find = false
		for _, old := range org {
			if old.Key == attr.Key {
				old.Value = attr.Value
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
