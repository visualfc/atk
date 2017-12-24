// Copyright 2017 visualfc. All rights reserved.

package interp

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"
)

var (
	interp *Interp
)

func init() {
	var err error
	interp, err = NewInterp()
	if err != nil {
		panic(err)
	}
	err = interp.InitTcl("")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("tcl_version", interp.TclVersion())
}

func TestInterp(t *testing.T) {
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

func TestObj(t *testing.T) {
	if NewStringObj("string", interp).ToString() != "string" {
		t.Fatalf("string obj")
	}
	if f := NewFloat64Obj(1e16, interp).ToFloat64(); f != 1e16 {
		t.Fatal("float64 obj", f)
	}
	if f := NewFloat64Obj(1.123456789123456789, interp).ToFloat64(); f != 1.123456789123456789 {
		t.Fatal("float64 obj", f)
	}
	if NewInt64Obj(1e12, interp).ToInt64() != 1e12 {
		t.Fatalf("int64 obj")
	}
	if NewIntObj(1024, interp).ToInt() != 1024 {
		t.Fatalf("int obj")
	}
	if NewBoolObj(true, interp).ToBool() != true {
		t.Fatalf("bool boj")
	}
}

func TestTkSync(t *testing.T) {
	err := interp.InitTk("")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("tk_version", interp.TkVersion())
	MainLoop(func() {
		go func() {
			fmt.Println("run tk mainloop wait 1 sec async destroy")
			<-time.After(1e9)
			Async(func() {
				interp.Destroy()
			})
		}()
	})
}
