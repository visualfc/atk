// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

// scrollbar
type ScrollBar struct {
	BaseWidget
	command *CommandEx
}

func NewScrollBar(parent Widget, orient Orient, attributes ...*WidgetAttr) *ScrollBar {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_scrollbar")
	attributes = append(attributes, &WidgetAttr{"orient", orient})
	info := CreateWidgetInfo(iid, WidgetTypeScrollBar, theme, attributes)
	if info == nil {
		return nil
	}
	w := &ScrollBar{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *ScrollBar) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeScrollBar)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *ScrollBar) SetOrient(orient Orient) error {
	return eval(fmt.Sprintf("%v configure -orient {%v}", w.id, orient))
}

func (w *ScrollBar) Orient() Orient {
	r, err := evalAsString(fmt.Sprintf("%v cget -orient", w.id))
	return parserOrientResult(r, err)
}

func (w *ScrollBar) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *ScrollBar) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *ScrollBar) SetScroll(first float64, last float64) error {
	err := eval(fmt.Sprintf("%v set %v %v", w.id, first, last))
	return err
}

func (w *ScrollBar) SetScrollArgs(args []string) error {
	err := eval(fmt.Sprintf("%v set %v", w.id, strings.Join(args, " ")))
	return err
}

func (w *ScrollBar) OnCommandEx(fn func([]string) error) error {
	if fn == nil {
		return ErrInvalid
	}
	if w.command == nil {
		w.command = &CommandEx{}
		bindCommandEx(w.id, "command", w.command.Invoke)
	}
	w.command.Bind(fn)
	return nil
}

func ScrollBarAttrOrient(orient Orient) *WidgetAttr {
	return &WidgetAttr{"orient", orient}
}

func ScrollBarAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}
