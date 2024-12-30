package main

//go:generate go run . -type=_NameConvertTo -nameConvertToCase=snake_case

type _NameConvertTo int

const (
	_NameConvertToNone _NameConvertTo = iota
	_NameConvertToCamelCase
	_NameConvertToSnakeCase
	_NameConvertToKebabCase
)
