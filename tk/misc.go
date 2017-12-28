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

type Aligment int

const (
	AligmentCenter Aligment = 0
	AligmentLeft            = 1
	AligmentRight           = 2
)

func (v Aligment) String() string {
	switch v {
	case AligmentCenter:
		return "center"
	case AligmentLeft:
		return "left"
	case AligmentRight:
		return "right"
	}
	return ""
}
