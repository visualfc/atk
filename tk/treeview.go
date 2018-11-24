// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"strings"
)

// treeview
type TreeView struct {
	BaseWidget
	xscrollcommand *CommandEx
	yscrollcommand *CommandEx
}

func NewTreeView(parent Widget, attributes ...*WidgetAttr) *TreeView {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_treeview")
	info := CreateWidgetInfo(iid, WidgetTypeTreeView, theme, attributes)
	if info == nil {
		return nil
	}
	w := &TreeView{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *TreeView) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeTreeView)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *TreeView) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *TreeView) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *TreeView) SetHeight(row int) error {
	return eval(fmt.Sprintf("%v configure -height {%v}", w.id, row))
}

func (w *TreeView) Height() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -height", w.id))
	return r
}

func (w *TreeView) SetPaddingN(padx int, pady int) error {
	if w.info.IsTtk {
		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
	}
	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
}

func (w *TreeView) PaddingN() (int, int) {
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

func (w *TreeView) SetPadding(pad Pad) error {
	return w.SetPaddingN(pad.X, pad.Y)
}

func (w *TreeView) Padding() Pad {
	x, y := w.PaddingN()
	return Pad{x, y}
}

func (w *TreeView) SetTreeSelectMode(mode TreeSelectMode) error {
	return eval(fmt.Sprintf("%v configure -selectmode {%v}", w.id, mode))
}

func (w *TreeView) TreeSelectMode() TreeSelectMode {
	r, err := evalAsString(fmt.Sprintf("%v cget -selectmode", w.id))
	return parserTreeSelectModeResult(r, err)
}

func (w *TreeView) SetHeaderHidden(hide bool) error {
	var value string
	if hide {
		value = "tree"
	} else {
		value = "tree headings"
	}
	return eval(fmt.Sprintf("%v configure -show {%v}", w.id, value))
}

func (w *TreeView) IsHeaderHidden() bool {
	r, _ := evalAsString(fmt.Sprintf("%v cget -show", w.id))
	return r == "tree"
}

func (w *TreeView) SetColumnCount(columns int) error {
	if columns < 1 {
		return ErrInvalid
	}
	columns--
	if columns < 1 {
		return nil
	}
	var ids []string
	for i := 0; i < columns; i++ {
		ids = append(ids, fmt.Sprintf("column%v", i))
	}
	return eval(fmt.Sprintf("%v configure -columns {%v}", w.id, strings.Join(ids, " ")))
}

func (w *TreeView) ColumnCount() int {
	list, err := evalAsStringList(fmt.Sprintf("%v cget -columns", w.id))
	if err != nil {
		return 1
	}
	return len(list) + 1
}

func (w *TreeView) SetHeaderLabels(labels []string) error {
	for n, label := range labels {
		setObjText("atk_heading_label", label)
		err := eval(fmt.Sprintf("%v heading #%v -text $atk_heading_label", w.id, n))
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *TreeView) SetHeaderLabel(column int, label string) error {
	setObjText("atk_heading_label", label)
	return eval(fmt.Sprintf("%v heading #%v -text $atk_heading_label", w.id, column))
}

func (w *TreeView) HeaderLabel(column int) string {
	r, _ := evalAsString(fmt.Sprintf("%v heading #%v -text", w.id, column))
	return r
}

func (w *TreeView) SetHeaderImage(column int, img *Image) error {
	var iid string
	if img != nil {
		iid = img.Id()
	}
	return eval(fmt.Sprintf("%v heading #%v -image {%v}", w.id, column, iid))
}

func (w *TreeView) HeaderImage(column int) *Image {
	r, err := evalAsString(fmt.Sprintf("%v heading #%v -image", w.id, column))
	return parserImageResult(r, err)
}

func (w *TreeView) SetHeaderAnchor(column int, anchor Anchor) error {
	return eval(fmt.Sprintf("%v heading #%v -anchor %v", w.id, column, anchor))
}

func (w *TreeView) HeaderAnchor(column int) Anchor {
	r, err := evalAsString(fmt.Sprintf("%v heading #%v -anchor", w.id, column))
	return parserAnchorResult(r, err)
}

func (w *TreeView) SetColumnWidth(column int, width int) error {
	return eval(fmt.Sprintf("%v column #%v -width %v", w.id, column, width))
}

func (w *TreeView) ColumnWidth(column int) int {
	r, _ := evalAsInt(fmt.Sprintf("%v column #%v -width", w.id, column))
	return r
}

func (w *TreeView) SetColumnMinimumWidth(column int, width int) error {
	return eval(fmt.Sprintf("%v column #%v -minwidth %v", w.id, column, width))
}

func (w *TreeView) ColumnMinimumWidth(column int) int {
	r, _ := evalAsInt(fmt.Sprintf("%v column #%v -minwidth", w.id, column))
	return r
}

func (w *TreeView) SetColumnAnchor(column int, anchor Anchor) error {
	return eval(fmt.Sprintf("%v column #%v -anchor %v", w.id, column, anchor))
}

func (w *TreeView) ColumnAnchor(column int) Anchor {
	r, err := evalAsString(fmt.Sprintf("%v column #%v -anchor", w.id, column))
	return parserAnchorResult(r, err)
}

// default all column stretch 1
func (w *TreeView) SetColumnStretch(column int, stretch bool) error {
	return eval(fmt.Sprintf("%v column #%v -stretch %v", w.id, column, stretch))
}

// default all column stretch 1
func (w *TreeView) ColumnStretch(column int) bool {
	r, _ := evalAsBool(fmt.Sprintf("%v column #%v -stretch", w.id, column))
	return r
}

func (w *TreeView) IsValidItem(item *TreeItem) bool {
	return item != nil && item.tree != nil && item.tree.id == w.id
}

func (w *TreeView) RootItem() *TreeItem {
	return &TreeItem{w, ""}
}

func (w *TreeView) ToplevelItems() []*TreeItem {
	return w.RootItem().Children()
}

func (w *TreeView) InsertItem(parent *TreeItem, index int, text string, values []string) *TreeItem {
	var pid string
	if parent != nil {
		if !w.IsValidItem(parent) {
			return nil
		}
		pid = parent.id
	}
	setObjText("atk_tree_item", text)
	setObjTextList("atk_tree_values", values)
	cid := makeTreeItemId(w.id, pid)
	err := eval(fmt.Sprintf("%v insert {%v} %v -id {%v} -text $atk_tree_item -values $atk_tree_values", w.id, pid, index, cid))
	if err != nil {
		return nil
	}
	return &TreeItem{w, cid}
}

func (w *TreeView) DeleteItem(item *TreeItem) error {
	if !w.IsValidItem(item) || item.IsRoot() {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v delete {%v}", w.id, item.id))
}

func (w *TreeView) DeleteAllItems() error {
	var ids []string
	for _, item := range w.RootItem().Children() {
		ids = append(ids, item.Id())
	}
	if len(ids) == 0 {
		return ErrInvalid
	}
	setObjTextList("atk_tmp_items", ids)
	return eval(fmt.Sprintf("%v delete $atk_tmp_items", w.id))
}

func (w *TreeView) MoveItem(item *TreeItem, parent *TreeItem, index int) error {
	if !w.IsValidItem(item) || item.IsRoot() {
		return ErrInvalid
	}
	var pid string
	if parent != nil {
		if !w.IsValidItem(parent) {
			return ErrInvalid
		}
		pid = parent.id
	}
	return eval(fmt.Sprintf("%v move {%v} {%v} %v", w.id, item.id, pid, index))
}

func (w *TreeView) ScrollTo(item *TreeItem) error {
	if !w.IsValidItem(item) || item.IsRoot() {
		return ErrInvalid
	}
	children := w.RootItem().Children()
	if len(children) == 0 {
		return ErrInvalid
	}
	//fix see bug: first scroll to root
	eval(fmt.Sprintf("%v see %v", w.id, children[0].id))
	return eval(fmt.Sprintf("%v see %v", w.id, item.id))
}

func (w *TreeView) CurrentIndex() *TreeItem {
	lst := w.SelectionList()
	if len(lst) == 0 {
		return nil
	}
	return lst[0]
}

func (w *TreeView) SetCurrentIndex(item *TreeItem) error {
	return w.SetSelections(item)
}

func (w *TreeView) SelectionList() (lst []*TreeItem) {
	ids, err := evalAsStringList(fmt.Sprintf("%v selection", w.id))
	if err != nil {
		return
	}
	for _, id := range ids {
		lst = append(lst, &TreeItem{w, id})
	}
	return lst
}

func (w *TreeView) SetSelections(items ...*TreeItem) error {
	return w.SetSelectionList(items)
}

func (w *TreeView) RemoveSelections(items ...*TreeItem) error {
	return w.RemoveSelectionList(items)
}

func (w *TreeView) AddSelections(items ...*TreeItem) error {
	return w.AddSelectionList(items)
}

func (w *TreeView) ToggleSelections(items ...*TreeItem) error {
	return w.ToggleSelectionList(items)
}

func (w *TreeView) SetSelectionList(items []*TreeItem) error {
	var ids []string
	for _, item := range items {
		if w.IsValidItem(item) && !item.IsRoot() {
			ids = append(ids, item.id)
		}
	}
	if len(ids) == 0 {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v selection set {%v}", w.id, strings.Join(ids, " ")))
}

func (w *TreeView) RemoveSelectionList(items []*TreeItem) error {
	var ids []string
	for _, item := range items {
		if w.IsValidItem(item) && !item.IsRoot() {
			ids = append(ids, item.id)
		}
	}
	if len(ids) == 0 {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v selection remove {%v}", w.id, strings.Join(ids, " ")))
}

func (w *TreeView) AddSelectionList(items []*TreeItem) error {
	var ids []string
	for _, item := range items {
		if w.IsValidItem(item) && !item.IsRoot() {
			ids = append(ids, item.id)
		}
	}
	if len(ids) == 0 {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v selection add {%v}", w.id, strings.Join(ids, " ")))
}

func (w *TreeView) ToggleSelectionList(items []*TreeItem) error {
	var ids []string
	for _, item := range items {
		if w.IsValidItem(item) && !item.IsRoot() {
			ids = append(ids, item.id)
		}
	}
	if len(ids) == 0 {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v selection toggle {%v}", w.id, strings.Join(ids, " ")))
}

func (w *TreeView) ExpandAll() error {
	return w.RootItem().ExpandAll()
}

func (w *TreeView) CollepseAll() error {
	return w.RootItem().CollapseAll()
}

func (w *TreeView) Expand(item *TreeItem) error {
	return w.SetExpanded(item, true)
}

func (w *TreeView) Collapse(item *TreeItem) error {
	return w.SetExpanded(item, false)
}

func (w *TreeView) SetExpanded(item *TreeItem, expand bool) error {
	if !w.IsValidItem(item) || item.IsRoot() {
		return ErrInvalid
	}
	return item.SetExpanded(expand)
}

func (w *TreeView) IsExpanded(item *TreeItem) bool {
	if !w.IsValidItem(item) || item.IsRoot() {
		return false
	}
	return item.IsExpanded()
}

func (w *TreeView) FocusItem() *TreeItem {
	r, _ := evalAsString(fmt.Sprintf("%v focus", w.id))
	if r == "" {
		return nil
	}
	return &TreeItem{w, r}
}

func (w *TreeView) SetFocusItem(item *TreeItem) error {
	if !w.IsValidItem(item) || item.IsRoot() {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v focus %v", w.id, item.id))
}

func (w *TreeView) OnSelectionChanged(fn func()) error {
	if fn == nil {
		return ErrInvalid
	}
	return w.BindEvent("<<TreeviewSelect>>", func(e *Event) {
		fn()
	})
}

func (w *TreeView) OnItemExpanded(fn func()) error {
	if fn == nil {
		return ErrInvalid
	}
	return w.BindEvent("<<TreeviewOpen>>", func(e *Event) {
		fn()
	})
}

func (w *TreeView) OnItemCollapsed(fn func()) error {
	if fn == nil {
		return ErrInvalid
	}
	return w.BindEvent("<<TreeviewSelect>>", func(e *Event) {
		fn()
	})
}

func (w *TreeView) ItemAt(x int, y int) *TreeItem {
	id, err := evalAsString(fmt.Sprintf("%v identify item %v %v", w.id, x, y))
	if err != nil {
		return nil
	}
	if id == "" {
		return nil
	}
	return &TreeItem{w, id}
}

func (w *TreeView) OnDoubleClickedItem(fn func(item *TreeItem)) {
	if fn == nil {
		return
	}
	w.BindEvent("<Double-ButtonPress-1>", func(e *Event) {
		item := w.ItemAt(e.PosX, e.PosY)
		fn(item)
	})
}

func (w *TreeView) SetXViewArgs(args []string) error {
	return eval(fmt.Sprintf("%v xview %v", w.id, strings.Join(args, " ")))
}

func (w *TreeView) SetYViewArgs(args []string) error {
	return eval(fmt.Sprintf("%v yview %v", w.id, strings.Join(args, " ")))
}

func (w *TreeView) OnXScrollEx(fn func([]string) error) error {
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

func (w *TreeView) OnYScrollEx(fn func([]string) error) error {
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

func (w *TreeView) BindXScrollBar(bar *ScrollBar) error {
	if !IsValidWidget(bar) {
		return ErrInvalid
	}
	w.OnXScrollEx(bar.SetScrollArgs)
	bar.OnCommandEx(w.SetXViewArgs)
	return nil
}

func (w *TreeView) BindYScrollBar(bar *ScrollBar) error {
	if !IsValidWidget(bar) {
		return ErrInvalid
	}
	w.OnYScrollEx(bar.SetScrollArgs)
	bar.OnCommandEx(w.SetYViewArgs)
	return nil
}

type TreeViewEx struct {
	*ScrollLayout
	*TreeView
}

func NewTreeViewEx(parent Widget, attributs ...*WidgetAttr) *TreeViewEx {
	w := &TreeViewEx{}
	w.ScrollLayout = NewScrollLayout(parent)
	w.TreeView = NewTreeView(parent, attributs...)
	w.SetWidget(w.TreeView)
	w.TreeView.BindXScrollBar(w.XScrollBar)
	w.TreeView.BindYScrollBar(w.YScrollBar)
	RegisterWidget(w)
	return w
}

func TreeViewAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}

func TreeViewAttrHeight(row int) *WidgetAttr {
	return &WidgetAttr{"height", row}
}

func TreeViewAttrPadding(padding Pad) *WidgetAttr {
	return &WidgetAttr{"padding", padding}
}

func TreeViewAttrTreeSelectMode(mode TreeSelectMode) *WidgetAttr {
	return &WidgetAttr{"selectmode", mode}
}
