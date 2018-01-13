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
		return fmt.Sprintf("%v{%v}", iw.Type(), w.id)
	} else {
		return fmt.Sprintf("Widget{%v}", w.id)
	}
}

func (w *BaseWidget) Id() string {
	return w.id
}

func (w *BaseWidget) Info() *WidgetInfo {
	return w.info
}

func (w *BaseWidget) Type() string {
	return "BaseWidget"
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

func (w *BaseWidget) NativeOption(key string) string {
	if !IsValidWidget(w) {
		return ""
	}
	if !w.info.MetaClass.HasOption(key) {
		return ""
	}
	r, _ := evalAsString(fmt.Sprintf("%v cget -%v", w.id, key))
	return r
}

func (w *BaseWidget) NativeOptions(keys ...string) (opts []NativeOpt) {
	if !IsValidWidget(w) {
		return nil
	}
	if keys == nil {
		for _, key := range w.info.MetaClass.Options {
			r, _ := evalAsString(fmt.Sprintf("%v cget -%v", w.id, key))
			opts = append(opts, NativeOpt{key, r})
		}
	} else {
		for _, key := range keys {
			if w.info.MetaClass.HasOption(key) {
				r, _ := evalAsString(fmt.Sprintf("%v cget -%v", w.id, key))
				opts = append(opts, NativeOpt{key, r})
			}
		}
	}
	return
}

func (w *BaseWidget) SetNativeOption(key string, value string) error {
	return w.SetNativeOptions([]NativeOpt{NativeOpt{key, value}}...)
}

func (w *BaseWidget) SetNativeOptions(options ...NativeOpt) error {
	if !IsValidWidget(w) {
		return os.ErrInvalid
	}
	var optList []string
	for _, opt := range options {
		if !w.info.MetaClass.HasOption(opt.Key) {
			continue
		}
		optList = append(optList, fmt.Sprintf("-%v {%v}", opt.Key, opt.Value))
	}
	if len(optList) > 0 {
		return eval(fmt.Sprintf("%v configure %v", w.id, strings.Join(optList, " ")))
	}
	return nil
}

func (w *BaseWidget) SetWidgetOptions(options ...*WidgetOpt) error {
	if !IsValidWidget(w) {
		return os.ErrInvalid
	}
	var optList []string
	for _, opt := range options {
		if opt == nil {
			continue
		}
		if !w.info.MetaClass.HasOption(opt.Key) {
			continue
		}
		optList = append(optList, fmt.Sprintf("-%v {%v}", opt.Key, opt.Value))
	}
	if len(optList) > 0 {
		return eval(fmt.Sprintf("%v configure %v", w.id, strings.Join(optList, " ")))
	}
	return nil
}
