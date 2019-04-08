// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

// Create and manipulate 'text' hypertext editing widgets
type Text struct {
	BaseWidget
	xscrollcommand *CommandEx
	yscrollcommand *CommandEx
}

func NewText(parent Widget, attributes ...*WidgetAttr) *Text {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_text")
	info := CreateWidgetInfo(iid, WidgetTypeText, theme, attributes)
	if info == nil {
		return nil
	}
	w := &Text{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *Text) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeText)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *Text) SetBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -background $atk_tmp_text", w.id))
}

func (w *Text) Background() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -background", w.id))
	return r
}

func (w *Text) SetBorderWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -borderwidth {%v}", w.id, width))
}

func (w *Text) BorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -borderwidth", w.id))
	return r
}

func (w *Text) SetFont(font Font) error {
	if font == nil {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v configure -font {%v}", w.id, font.Id()))
}

func (w *Text) Font() Font {
	r, err := evalAsString(fmt.Sprintf("%v cget -font", w.id))
	return parserFontResult(r, err)
}

func (w *Text) SetForeground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -foreground $atk_tmp_text", w.id))
}

func (w *Text) Foreground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -foreground", w.id))
	return r
}

func (w *Text) SetHighlightBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -highlightbackground $atk_tmp_text", w.id))
}

func (w *Text) HighlightBackground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -highlightbackground", w.id))
	return r
}

func (w *Text) SetHighlightColor(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -highlightcolor $atk_tmp_text", w.id))
}

func (w *Text) HighlightColor() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -highlightcolor", w.id))
	return r
}

func (w *Text) SetHighlightthickness(width int) error {
	return eval(fmt.Sprintf("%v configure -highlightthickness {%v}", w.id, width))
}

func (w *Text) Highlightthickness() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -highlightthickness", w.id))
	return r
}

func (w *Text) SetInsertBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -insertbackground $atk_tmp_text", w.id))
}

func (w *Text) InsertBackground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -insertbackground", w.id))
	return r
}

func (w *Text) SetInsertBorderWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -insertborderwidth {%v}", w.id, width))
}

func (w *Text) InsertBorderWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -insertborderwidth", w.id))
	return r
}

func (w *Text) SetInsertOffTime(offtime int) error {
	return eval(fmt.Sprintf("%v configure -insertofftime {%v}", w.id, offtime))
}

func (w *Text) InsertOffTime() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -insertofftime", w.id))
	return r
}

func (w *Text) SetInsertOnTime(ontime int) error {
	return eval(fmt.Sprintf("%v configure -insertontime {%v}", w.id, ontime))
}

func (w *Text) InsertOnTime() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -insertontime", w.id))
	return r
}

func (w *Text) SetInsertWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -insertwidth {%v}", w.id, width))
}

func (w *Text) InsertWidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -insertwidth", w.id))
	return r
}

func (w *Text) SetPaddingN(padx int, pady int) error {
	if w.info.IsTtk {
		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
	}
	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
}

func (w *Text) PaddingN() (int, int) {
	var r string
	var err error
	if w.info.IsTtk {
		r, err = evalAsString(fmt.Sprintf("%v cget -padding", w.id))
	} else {
		r1, _ := evalAsString(fmt.Sprintf("%v cget -padx", w.id))
		r2, _ := evalAsString(fmt.Sprintf("%v cget -pady", w.id))
		r = r1 + " " + r2
	}
	return parserPaddingResult(r, err)
}

func (w *Text) SetPadding(pad Pad) error {
	return w.SetPaddingN(pad.X, pad.Y)
}

func (w *Text) Padding() Pad {
	x, y := w.PaddingN()
	return Pad{x, y}
}

func (w *Text) SetReliefStyle(relief ReliefStyle) error {
	return eval(fmt.Sprintf("%v configure -relief {%v}", w.id, relief))
}

func (w *Text) ReliefStyle() ReliefStyle {
	r, err := evalAsString(fmt.Sprintf("%v cget -relief", w.id))
	return parserReliefStyleResult(r, err)
}

func (w *Text) SetSelectBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -selectbackground $atk_tmp_text", w.id))
}

func (w *Text) SelectBackground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -selectbackground", w.id))
	return r
}

func (w *Text) SetSelectborderwidth(width int) error {
	return eval(fmt.Sprintf("%v configure -selectborderwidth {%v}", w.id, width))
}

func (w *Text) Selectborderwidth() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -selectborderwidth", w.id))
	return r
}

func (w *Text) SetSelectforeground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -selectforeground $atk_tmp_text", w.id))
}

func (w *Text) Selectforeground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -selectforeground", w.id))
	return r
}

func (w *Text) SetInactiveSelectBackground(color string) error {
	setObjText("atk_tmp_text", color)
	return eval(fmt.Sprintf("%v configure -inactiveselectbackground $atk_tmp_text", w.id))
}

func (w *Text) InactiveSelectBackground() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -inactiveselectbackground", w.id))
	return r
}

func (w *Text) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *Text) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *Text) SetAutoSeparatorsOnUndo(autoseparators bool) error {
	return eval(fmt.Sprintf("%v configure -autoseparators {%v}", w.id, boolToInt(autoseparators)))
}

func (w *Text) IsAutoSeparatorsOnUndo() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -autoseparators", w.id))
	return r
}

func (w *Text) SetBlockCursor(blockcursor bool) error {
	return eval(fmt.Sprintf("%v configure -blockcursor {%v}", w.id, boolToInt(blockcursor)))
}

func (w *Text) IsBlockCursor() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -blockcursor", w.id))
	return r
}

func (w *Text) SetStartLine(startline int) error {
	return eval(fmt.Sprintf("%v configure -startline {%v}", w.id, startline))
}

func (w *Text) StartLine() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -startline", w.id))
	return r
}

func (w *Text) SetEndLine(endline int) error {
	return eval(fmt.Sprintf("%v configure -endline {%v}", w.id, endline))
}

func (w *Text) EndLine() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -endline", w.id))
	return r
}

func (w *Text) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *Text) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *Text) SetHeight(height int) error {
	return eval(fmt.Sprintf("%v configure -height {%v}", w.id, height))
}

func (w *Text) Height() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -height", w.id))
	return r
}

func (w *Text) SetInsertUnfocussed(style DisplyCursor) error {
	if !mainInterp.SupportVer86() {
		return ErrUnsupport
	}
	return eval(fmt.Sprintf("%v configure -insertunfocussed {%v}", w.id, style))
}

func (w *Text) InsertUnfocussed() DisplyCursor {
	if !mainInterp.SupportVer86() {
		return DisplyCursorHollow
	}
	r, err := evalAsString(fmt.Sprintf("%v cget -insertunfocussed", w.id))
	return parserDisplyCursorResult(r, err)
}

func (w *Text) SetMaxUndo(maxundo int) error {
	return eval(fmt.Sprintf("%v configure -maxundo {%v}", w.id, maxundo))
}

func (w *Text) MaxUndo() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -maxundo", w.id))
	return r
}

func (w *Text) SetLineAboveSpace(spacing int) error {
	return eval(fmt.Sprintf("%v configure -spacing1 {%v}", w.id, spacing))
}

func (w *Text) LineAboveSpace() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -spacing1", w.id))
	return r
}

func (w *Text) SetLineWrapSpace(spacing int) error {
	return eval(fmt.Sprintf("%v configure -spacing2 {%v}", w.id, spacing))
}

func (w *Text) LineWrapSpace() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -spacing2", w.id))
	return r
}

func (w *Text) SetLineBelowSpace(spacing int) error {
	return eval(fmt.Sprintf("%v configure -spacing3 {%v}", w.id, spacing))
}

func (w *Text) LineBelowSpace() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -spacing3", w.id))
	return r
}

func (w *Text) SetLineWrap(wrap LineWrapMode) error {
	return eval(fmt.Sprintf("%v configure -wrap {%v}", w.id, wrap))
}

func (w *Text) LineWrap() LineWrapMode {
	r, err := evalAsString(fmt.Sprintf("%v cget -wrap", w.id))
	return parserLineWrapModeResult(r, err)
}

func (w *Text) SetEnableUndo(undo bool) error {
	return eval(fmt.Sprintf("%v configure -undo {%v}", w.id, boolToInt(undo)))
}

func (w *Text) IsEnableUndo() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -undo", w.id))
	return r
}

func (w *Text) SetReadOnly(b bool) error {
	var script string
	if b {
		script = fmt.Sprintf("%v configure -state disable", w.id)
	} else {
		script = fmt.Sprintf("%v configure -state normal", w.id)
	}
	return eval(script)
}

func (w *Text) IsReadOnly() bool {
	r, _ := evalAsString(fmt.Sprintf("%v cget -state", w.id))
	return r != "normal"
}

type TextIndex struct {
	index string
}

func (t TextIndex) String() string {
	return t.index
}

func (w *Text) TextLength() int {
	r, _ := evalAsInt(fmt.Sprintf("%v count -chars 1.0 end", w.id))
	return r
}

func (w *Text) PlainText() string {
	r, _ := evalAsString(fmt.Sprintf("%v get 1.0 end", w.id))
	return r
}

func (w *Text) LineCount() int {
	r, _ := evalAsInt(fmt.Sprintf("%v count -lines 1.0 end", w.id))
	return r
}

func (w *Text) InsertText(pos int, text string) error {
	setObjText("atk_text_insert", text)
	return eval(fmt.Sprintf("%v insert {0.0 + %v chars} $atk_text_insert", w.id, pos))
}

func (w *Text) AppendText(text string) error {
	setObjText("atk_text_insert", text)
	return eval(fmt.Sprintf("%v insert end $atk_text_insert", w.id))
}

func (w *Text) Length() int {
	r, _ := evalAsInt(fmt.Sprintf("%v index end", w.id))
	return r
}

func (w *Text) Clear() error {
	return eval(fmt.Sprintf("%v delete 1.0 end", w.id))
}

func (w *Text) SetText(text string) error {
	setObjText("atk_text_insert", text)
	return eval(fmt.Sprintf("%v delete 1.0 end\n%v insert end $atk_text_insert", w.id, w.id))
}

func (w *Text) SetTabSize(size int) error {
	return eval(fmt.Sprintf("%v configure -tabs %v", w.id, size))
}

func (w *Text) SetXViewArgs(args []string) error {
	return eval(fmt.Sprintf("%v xview %v", w.id, strings.Join(args, " ")))
}

func (w *Text) SetYViewArgs(args []string) error {
	return eval(fmt.Sprintf("%v yview %v", w.id, strings.Join(args, " ")))
}

func (w *Text) OnXScrollEx(fn func([]string) error) error {
	if fn == nil {
		return ErrInvalid
	}
	if w.xscrollcommand == nil {
		w.xscrollcommand = &CommandEx{}
		bindCommandEx(w.id, "xscrollcommand", w.xscrollcommand.Invoke)
	}
	w.xscrollcommand.Bind(fn)
	return nil
}

func (w *Text) OnYScrollEx(fn func([]string) error) error {
	if fn == nil {
		return ErrInvalid
	}
	if w.yscrollcommand == nil {
		w.yscrollcommand = &CommandEx{}
		bindCommandEx(w.id, "yscrollcommand", w.yscrollcommand.Invoke)
	}
	w.yscrollcommand.Bind(fn)
	return nil
}

func (w *Text) BindXScrollBar(bar *ScrollBar) error {
	if !IsValidWidget(bar) {
		return ErrInvalid
	}
	w.OnXScrollEx(bar.SetScrollArgs)
	bar.OnCommandEx(w.SetXViewArgs)
	return nil
}

func (w *Text) BindYScrollBar(bar *ScrollBar) error {
	if !IsValidWidget(bar) {
		return ErrInvalid
	}
	w.OnYScrollEx(bar.SetScrollArgs)
	bar.OnCommandEx(w.SetYViewArgs)
	return nil
}

type TextEx struct {
	*ScrollLayout
	*Text
}

func NewTextEx(parent Widget, attributs ...*WidgetAttr) *TextEx {
	w := &TextEx{}
	w.ScrollLayout = NewScrollLayout(parent)
	w.Text = NewText(parent, attributs...)
	w.SetWidget(w.Text)
	w.Text.BindXScrollBar(w.XScrollBar)
	w.Text.BindYScrollBar(w.YScrollBar)
	RegisterWidget(w)
	return w
}

func TextAttrBackground(color string) *WidgetAttr {
	return &WidgetAttr{"background", color}
}

func TextAttrBorderWidth(width int) *WidgetAttr {
	return &WidgetAttr{"borderwidth", width}
}

func TextAttrFont(font Font) *WidgetAttr {
	if font == nil {
		return nil
	}
	return &WidgetAttr{"font", font.Id()}
}

func TextAttrForeground(color string) *WidgetAttr {
	return &WidgetAttr{"foreground", color}
}

func TextAttrHighlightBackground(color string) *WidgetAttr {
	return &WidgetAttr{"highlightbackground", color}
}

func TextAttrHighlightColor(color string) *WidgetAttr {
	return &WidgetAttr{"highlightcolor", color}
}

func TextAttrHighlightthickness(width int) *WidgetAttr {
	return &WidgetAttr{"highlightthickness", width}
}

func TextAttrInsertBackground(color string) *WidgetAttr {
	return &WidgetAttr{"insertbackground", color}
}

func TextAttrInsertBorderWidth(width int) *WidgetAttr {
	return &WidgetAttr{"insertborderwidth", width}
}

func TextAttrInsertOffTime(offtime int) *WidgetAttr {
	return &WidgetAttr{"insertofftime", offtime}
}

func TextAttrInsertOnTime(ontime int) *WidgetAttr {
	return &WidgetAttr{"insertontime", ontime}
}

func TextAttrInsertWidth(width int) *WidgetAttr {
	return &WidgetAttr{"insertwidth", width}
}

func TextAttrPadding(padding Pad) *WidgetAttr {
	return &WidgetAttr{"padding", padding}
}

func TextAttrReliefStyle(relief ReliefStyle) *WidgetAttr {
	return &WidgetAttr{"relief", relief}
}

func TextAttrSelectBackground(color string) *WidgetAttr {
	return &WidgetAttr{"selectbackground", color}
}

func TextAttrSelectborderwidth(width int) *WidgetAttr {
	return &WidgetAttr{"selectborderwidth", width}
}

func TextAttrSelectforeground(color string) *WidgetAttr {
	return &WidgetAttr{"selectforeground", color}
}

func TextAttrInactiveSelectBackground(color string) *WidgetAttr {
	return &WidgetAttr{"inactiveselectbackground", color}
}

func TextAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}

func TextAttrAutoSeparatorsOnUndo(autoseparators bool) *WidgetAttr {
	return &WidgetAttr{"autoseparators", boolToInt(autoseparators)}
}

func TextAttrBlockCursor(blockcursor bool) *WidgetAttr {
	return &WidgetAttr{"blockcursor", boolToInt(blockcursor)}
}

func TextAttrStartLine(startline int) *WidgetAttr {
	return &WidgetAttr{"startline", startline}
}

func TextAttrEndLine(endline int) *WidgetAttr {
	return &WidgetAttr{"endline", endline}
}

func TextAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func TextAttrHeight(height int) *WidgetAttr {
	return &WidgetAttr{"height", height}
}

func TextAttrInsertUnfocussed(style DisplyCursor) *WidgetAttr {
	if !mainInterp.SupportVer86() {
		return nil
	}
	return &WidgetAttr{"insertunfocussed", style}
}

func TextAttrMaxUndo(maxundo int) *WidgetAttr {
	return &WidgetAttr{"maxundo", maxundo}
}

func TextAttrLineAboveSpace(spacing int) *WidgetAttr {
	return &WidgetAttr{"spacing1", spacing}
}

func TextAttrLineWrapSpace(spacing int) *WidgetAttr {
	return &WidgetAttr{"spacing2", spacing}
}

func TextAttrLineBelowSpace(spacing int) *WidgetAttr {
	return &WidgetAttr{"spacing3", spacing}
}

func TextAttrLineWrap(wrap LineWrapMode) *WidgetAttr {
	return &WidgetAttr{"wrap", wrap}
}

func TextAttrEnableUndo(undo bool) *WidgetAttr {
	return &WidgetAttr{"undo", boolToInt(undo)}
}
