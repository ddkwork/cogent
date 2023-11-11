// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gear

import (
	"fmt"

	"goki.dev/xe"
)

// Parse uses the help messages of the app to fill in its data fields.
func (a *App) Parse() error {
	rh, err := a.GetHelp("")
	if err != nil {
		return err
	}
	fmt.Println(rh)
	return nil
}

// GetHelp gets the help information for the given command, or the root command
// if the command is unspecified. It tries various different commands and flags
// to get the help information and only returns an error if all of them fail.
func (a *App) GetHelp(cmd string) (string, error) {
	hcmds := []string{"help", "--help", "-h"}
	for _, hcmd := range hcmds {
		args := []string{hcmd}
		if cmd != "" {
			args = append(args, cmd)
		}
		out, err := xe.Silent().Output(a.Command, args...)
		if err == nil {
			return out, nil
		}
		if cmd != "" {
			// try both orders
			out, err = xe.Silent().Output(a.Command, cmd, hcmd)
			if err == nil {
				return out, nil
			}
		}
	}
	return "", fmt.Errorf("unable to get help information for command %q of app %q (%q)", cmd, a.Name, a.Command+" "+cmd)
}
