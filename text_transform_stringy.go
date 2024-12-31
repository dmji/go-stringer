// Copyright 2014 daimonji. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"strings"

	"github.com/gobeam/stringy"
)

//go:generate go run . -type=_TextConvertTo -nametransform=snake_case_lower -trimprefix=_TextConvertTo -output=text_transform_stringy_string.go -genfromstringfn=true

type _TextConvertTo int

const (
	_TextConvertToNone _TextConvertTo = iota

	_TextConvertToCamelCase
	_TextConvertToSnakeCase
	_TextConvertToKebabCase

	_TextConvertToCamelCaseLower
	_TextConvertToSnakeCaseLower
	_TextConvertToKebabCaseLower

	_TextConvertToCamelCaseUpper
	_TextConvertToSnakeCaseUpper
	_TextConvertToKebabCaseUpper

	_TextConvertToCount
)

func transformTextTo(nametransform _TextConvertTo, name string) string {

	switch nametransform {
	case _TextConvertToCamelCase:
		name = stringy.New(name).CamelCase().Get()
	case _TextConvertToKebabCase:
		name = stringy.New(name).KebabCase().Get()
	case _TextConvertToSnakeCase:
		name = stringy.New(name).SnakeCase().Get()

	case _TextConvertToCamelCaseLower:
		name = stringy.New(name).CamelCase().ToLower()
	case _TextConvertToKebabCaseLower:
		name = stringy.New(name).KebabCase().ToLower()
	case _TextConvertToSnakeCaseLower:
		name = stringy.New(name).SnakeCase().ToLower()

	case _TextConvertToCamelCaseUpper:
		name = stringy.New(name).CamelCase().ToUpper()
	case _TextConvertToKebabCaseUpper:
		name = stringy.New(name).KebabCase().ToUpper()
	case _TextConvertToSnakeCaseUpper:
		name = stringy.New(name).SnakeCase().ToUpper()

	case _TextConvertToNone:
	default:
		panic(fmt.Sprintf("unexpected _TextConvertTo: %#v", nametransform))
	}

	return name
}

func availableValuesForFlagsDefault() string {
	res := make([]string, 0, _TextConvertToCount)
	for i := _TextConvertToNone; i <= _TextConvertToCount; i++ {
		res = append(res, _TextConvertTo(i).String())
	}
	return strings.Join(res, ", ")
}
