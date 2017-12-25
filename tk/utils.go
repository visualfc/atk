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

func NewGenInt64Func(id int64) func() <-chan int64 {
	ch := make(chan int64)
	go func(i int64) {
		for {
			i++
			ch <- i
		}
	}(id)
	return func() <-chan int64 {
		return ch
	}
}

func NewGenIntFunc(id int) func() <-chan int {
	ch := make(chan int)
	go func(i int) {
		for {
			i++
			ch <- i
		}
	}(id)
	return func() <-chan int {
		return ch
	}
}
