// Copyright 2017 visualfc. All rights reserved.

package interp

import (
	"errors"
	"sync"
)

const (
	TCL_OK    = 0
	TCL_ERROR = 1
)

type Tcl_QueuePosition int

const (
	TCL_QUEUE_TAIL Tcl_QueuePosition = 0
	TCL_QUEUE_HEAD
	TCL_QUEUE_MARK
)

const (
	TCL_DONT_WAIT     = (1 << 1)
	TCL_WINDOW_EVENTS = 1 << 2
	TCL_FILE_EVENTS   = 1 << 3
	TCL_TIMER_EVENTS  = 1 << 4
	TCL_IDLE_EVENTS   = 1 << 5
	TCL_ALL_EVENTS    = ^TCL_DONT_WAIT
)

var (
	globalAsyncEvent sync.Map
	globalCommandMap = NewCommandMap()
	globalActionMap  = NewActionMap()
)

type ActionMap struct {
	fnMap map[uintptr]func()
	id    uintptr
}

func NewActionMap() *ActionMap {
	return &ActionMap{make(map[uintptr]func()), 1}
}

func (m *ActionMap) Register(fn func()) uintptr {
	m.id = m.id + 1
	m.fnMap[m.id] = fn
	return m.id
}

func (m *ActionMap) UnRegister(id uintptr) {
	delete(m.fnMap, id)
}

func (m *ActionMap) Invoke(id uintptr) error {
	fn, ok := m.fnMap[id]
	if !ok {
		return errors.New("Not found action")
	}
	fn()
	return nil
}

type CommandMap struct {
	fnMap map[uintptr]func([]string) (string, error)
	id    uintptr
}

func (m *CommandMap) Register(fn func([]string) (string, error)) uintptr {
	m.id = m.id + 1
	m.fnMap[m.id] = fn
	return m.id
}

func (m *CommandMap) UnRegister(id uintptr) {
	delete(m.fnMap, id)
}

func (m *CommandMap) Find(id uintptr) func([]string) (string, error) {
	return m.fnMap[id]
}

func (m *CommandMap) Invoke(id uintptr, args []string) (string, error) {
	fn, ok := m.fnMap[id]
	if !ok {
		return "", errors.New("Not found action")
	}
	return fn(args)
}

func NewCommandMap() *CommandMap {
	return &CommandMap{make(map[uintptr]func([]string) (string, error)), 1}
}

func (interp *Interp) EvalAsString(script string) (string, error) {
	err := interp.Eval(script)
	if err != nil {
		return "", err
	}
	return interp.GetStringResult(), nil
}

func (interp *Interp) EvalAsInt64(script string) (int64, error) {
	err := interp.Eval(script)
	if err != nil {
		return 0, err
	}
	return interp.GetInt64Result(), nil
}

func (interp *Interp) EvalAsInt(script string) (int, error) {
	err := interp.Eval(script)
	if err != nil {
		return 0, err
	}
	return interp.GetIntResult(), nil
}

func (interp *Interp) EvalAsFloat64(script string) (float64, error) {
	err := interp.Eval(script)
	if err != nil {
		return 0, err
	}
	return interp.GetFloat64Result(), nil
}
