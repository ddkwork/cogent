// Copyright (c) 2021, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"

	"cogentcore.org/cogent/vector"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/grr"
)

func main() {
	vector.InitPrefs()

	ofs := gi.TheApp.OpenFiles()

	var fnms []string
	if len(ofs) > 0 {
		fnms = ofs
	} else if len(os.Args) > 1 {
		fnms = os.Args[1:]
	}

	if len(fnms) == 0 {
		grr.Log(os.Chdir(gi.SystemSettings.User.HomeDir))
		vector.NewDrawing(vector.Prefs.Size)
	} else {
		fdir, _ := filepath.Split(fnms[0])
		os.Chdir(fdir)
		for _, fnm := range fnms {
			vector.NewVectorWindow(fnm)
		}
	}
	gi.Wait()
}
