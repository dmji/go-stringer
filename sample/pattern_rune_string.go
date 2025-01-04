// Code generated by "go-stringer -type=PatternRune -trimprefix=PatternRune -linecomment=true -nametransform=snake_case_lower -outputtransform=snake_case_lower -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=snake_case_lower -extraconstsvaluesuffix=Presentation"; DO NOT EDIT.

package sample

import (
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PatternRuneFirst-0]
	_ = x[PatternRuneSecond-2]
	_ = x[PatternRuneThird-4]
	_ = x[PatternRuneFourth-6]
	_ = x[PatternRuneFifth-8]
	_ = x[PatternRuneSixth-10]
	_ = x[PatternRuneSeventh-12]
	_ = x[PatternRuneEighth-14]
	_ = x[PatternRuneNinth-16]
	_ = x[PatternRuneTenth-18]
}

const (
	_PatternRune_name_0 = "first"
	_PatternRune_name_1 = "second"
	_PatternRune_name_2 = "third"
	_PatternRune_name_3 = "fourth"
	_PatternRune_name_4 = "fifth"
	_PatternRune_name_5 = "sixth"
	_PatternRune_name_6 = "seventh"
	_PatternRune_name_7 = "eighth"
	_PatternRune_name_8 = "ninth"
	_PatternRune_name_9 = "tenth"
)

func (i PatternRune) String() string {
	switch {
	case i == 0:
		return _PatternRune_name_0
	case i == 2:
		return _PatternRune_name_1
	case i == 4:
		return _PatternRune_name_2
	case i == 6:
		return _PatternRune_name_3
	case i == 8:
		return _PatternRune_name_4
	case i == 10:
		return _PatternRune_name_5
	case i == 12:
		return _PatternRune_name_6
	case i == 14:
		return _PatternRune_name_7
	case i == 16:
		return _PatternRune_name_8
	case i == 18:
		return _PatternRune_name_9
	default:
		return "PatternRune(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

const (
	PatternRuneFirst_i18n_ID   = "sample_pattern_rune_first_presentation"
	PatternRuneSecond_i18n_ID  = "sample_pattern_rune_second_presentation"
	PatternRuneThird_i18n_ID   = "sample_pattern_rune_third_presentation"
	PatternRuneFourth_i18n_ID  = "sample_pattern_rune_fourth_presentation"
	PatternRuneFifth_i18n_ID   = "sample_pattern_rune_fifth_presentation"
	PatternRuneSixth_i18n_ID   = "sample_pattern_rune_sixth_presentation"
	PatternRuneSeventh_i18n_ID = "sample_pattern_rune_seventh_presentation"
	PatternRuneEighth_i18n_ID  = "sample_pattern_rune_eighth_presentation"
	PatternRuneNinth_i18n_ID   = "sample_pattern_rune_ninth_presentation"
	PatternRuneTenth_i18n_ID   = "sample_pattern_rune_tenth_presentation"
)
