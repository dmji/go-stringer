# Aboud go-stringer

Fork of https://github.com/golang/tools/blob/master/cmd/stringer/stringer.go

Forked on commit: 28cd7188fd2c8f3ede85d4b7a6734cc426422094 [ref](https://github.com/golang/tools/commit/28cd7188fd2c8f3ede85d4b7a6734cc426422094 "GitHub Golang/Tools Commit 28cd718") `Unfortunately, I don't know how to fork only cmd/stringer, so there is no direct connection and it seems impossible to synchronize changes automatically.`

## Motivation

To extend the generator functional to get automatically the text in the format I need and an additional function that returns enum value by text.

## Changes from original

- For my comfort, the monolithic stringer.go is divided into component files
- Added samples for each generator builder case
- Added flag target to make it easier to run debug on samples
- Added flag nametransform to convert strings from the original name to the one I need using github.com/gobeam/stringy `my pet project required a snake case and I didn't want to do it by hand`
- Added flag genfromstringfn to activate generation of ${TypeName}FromString() functions
- Added flag outputtransform to convert strings of file names in similar way as names `example: originally it generate type PatternRuneWithRange by default as 'patternrunewithrange_string.go' but i want to get it as 'pattern_rune_with_range_string.go'`

## TODO:

- [ ] add some generation for go-i18n to eleminate thinking of DefaultMessage.ID
