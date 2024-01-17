// Code generated by "core generate"; DO NOT EDIT.

package cdebug

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"cogentcore.org/core/enums"
)

var _StatusValues = []Status{0, 1, 2, 3, 4, 5, 6, 7}

// StatusN is the highest valid value
// for type Status, plus one.
const StatusN Status = 8

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _StatusNoOp() {
	var x [1]struct{}
	_ = x[NotInit-(0)]
	_ = x[Error-(1)]
	_ = x[Building-(2)]
	_ = x[Ready-(3)]
	_ = x[Running-(4)]
	_ = x[Stopped-(5)]
	_ = x[Breakpoint-(6)]
	_ = x[Finished-(7)]
}

var _StatusNameToValueMap = map[string]Status{
	`NotInit`:    0,
	`notinit`:    0,
	`Error`:      1,
	`error`:      1,
	`Building`:   2,
	`building`:   2,
	`Ready`:      3,
	`ready`:      3,
	`Running`:    4,
	`running`:    4,
	`Stopped`:    5,
	`stopped`:    5,
	`Breakpoint`: 6,
	`breakpoint`: 6,
	`Finished`:   7,
	`finished`:   7,
}

var _StatusDescMap = map[Status]string{
	0: `NotInit is not initialized`,
	1: `Error means the debugger has an error -- usually from building`,
	2: `Building is building the exe for debugging`,
	3: `Ready means executable is built and ready to start (or restarted)`,
	4: `Running means the process is running`,
	5: `Stopped means the process has stopped (at a breakpoint, crash, or from single stepping)`,
	6: `Breakpoint means the process has stopped at a breakpoint`,
	7: `Finished means the process has finished running. See console for output and return value etc`,
}

var _StatusMap = map[Status]string{
	0: `NotInit`,
	1: `Error`,
	2: `Building`,
	3: `Ready`,
	4: `Running`,
	5: `Stopped`,
	6: `Breakpoint`,
	7: `Finished`,
}

// String returns the string representation
// of this Status value.
func (i Status) String() string {
	if str, ok := _StatusMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Status value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Status) SetString(s string) error {
	if val, ok := _StatusNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _StatusNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Status")
}

// Int64 returns the Status value as an int64.
func (i Status) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Status value from an int64.
func (i *Status) SetInt64(in int64) {
	*i = Status(in)
}

// Desc returns the description of the Status value.
func (i Status) Desc() string {
	if str, ok := _StatusDescMap[i]; ok {
		return str
	}
	return i.String()
}

// StatusValues returns all possible values
// for the type Status.
func StatusValues() []Status {
	return _StatusValues
}

// Values returns all possible values
// for the type Status.
func (i Status) Values() []enums.Enum {
	res := make([]enums.Enum, len(_StatusValues))
	for i, d := range _StatusValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Status.
func (i Status) IsValid() bool {
	_, ok := _StatusMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Status) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Status) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}
