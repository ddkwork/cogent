// Code generated by "goki generate"; DO NOT EDIT.

package gide

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _ArgVarTypesValues = []ArgVarTypes{0, 1, 2, 3, 4, 5}

// ArgVarTypesN is the highest valid value
// for type ArgVarTypes, plus one.
const ArgVarTypesN ArgVarTypes = 6

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ArgVarTypesNoOp() {
	var x [1]struct{}
	_ = x[ArgVarFile-(0)]
	_ = x[ArgVarDir-(1)]
	_ = x[ArgVarExt-(2)]
	_ = x[ArgVarPos-(3)]
	_ = x[ArgVarText-(4)]
	_ = x[ArgVarPrompt-(5)]
}

var _ArgVarTypesNameToValueMap = map[string]ArgVarTypes{
	`File`:   0,
	`file`:   0,
	`Dir`:    1,
	`dir`:    1,
	`Ext`:    2,
	`ext`:    2,
	`Pos`:    3,
	`pos`:    3,
	`Text`:   4,
	`text`:   4,
	`Prompt`: 5,
	`prompt`: 5,
}

var _ArgVarTypesDescMap = map[ArgVarTypes]string{
	0: `ArgVarFile is a file name, not a directory`,
	1: `ArgVarDir is a directory name, not a file`,
	2: `ArgVarExt is a file extension`,
	3: `ArgVarPos is a text position`,
	4: `ArgVarText is text from a buffer`,
	5: `ArgVarPrompt is a user-prompted variable`,
}

var _ArgVarTypesMap = map[ArgVarTypes]string{
	0: `File`,
	1: `Dir`,
	2: `Ext`,
	3: `Pos`,
	4: `Text`,
	5: `Prompt`,
}

// String returns the string representation
// of this ArgVarTypes value.
func (i ArgVarTypes) String() string {
	if str, ok := _ArgVarTypesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the ArgVarTypes value from its
// string representation, and returns an
// error if the string is invalid.
func (i *ArgVarTypes) SetString(s string) error {
	if val, ok := _ArgVarTypesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _ArgVarTypesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type ArgVarTypes")
}

// Int64 returns the ArgVarTypes value as an int64.
func (i ArgVarTypes) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the ArgVarTypes value from an int64.
func (i *ArgVarTypes) SetInt64(in int64) {
	*i = ArgVarTypes(in)
}

// Desc returns the description of the ArgVarTypes value.
func (i ArgVarTypes) Desc() string {
	if str, ok := _ArgVarTypesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ArgVarTypesValues returns all possible values
// for the type ArgVarTypes.
func ArgVarTypesValues() []ArgVarTypes {
	return _ArgVarTypesValues
}

// Values returns all possible values
// for the type ArgVarTypes.
func (i ArgVarTypes) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ArgVarTypesValues))
	for i, d := range _ArgVarTypesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type ArgVarTypes.
func (i ArgVarTypes) IsValid() bool {
	_, ok := _ArgVarTypesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i ArgVarTypes) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *ArgVarTypes) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}

var _FindLocValues = []FindLoc{0, 1, 2, 3}

// FindLocN is the highest valid value
// for type FindLoc, plus one.
const FindLocN FindLoc = 4

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _FindLocNoOp() {
	var x [1]struct{}
	_ = x[FindLocAll-(0)]
	_ = x[FindLocFile-(1)]
	_ = x[FindLocDir-(2)]
	_ = x[FindLocNotTop-(3)]
}

var _FindLocNameToValueMap = map[string]FindLoc{
	`All`:    0,
	`all`:    0,
	`File`:   1,
	`file`:   1,
	`Dir`:    2,
	`dir`:    2,
	`NotTop`: 3,
	`nottop`: 3,
}

var _FindLocDescMap = map[FindLoc]string{
	0: `FindLocAll finds in all open folders in the left file browser`,
	1: `FindLocFile only finds in the current active file`,
	2: `FindLocDir only finds in the directory of the current active file`,
	3: `FindLocNotTop finds in all open folders *except* the top-level folder`,
}

var _FindLocMap = map[FindLoc]string{
	0: `All`,
	1: `File`,
	2: `Dir`,
	3: `NotTop`,
}

// String returns the string representation
// of this FindLoc value.
func (i FindLoc) String() string {
	if str, ok := _FindLocMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the FindLoc value from its
// string representation, and returns an
// error if the string is invalid.
func (i *FindLoc) SetString(s string) error {
	if val, ok := _FindLocNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _FindLocNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type FindLoc")
}

// Int64 returns the FindLoc value as an int64.
func (i FindLoc) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the FindLoc value from an int64.
func (i *FindLoc) SetInt64(in int64) {
	*i = FindLoc(in)
}

// Desc returns the description of the FindLoc value.
func (i FindLoc) Desc() string {
	if str, ok := _FindLocDescMap[i]; ok {
		return str
	}
	return i.String()
}

// FindLocValues returns all possible values
// for the type FindLoc.
func FindLocValues() []FindLoc {
	return _FindLocValues
}

// Values returns all possible values
// for the type FindLoc.
func (i FindLoc) Values() []enums.Enum {
	res := make([]enums.Enum, len(_FindLocValues))
	for i, d := range _FindLocValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type FindLoc.
func (i FindLoc) IsValid() bool {
	_, ok := _FindLocMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i FindLoc) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *FindLoc) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}

var _KeyFunsValues = []KeyFuns{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}

// KeyFunsN is the highest valid value
// for type KeyFuns, plus one.
const KeyFunsN KeyFuns = 22

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _KeyFunsNoOp() {
	var x [1]struct{}
	_ = x[KeyFunNil-(0)]
	_ = x[KeyFunNeeds2-(1)]
	_ = x[KeyFunNextPanel-(2)]
	_ = x[KeyFunPrevPanel-(3)]
	_ = x[KeyFunFileOpen-(4)]
	_ = x[KeyFunBufSelect-(5)]
	_ = x[KeyFunBufClone-(6)]
	_ = x[KeyFunBufSave-(7)]
	_ = x[KeyFunBufSaveAs-(8)]
	_ = x[KeyFunBufClose-(9)]
	_ = x[KeyFunExecCmd-(10)]
	_ = x[KeyFunRectCopy-(11)]
	_ = x[KeyFunRectCut-(12)]
	_ = x[KeyFunRectPaste-(13)]
	_ = x[KeyFunRegCopy-(14)]
	_ = x[KeyFunRegPaste-(15)]
	_ = x[KeyFunCommentOut-(16)]
	_ = x[KeyFunIndent-(17)]
	_ = x[KeyFunJump-(18)]
	_ = x[KeyFunSetSplit-(19)]
	_ = x[KeyFunBuildProj-(20)]
	_ = x[KeyFunRunProj-(21)]
}

var _KeyFunsNameToValueMap = map[string]KeyFuns{
	`Nil`:        0,
	`nil`:        0,
	`Needs2`:     1,
	`needs2`:     1,
	`NextPanel`:  2,
	`nextpanel`:  2,
	`PrevPanel`:  3,
	`prevpanel`:  3,
	`FileOpen`:   4,
	`fileopen`:   4,
	`BufSelect`:  5,
	`bufselect`:  5,
	`BufClone`:   6,
	`bufclone`:   6,
	`BufSave`:    7,
	`bufsave`:    7,
	`BufSaveAs`:  8,
	`bufsaveas`:  8,
	`BufClose`:   9,
	`bufclose`:   9,
	`ExecCmd`:    10,
	`execcmd`:    10,
	`RectCopy`:   11,
	`rectcopy`:   11,
	`RectCut`:    12,
	`rectcut`:    12,
	`RectPaste`:  13,
	`rectpaste`:  13,
	`RegCopy`:    14,
	`regcopy`:    14,
	`RegPaste`:   15,
	`regpaste`:   15,
	`CommentOut`: 16,
	`commentout`: 16,
	`Indent`:     17,
	`indent`:     17,
	`Jump`:       18,
	`jump`:       18,
	`SetSplit`:   19,
	`setsplit`:   19,
	`BuildProj`:  20,
	`buildproj`:  20,
	`RunProj`:    21,
	`runproj`:    21,
}

var _KeyFunsDescMap = map[KeyFuns]string{
	0:  ``,
	1:  `special internal signal returned by KeyFun indicating need for second key`,
	2:  `move to next panel to the right`,
	3:  `move to prev panel to the left`,
	4:  `open a new file in active textview`,
	5:  `select an open buffer to edit in active textview`,
	6:  `open active file in other view`,
	7:  `save active textview buffer to its file`,
	8:  `save as active textview buffer to its file`,
	9:  `close active textview buffer`,
	10: `execute a command on active textview buffer`,
	11: `copy rectangle`,
	12: `cut rectangle`,
	13: `paste rectangle`,
	14: `copy selection to named register`,
	15: `paste selection from named register`,
	16: `comment out region`,
	17: `indent region`,
	18: `jump to line (same as keyfun.Jump)`,
	19: `set named splitter config`,
	20: `build overall project`,
	21: `run overall project`,
}

var _KeyFunsMap = map[KeyFuns]string{
	0:  `Nil`,
	1:  `Needs2`,
	2:  `NextPanel`,
	3:  `PrevPanel`,
	4:  `FileOpen`,
	5:  `BufSelect`,
	6:  `BufClone`,
	7:  `BufSave`,
	8:  `BufSaveAs`,
	9:  `BufClose`,
	10: `ExecCmd`,
	11: `RectCopy`,
	12: `RectCut`,
	13: `RectPaste`,
	14: `RegCopy`,
	15: `RegPaste`,
	16: `CommentOut`,
	17: `Indent`,
	18: `Jump`,
	19: `SetSplit`,
	20: `BuildProj`,
	21: `RunProj`,
}

// String returns the string representation
// of this KeyFuns value.
func (i KeyFuns) String() string {
	if str, ok := _KeyFunsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the KeyFuns value from its
// string representation, and returns an
// error if the string is invalid.
func (i *KeyFuns) SetString(s string) error {
	if val, ok := _KeyFunsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _KeyFunsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type KeyFuns")
}

// Int64 returns the KeyFuns value as an int64.
func (i KeyFuns) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the KeyFuns value from an int64.
func (i *KeyFuns) SetInt64(in int64) {
	*i = KeyFuns(in)
}

// Desc returns the description of the KeyFuns value.
func (i KeyFuns) Desc() string {
	if str, ok := _KeyFunsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// KeyFunsValues returns all possible values
// for the type KeyFuns.
func KeyFunsValues() []KeyFuns {
	return _KeyFunsValues
}

// Values returns all possible values
// for the type KeyFuns.
func (i KeyFuns) Values() []enums.Enum {
	res := make([]enums.Enum, len(_KeyFunsValues))
	for i, d := range _KeyFunsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type KeyFuns.
func (i KeyFuns) IsValid() bool {
	_, ok := _KeyFunsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i KeyFuns) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *KeyFuns) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}

var _SymScopesValues = []SymScopes{0, 1}

// SymScopesN is the highest valid value
// for type SymScopes, plus one.
const SymScopesN SymScopes = 2

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _SymScopesNoOp() {
	var x [1]struct{}
	_ = x[SymScopePackage-(0)]
	_ = x[SymScopeFile-(1)]
}

var _SymScopesNameToValueMap = map[string]SymScopes{
	`Package`: 0,
	`package`: 0,
	`File`:    1,
	`file`:    1,
}

var _SymScopesDescMap = map[SymScopes]string{
	0: `SymScopePackage scopes list of symbols to the package of the active file`,
	1: `SymScopeFile restricts the list of symbols to the active file`,
}

var _SymScopesMap = map[SymScopes]string{
	0: `Package`,
	1: `File`,
}

// String returns the string representation
// of this SymScopes value.
func (i SymScopes) String() string {
	if str, ok := _SymScopesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the SymScopes value from its
// string representation, and returns an
// error if the string is invalid.
func (i *SymScopes) SetString(s string) error {
	if val, ok := _SymScopesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _SymScopesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type SymScopes")
}

// Int64 returns the SymScopes value as an int64.
func (i SymScopes) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the SymScopes value from an int64.
func (i *SymScopes) SetInt64(in int64) {
	*i = SymScopes(in)
}

// Desc returns the description of the SymScopes value.
func (i SymScopes) Desc() string {
	if str, ok := _SymScopesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// SymScopesValues returns all possible values
// for the type SymScopes.
func SymScopesValues() []SymScopes {
	return _SymScopesValues
}

// Values returns all possible values
// for the type SymScopes.
func (i SymScopes) Values() []enums.Enum {
	res := make([]enums.Enum, len(_SymScopesValues))
	for i, d := range _SymScopesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type SymScopes.
func (i SymScopes) IsValid() bool {
	_, ok := _SymScopesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i SymScopes) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *SymScopes) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}
