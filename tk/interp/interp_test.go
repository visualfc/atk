// Copyright 2017 visualfc. All rights reserved.

package interp

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestInterp(t *testing.T) {
	interp, err := NewInterp()
	defer interp.Destroy()
	if err != nil {
		t.Fatal(err)
	}
	tcl_ver, _ := interp.EvalAsString("set tcl_version")
	fmt.Println("tcl_version", tcl_ver)
	tk_ver, _ := interp.EvalAsString("set tk_version")
	fmt.Println("tk_version", tk_ver)

	a, err := interp.EvalAsString("set a {hello}\nset a")
	if err != nil {
		t.Fatal(err)
	}
	if a != "hello" {
		t.Fatalf("EvalAsString")
	}
	b, err := interp.EvalAsInt64("set b 1000000000000\nexpr $b")
	if err != nil {
		t.Fatal(err)
	}
	if b != 1e12 {
		t.Fatalf("EvalAsInt64 %v", b)
	}
	c, err := interp.EvalAsInt("set c 100\nexpr $c")
	if err != nil {
		t.Fatal(err)
	}
	if c != 100 {
		t.Fatalf("EvalAsInt")
	}
	d, err := interp.EvalAsFloat64("set d 1e12\nexpr $d")
	if err != nil {
		t.Fatal(err)
	}
	if d != 1e12 {
		t.Fatalf("EvalAsFloat64 %v", d)
	}
}

func TestCommand(t *testing.T) {
	interp, err := NewInterp()
	defer interp.Destroy()
	if err != nil {
		t.Fatal(err)
	}
	interp.CreateCommand("go::join", func(args []string) (string, error) {
		return strings.Join(args, ","), nil
	})
	s, err := interp.EvalAsString("go::join hello world")
	if err != nil {
		t.Fatal(err, s)
	}
	if s != "hello,world" {
		t.Fatal(s)
	}
	interp.CreateCommand("go::sum", func(args []string) (string, error) {
		var sum int
		for _, arg := range args {
			i, err := strconv.Atoi(arg)
			if err != nil {
				return "", err
			}
			sum += i
		}
		return strconv.Itoa(sum), nil
	})
	sum, err := interp.EvalAsInt("expr [go::sum 100 200 300]")
	if err != nil {
		t.Fatal(err)
	}
	if sum != 600 {
		t.Fatal("CreateCommand")
	}
	var check_success bool
	interp.CreateAction("go::action", func() {
		check_success = true
	})
	err = interp.Eval("go::action")
	if err != nil {
		t.Fatal(err)
	}
	if !check_success {
		t.Fatal("CreateAction")
	}
}
