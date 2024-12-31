package sample

//go:generate go run ./../. -type=PatternRune -linecomment=true -trimprefix=PatternRune -nametransform=snake_case

type PatternRune uint8

const (
	PatternRuneFirst PatternRune = iota * 2
	PatternRuneSecond
	PatternRuneThird
	PatternRuneFourth
	PatternRuneFifth
	PatternRuneSixth
	PatternRuneSeventh
	PatternRuneEighth
	PatternRuneNinth
	PatternRuneTenth
)
