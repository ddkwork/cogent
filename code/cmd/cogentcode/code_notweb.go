// Copyright (c) 2020, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js

package main

import (
	"cogentcore.org/cogent/code"
	_ "cogentcore.org/cogent/code/cdebug/cdelve"
)

func InitConsole(logfile string) {
	code.TheConsole.Init(logfile)
}
