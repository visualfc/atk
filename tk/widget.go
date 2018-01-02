// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"os"
	"strings"
)

type Widget interface {
	Id() string
	Type() string
	Parent() Widget
	Children() []Widget
	IsValid() bool
	Destroy() error
	DestroyChildren() error
}

type BaseWidget struct {
	id string
}

func (w *BaseWidget) String() string {
	iw := globalWidgetMap[w.id]
	if iw != nil {
		return fmt.Sprintf("%v{%v}", iw.Type(), w.id)
	} else {
		return fmt.Sprintf("Widget{%v}", w.id)
	}
}

func (w *BaseWidget) SetInternalId(id string) {
	w.id = id
}

func (w *BaseWidget) Id() string {
	return w.id
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

var (
	globalWidgetMap = make(map[string]Widget)
)

func RegisterWidget(w Widget) {
	if w == nil {
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
	if w == nil {
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
	if w == nil {
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
	if w == nil {
		return false
	}
	_, ok := globalWidgetMap[w.Id()]
	return ok
}

func DestroyWidget(w Widget) error {
	if !IsValidWidget(w) {
		return os.ErrInvalid
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

func DumpWidget(w Widget, offset string) string {
	var ar []string
	dumpWidgetHelp(w, offset, "", &ar)
	return strings.Join(ar, "\n")
}
