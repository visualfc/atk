// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"sync"
)

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

func parserTwoInt(s string) (n1 int, n2 int) {
	var p = &n1
	for _, r := range s {
		if r == ' ' {
			p = &n2
		} else {
			*p = *p*10 + int(r-'0')
		}
	}
	return
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
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
	})
}
