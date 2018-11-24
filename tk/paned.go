// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// panedwindow
type Paned struct {
	BaseWidget
}

func NewPaned(parent Widget, orient Orient, attributes ...*WidgetAttr) *Paned {
	iid := makeNamedWidgetId(parent, "atk_paned")
	attributes = append(attributes, &WidgetAttr{"orient", orient})
	info := CreateWidgetInfo(iid, WidgetTypePaned, true, attributes)
	if info == nil {
		return nil
	}
	w := &Paned{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *Paned) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypePaned)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *Paned) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *Paned) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *Paned) SetHeight(height int) error {
	return eval(fmt.Sprintf("%v configure -height {%v}", w.id, height))
}

func (w *Paned) Height() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -height", w.id))
	return r
}

func (w *Paned) AddWidget(widget Widget, weight int) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v add %v -weight %v", w.id, widget.Id(), weight))
}

func (w *Paned) InsertWidget(pane int, widget Widget, weight int) error {
	if !IsValidWidget(widget) {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v insert %v %v -weight %v", w.id, pane, widget.Id(), weight))
}

func (w *Paned) SetPane(pane int, weight int) error {
	return eval(fmt.Sprintf("%v pane %v -weight %v", w.id, pane, weight))
}

func (w *Paned) RemovePane(pane int) error {
	return eval(fmt.Sprintf("%v forget %v", w.id, pane))
}

func PanedAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func PanedAttrHeight(height int) *WidgetAttr {
	return &WidgetAttr{"height", height}
}
