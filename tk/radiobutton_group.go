// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
)

type _RadioData struct {
	btn   *RadioButton
	value string
	data  interface{}
}

type RadioGroup struct {
	id      string
	rds     []*_RadioData
	command *Command
}

func NewRadioGroup() *RadioGroup {
	id := makeNamedId("atk_radiogroup")
	evalSetValue(id, "")
	return &RadioGroup{id, nil, &Command{}}
}

func (w *RadioGroup) IsValid() bool {
	return w.id != ""
}

func (w *RadioGroup) findRadio(btn *RadioButton) *_RadioData {
	for _, rd := range w.rds {
		if rd.btn == btn {
			return rd
		}
	}
	return nil
}

func (w *RadioGroup) findRadioByValue(value string) *_RadioData {
	for _, rd := range w.rds {
		if rd.value == value {
			return rd
		}
	}
	return nil
}

func (w *RadioGroup) AddNewRadio(parent Widget, text string, data interface{}, attributes ...*WidgetAttr) *RadioButton {
	btn := NewRadioButton(parent, text, attributes...)
	w.AddRadio(btn, data)
	return btn
}

func (w *RadioGroup) AddRadios(btns ...*RadioButton) error {
	for _, btn := range btns {
		w.AddRadio(btn, nil)
	}
	return nil
}

func (w *RadioGroup) AddRadio(btn *RadioButton, data interface{}) error {
	if w.findRadio(btn) != nil {
		return ErrExist
	}
	if !IsValidWidget(btn) {
		return ErrInvalid
	}
	value := makeNamedId("atk_radiovalue")
	err := eval(fmt.Sprintf("%v configure -variable {%v} -value {%v}", btn.Id(), w.id, value))
	if err != nil {
		return err
	}
	w.rds = append(w.rds, &_RadioData{btn, value, data})
	btn.OnCommand(w.command.Invoke)
	return nil
}

func (w *RadioGroup) SetRadioData(btn *RadioButton, data interface{}) error {
	rd := w.findRadio(btn)
	if rd == nil {
		return ErrNotExist
	}
	rd.data = data
	return nil
}

func (w *RadioGroup) RadioList() (lst []*RadioButton) {
	for _, v := range w.rds {
		lst = append(lst, v.btn)
	}
	return
}

func (w *RadioGroup) WidgetList() (lst []Widget) {
	for _, v := range w.rds {
		lst = append(lst, v.btn)
	}
	return
}

func (w *RadioGroup) SetCheckedRadio(btn *RadioButton) error {
	rd := w.findRadio(btn)
	if rd == nil {
		return ErrInvalid
	}
	evalSetValue(w.id, rd.value)
	return nil
}

func (w *RadioGroup) CheckedRadio() *RadioButton {
	s := w.checkedValue()
	rd := w.findRadioByValue(s)
	if rd != nil {
		return rd.btn
	}
	return nil
}

func (w *RadioGroup) SetCheckedIndex(index int) error {
	if index < 0 || index > len(w.rds) {
		return ErrInvalid
	}
	evalSetValue(w.id, w.rds[index].value)
	return nil
}

func (w *RadioGroup) CheckedIndex() int {
	s := w.checkedValue()
	for n, rd := range w.rds {
		if rd.value == s {
			return n
		}
	}
	return -1
}

func (w *RadioGroup) checkedValue() string {
	return evalGetValue(w.id)
}

func (w *RadioGroup) CheckedData() interface{} {
	s := w.checkedValue()
	rd := w.findRadioByValue(s)
	if rd != nil {
		return rd.data
	}
	return nil
}

func (w *RadioGroup) RadioData(btn *RadioButton) interface{} {
	rd := w.findRadio(btn)
	if rd == nil {
		return nil
	}
	return rd.data
}

func (w *RadioGroup) OnRadioChanged(fn func()) error {
	if fn == nil {
		return ErrInvalid
	}
	w.command.Bind(fn)
	return nil
}
