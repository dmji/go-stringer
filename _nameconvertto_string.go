// Code generated by "stringer -type=_NameConvertTo -nameConvertToCase=snake_case"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[_NameConvertToNone-0]
	_ = x[_NameConvertToCamelCase-1]
	_ = x[_NameConvertToSnakeCase-2]
	_ = x[_NameConvertToKebabCase-3]
}

const __NameConvertTo_name = "name_convert_to_nonename_convert_to_camel_casename_convert_to_snake_casename_convert_to_kebab_case"

var __NameConvertTo_index = [...]uint8{0, 20, 46, 72, 98}

func (i _NameConvertTo) String() string {
	if i < 0 || i >= _NameConvertTo(len(__NameConvertTo_index)-1) {
		return "_NameConvertTo(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return __NameConvertTo_name[__NameConvertTo_index[i]:__NameConvertTo_index[i+1]]
}