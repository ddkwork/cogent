// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//go:generate core generate

import (
	"time"

	"cogentcore.org/core/core"
	"cogentcore.org/core/errors"
	"cogentcore.org/core/events"
	"cogentcore.org/core/gox/datasize"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/views"

	"github.com/shirou/gopsutil/v3/process"
)

type Task struct { //types:add
	*process.Process `view:"-"`

	// The name of this task
	Name string `grow:"1"`

	// The percentage of the CPU time this task uses
	CPU float64 `format:"%.3g%%"`

	// The actual number of bytes of RAM this task uses (RSS)
	RAM datasize.Size

	// The percentage of total RAM this task uses
	RAMPct float32 `label:"RAM %" format:"%.3g%%"`

	// The number of threads this task uses
	Threads int32

	// The user that started this task
	User string

	// The Process ID (PID) of this task
	PID int32
}

func main() {
	b := core.NewBody("Cogent Task Manager")

	ts := getTasks(b)
	tv := views.NewTableView(b)
	tv.SetReadOnly(true)
	tv.SetSlice(&ts)
	tv.SortSliceAction(1)
	tv.SortSliceAction(1)

	tv.OnDoubleClick(func(e events.Event) {
		t := ts[tv.SelectedIndex]
		d := core.NewBody().AddTitle("Task info")
		views.NewStructView(d).SetStruct(&t).SetReadOnly(true)
		d.AddOKOnly().RunDialog(b)
	})

	tick := time.NewTicker(time.Second)
	paused := false

	b.OnShow(func(e events.Event) {
		go func() {
			for range tick.C {
				if paused {
					continue
				}
				ts = getTasks(b)
				tv.AsyncLock()
				tv.SortSlice()
				tv.Update()
				tv.AsyncUnlock()
			}
		}()
	})

	b.AddAppBar(func(tb *core.Toolbar) {
		core.NewButton(tb).SetText("End task").SetIcon(icons.Cancel).
			SetTooltip("Stop the currently selected task").
			OnClick(func(e events.Event) {
				t := ts[tv.SelectedIndex]
				core.ErrorSnackbar(tv, t.Kill(), "Error ending task")
			})
		pause := core.NewButton(tb).SetText("Pause").SetIcon(icons.Pause).
			SetTooltip("Stop updating the list of tasks")
		pause.OnClick(func(e events.Event) {
			paused = !paused
			if paused {
				pause.SetText("Resume").SetIcon(icons.Resume).
					SetTooltip("Resume updating the list of tasks")
			} else {
				pause.SetText("Pause").SetIcon(icons.Pause).
					SetTooltip("Stop updating the list of tasks")
			}
			pause.Update()
		})
	})

	b.RunMainWindow()
}

func getTasks(b *core.Body) []*Task {
	ps, err := process.Processes()
	core.ErrorDialog(b, err, "Error getting system processes")

	ts := make([]*Task, len(ps))
	for i, p := range ps {
		t := &Task{
			Process: p,
			Name:    errors.Ignore1(p.Name()),
			CPU:     errors.Ignore1(p.CPUPercent()),
			RAMPct:  errors.Ignore1(p.MemoryPercent()),
			Threads: errors.Ignore1(p.NumThreads()),
			User:    errors.Ignore1(p.Username()),
			PID:     p.Pid,
		}
		mi := errors.Ignore1(p.MemoryInfo())
		if mi != nil {
			t.RAM = datasize.Size(mi.RSS)
		}
		ts[i] = t
	}
	return ts
}
