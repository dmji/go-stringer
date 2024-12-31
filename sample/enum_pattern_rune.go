package sample

//go:generate go run ./../. -type=PatternRune -trimprefix=PatternRune -linecomment=true -nametransform=snake_case_lower -outputtransform=snake_case_lower -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=snake_case_lower -extraconstsvaluesuffix=Presentation

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
