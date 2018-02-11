// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"sync"
)

type Command struct {
	sync.RWMutex
	cmds []func()
}

func (c *Command) Bind(fn func()) {
	if fn == nil {
		return
	}
	c.Lock()
	c.cmds = append(c.cmds, fn)
	c.Unlock()
}

func (c *Command) Clear() {
	c.Lock()
	c.cmds = nil
	c.Unlock()
}

func (c *Command) Invoke() {
	c.RLock()
	for _, cmd := range c.cmds {
		if cmd != nil {
			cmd()
		}
	}
	c.RUnlock()
}

type CommandEx struct {
	sync.RWMutex
	cmds []func([]string) error
}

func (c *CommandEx) Bind(fn func([]string) error) {
	if fn == nil {
		return
	}
	c.Lock()
	c.cmds = append(c.cmds, fn)
	c.Unlock()
}

func (c *CommandEx) Clear() {
	c.Lock()
	c.cmds = nil
	c.Unlock()
}

func (c *CommandEx) Invoke(args []string) {
	c.RLock()
	for _, cmd := range c.cmds {
		if cmd != nil {
			cmd(args)
		}
	}
	c.RUnlock()
}

func bindCommand(id string, command string, fn func()) error {
	actName := makeActionId()
	err := eval(fmt.Sprintf("%v configure -%v {%v}", id, command, actName))
	if err != nil {
		return err
	}
	mainInterp.CreateAction(actName, func(args []string) {
		fn()
	})
	return nil
}

func bindCommandEx(id string, command string, fn func([]string)) error {
	actName := makeActionId()
	err := eval(fmt.Sprintf("%v configure -%v {%v}", id, command, actName))
	if err != nil {
		return err
	}
	mainInterp.CreateAction(actName, func(args []string) {
		fn(args)
	})
	return nil
}
