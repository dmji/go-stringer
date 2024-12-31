// Code generated by "stringer -type=PatternRuneWithRange -linecomment=true -trimprefix=PatternRuneWithRange -nametransform=snake_case_lower -outputtransform=snake_case_lower"; DO NOT EDIT.

package sample

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PatternRuneWithRangeFirst-10]
	_ = x[PatternRuneWithRangeSecond-11]
	_ = x[PatternRuneWithRangeFourth-12]
	_ = x[PatternRuneWithRangeFifth-22]
	_ = x[PatternRuneWithRangeSixth-23]
	_ = x[PatternRuneWithRangeSeventh-99]
	_ = x[PatternRuneWithRangeEighth-101]
	_ = x[PatternRuneWithRangeNinth-202]
	_ = x[PatternRuneWithRangeTenth-203]
}

const (
	_PatternRuneWithRange_name_0 = "firstsecondfourth"
	_PatternRuneWithRange_name_1 = "fifthsixth"
	_PatternRuneWithRange_name_2 = "seventh"
	_PatternRuneWithRange_name_3 = "eighth"
	_PatternRuneWithRange_name_4 = "ninthtenth"
)

var (
	_PatternRuneWithRange_index_0 = [...]uint8{0, 5, 11, 17}
	_PatternRuneWithRange_index_1 = [...]uint8{0, 5, 10}
	_PatternRuneWithRange_index_4 = [...]uint8{0, 5, 10}
)

func (i PatternRuneWithRange) String() string {
	switch {
	case 10 <= i && i <= 12:
		i -= 10
		return _PatternRuneWithRange_name_0[_PatternRuneWithRange_index_0[i]:_PatternRuneWithRange_index_0[i+1]]
	case 22 <= i && i <= 23:
		i -= 22
		return _PatternRuneWithRange_name_1[_PatternRuneWithRange_index_1[i]:_PatternRuneWithRange_index_1[i+1]]
	case i == 99:
		return _PatternRuneWithRange_name_2
	case i == 101:
		return _PatternRuneWithRange_name_3
	case 202 <= i && i <= 203:
		i -= 202
		return _PatternRuneWithRange_name_4[_PatternRuneWithRange_index_4[i]:_PatternRuneWithRange_index_4[i+1]]
	default:
		return "PatternRuneWithRange(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
