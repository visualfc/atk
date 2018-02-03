// Copyright 2018 visualfc. All rights reserved.

package interp

import (
	"errors"
	"strconv"
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
	TCL_DONT_WAIT     = 1 << 1
	TCL_WINDOW_EVENTS = 1 << 2
	TCL_FILE_EVENTS   = 1 << 3
	TCL_TIMER_EVENTS  = 1 << 4
	TCL_IDLE_EVENTS   = 1 << 5
	TCL_ALL_EVENTS    = ^TCL_DONT_WAIT
)

var (
	globalCommandMap = NewCommandMap()
	globalActionMap  = NewActionMap()
)

type ActionMap struct {
	fnMap map[uintptr]func([]string)
	id    uintptr
}

func NewActionMap() *ActionMap {
	return &ActionMap{make(map[uintptr]func([]string)), 1}
}

func (m *ActionMap) Register(fn func([]string)) uintptr {
	m.id = m.id + 1
	m.fnMap[m.id] = fn
	return m.id
}

func (m *ActionMap) UnRegister(id uintptr) {
	delete(m.fnMap, id)
}

func (m *ActionMap) Invoke(id uintptr, args []string) error {
	fn, ok := m.fnMap[id]
	if !ok {
		return errors.New("Not found action")
	}
	fn(args)
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
		return "", errors.New("Not found command")
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

func (interp *Interp) EvalAsBool(script string) (bool, error) {
	err := interp.Eval(script)
	if err != nil {
		return false, err
	}
	return interp.GetBoolResult(), nil
}

func (interp *Interp) EvalAsObj(script string) (*Obj, error) {
	err := interp.Eval(script)
	if err != nil {
		return nil, err
	}
	return interp.GetObjResult(), nil
}

func (interp *Interp) EvalAsListObj(script string) (*ListObj, error) {
	err := interp.Eval(script)
	if err != nil {
		return nil, err
	}
	return interp.GetListObjResult(), nil
}

func (interp *Interp) EvalAsStringList(script string) ([]string, error) {
	err := interp.Eval(script)
	if err != nil {
		return nil, err
	}
	return interp.GetListObjResult().ToStringList(), nil
}

func (interp *Interp) EvalAsIntList(script string) ([]int, error) {
	err := interp.Eval(script)
	if err != nil {
		return nil, err
	}
	return interp.GetListObjResult().ToIntList(), nil
}

func (interp *Interp) TclVersion() string {
	ver, _ := interp.EvalAsString("set tcl_version")
	return ver
}

func (interp *Interp) TkVersion() string {
	ver, _ := interp.EvalAsString("set tk_version")
	return ver
}

func (p *Interp) GetStringResult() string {
	return p.GetObjResult().ToString()
}

func (p *Interp) GetIntResult() int {
	return p.GetObjResult().ToInt()
}

func (p *Interp) GetInt64Result() int64 {
	return p.GetObjResult().ToInt64()
}

func (p *Interp) GetFloat64Result() float64 {
	return p.GetObjResult().ToFloat64()
}

func (p *Interp) GetBoolResult() bool {
	return p.GetObjResult().ToBool()
}

func (p *Interp) GetErrorResult() error {
	return errors.New(p.GetObjResult().ToString())
}

func (p *Interp) GetStringVar(name string, global bool) string {
	obj := p.GetVar(name, global)
	if obj == nil {
		return ""
	}
	return obj.ToString()
}

func (p *Interp) GetIntVar(name string, global bool) int {
	obj := p.GetVar(name, global)
	if obj == nil {
		return 0
	}
	return obj.ToInt()
}

func (p *Interp) GetInt64Var(name string, global bool) int64 {
	obj := p.GetVar(name, global)
	if obj == nil {
		return 0
	}
	return obj.ToInt64()
}

func (p *Interp) GetFloadt64Var(name string, global bool) float64 {
	obj := p.GetVar(name, global)
	if obj == nil {
		return 0
	}
	return obj.ToFloat64()
}

func (p *Interp) GetBoolVar(name string, global bool) bool {
	obj := p.GetVar(name, global)
	if obj == nil {
		return false
	}
	return obj.ToBool()
}

func (p *Interp) SetIntVar(name string, value int, global bool) error {
	return p.SetStringVar(name, strconv.Itoa(value), global)
}

func (p *Interp) SetInt64Var(name string, value int64, global bool) error {
	return p.SetStringVar(name, strconv.FormatInt(value, 10), global)
}

func (p *Interp) SetFloat64Var(name string, value float64, global bool) error {
	return p.SetStringVar(name, strconv.FormatFloat(value, 'E', -1, 64), global)
}

func (p *Interp) SetBoolVar(name string, b bool, global bool) error {
	if b {
		return p.SetStringVar(name, "1", global)
	}
	return p.SetStringVar(name, "0", global)
}
