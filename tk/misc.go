// Copyright 2018 visualfc. All rights reserved.

package tk

type Pos struct {
	X int
	Y int
}

type Pad struct {
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

type Anchor int

const (
	AnchorCenter = iota
	AnchorNorth
	AnchorEast
	AnchorSouth
	AnchorWest
	AnchorNorthEast
	AnchorNorthWest
	AnchorSouthEast
	AnchorSouthWest
)

var (
	anchorName = []string{"center", "n", "e", "s", "w", "ne", "nw", "se", "sw"}
)

func (v Anchor) String() string {
	if v >= 0 && int(v) < len(anchorName) {
		return anchorName[v]
	}
	return ""
}

func parserAnchorResult(r string, err error) Anchor {
	if err != nil {
		return -1
	}
	for n, s := range anchorName {
		if s == r {
			return Anchor(n)
		}
	}
	return -1
}

type Compound int

const (
	CompoundNone = iota
	CompoundTop
	CompoundBottom
	CompoundLeft
	CompoundRight
	CompoundCenter
)

var (
	compoundName = []string{"none", "top", "bottom", "left", "right", "center"}
)

func (v Compound) String() string {
	if v >= 0 && int(v) < len(compoundName) {
		return compoundName[v]
	}
	return ""
}

func parserCompoundResult(r string, err error) Compound {
	if err != nil {
		return 0
	}
	for n, s := range compoundName {
		if s == r {
			return Compound(n)
		}
	}
	return 0
}

type State int

const (
	StateNormal = iota
	StateActive
	StateDisable
	StateReadOnly
)

var (
	stateName = []string{"normal", "active", "disabled", "readonly"}
)

func (v State) String() string {
	if v >= 0 && int(v) < len(stateName) {
		return stateName[v]
	}
	return ""
}

func parserStateResult(r string, err error) State {
	if err != nil {
		return 0
	}
	for n, s := range stateName {
		if s == r {
			return State(n)
		}
	}
	return 0
}

func parserPaddingResult(r string, err error) (int, int) {
	if err != nil {
		return 0, 0
	}
	return parserTwoInt(r)
}
