// Copyright 2023 The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"goki.dev/events"
	"goki.dev/gi"
	"goki.dev/giv"
)

func main() {
	b := gi.NewAppBody("Goki Files")

	fv := giv.NewFileView(b)
	fv.OnDoubleClick(func(e events.Event) {
		if fv.SelectedDoubleClick {
			gi.OpenURL(fv.SelectedFile())
		}
	})

	b.NewWindow().Run().Wait()
}
