// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"sync"
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

type NamedId interface {
	GetId(name string) string
}

type baseNamedId struct {
	m map[string]int
}

func (m *baseNamedId) GetId(name string) (r string) {
	m.m[name]++
	r = fmt.Sprintf("%v%v", name, m.m[name])
	return
}

type safeNamedId struct {
	sync.Mutex
	m map[string]int
}

func (m *safeNamedId) GetId(name string) (r string) {
	m.Lock()
	m.m[name]++
	r = fmt.Sprintf("%v%v", name, m.m[name])
	m.Unlock()
	return
}

func NewNamedId(safe bool) NamedId {
	if safe {
		return &safeNamedId{m: make(map[string]int)}
	}
	return &baseNamedId{make(map[string]int)}
}

var (
	atkNamedId = NewNamedId(false)
)

func makeNamedId(name string) string {
	return atkNamedId.GetId(name)
}

func makeNamedWidgetId(parent Widget, typ string) string {
	if parent == nil || parent.Id() == "." {
		return makeNamedId("." + typ)
	}
	return makeNamedId(parent.Id() + "." + typ)
}

func makeActionId() string {
	return makeNamedId("atk_action")
}

func makeBindEventId() string {
	return makeNamedId("atk_bindevent")
}

func makeTreeItemId(treeid string, pid string) string {
	if pid != "" {
		return makeNamedId(pid + ".I")
	}
	return makeNamedId(treeid + ".I")
}

func variableId(id string) string {
	return "::atk" + id + "_variable"
}

func evalSetValue(id string, value string) error {
	return eval(fmt.Sprintf("set %v {%v}", id, value))
}

func evalGetValue(id string) string {
	r, _ := evalAsString(fmt.Sprintf("set %v", id))
	return r
}

func traceVariable(id string, fn func()) error {
	act := makeActionId()
	mainInterp.CreateAction(act, func(args []string) {
		if fn != nil {
			fn()
		}
	})
	return eval(fmt.Sprintf("trace add variable %v write %v", id, act))
}
