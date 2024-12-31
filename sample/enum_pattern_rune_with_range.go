package sample

//go:generate go run ./../. -type=PatternRuneWithRange -linecomment=true -trimprefix=PatternRuneWithRange -nametransform=snake_case

type PatternRuneWithRange uint8

const (
	PatternRuneWithRangeFirst   PatternRuneWithRange = 10
	PatternRuneWithRangeSecond  PatternRuneWithRange = 11
	PatternRuneWithRangeFourth  PatternRuneWithRange = 12
	PatternRuneWithRangeFifth   PatternRuneWithRange = 22
	PatternRuneWithRangeSixth   PatternRuneWithRange = 23
	PatternRuneWithRangeSeventh PatternRuneWithRange = 99
	PatternRuneWithRangeEighth  PatternRuneWithRange = 101
	PatternRuneWithRangeNinth   PatternRuneWithRange = 202
	PatternRuneWithRangeTenth   PatternRuneWithRange = 203
)
