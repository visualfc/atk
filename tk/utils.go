// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"strings"
)

var (
	rep1 = strings.NewReplacer("\\", "\\\\", "{", "\\{", "}", "\\}")
	rep2 = strings.NewReplacer(string([]byte{0x00}), "")
)

func ToTkList(ar []string) string {
	var list []string
	for _, v := range ar {
		list = append(list, "{"+rep1.Replace(v)+"}")
	}
	return strings.Join(list, " ")
}

func buildTkString(gostring string) string {
	return rep1.Replace(gostring)
}

// TODO extract string and remove backslash, no optimization
func extractGoString(tklist string, start int, n int, bsList []int) string {
	s := tklist[start:n]
	if bsList == nil {
		return s
	}
	data := []byte(s)
	for _, n := range bsList {
		data[n-start] = 0x00
	}
	return rep2.Replace(string(data))
}

func FromTkList(tklist string) (ar []string) {
	lastIndex := -1
	inBrace := false
	inString := false
	nBrace := 0
	firstBackslash := false
	var bsList []int
	for n, v := range tklist {
		if firstBackslash {
			firstBackslash = false
			continue
		}
		if v == '\\' && !firstBackslash {
			bsList = append(bsList, n)
			firstBackslash = true
			if !inBrace && !inString {
				inString = true
				lastIndex = n
			}
		} else if v == '{' {
			nBrace++
			if nBrace == 1 {
				inBrace = true
				inString = false
				lastIndex = n
			}
		} else if v == '}' {
			nBrace--
			if nBrace == 0 {
				ar = append(ar, extractGoString(tklist, lastIndex+1, n, bsList))
				bsList = nil
				inBrace = false
				inString = false
			}
		} else if !inBrace {
			if v == ' ' {
				if inString {
					ar = append(ar, extractGoString(tklist, lastIndex+1, n, bsList))
					bsList = nil
				}
				lastIndex = n
				inString = false
			} else {
				inString = true
			}
		}
	}
	if inString {
		ar = append(ar, extractGoString(tklist, lastIndex+1, len(tklist), bsList))
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
