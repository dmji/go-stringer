# Aboud go-stringer

Fork of https://github.com/golang/tools/blob/master/cmd/stringer/stringer.go

Forked on commit: 28cd7188fd2c8f3ede85d4b7a6734cc426422094 [ref](https://github.com/golang/tools/commit/28cd7188fd2c8f3ede85d4b7a6734cc426422094 "GitHub Golang/Tools Commit 28cd718") `Unfortunately, I don't know how to fork only cmd/stringer, so there is no direct connection and it seems impossible to synchronize changes automatically.`

## Motivation

To extend the generator functional to get automatically the text in the format I need and an additional function that returns enum value by text.

## Changes from original

- For my comfort, the monolithic stringer.go is divided into component files
- Added samples for each generator builder case
- Added flag _target_ to make it easier to run debug on samples
- Added flag _nametransform_ to convert strings from the original name to the one I need using github.com/gobeam/stringy `my pet project required a snake case and I didn't want to do it by hand`
- Added flag _fromstringgenfn_ to activate generation of ${TypeName}FromString() functions
- Added flag _outputtransform_ to convert strings of file names in similar way as names `example: originally it generate type PatternRuneWithRange by default as 'patternrunewithrange_string.go' but i want to get it as 'pattern_rune_with_range_string.go'`
- Added flag _i18ngenconst_ to activate generation consts with ID string {PackageName}{TypeValue}{Suffix}
- Added flag _i18ntransformto_ convert string id format similar to nametransform
- Added flag _i18nidsuffix_ to append custom text to the end of string id
- Added flag _i18nidnameprefix_ to append custom text before the const name
- Replaced flag _i18ngenconst_ to _i18nidnameprefix_ to append custom text afterthe const name, if not set the consts would not create

## How to use

Install:

```bash
go install github.com/dmji/go-stringer@latest
```

Information about flags can be query with execution without parameters:

```bash
go-stringer
```

Sample of generage comment:

```go
package main

//go:generate go-stringer -type=SimpleRune -trimprefix=SimpleRune -linecomment=true -nametransform=snake_case_lower -outputtransform=snake_case_lower -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=snake_case_lower -extraconstsvaluesuffix=Presentation

type SimpleRune int16

const (
	SimpleRuneFirst SimpleRune = iota * -1
	SimpleRuneSecond
	SimpleRuneThird
	SimpleRuneFourth
)
```
