// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"os"
	"strings"
)

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
		return os.ErrInvalid
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
		return os.ErrInvalid
	}
	var attrList []string
	for _, attr := range attributes {
		if !w.info.MetaClass.HasAttribute(attr.Key) {
			continue
		}
		attrList = append(attrList, fmt.Sprintf("-%v {%v}", attr.Key, attr.Value))
	}
	if len(attrList) > 0 {
		return eval(fmt.Sprintf("%v configure %v", w.id, strings.Join(attrList, " ")))
	}
	return nil
}

func (w *BaseWidget) SetAttributes(attributes ...*WidgetAttr) error {
	if !IsValidWidget(w) {
		return os.ErrInvalid
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
