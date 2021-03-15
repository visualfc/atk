// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
)

type ttkTheme struct {
}

func (t *ttkTheme) Name() string {
	return "ttk theme"
}

func (t *ttkTheme) IsTtk() bool {
	return true
}

func (t *ttkTheme) InitAttributes(typ WidgetType) []NativeAttr {
	return nil
}

func (t *ttkTheme) ThemeIdList() []string {
	return ttk_theme_list
}

func (t *ttkTheme) SetThemeId(id string) error {
	for _, v := range ttk_theme_list {
		if v == id {
			err := eval(fmt.Sprintf("ttk::setTheme %v", id))
			return err
		}
	}
	err := fmt.Errorf("not found ttk_theme id:%v", id)
	dumpError(err)
	return err
}

func (t *ttkTheme) ThemeId() string {
	r, _ := evalAsString("ttk::style theme use")
	return r
}

var (
	ttk_theme_list []string
	TtkTheme       = &ttkTheme{}
)

func (w *BaseWidget) StyleName() string {
    // ttk::button .b; winfo class .b  // ==> TButton
	r, _ := evalAsString(fmt.Sprintf("winfo class %v", w.id))
	return r
}

func (w *BaseWidget) StyleLookUp(name, option string) string {
    // ttk::style lookup style -option
    // ttk::style lookup 1.TButton -font  // [---> helvetica 24]
    r1, _ := evalAsString(fmt.Sprintf("ttk::style lookup %v -%v", name,option))
    return r1
}


func StyleConfigure(name string, options map[string]string) error {
    // ttk::style configure style ?-option ?value option value...? ?
    // ttk::style configure Emergency.TButton -foreground red -padding 10
    // ttk::button .b -text "Hello" -style "Fun.TButton"
    var tmp = ""
    for k,v := range options {
        tmp = tmp + "-" + k + " " + v + " "
    }
    
    return eval(fmt.Sprintf("ttk::style configure %v %v", name,tmp))
}

func StyleMap(name string, options map[string]map[string]string) error{
    // ttk::style map style ?-option { statespec value... }?
    // ttk::style map TRadiobutton -foreground [list !pressed blue pressed yellow] -background [list selected black !selected white]
    var tmp1 = ""
    var tmp2 = ""
    for k1,v1 := range options {
        tmp2 = "[list "
        tmp1 = tmp1 + "-" + k1 + " "
        for k2,v2 := range v1 {
            tmp2 = tmp2 + k2 + " " + v2 + " "
        }
        tmp2 = tmp2 + "] "
        tmp1 = tmp1 + tmp2
    }
    // fmt.Println(fmt.Sprintf("ttk::style map %v %v", name,tmp1))
    return eval(fmt.Sprintf("ttk::style map %v %v", name,tmp1))
}

// 参考：
// https://tkdocs.com/tutorial/styles.html
// http://www.tcl-lang.org/man/tcl8.6/TkCmd/ttk_style.htm
// https://tkdocs.com/shipman/ttk-map.html

func init() {
	registerInit(func() {
		ttk_theme_list, _ = evalAsStringList("ttk::themes")
	})
	SetMainTheme(TtkTheme)
}
