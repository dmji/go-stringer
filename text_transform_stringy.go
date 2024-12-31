// Copyright 2014 daimonji. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"

	"github.com/gobeam/stringy"
)

//go:generate go run . -type=_TextConvertTo -nametransform=snake_case -trimprefix=_TextConvertTo -output=text_transform_stringy_string.go -genfromstringfn=true

type _TextConvertTo int

const (
	_TextConvertToNone _TextConvertTo = iota
	_TextConvertToCamelCase
	_TextConvertToSnakeCase
	_TextConvertToKebabCase
)

func transformTextTo(nametransform _TextConvertTo, name string) string {

	switch nametransform {
	case _TextConvertToCamelCase:
		name = stringy.New(name).CamelCase().ToLower()
	case _TextConvertToKebabCase:
		name = stringy.New(name).KebabCase().ToLower()
	case _TextConvertToSnakeCase:
		name = stringy.New(name).SnakeCase().ToLower()
	case _TextConvertToNone:
	default:
		panic(fmt.Sprintf("unexpected _TextConvertTo: %#v", nametransform))
	}

	return name
}
