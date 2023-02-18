// Copyright (c) 2018, The Gide Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gide

import "github.com/goki/pi/filecat"

// Use these for more obvious command options
const (
	CmdWait      = true
	CmdNoWait    = false
	CmdFocus     = true
	CmdNoFocus   = false
	CmdConfirm   = true
	CmdNoConfirm = false
)

// StdCmds is the original compiled-in set of standard commands.
var StdCmds = Commands{
	{Cat: "Build", Name: "Run Proj",
		Desc: "run RunExec executable set in project",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "{RunExecPath}"}},
		Dir:  "{RunExecDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},
	{Cat: "Build", Name: "Run Prompt",
		Desc: "run any command you enter at the prompt",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "{PromptString1}"}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},
	// Make
	{Cat: "Build", Name: "Make",
		Desc: "run make with no args",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "make"}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Build", Name: "Make Prompt",
		Desc: "run make with prompted make target",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "make",
			Args: []string{"{PromptString1}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "File", Name: "Grep",
		Desc: "recursive grep of all files for prompted value",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "grep",
			Args: []string{"-R", "-e", "{PromptString1}", "{FileDirPath}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "File", Name: "Open",
		Desc: "open file using OS 'open' command",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "open",
			Args: []string{"{FilePath}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "File", Name: "Open Target",
		Desc: "open project target file using OS 'open' command",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "open",
			Args: []string{"{RunExecPath}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	// Go
	{Cat: "Go", Name: "Imports File",
		Desc: "run goimports on file",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "goimports",
			Args: []string{"-w", "{FilePath}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Go", Name: "Fmt File",
		Desc: "run go fmt on file",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "gofmt",
			Args: []string{"-w", "{FilePath}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Go", Name: "Build Dir",
		Desc: "run go build to build in current dir",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "go",
			Args: []string{"build", "-v"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Go", Name: "Build Proj",
		Desc: "run go build for project BuildDir",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "go",
			Args: []string{"build", "-v"}}},
		Dir:  "{BuildDir}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Go", Name: "Install Dir",
		Desc: "run go install in current dir",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "go",
			Args: []string{"install", "-v"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Go", Name: "Generate",
		Desc: "run go generate in current dir",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "go",
			Args: []string{"generate"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Go", Name: "Test",
		Desc: "run go test in current dir.  Options include: -run TestName or -bench",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "go",
			Args: []string{"test", "-v", "{PromptString1}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Go", Name: "Vet",
		Desc: "run go vet in current dir",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "go",
			Args: []string{"vet"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Go", Name: "Mod Tidy",
		Desc: "run go mod tidy in current dir",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "go",
			Args: []string{"mod", "tidy"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Go", Name: "Mod Init",
		Desc: "run go mod init in current dir with module path from prompt",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "go",
			Args: []string{"mod", "init", "{PromptString1}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Go", Name: "Get",
		Desc: "run go get on package(s) you enter at prompt",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "go",
			Args: []string{"get", "{PromptString1}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Go", Name: "Get Updt",
		Desc: "run go get -u (updt) on package(s) you enter at prompt.  use ./... for all.",
		Lang: filecat.Go,
		Cmds: []CmdAndArgs{{Cmd: "go",
			Args:    []string{"get", "-u", "{PromptString1}"},
			Default: "./..."}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	// Git
	{Cat: "Git", Name: "Add",
		Desc: "git add file",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args: []string{"add", "{FilePath}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Checkout Branch",
		Desc: "git checkout: switch to an existing branch",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args: []string{"checkout", "{PromptBranch}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Checkout",
		Desc: "git checkout: file, directory, branch; -b <branch> creates a new branch",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args: []string{"checkout", "{PromptString1}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Status",
		Desc: "git status",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args: []string{"status"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Diff",
		Desc: "git diff -- see changes since last checkin",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args: []string{"diff"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Log",
		Desc: "git log",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args: []string{"log"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Commit",
		Desc: "git commit",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args: []string{"commit", "-am", "{PromptString1}"}, PromptIsString: true}},
		Dir:  "{FileDirPath}",
		Wait: CmdWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm}, // promptstring1 provided during normal commit process, MUST be wait!

	{Cat: "Git", Name: "Pull",
		Desc: "git pull",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args:    []string{"pull", "{PromptString1}"},
			Default: "origin"}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Push",
		Desc: "git push",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args:    []string{"push", "origin", "{PromptBranch}"},
			Default: "origin"}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Stash",
		Desc: "git stash: push / pop local changes for later without committing, resetting to HEAD: also list, show, drop, clear",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args:    []string{"stash", "{PromptString1}"},
			Default: "push"}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Branch",
		Desc: "git branch: -a shows all; <branchname> makes a new one, optionally given sha, --delete to delete",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args:    []string{"branch", "{PromptString1}"},
			Default: "-a"}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Branch Delete",
		Desc: "git branch --delete selected branch name",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args: []string{"branch", "--delete", "{PromptBranch}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Rebase",
		Desc: "git rebase: updates current branch to given branch, often master",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args:    []string{"rebase", "{PromptBranch}"},
			Default: "master"}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "Git", Name: "Reset",
		Desc: "git reset: resets current state to given source -- use --hard to override -- CAN RESULT IN LOSS OF CHANGES -- make sure everything is committed",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "git",
			Args:    []string{"reset", "{PromptString1}"},
			Default: "--hard origin/master"}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	// SVN
	{Cat: "SVN", Name: "Add",
		Desc: "svn add file",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "svn",
			Args: []string{"add", "{FilePath}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "SVN", Name: "Status",
		Desc: "svn status",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "svn",
			Args: []string{"status"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "SVN", Name: "Info",
		Desc: "svn info",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "svn",
			Args: []string{"info"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "SVN", Name: "Log",
		Desc: "svn log",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "svn",
			Args: []string{"log", "-v"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "SVN", Name: "Commit Proj",
		Desc: "svn commit for entire project directory",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "svn",
			Args: []string{"commit", "-m", "{PromptString1}"}, PromptIsString: true}},
		Dir:  "{ProjPath}",
		Wait: CmdWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm}, // promptstring1 provided during normal commit process

	{Cat: "SVN", Name: "Commit Dir",
		Desc: "svn commit in directory of current file",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "svn",
			Args: []string{"commit", "-m", "{PromptString1}"}, PromptIsString: true}},
		Dir:  "{FileDirPath}",
		Wait: CmdWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm}, // promptstring1 provided during normal commit process

	{Cat: "SVN", Name: "Update",
		Desc: "svn update",
		Lang: filecat.Any,
		Cmds: []CmdAndArgs{{Cmd: "svn",
			Args: []string{"update"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	// LaTeX
	{Cat: "LaTeX", Name: "LaTeX PDF",
		Desc: "run PDFLaTeX on file",
		Lang: filecat.TeX,
		Cmds: []CmdAndArgs{{Cmd: "pdflatex",
			Args: []string{"-file-line-error", "-interaction=nonstopmode", "{FilePath}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "LaTeX", Name: "BibTeX",
		Desc: "run BibTeX on file",
		Lang: filecat.TeX,
		Cmds: []CmdAndArgs{{Cmd: "bibtex",
			Args: []string{"{FileNameNoExt}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "LaTeX", Name: "Biber",
		Desc: "run Biber on file",
		Lang: filecat.TeX,
		Cmds: []CmdAndArgs{{Cmd: "biber",
			Args: []string{"{FileNameNoExt}"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},

	{Cat: "LaTeX", Name: "CleanTeX",
		Desc: "remove aux LaTeX files",
		Lang: filecat.TeX,
		Cmds: []CmdAndArgs{{Cmd: "rm",
			Args: []string{"*.aux", "*.log", "*.blg", "*.bbl", "*.fff", "*.lof", "*.ttt", "*.toc", "*.spl"}}},
		Dir:  "{FileDirPath}",
		Wait: CmdNoWait, Focus: CmdNoFocus, Confirm: CmdNoConfirm},
}
