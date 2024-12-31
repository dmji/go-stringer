// Copyright 2014 daimonji. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"strings"

	"github.com/gobeam/stringy"
)

// generate two time as if we add new enum values we use prev state and loose order but generate correct value via linecomment and second run align other values to snake case lower

//go:generate go run . -type=_TextConvertTo -nametransform=snake_case_lower -trimprefix=_TextConvertTo -output=text_transform_stringy_string.go -fromstringgenfn=true -linecomment=true && go run . -type=_TextConvertTo -nametransform=snake_case_lower -trimprefix=_TextConvertTo -output=text_transform_stringy_string.go -fromstringgenfn=true -linecomment=true

type _TextConvertTo int

const (
	_TextConvertToNone _TextConvertTo = iota //none
	_TextConvertToLower
	_TextConvertToUpper

	_TextConvertToCamelCase
	_TextConvertToSnakeCase
	_TextConvertToKebabCase
	_TextConvertToPascalCase

	_TextConvertToCamelCaseLower
	_TextConvertToSnakeCaseLower //snake_case_lower
	_TextConvertToKebabCaseLower
	_TextConvertToPascalCaseLower

	_TextConvertToCamelCaseUpper
	_TextConvertToSnakeCaseUpper
	_TextConvertToKebabCaseUpper
	_TextConvertToPascalCaseUpper
)

func transformTextTo(nametransform _TextConvertTo, name string) string {

	switch nametransform {

	case _TextConvertToNone:
		return name
	case _TextConvertToLower:
		return strings.ToLower(name)
	case _TextConvertToUpper:
		return strings.ToUpper(name)

	case _TextConvertToCamelCase:
		return stringy.New(name).CamelCase().Get()
	case _TextConvertToKebabCase:
		return stringy.New(name).KebabCase().Get()
	case _TextConvertToSnakeCase:
		return stringy.New(name).SnakeCase().Get()
	case _TextConvertToPascalCase:
		return stringy.New(name).PascalCase().Get()

	case _TextConvertToCamelCaseLower:
		return stringy.New(name).CamelCase().ToLower()
	case _TextConvertToKebabCaseLower:
		return stringy.New(name).KebabCase().ToLower()
	case _TextConvertToSnakeCaseLower:
		return stringy.New(name).SnakeCase().ToLower()
	case _TextConvertToPascalCaseLower:
		return stringy.New(name).PascalCase().ToLower()

	case _TextConvertToCamelCaseUpper:
		return stringy.New(name).CamelCase().ToUpper()
	case _TextConvertToKebabCaseUpper:
		return stringy.New(name).KebabCase().ToUpper()
	case _TextConvertToSnakeCaseUpper:
		return stringy.New(name).SnakeCase().ToUpper()
	case _TextConvertToPascalCaseUpper:
		return stringy.New(name).PascalCase().ToUpper()

	default:
		panic(fmt.Sprintf("unexpected _TextConvertTo: %#v", nametransform))
	}
}

func availableValuesForFlagsDefault() string {
	n := _TextConvertTo(len(__TextConvertTo_index) - 1)
	res := make([]string, 0, n)
	for i := _TextConvertToNone; i <= n; i++ {
		res = append(res, i.String())
	}
	return strings.Join(res, ", ")
}
