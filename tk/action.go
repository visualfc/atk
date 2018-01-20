// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

type Action struct {
	actid   string
	label   string
	checkid string
	groupid string
	radioid string
	fncmd   func()
	data    interface{}
}

func (a *Action) IsSeparator() bool {
	return a.actid == ""
}

func (a *Action) IsRadioAction() bool {
	return a.groupid != ""
}

func (a *Action) IsCheckAction() bool {
	return a.checkid != ""
}

func (a *Action) SetChecked(b bool) {
	if a.groupid != "" {
		eval(fmt.Sprintf("set %v {%v}", a.groupid, a.radioid))
	} else if a.checkid != "" {
		eval(fmt.Sprintf("set %v {%v}", a.checkid, boolToInt(b)))
	}
}

func (a *Action) IsChecked() bool {
	if a.groupid != "" {
		r, _ := evalAsString(fmt.Sprintf("set %v", a.groupid))
		return r == a.radioid
	} else if a.checkid != "" {
		b, _ := evalAsBool(fmt.Sprintf("set %v", a.checkid))
		return b
	}
	return false
}

func (a *Action) Label() string {
	return a.label
}

func (a *Action) SetData(data interface{}) {
	a.data = data
}

func (a *Action) Data() interface{} {
	return a.data
}

func (a *Action) OnCommand(fn func()) {
	a.fncmd = fn
}

func NewAction(label string, fn func()) *Action {
	action := &Action{}
	action.label = label
	action.fncmd = fn
	action.actid = MakeActionId()
	mainInterp.CreateAction(action.actid, func([]string) {
		if action.fncmd != nil {
			action.fncmd()
		}
	})
	return action
}

func NewCheckAction(label string, fn func()) *Action {
	action := NewAction(label, fn)
	action.checkid = MakeCustomId("atk_checkaction")
	return action
}

func NewSeparatorAction() *Action {
	action := &Action{}
	return action
}

type RadioActionGroup struct {
	groupid        string
	actions        []*Action
	fnRadioCommand func()
}

func NewActionGroup() *RadioActionGroup {
	id := MakeCustomId("atk_actiongroup")
	return &RadioActionGroup{id, nil, nil}
}

func (a *RadioActionGroup) findAction(act *Action) bool {
	for _, v := range a.actions {
		if v == act {
			return true
		}
	}
	return false
}

func (a *RadioActionGroup) radioCommand() {
	if a.fnRadioCommand != nil {
		a.fnRadioCommand()
	}
}

func (a *RadioActionGroup) AddRadioAction(act *Action) {
	if a.findAction(act) {
		return
	}
	act.groupid = a.groupid
	act.radioid = MakeCustomId("action_radio_value")
	act.fncmd = a.radioCommand
	a.actions = append(a.actions, act)
}

func (a *RadioActionGroup) AddNewRadioAction(label string) *Action {
	act := NewCheckAction(label, nil)
	act.groupid = a.groupid
	act.radioid = MakeCustomId("action_radio_value")
	act.fncmd = a.radioCommand
	a.actions = append(a.actions, act)
	return act
}

func (a *RadioActionGroup) OnCommand(fn func()) {
	a.fnRadioCommand = fn
}

func (a *RadioActionGroup) checkedValue() string {
	r, _ := evalAsString(fmt.Sprintf("set %v", a.groupid))
	return r
}

func (a *RadioActionGroup) CheckedAction() *Action {
	s := a.checkedValue()
	for _, act := range a.actions {
		if act.radioid == s {
			return act
		}
	}
	return nil
}

func (a *RadioActionGroup) Actions() []*Action {
	return a.actions
}
