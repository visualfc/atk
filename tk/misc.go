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
	AlignmentCenter  Alignment = 0
	AlignmentLeft              = 1
	AlignmentRight             = 2
	AlignmentInvalid Alignment = -1
)

func (v Alignment) String() string {
	switch v {
	case AlignmentCenter:
		return "center"
	case AlignmentLeft:
		return "left"
	case AlignmentRight:
		return "right"
	}
	return ""
}

func parserAlignmentResult(r string, err error) Alignment {
	if err != nil {
		return -1
	}
	switch r {
	case "center":
		return AlignmentCenter
	case "left":
		return AlignmentLeft
	case "right":
		return AlignmentRight
	}
	return -1
}
