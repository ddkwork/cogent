// Copyright (c) 2023, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"cogentcore.org/core/gi/v2/gi"
	"cogentcore.org/core/gi/v2/giv"
)

func main() {
	b := gi.NewAppBody("goki-settings")
	giv.SettingsView(b)
	b.NewWindow().Run().Wait()
}
