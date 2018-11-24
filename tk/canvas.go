// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// Create and manipulate 'canvas' hypergraphics drawing surface widgets
type Canvas struct {
	BaseWidget
	xscrollcommand *CommandEx
	yscrollcommand *CommandEx
}

func NewCanvas(parent Widget, attributes ...*WidgetAttr) *Canvas {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_canvas")
	info := CreateWidgetInfo(iid, WidgetTypeCanvas, theme, attributes)
	if info == nil {
		return nil
	}
	w := &Canvas{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *Canvas) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeCanvas)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *Canvas) SetBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -background $atk_tmp_text", w.id))
}

func (w *Canvas) Background() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -background", w.id))
	return r
}

func (w *Canvas) SetBorderWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -borderwidth {%v}", w.id, width))
}

func (w *Canvas) BorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.id))
	return r
}

func (w *Canvas) SetHighlightBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -highlightbackground $atk_tmp_text", w.id))
}

func (w *Canvas) HighlightBackground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -highlightbackground", w.id))
	return r
}

func (w *Canvas) SetHighlightColor(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -highlightcolor $atk_tmp_text", w.id))
}

func (w *Canvas) HighlightColor() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -highlightcolor", w.id))
	return r
}

func (w *Canvas) SetHighlightthickness(width int) error {
	return eval(fmt.Sprintf("%v configure -highlightthickness {%v}", w.id, width))
}

func (w *Canvas) Highlightthickness() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -highlightthickness", w.id))
	return r
}

func (w *Canvas) SetInsertBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -insertbackground $atk_tmp_text", w.id))
}

func (w *Canvas) InsertBackground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -insertbackground", w.id))
	return r
}

func (w *Canvas) SetInsertBorderWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -insertborderwidth {%v}", w.id, width))
}

func (w *Canvas) InsertBorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -insertborderwidth", w.id))
	return r
}

func (w *Canvas) SetInsertOffTime(offtime int) error {
	return eval(fmt.Sprintf("%v configure -insertofftime {%v}", w.id, offtime))
}

func (w *Canvas) InsertOffTime() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -insertofftime", w.id))
	return r
}

func (w *Canvas) SetInsertOnTime(ontime int) error {
	return eval(fmt.Sprintf("%v configure -insertontime {%v}", w.id, ontime))
}

func (w *Canvas) InsertOnTime() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -insertontime", w.id))
	return r
}

func (w *Canvas) SetInsertWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -insertwidth {%v}", w.id, width))
}

func (w *Canvas) InsertWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -insertwidth", w.id))
	return r
}

func (w *Canvas) SetReliefStyle(relief ReliefStyle) error {
	return eval(fmt.Sprintf("%v configure -relief {%v}", w.id, relief))
}

func (w *Canvas) ReliefStyle() ReliefStyle {
	r, err := evalAsString(fmt.Sprintf("%v cget -relief", w.id))
	return parserReliefStyleResult(r, err)
}

func (w *Canvas) SetSelectBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -selectbackground $atk_tmp_text", w.id))
}

func (w *Canvas) SelectBackground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -selectbackground", w.id))
	return r
}

func (w *Canvas) SetSelectborderwidth(width int) error {
	return eval(fmt.Sprintf("%v configure -selectborderwidth {%v}", w.id, width))
}

func (w *Canvas) Selectborderwidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -selectborderwidth", w.id))
	return r
}

func (w *Canvas) SetSelectforeground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -selectforeground $atk_tmp_text", w.id))
}

func (w *Canvas) Selectforeground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -selectforeground", w.id))
	return r
}

func (w *Canvas) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *Canvas) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *Canvas) SetCloseEnough(closeenough float64) error {
	return eval(fmt.Sprintf("%v configure -closeenough {%v}", w.id, closeenough))
}

func (w *Canvas) CloseEnough() float64 {
	r, _ := evalAsFloat64(fmt.Sprintf("%v cget -closeenough", w.id))
	return r
}

func (w *Canvas) SetConfine(confine bool) error {
	return eval(fmt.Sprintf("%v configure -confine {%v}", w.id, boolToInt(confine)))
}

func (w *Canvas) IsConfine() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -confine", w.id))
	return r
}

func (w *Canvas) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *Canvas) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *Canvas) SetHeight(height int) error {
	return eval(fmt.Sprintf("%v configure -height {%v}", w.id, height))
}

func (w *Canvas) Height() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -height", w.id))
	return r
}

func (w *Canvas) SetState(state State) error {
	return eval(fmt.Sprintf("%v configure -state {%v}", w.id, state))
}

func (w *Canvas) State() State {
	r, err := evalAsString(fmt.Sprintf("%v cget -state", w.id))
	return parserStateResult(r, err)
}

func (w *Canvas) SetXScrollIncrement(value int) error {
	return eval(fmt.Sprintf("%v configure -xscrollincrement {%v}", w.id, value))
}

func (w *Canvas) XScrollIncrement() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -xscrollincrement", w.id))
	return r
}

func (w *Canvas) SetYScrollIncrement(value int) error {
	return eval(fmt.Sprintf("%v configure -yscrollincrement {%v}", w.id, value))
}

func (w *Canvas) YScrollIncrement() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -yscrollincrement", w.id))
	return r
}

func CanvasAttrBackground(color string) *WidgetAttr {
	return &WidgetAttr{"background", color}
}

func CanvasAttrBorderWidth(width int) *WidgetAttr {
	return &WidgetAttr{"borderwidth", width}
}

func CanvasAttrHighlightBackground(color string) *WidgetAttr {
	return &WidgetAttr{"highlightbackground", color}
}

func CanvasAttrHighlightColor(color string) *WidgetAttr {
	return &WidgetAttr{"highlightcolor", color}
}

func CanvasAttrHighlightthickness(width int) *WidgetAttr {
	return &WidgetAttr{"highlightthickness", width}
}

func CanvasAttrInsertBackground(color string) *WidgetAttr {
	return &WidgetAttr{"insertbackground", color}
}

func CanvasAttrInsertBorderWidth(width int) *WidgetAttr {
	return &WidgetAttr{"insertborderwidth", width}
}

func CanvasAttrInsertOffTime(offtime int) *WidgetAttr {
	return &WidgetAttr{"insertofftime", offtime}
}

func CanvasAttrInsertOnTime(ontime int) *WidgetAttr {
	return &WidgetAttr{"insertontime", ontime}
}

func CanvasAttrInsertWidth(width int) *WidgetAttr {
	return &WidgetAttr{"insertwidth", width}
}

func CanvasAttrReliefStyle(relief ReliefStyle) *WidgetAttr {
	return &WidgetAttr{"relief", relief}
}

func CanvasAttrSelectBackground(color string) *WidgetAttr {
	return &WidgetAttr{"selectbackground", color}
}

func CanvasAttrSelectborderwidth(width int) *WidgetAttr {
	return &WidgetAttr{"selectborderwidth", width}
}

func CanvasAttrSelectforeground(color string) *WidgetAttr {
	return &WidgetAttr{"selectforeground", color}
}

func CanvasAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}

func CanvasAttrCloseEnough(closeenough float64) *WidgetAttr {
	return &WidgetAttr{"closeenough", closeenough}
}

func CanvasAttrConfine(confine bool) *WidgetAttr {
	return &WidgetAttr{"confine", boolToInt(confine)}
}

func CanvasAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func CanvasAttrHeight(height int) *WidgetAttr {
	return &WidgetAttr{"height", height}
}

func CanvasAttrState(state State) *WidgetAttr {
	return &WidgetAttr{"state", state}
}

func CanvasAttrXScrollIncrement(value int) *WidgetAttr {
	return &WidgetAttr{"xscrollincrement", value}
}

func CanvasAttrYScrollIncrement(value int) *WidgetAttr {
	return &WidgetAttr{"yscrollincrement", value}
}
