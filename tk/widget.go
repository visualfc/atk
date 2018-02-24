// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"reflect"
	"strings"
)

type Widget interface {
	Id() string
	Info() *WidgetInfo
	Type() WidgetType
	TypeName() string
	Parent() Widget
	Children() []Widget
	IsValid() bool
	Destroy() error
	DestroyChildren() error
}

var (
	globalWidgetMap = make(map[string]Widget)
)

func IsNilInterface(w Widget) bool {
	if w == nil {
		return true
	}
	return reflect.ValueOf(w).IsNil()
}

func RegisterWidget(w Widget) {
	if IsNilInterface(w) {
		return
	}
	globalWidgetMap[w.Id()] = w
}

func FindWidget(id string) Widget {
	return globalWidgetMap[id]
}

func LookupWidget(id string) (w Widget, ok bool) {
	w, ok = globalWidgetMap[id]
	return
}

func ParentOfWidget(w Widget) Widget {
	if IsNilInterface(w) {
		return nil
	}
	id := w.Id()
	if id == "." {
		return nil
	}
	pos := strings.LastIndex(id, ".")
	if pos == -1 {
		return nil
	} else if pos == 0 {
		return mainWindow
	}
	return globalWidgetMap[id[:pos]]
}

func ChildrenOfWidget(w Widget) (list []Widget) {
	if IsNilInterface(w) {
		return nil
	}
	id := w.Id()
	if id == "." {
		for k, v := range globalWidgetMap {
			if strings.HasPrefix(k, id) {
				if k == "." {
					continue
				} else if strings.Index(k[1:], ".") >= 0 {
					continue
				}
				list = append(list, v)
			}
		}
	} else {
		id = id + "."
		offset := len(id)
		for k, v := range globalWidgetMap {
			if strings.HasPrefix(k, id) {
				if strings.Index(k[offset:], ".") >= 0 {
					continue
				}
				list = append(list, v)
			}
		}
	}
	return
}

func removeWidget(id string) {
	if id == "." {
		globalWidgetMap = make(map[string]Widget)
	} else {
		delete(globalWidgetMap, id)
		id = id + "."
		var list []string
		for k, _ := range globalWidgetMap {
			if strings.HasPrefix(k, id) {
				list = append(list, k)
			}
		}
		for _, k := range list {
			delete(globalWidgetMap, k)
		}
	}
}

func IsValidWidget(w Widget) bool {
	if IsNilInterface(w) {
		return false
	}
	_, ok := globalWidgetMap[w.Id()]
	return ok
}

func DestroyWidget(w Widget) error {
	if !IsValidWidget(w) {
		return ErrInvalid
	}
	id := w.Id()
	eval(fmt.Sprintf("destroy %v", id))
	removeWidget(id)
	return nil
}

func dumpWidgetHelp(w Widget, offset string, space string, ar *[]string) {
	s := fmt.Sprintf("%v%v", space, w)
	*ar = append(*ar, s)
	for _, child := range w.Children() {
		dumpWidgetHelp(child, offset, space+offset, ar)
	}
}

func DumpWidget(w Widget) string {
	return DumpWidgetEx(w, "\t")
}

func DumpWidgetEx(w Widget, offset string) string {
	var ar []string
	dumpWidgetHelp(w, offset, "", &ar)
	return strings.Join(ar, "\n")
}
