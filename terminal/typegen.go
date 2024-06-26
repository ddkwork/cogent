// Code generated by "core generate -add-types"; DO NOT EDIT.

package terminal

import (
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/cogent/terminal.App", IDName: "app", Doc: "App is a GUI view of a terminal command.", Embeds: []types.Field{{Name: "Frame"}}, Fields: []types.Field{{Name: "Cmd", Doc: "Cmd is the root command associated with this app."}, {Name: "CurCmd", Doc: "CurCmd is the current root command being typed in."}, {Name: "Dir", Doc: "Dir is the current directory of the app."}}})

// NewApp returns a new [App] with the given optional parent:
// App is a GUI view of a terminal command.
func NewApp(parent ...tree.Node) *App { return tree.New[App](parent...) }

// SetCmd sets the [App.Cmd]:
// Cmd is the root command associated with this app.
func (t *App) SetCmd(v *Cmd) *App { t.Cmd = v; return t }

// SetCurCmd sets the [App.CurCmd]:
// CurCmd is the current root command being typed in.
func (t *App) SetCurCmd(v string) *App { t.CurCmd = v; return t }

// SetDir sets the [App.Dir]:
// Dir is the current directory of the app.
func (t *App) SetDir(v string) *App { t.Dir = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/cogent/terminal.Cmd", IDName: "cmd", Doc: "Cmd contains all of the data for a parsed command line command.", Fields: []types.Field{{Name: "Cmd", Doc: "Cmd is the actual name of the command (eg: \"git\", \"go build\")"}, {Name: "Name", Doc: "Name is the formatted name of the command (eg: \"Git\", \"Go build\")"}, {Name: "Doc", Doc: "Doc is the documentation for the command (eg: \"compile packages and dependencies\")"}, {Name: "Flags", Doc: "Flags contains the flags for the command"}, {Name: "Cmds", Doc: "Cmds contains the subcommands of the command"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/cogent/terminal.Flag", IDName: "flag", Doc: "Flag contains the information for a parsed command line flag.", Fields: []types.Field{{Name: "Name", Doc: "Name is the canonical (longest) name of the flag.\nIt includes the leading dashes of the flag."}, {Name: "Names", Doc: "Names are the different names the flag can go by.\nThey include the leading dashes of the flag."}, {Name: "Type", Doc: "Type is the type or value hint for the flag."}, {Name: "Doc", Doc: "Doc is the documentation for the flag."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/cogent/terminal.ParseBlock", IDName: "parse-block", Doc: "ParseBlock is a block of parsed content containing the name of something and\nthe documentation for it.", Fields: []types.Field{{Name: "Name"}, {Name: "Doc"}}})
