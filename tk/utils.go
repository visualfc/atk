// Copyright 2018 visualfc. All rights reserved.

package tk

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
