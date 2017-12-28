// Copyright 2017 visualfc. All rights reserved.

package tk

type Pos struct {
	X int
	Y int
}

type Size struct {
	Width  int
	Height int
}

type Geometry struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Alignment int

const (
	AlignmentCenter Alignment = iota
	AlignmentLeft
	AlignmentRight
)

var (
	alignmentName = []string{"center", "left", "right"}
)

func (v Alignment) String() string {
	if v >= 0 && int(v) < len(alignmentName) {
		return alignmentName[v]
	}
	return ""
}

func parserAlignmentResult(r string, err error) Alignment {
	if err != nil {
		return -1
	}
	for n, s := range alignmentName {
		if s == r {
			return Alignment(n)
		}
	}
	return -1
}

type BorderStyle int

const (
	BorderStyleFlat BorderStyle = iota
	BorderStyleGroove
	BorderStyleRaised
	BorderStyleRidge
	BorderStyleSolid
	BorderStyleSunken
)

var (
	borderStyleName = []string{"flat", "groove", "raised", "ridge", "solid", "sunken"}
)

func (v BorderStyle) String() string {
	if v >= 0 && int(v) < len(borderStyleName) {
		return borderStyleName[v]
	}
	return ""
}

func parserBorderStyleResult(r string, err error) BorderStyle {
	if err != nil {
		return -1
	}
	for n, s := range borderStyleName {
		if s == r {
			return BorderStyle(n)
		}
	}
	return -1
}
