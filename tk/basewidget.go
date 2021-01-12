// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

var _ Widget = &BaseWidget{}

type BaseWidget struct {
	id   string
	info *WidgetInfo
}

func (w *BaseWidget) String() string {
	iw := globalWidgetMap[w.id]
	if iw != nil {
		return fmt.Sprintf("%v{%v}", iw.TypeName(), w.id)
	} else {
		return fmt.Sprintf("Invalid{%v}", w.id)
	}
}

func (w *BaseWidget) Id() string {
	return w.id
}

func (w *BaseWidget) Info() *WidgetInfo {
	return w.info
}

func (w *BaseWidget) Type() WidgetType {
	if w.info != nil {
		return w.info.Type
	}
	return WidgetTypeNone
}

func (w *BaseWidget) TypeName() string {
	if w.info != nil {
		return w.info.TypeName
	}
	return "Invalid"
}

func (w *BaseWidget) Parent() Widget {
	return ParentOfWidget(w)
}

func (w *BaseWidget) Children() []Widget {
	return ChildrenOfWidget(w)
}

func (w *BaseWidget) IsValid() bool {
	return IsValidWidget(w)
}

func (w *BaseWidget) Destroy() error {
	return DestroyWidget(w)
}

func (w *BaseWidget) DestroyChildren() error {
	if !IsValidWidget(w) {
		return ErrInvalid
	}
	for _, child := range w.Children() {
		DestroyWidget(child)
	}
	return nil
}

func (w *BaseWidget) NativeAttribute(key string) string {
	if !IsValidWidget(w) {
		return ""
	}
	if !w.info.MetaClass.HasAttribute(key) {
		return ""
	}
	r, _ := evalAsString(fmt.Sprintf("%v cget -%v", w.id, key))
	return r
}

func (w *BaseWidget) NativeAttributes(keys ...string) (attributes []NativeAttr) {
	if !IsValidWidget(w) {
		return nil
	}
	if keys == nil {
		for _, key := range w.info.MetaClass.Attributes {
			r, _ := evalAsString(fmt.Sprintf("%v cget -%v", w.id, key))
			attributes = append(attributes, NativeAttr{key, r})
		}
	} else {
		for _, key := range keys {
			if w.info.MetaClass.HasAttribute(key) {
				r, _ := evalAsString(fmt.Sprintf("%v cget -%v", w.id, key))
				attributes = append(attributes, NativeAttr{key, r})
			}
		}
	}
	return
}

func (w *BaseWidget) SetNativeAttribute(key string, value string) error {
	return w.SetNativeAttributes([]NativeAttr{NativeAttr{key, value}}...)
}

func (w *BaseWidget) SetNativeAttributes(attributes ...NativeAttr) error {
	if !IsValidWidget(w) {
		return ErrInvalid
	}
	var attrList []string
	for _, attr := range attributes {
		if !w.info.MetaClass.HasAttribute(attr.Key) {
			continue
		}
		pname := "atk_tmp_" + attr.Key
		setObjText(pname, attr.Value)
		attrList = append(attrList, fmt.Sprintf("-%v $%v", attr.Key, pname))
	}
	if len(attrList) > 0 {
		return eval(fmt.Sprintf("%v configure %v", w.id, strings.Join(attrList, " ")))
	}
	return nil
}

func (w *BaseWidget) SetAttributes(attributes ...*WidgetAttr) error {
	if !IsValidWidget(w) {
		return ErrInvalid
	}
	extra := buildWidgetAttributeScript(w.info.MetaClass, w.info.IsTtk, attributes)
	if len(extra) > 0 {
		return eval(fmt.Sprintf("%v configure %v", w.id, extra))
	}
	return nil
}

func (w *BaseWidget) BindEvent(event string, fn func(e *Event)) error {
	return BindEvent(w.id, event, fn)
}

func (w *BaseWidget) BindKeyEvent(fn func(e *KeyEvent)) error {
	return BindKeyEventEx(w.id, fn, nil)
}

func (w *BaseWidget) BindKeyEventEx(fnPress func(e *KeyEvent), fnRelease func(e *KeyEvent)) error {
	return BindKeyEventEx(w.id, fnPress, fnRelease)
}

func (w *BaseWidget) BindInfo() []string {
	return BindInfo(w.id)
}

func (w *BaseWidget) ClearBind(event string) error {
	return ClearBindEvent(w.id, event)
}

func (w *BaseWidget) Lower(below Widget) error {
	script := fmt.Sprintf("lower %v", w.id)
	if IsValidWidget(below) {
		script += " " + below.Id()
	}
	return eval(script)
}

func (w *BaseWidget) Raise(above Widget) error {
	script := fmt.Sprintf("raise %v", w.id)
	if IsValidWidget(above) {
		script += " " + above.Id()
	}
	return eval(script)
}

func (w *BaseWidget) SetFocus() error {
	return eval(fmt.Sprintf("focus %v", w.id))
}

func (w *BaseWidget) IsFocus() bool {
	id, err := evalAsString("focus")
	if err != nil || id == "" {
		return false
	}
	return w.id == id
}

func (w *BaseWidget) FocusNextWidget() Widget {
	id, err := evalAsString("tk_focusNext " + w.id)
	if err != nil || id == "" {
		return nil
	}
	return FindWidget(id)
}

func (w *BaseWidget) FocusPrevWidget() Widget {
	id, err := evalAsString("tk_focusPrev " + w.id)
	if err != nil || id == "" {
		return nil
	}
	return FindWidget(id)
}

func SetFocusFollowsMouse() error {
	return eval("tk_focusFollowsMouse")
}

func FocusWidget() Widget {
	id, err := evalAsString("focus")
	if err != nil || id == "" {
		return nil
	}
	return FindWidget(id)
}
