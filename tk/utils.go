// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"strings"
)

func escapeTS(s string) string {
	return strings.Replace(s, `"`, `\"`, -1)
}

func ToTkList(ar []string) string {
	var list []string
	for _, v := range ar {
		list = append(list, "{"+v+"}")
	}
	return strings.Join(list, " ")
}

func FromTkList(tklist string) (ar []string) {
	lastIndex := -1
	inBrace := false
	inString := false
	nBrace := 0
	for n, v := range tklist {
		if v == '{' {
			nBrace++
			if nBrace == 1 {
				inBrace = true
				inString = false
				lastIndex = n
			}
		} else if v == '}' {
			nBrace--
			if nBrace == 0 {
				ar = append(ar, tklist[lastIndex+1:n])
				inBrace = false
				inString = false
			}
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

func isValidKey(key string, keys []string) bool {
	for _, v := range keys {
		if v == key {
			return true
		}
	}
	return false
}

func SubString(text string, start int, end int) string {
	var n int = -1
	var r string
	for _, v := range text {
		n++
		if n < start {
			continue
		}
		if n >= end {
			break
		}
		r += string(v)
	}
	return r
}
