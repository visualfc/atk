// Copyright 2018 visualfc. All rights reserved.

package tk

import "errors"

var (
	ErrInvalid   = errors.New("invalid argument")
	ErrExist     = errors.New("already exists")
	ErrNotExist  = errors.New("does not exist")
	ErrClosed    = errors.New("already closed")
	ErrUnsupport = errors.New("unsupport")
)
