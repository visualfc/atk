// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"log"
	"sync"
)

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

var (
	makeActionFunc = NewGenIntFunc(1024)
)

func MakeActionName() string {
	return fmt.Sprintf("go_action_%v", <-makeActionFunc())
}

func SplitTkList(tklist string) (ar []string) {
	lastIndex := 0
	inBrace := false
	inString := false
	for n, v := range tklist {
		if v == '{' {
			inBrace = true
			inString = false
			lastIndex = n
		} else if v == '}' {
			ar = append(ar, tklist[lastIndex+1:n])
			inBrace = false
			inString = false
		} else if !inBrace {
			if v == ' ' {
				if inString {
					ar = append(ar, tklist[lastIndex+1:n])
				}
				lastIndex = n
				inString = false
			} else {
				inString = true
			}
		}
	}
	if inString {
		ar = append(ar, tklist[lastIndex+1:])
	}
	return
}

var (
	testOnce sync.Once
)

func InitTest() {
	testOnce.Do(func() {
		err := Init()
		if err != nil {
			panic(err)
		}
		SetErrorHandle(func(err error) {
			log.Println(err)
		})
	})
}
