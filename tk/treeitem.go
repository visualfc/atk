// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

type TreeItem struct {
	tree *TreeView
	id   string
}

func (t *TreeItem) Id() string {
	return t.id
}

func (t *TreeItem) IsValid() bool {
	return t != nil && t.tree != nil
}

func (t *TreeItem) InsertItem(index int, text string, values []string) *TreeItem {
	if !t.IsValid() {
		return nil
	}
	return t.tree.InsertItem(t, index, text, values)
}

func (t *TreeItem) Index() int {
	if !t.IsValid() || t.IsRoot() {
		return -1
	}
	r, err := evalAsIntEx(fmt.Sprintf("%v index {%v}", t.tree.id, t.id), false)
	if err != nil {
		return -1
	}
	return r
}

func (t *TreeItem) IsRoot() bool {
	return t.id == ""
}

func (t *TreeItem) Parent() *TreeItem {
	if !t.IsValid() || t.IsRoot() {
		return nil
	}
	r, err := evalAsStringEx(fmt.Sprintf("%v parent {%v}", t.tree.id, t.id), false)
	if err != nil {
		return nil
	}
	return &TreeItem{t.tree, r}
}

func (t *TreeItem) Next() *TreeItem {
	if !t.IsValid() || t.IsRoot() {
		return nil
	}
	r, err := evalAsStringEx(fmt.Sprintf("%v next {%v}", t.tree.id, t.id), false)
	if err != nil || r == "" {
		return nil
	}
	return &TreeItem{t.tree, r}
}

func (t *TreeItem) Prev() *TreeItem {
	if !t.IsValid() || t.IsRoot() {
		return nil
	}
	r, err := evalAsStringEx(fmt.Sprintf("%v prev {%v}", t.tree.id, t.id), false)
	if err != nil || r == "" {
		return nil
	}
	return &TreeItem{t.tree, r}
}

func (t *TreeItem) Children() (lst []*TreeItem) {
	if !t.IsValid() {
		return
	}
	ids, err := evalAsStringList(fmt.Sprintf("%v children {%v}", t.tree.id, t.id))
	if err != nil {
		return
	}
	for _, id := range ids {
		lst = append(lst, &TreeItem{t.tree, id})
	}
	return
}

func (t *TreeItem) SetExpanded(expand bool) error {
	if !t.IsValid() || t.IsRoot() {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v item {%v} -open %v", t.tree.id, t.id, expand))
}

func (t *TreeItem) IsExpanded() bool {
	if !t.IsValid() || t.IsRoot() {
		return false
	}
	r, _ := evalAsBool(fmt.Sprintf("%v item {%v} -open", t.tree.id, t.id))
	return r
}

func (t *TreeItem) expandAll(item *TreeItem) error {
	for _, child := range item.Children() {
		child.SetExpanded(true)
		t.expandAll(child)
	}
	return nil
}

func (t *TreeItem) ExpandAll() error {
	return t.expandAll(t)
}

func (t *TreeItem) collapseAll(item *TreeItem) error {
	for _, child := range item.Children() {
		child.SetExpanded(false)
		t.collapseAll(child)
	}
	return nil
}

func (t *TreeItem) CollapseAll() error {
	return t.collapseAll(t)
}

func (t *TreeItem) Expand() error {
	return t.SetExpanded(true)
}

func (t *TreeItem) Collapse() error {
	return t.SetExpanded(false)
}

func (t *TreeItem) SetText(text string) error {
	if !t.IsValid() || t.IsRoot() {
		return ErrInvalid
	}
	setObjText("atk_tree_item", text)
	return eval(fmt.Sprintf("%v item {%v} -text $atk_tree_item", t.tree.id, t.id))
}

func (t *TreeItem) Text() string {
	if !t.IsValid() || t.IsRoot() {
		return ""
	}
	r, _ := evalAsString(fmt.Sprintf("%v item {%v} -text", t.tree.id, t.id))
	return r
}

func (t *TreeItem) SetValues(values []string) error {
	if !t.IsValid() || t.IsRoot() {
		return ErrInvalid
	}
	setObjTextList("atk_tree_values", values)
	return eval(fmt.Sprintf("%v item {%v} -values $atk_tree_values", t.tree.id, t.id))
}

func (t *TreeItem) Values() []string {
	if !t.IsValid() || t.IsRoot() {
		return nil
	}
	r, _ := evalAsStringList(fmt.Sprintf("%v item {%v} -values", t.tree.id, t.id))
	return r
}

func (t *TreeItem) SetImage(img *Image) error {
	if !t.IsValid() || t.IsRoot() {
		return ErrInvalid
	}
	var iid string
	if img != nil {
		iid = img.Id()
	}
	return eval(fmt.Sprintf("%v item {%v} -image {%v}", t.tree.id, t.id, iid))
}

func (t *TreeItem) Image() *Image {
	if !t.IsValid() || t.IsRoot() {
		return nil
	}
	r, err := evalAsString(fmt.Sprintf("%v item {%v} -image", t.tree.id, t.id))
	return parserImageResult(r, err)
}

func (t *TreeItem) SetColumnText(column int, text string) error {
	if column < 0 {
		return ErrInvalid
	} else if column == 0 {
		return t.SetText(text)
	}
	if !t.IsValid() || t.IsRoot() {
		return ErrInvalid
	}
	setObjText("atk_tree_column", text)
	return eval(fmt.Sprintf("%v set {%v} %v $atk_tree_column", t.tree.id, t.id, column-1))
}

func (t *TreeItem) ColumnText(column int) string {
	if column < 0 {
		return ""
	} else if column == 0 {
		return t.Text()
	}
	if !t.IsValid() || t.IsRoot() {
		return ""
	}
	r, _ := evalAsString(fmt.Sprintf("%v set {%v} %v", t.tree.id, t.id, column-1))
	return r
}
