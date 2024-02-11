// Code generated by "core generate"; DO NOT EDIT.

package vector

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"sync/atomic"

	"cogentcore.org/core/enums"
	"cogentcore.org/core/gi"
)

var _AlignAnchorsValues = []AlignAnchors{0, 1, 2, 3}

// AlignAnchorsN is the highest valid value
// for type AlignAnchors, plus one.
const AlignAnchorsN AlignAnchors = 4

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _AlignAnchorsNoOp() {
	var x [1]struct{}
	_ = x[AlignFirst-(0)]
	_ = x[AlignLast-(1)]
	_ = x[AlignDrawing-(2)]
	_ = x[AlignSelectBox-(3)]
}

var _AlignAnchorsNameToValueMap = map[string]AlignAnchors{
	`AlignFirst`:     0,
	`AlignLast`:      1,
	`AlignDrawing`:   2,
	`AlignSelectBox`: 3,
}

var _AlignAnchorsDescMap = map[AlignAnchors]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
}

var _AlignAnchorsMap = map[AlignAnchors]string{
	0: `AlignFirst`,
	1: `AlignLast`,
	2: `AlignDrawing`,
	3: `AlignSelectBox`,
}

// String returns the string representation
// of this AlignAnchors value.
func (i AlignAnchors) String() string {
	if str, ok := _AlignAnchorsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the AlignAnchors value from its
// string representation, and returns an
// error if the string is invalid.
func (i *AlignAnchors) SetString(s string) error {
	if val, ok := _AlignAnchorsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type AlignAnchors")
}

// Int64 returns the AlignAnchors value as an int64.
func (i AlignAnchors) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the AlignAnchors value from an int64.
func (i *AlignAnchors) SetInt64(in int64) {
	*i = AlignAnchors(in)
}

// Desc returns the description of the AlignAnchors value.
func (i AlignAnchors) Desc() string {
	if str, ok := _AlignAnchorsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// AlignAnchorsValues returns all possible values
// for the type AlignAnchors.
func AlignAnchorsValues() []AlignAnchors {
	return _AlignAnchorsValues
}

// Values returns all possible values
// for the type AlignAnchors.
func (i AlignAnchors) Values() []enums.Enum {
	res := make([]enums.Enum, len(_AlignAnchorsValues))
	for i, d := range _AlignAnchorsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type AlignAnchors.
func (i AlignAnchors) IsValid() bool {
	_, ok := _AlignAnchorsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i AlignAnchors) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *AlignAnchors) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("AlignAnchors.UnmarshalText:", err)
	}
	return nil
}

var _AlignsValues = []Aligns{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

// AlignsN is the highest valid value
// for type Aligns, plus one.
const AlignsN Aligns = 12

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _AlignsNoOp() {
	var x [1]struct{}
	_ = x[AlignRightAnchor-(0)]
	_ = x[AlignLeft-(1)]
	_ = x[AlignCenter-(2)]
	_ = x[AlignRight-(3)]
	_ = x[AlignLeftAnchor-(4)]
	_ = x[AlignBaselineHoriz-(5)]
	_ = x[AlignBottomAnchor-(6)]
	_ = x[AlignTop-(7)]
	_ = x[AlignMiddle-(8)]
	_ = x[AlignBottom-(9)]
	_ = x[AlignTopAnchor-(10)]
	_ = x[AlignBaselineVert-(11)]
}

var _AlignsNameToValueMap = map[string]Aligns{
	`AlignRightAnchor`:   0,
	`AlignLeft`:          1,
	`AlignCenter`:        2,
	`AlignRight`:         3,
	`AlignLeftAnchor`:    4,
	`AlignBaselineHoriz`: 5,
	`AlignBottomAnchor`:  6,
	`AlignTop`:           7,
	`AlignMiddle`:        8,
	`AlignBottom`:        9,
	`AlignTopAnchor`:     10,
	`AlignBaselineVert`:  11,
}

var _AlignsDescMap = map[Aligns]string{
	0:  `align right edges to left edge of anchor item`,
	1:  `align left edges`,
	2:  `align horizontal centers`,
	3:  `align right edges`,
	4:  `align left edges to right edge of anchor item`,
	5:  `align left text baseline edges`,
	6:  `align bottom edges to top edge of anchor item`,
	7:  `align top edges`,
	8:  `align middle vertical point`,
	9:  `align bottom edges`,
	10: `align top edges to bottom edge of anchor item`,
	11: `align baseline points vertically`,
}

var _AlignsMap = map[Aligns]string{
	0:  `AlignRightAnchor`,
	1:  `AlignLeft`,
	2:  `AlignCenter`,
	3:  `AlignRight`,
	4:  `AlignLeftAnchor`,
	5:  `AlignBaselineHoriz`,
	6:  `AlignBottomAnchor`,
	7:  `AlignTop`,
	8:  `AlignMiddle`,
	9:  `AlignBottom`,
	10: `AlignTopAnchor`,
	11: `AlignBaselineVert`,
}

// String returns the string representation
// of this Aligns value.
func (i Aligns) String() string {
	if str, ok := _AlignsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Aligns value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Aligns) SetString(s string) error {
	if val, ok := _AlignsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Aligns")
}

// Int64 returns the Aligns value as an int64.
func (i Aligns) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Aligns value from an int64.
func (i *Aligns) SetInt64(in int64) {
	*i = Aligns(in)
}

// Desc returns the description of the Aligns value.
func (i Aligns) Desc() string {
	if str, ok := _AlignsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// AlignsValues returns all possible values
// for the type Aligns.
func AlignsValues() []Aligns {
	return _AlignsValues
}

// Values returns all possible values
// for the type Aligns.
func (i Aligns) Values() []enums.Enum {
	res := make([]enums.Enum, len(_AlignsValues))
	for i, d := range _AlignsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Aligns.
func (i Aligns) IsValid() bool {
	_, ok := _AlignsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Aligns) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Aligns) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Aligns.UnmarshalText:", err)
	}
	return nil
}

var _BBoxPointsValues = []BBoxPoints{0, 1, 2, 3, 4, 5}

// BBoxPointsN is the highest valid value
// for type BBoxPoints, plus one.
const BBoxPointsN BBoxPoints = 6

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _BBoxPointsNoOp() {
	var x [1]struct{}
	_ = x[BBLeft-(0)]
	_ = x[BBCenter-(1)]
	_ = x[BBRight-(2)]
	_ = x[BBTop-(3)]
	_ = x[BBMiddle-(4)]
	_ = x[BBBottom-(5)]
}

var _BBoxPointsNameToValueMap = map[string]BBoxPoints{
	`BBLeft`:   0,
	`BBCenter`: 1,
	`BBRight`:  2,
	`BBTop`:    3,
	`BBMiddle`: 4,
	`BBBottom`: 5,
}

var _BBoxPointsDescMap = map[BBoxPoints]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
	4: ``,
	5: ``,
}

var _BBoxPointsMap = map[BBoxPoints]string{
	0: `BBLeft`,
	1: `BBCenter`,
	2: `BBRight`,
	3: `BBTop`,
	4: `BBMiddle`,
	5: `BBBottom`,
}

// String returns the string representation
// of this BBoxPoints value.
func (i BBoxPoints) String() string {
	if str, ok := _BBoxPointsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the BBoxPoints value from its
// string representation, and returns an
// error if the string is invalid.
func (i *BBoxPoints) SetString(s string) error {
	if val, ok := _BBoxPointsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type BBoxPoints")
}

// Int64 returns the BBoxPoints value as an int64.
func (i BBoxPoints) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the BBoxPoints value from an int64.
func (i *BBoxPoints) SetInt64(in int64) {
	*i = BBoxPoints(in)
}

// Desc returns the description of the BBoxPoints value.
func (i BBoxPoints) Desc() string {
	if str, ok := _BBoxPointsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// BBoxPointsValues returns all possible values
// for the type BBoxPoints.
func BBoxPointsValues() []BBoxPoints {
	return _BBoxPointsValues
}

// Values returns all possible values
// for the type BBoxPoints.
func (i BBoxPoints) Values() []enums.Enum {
	res := make([]enums.Enum, len(_BBoxPointsValues))
	for i, d := range _BBoxPointsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type BBoxPoints.
func (i BBoxPoints) IsValid() bool {
	_, ok := _BBoxPointsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i BBoxPoints) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *BBoxPoints) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("BBoxPoints.UnmarshalText:", err)
	}
	return nil
}

var _MarkerColorsValues = []MarkerColors{0, 1, 2}

// MarkerColorsN is the highest valid value
// for type MarkerColors, plus one.
const MarkerColorsN MarkerColors = 3

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _MarkerColorsNoOp() {
	var x [1]struct{}
	_ = x[MarkerDef-(0)]
	_ = x[MarkerCopy-(1)]
	_ = x[MarkerCust-(2)]
}

var _MarkerColorsNameToValueMap = map[string]MarkerColors{
	`Def`:  0,
	`Copy`: 1,
	`Cust`: 2,
}

var _MarkerColorsDescMap = map[MarkerColors]string{
	0: `use the default color of marker (typically black)`,
	1: `copy color of object using marker (create separate marker object per element)`,
	2: `marker has its own separate custom color`,
}

var _MarkerColorsMap = map[MarkerColors]string{
	0: `Def`,
	1: `Copy`,
	2: `Cust`,
}

// String returns the string representation
// of this MarkerColors value.
func (i MarkerColors) String() string {
	if str, ok := _MarkerColorsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the MarkerColors value from its
// string representation, and returns an
// error if the string is invalid.
func (i *MarkerColors) SetString(s string) error {
	if val, ok := _MarkerColorsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type MarkerColors")
}

// Int64 returns the MarkerColors value as an int64.
func (i MarkerColors) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the MarkerColors value from an int64.
func (i *MarkerColors) SetInt64(in int64) {
	*i = MarkerColors(in)
}

// Desc returns the description of the MarkerColors value.
func (i MarkerColors) Desc() string {
	if str, ok := _MarkerColorsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// MarkerColorsValues returns all possible values
// for the type MarkerColors.
func MarkerColorsValues() []MarkerColors {
	return _MarkerColorsValues
}

// Values returns all possible values
// for the type MarkerColors.
func (i MarkerColors) Values() []enums.Enum {
	res := make([]enums.Enum, len(_MarkerColorsValues))
	for i, d := range _MarkerColorsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type MarkerColors.
func (i MarkerColors) IsValid() bool {
	_, ok := _MarkerColorsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i MarkerColors) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *MarkerColors) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("MarkerColors.UnmarshalText:", err)
	}
	return nil
}

var _PaintTypesValues = []PaintTypes{0, 1, 2, 3, 4}

// PaintTypesN is the highest valid value
// for type PaintTypes, plus one.
const PaintTypesN PaintTypes = 5

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _PaintTypesNoOp() {
	var x [1]struct{}
	_ = x[PaintOff-(0)]
	_ = x[PaintSolid-(1)]
	_ = x[PaintLinear-(2)]
	_ = x[PaintRadial-(3)]
	_ = x[PaintInherit-(4)]
}

var _PaintTypesNameToValueMap = map[string]PaintTypes{
	`Off`:     0,
	`Solid`:   1,
	`Linear`:  2,
	`Radial`:  3,
	`Inherit`: 4,
}

var _PaintTypesDescMap = map[PaintTypes]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
	4: ``,
}

var _PaintTypesMap = map[PaintTypes]string{
	0: `Off`,
	1: `Solid`,
	2: `Linear`,
	3: `Radial`,
	4: `Inherit`,
}

// String returns the string representation
// of this PaintTypes value.
func (i PaintTypes) String() string {
	if str, ok := _PaintTypesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the PaintTypes value from its
// string representation, and returns an
// error if the string is invalid.
func (i *PaintTypes) SetString(s string) error {
	if val, ok := _PaintTypesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type PaintTypes")
}

// Int64 returns the PaintTypes value as an int64.
func (i PaintTypes) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the PaintTypes value from an int64.
func (i *PaintTypes) SetInt64(in int64) {
	*i = PaintTypes(in)
}

// Desc returns the description of the PaintTypes value.
func (i PaintTypes) Desc() string {
	if str, ok := _PaintTypesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// PaintTypesValues returns all possible values
// for the type PaintTypes.
func PaintTypesValues() []PaintTypes {
	return _PaintTypesValues
}

// Values returns all possible values
// for the type PaintTypes.
func (i PaintTypes) Values() []enums.Enum {
	res := make([]enums.Enum, len(_PaintTypesValues))
	for i, d := range _PaintTypesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type PaintTypes.
func (i PaintTypes) IsValid() bool {
	_, ok := _PaintTypesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i PaintTypes) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *PaintTypes) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("PaintTypes.UnmarshalText:", err)
	}
	return nil
}

var _StdSizesValues = []StdSizes{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}

// StdSizesN is the highest valid value
// for type StdSizes, plus one.
const StdSizesN StdSizes = 22

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _StdSizesNoOp() {
	var x [1]struct{}
	_ = x[CustomSize-(0)]
	_ = x[Img1280x720-(1)]
	_ = x[Img1920x1080-(2)]
	_ = x[Img3840x2160-(3)]
	_ = x[Img7680x4320-(4)]
	_ = x[Img1024x768-(5)]
	_ = x[Img720x480-(6)]
	_ = x[Img640x480-(7)]
	_ = x[Img320x240-(8)]
	_ = x[A4-(9)]
	_ = x[USLetter-(10)]
	_ = x[USLegal-(11)]
	_ = x[A0-(12)]
	_ = x[A1-(13)]
	_ = x[A2-(14)]
	_ = x[A3-(15)]
	_ = x[A5-(16)]
	_ = x[A6-(17)]
	_ = x[A7-(18)]
	_ = x[A8-(19)]
	_ = x[A9-(20)]
	_ = x[A10-(21)]
}

var _StdSizesNameToValueMap = map[string]StdSizes{
	`CustomSize`:   0,
	`Img1280x720`:  1,
	`Img1920x1080`: 2,
	`Img3840x2160`: 3,
	`Img7680x4320`: 4,
	`Img1024x768`:  5,
	`Img720x480`:   6,
	`Img640x480`:   7,
	`Img320x240`:   8,
	`A4`:           9,
	`USLetter`:     10,
	`USLegal`:      11,
	`A0`:           12,
	`A1`:           13,
	`A2`:           14,
	`A3`:           15,
	`A5`:           16,
	`A6`:           17,
	`A7`:           18,
	`A8`:           19,
	`A9`:           20,
	`A10`:          21,
}

var _StdSizesDescMap = map[StdSizes]string{
	0:  `CustomSize = nonstandard`,
	1:  `Image 1280x720 Px = 720p`,
	2:  `Image 1920x1080 Px = 1080p HD`,
	3:  `Image 3840x2160 Px = 4K`,
	4:  `Image 7680x4320 Px = 8K`,
	5:  `Image 1024x768 Px = XGA`,
	6:  `Image 720x480 Px = DVD`,
	7:  `Image 640x480 Px = VGA`,
	8:  `Image 320x240 Px = old CRT`,
	9:  `A4 = 210 x 297 mm`,
	10: `USLetter = 8.5 x 11 in = 612 x 792 pt`,
	11: `USLegal = 8.5 x 14 in = 612 x 1008 pt`,
	12: `A0 = 841 x 1189 mm`,
	13: `A1 = 594 x 841 mm`,
	14: `A2 = 420 x 594 mm`,
	15: `A3 = 297 x 420 mm`,
	16: `A5 = 148 x 210 mm`,
	17: `A6 = 105 x 148 mm`,
	18: `A7 = 74 x 105`,
	19: `A8 = 52 x 74 mm`,
	20: `A9 = 37 x 52`,
	21: `A10 = 26 x 37`,
}

var _StdSizesMap = map[StdSizes]string{
	0:  `CustomSize`,
	1:  `Img1280x720`,
	2:  `Img1920x1080`,
	3:  `Img3840x2160`,
	4:  `Img7680x4320`,
	5:  `Img1024x768`,
	6:  `Img720x480`,
	7:  `Img640x480`,
	8:  `Img320x240`,
	9:  `A4`,
	10: `USLetter`,
	11: `USLegal`,
	12: `A0`,
	13: `A1`,
	14: `A2`,
	15: `A3`,
	16: `A5`,
	17: `A6`,
	18: `A7`,
	19: `A8`,
	20: `A9`,
	21: `A10`,
}

// String returns the string representation
// of this StdSizes value.
func (i StdSizes) String() string {
	if str, ok := _StdSizesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the StdSizes value from its
// string representation, and returns an
// error if the string is invalid.
func (i *StdSizes) SetString(s string) error {
	if val, ok := _StdSizesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type StdSizes")
}

// Int64 returns the StdSizes value as an int64.
func (i StdSizes) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the StdSizes value from an int64.
func (i *StdSizes) SetInt64(in int64) {
	*i = StdSizes(in)
}

// Desc returns the description of the StdSizes value.
func (i StdSizes) Desc() string {
	if str, ok := _StdSizesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// StdSizesValues returns all possible values
// for the type StdSizes.
func StdSizesValues() []StdSizes {
	return _StdSizesValues
}

// Values returns all possible values
// for the type StdSizes.
func (i StdSizes) Values() []enums.Enum {
	res := make([]enums.Enum, len(_StdSizesValues))
	for i, d := range _StdSizesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type StdSizes.
func (i StdSizes) IsValid() bool {
	_, ok := _StdSizesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i StdSizes) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *StdSizes) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("StdSizes.UnmarshalText:", err)
	}
	return nil
}

var _SpritesValues = []Sprites{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

// SpritesN is the highest valid value
// for type Sprites, plus one.
const SpritesN Sprites = 15

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _SpritesNoOp() {
	var x [1]struct{}
	_ = x[SpUnk-(0)]
	_ = x[SpReshapeBBox-(1)]
	_ = x[SpSelBBox-(2)]
	_ = x[SpNodePoint-(3)]
	_ = x[SpNodeCtrl-(4)]
	_ = x[SpRubberBand-(5)]
	_ = x[SpAlignMatch-(6)]
	_ = x[SpBBoxUpL-(7)]
	_ = x[SpBBoxUpC-(8)]
	_ = x[SpBBoxUpR-(9)]
	_ = x[SpBBoxDnL-(10)]
	_ = x[SpBBoxDnC-(11)]
	_ = x[SpBBoxDnR-(12)]
	_ = x[SpBBoxLfM-(13)]
	_ = x[SpBBoxRtM-(14)]
}

var _SpritesNameToValueMap = map[string]Sprites{
	`SpUnk`:         0,
	`SpReshapeBBox`: 1,
	`SpSelBBox`:     2,
	`SpNodePoint`:   3,
	`SpNodeCtrl`:    4,
	`SpRubberBand`:  5,
	`SpAlignMatch`:  6,
	`SpBBoxUpL`:     7,
	`SpBBoxUpC`:     8,
	`SpBBoxUpR`:     9,
	`SpBBoxDnL`:     10,
	`SpBBoxDnC`:     11,
	`SpBBoxDnR`:     12,
	`SpBBoxLfM`:     13,
	`SpBBoxRtM`:     14,
}

var _SpritesDescMap = map[Sprites]string{
	0:  `SpUnk is an unknown sprite type`,
	1:  `SpReshapeBBox is a reshape bbox -- the overall active selection BBox for active manipulation`,
	2:  `SpSelBBox is a selection bounding box -- display only`,
	3:  `SpNodePoint is a main coordinate point for path node`,
	4:  `SpNodeCtrl is a control coordinate point for path node`,
	5:  `SpRubberBand is the draggable sel box subtyp = UpC, LfM, RtM, DnC for sides`,
	6:  `SpAlignMatch is an alignment match (n of these), subtyp is actually BBoxPoints so we just hack cast that`,
	7:  `Sprite bounding boxes are set as a &#34;bbox&#34; property on sprites`,
	8:  ``,
	9:  ``,
	10: ``,
	11: ``,
	12: ``,
	13: ``,
	14: ``,
}

var _SpritesMap = map[Sprites]string{
	0:  `SpUnk`,
	1:  `SpReshapeBBox`,
	2:  `SpSelBBox`,
	3:  `SpNodePoint`,
	4:  `SpNodeCtrl`,
	5:  `SpRubberBand`,
	6:  `SpAlignMatch`,
	7:  `SpBBoxUpL`,
	8:  `SpBBoxUpC`,
	9:  `SpBBoxUpR`,
	10: `SpBBoxDnL`,
	11: `SpBBoxDnC`,
	12: `SpBBoxDnR`,
	13: `SpBBoxLfM`,
	14: `SpBBoxRtM`,
}

// String returns the string representation
// of this Sprites value.
func (i Sprites) String() string {
	if str, ok := _SpritesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Sprites value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Sprites) SetString(s string) error {
	if val, ok := _SpritesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Sprites")
}

// Int64 returns the Sprites value as an int64.
func (i Sprites) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Sprites value from an int64.
func (i *Sprites) SetInt64(in int64) {
	*i = Sprites(in)
}

// Desc returns the description of the Sprites value.
func (i Sprites) Desc() string {
	if str, ok := _SpritesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// SpritesValues returns all possible values
// for the type Sprites.
func SpritesValues() []Sprites {
	return _SpritesValues
}

// Values returns all possible values
// for the type Sprites.
func (i Sprites) Values() []enums.Enum {
	res := make([]enums.Enum, len(_SpritesValues))
	for i, d := range _SpritesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Sprites.
func (i Sprites) IsValid() bool {
	_, ok := _SpritesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Sprites) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Sprites) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Sprites.UnmarshalText:", err)
	}
	return nil
}

var _ToolsValues = []Tools{0, 1, 2, 3, 4, 5}

// ToolsN is the highest valid value
// for type Tools, plus one.
const ToolsN Tools = 6

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ToolsNoOp() {
	var x [1]struct{}
	_ = x[SelectTool-(0)]
	_ = x[NodeTool-(1)]
	_ = x[RectTool-(2)]
	_ = x[EllipseTool-(3)]
	_ = x[BezierTool-(4)]
	_ = x[TextTool-(5)]
}

var _ToolsNameToValueMap = map[string]Tools{
	`SelectTool`:  0,
	`NodeTool`:    1,
	`RectTool`:    2,
	`EllipseTool`: 3,
	`BezierTool`:  4,
	`TextTool`:    5,
}

var _ToolsDescMap = map[Tools]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
	4: ``,
	5: ``,
}

var _ToolsMap = map[Tools]string{
	0: `SelectTool`,
	1: `NodeTool`,
	2: `RectTool`,
	3: `EllipseTool`,
	4: `BezierTool`,
	5: `TextTool`,
}

// String returns the string representation
// of this Tools value.
func (i Tools) String() string {
	if str, ok := _ToolsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Tools value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Tools) SetString(s string) error {
	if val, ok := _ToolsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Tools")
}

// Int64 returns the Tools value as an int64.
func (i Tools) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Tools value from an int64.
func (i *Tools) SetInt64(in int64) {
	*i = Tools(in)
}

// Desc returns the description of the Tools value.
func (i Tools) Desc() string {
	if str, ok := _ToolsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ToolsValues returns all possible values
// for the type Tools.
func ToolsValues() []Tools {
	return _ToolsValues
}

// Values returns all possible values
// for the type Tools.
func (i Tools) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ToolsValues))
	for i, d := range _ToolsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Tools.
func (i Tools) IsValid() bool {
	_, ok := _ToolsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Tools) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Tools) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Tools.UnmarshalText:", err)
	}
	return nil
}

var _VectorViewFlagsValues = []VectorViewFlags{8}

// VectorViewFlagsN is the highest valid value
// for type VectorViewFlags, plus one.
const VectorViewFlagsN VectorViewFlags = 9

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _VectorViewFlagsNoOp() {
	var x [1]struct{}
	_ = x[VectorViewAutoSaving-(8)]
}

var _VectorViewFlagsNameToValueMap = map[string]VectorViewFlags{
	`VectorViewAutoSaving`: 8,
}

var _VectorViewFlagsDescMap = map[VectorViewFlags]string{
	8: `VectorViewAutoSaving means`,
}

var _VectorViewFlagsMap = map[VectorViewFlags]string{
	8: `VectorViewAutoSaving`,
}

// String returns the string representation
// of this VectorViewFlags value.
func (i VectorViewFlags) String() string {
	str := ""
	for _, ie := range gi.WidgetFlagsValues() {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	for _, ie := range _VectorViewFlagsValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string
// representation of this VectorViewFlags value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i VectorViewFlags) BitIndexString() string {
	if str, ok := _VectorViewFlagsMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).BitIndexString()
}

// SetString sets the VectorViewFlags value from its
// string representation, and returns an
// error if the string is invalid.
func (i *VectorViewFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the VectorViewFlags value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *VectorViewFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _VectorViewFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			err := (*gi.WidgetFlags)(i).SetStringOr(flg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Int64 returns the VectorViewFlags value as an int64.
func (i VectorViewFlags) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the VectorViewFlags value from an int64.
func (i *VectorViewFlags) SetInt64(in int64) {
	*i = VectorViewFlags(in)
}

// Desc returns the description of the VectorViewFlags value.
func (i VectorViewFlags) Desc() string {
	if str, ok := _VectorViewFlagsDescMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).Desc()
}

// VectorViewFlagsValues returns all possible values
// for the type VectorViewFlags.
func VectorViewFlagsValues() []VectorViewFlags {
	es := gi.WidgetFlagsValues()
	res := make([]VectorViewFlags, len(es))
	for i, e := range es {
		res[i] = VectorViewFlags(e)
	}
	res = append(res, _VectorViewFlagsValues...)
	return res
}

// Values returns all possible values
// for the type VectorViewFlags.
func (i VectorViewFlags) Values() []enums.Enum {
	es := gi.WidgetFlagsValues()
	les := len(es)
	res := make([]enums.Enum, les+len(_VectorViewFlagsValues))
	for i, d := range es {
		res[i] = d
	}
	for i, d := range _VectorViewFlagsValues {
		res[i+les] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type VectorViewFlags.
func (i VectorViewFlags) IsValid() bool {
	_, ok := _VectorViewFlagsMap[i]
	if !ok {
		return gi.WidgetFlags(i).IsValid()
	}
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i VectorViewFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *VectorViewFlags) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i VectorViewFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *VectorViewFlags) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("VectorViewFlags.UnmarshalText:", err)
	}
	return nil
}
