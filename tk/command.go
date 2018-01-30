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

func bindCommand(id string, command string, fn func()) error {
	actName := makeActionId()
	err := eval(fmt.Sprintf("%v configure -%v {%v}", id, command, actName))
	if err != nil {
		return err
	}
	mainInterp.CreateAction(actName, func([]string) {
		fn()
	})
	return nil
}
