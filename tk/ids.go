// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

func NewGenInt64Func(id int64) func() <-chan int64 {
	ch := make(chan int64)
	go func(i int64) {
		for {
			i++
			ch <- i
		}
	}(id)
	return func() <-chan int64 {
		return ch
	}
}

func NewGenIntFunc(id int) func() <-chan int {
	ch := make(chan int)
	go func(i int) {
		for {
			i++
			ch <- i
		}
	}(id)
	return func() <-chan int {
		return ch
	}
}

var (
	fnGenFontId   = NewGenIntFunc(0)
	fnGenActionId = NewGenIntFunc(0)
	fnGenWidgetId = NewGenIntFunc(0)
	fnGetWindowId = NewGenIntFunc(0)
)

func MakeActionId() string {
	return fmt.Sprintf("action_%v", <-fnGenActionId())
}

func MakeWindowId(parent Widget, id string) string {
	if len(id) == 0 {
		id = fmt.Sprintf("window_%v", <-fnGetWindowId())
	} else if id[0] == '.' {
		return id
	}
	id = strings.ToLower(id)
	id = strings.Replace(id, " ", "_", -1)
	if parent != nil {
		pid := parent.Id()
		if pid == "." {
			return "." + id
		} else {
			return parent.Id() + "." + id
		}
	}
	return "." + id
}

func MakeWidgetId(parent Widget, id string) string {
	if len(id) == 0 {
		id = fmt.Sprintf("widget_%v", <-fnGenWidgetId())
	} else if id[0] == '.' {
		return id
	}
	id = strings.ToLower(id)
	id = strings.Replace(id, " ", "_", -1)
	if parent != nil {
		pid := parent.Id()
		if pid == "." {
			return "." + id
		} else {
			return parent.Id() + "." + id
		}
	}
	return "." + id
}

func MakeFontId() string {
	return fmt.Sprintf(".font_%v", <-fnGenActionId())
}
