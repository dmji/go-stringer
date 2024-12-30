package sample

//go:generate go run ./../. -type=PatternMapRune -linecomment=true -trimprefix=PatternMapRune -nameConvertToCase=snake_case

type PatternMapRune uint16

const (
	PatternMapRuneFirst PatternMapRune = iota * 2
	PatternMapRuneSecond
	PatternMapRuneThird
	PatternMapRuneFourth
	PatternMapRuneFifth
	PatternMapRuneSixth
	PatternMapRuneSeventh
	PatternMapRuneEighth
	PatternMapRuneNinth
	PatternMapRuneTenth
	PatternMapRuneEleventh
	PatternMapRuneTwelfth
	PatternMapRuneThirteenth
	PatternMapRuneFourteenth
	PatternMapRuneFifteenth
	PatternMapRuneSixteenth
	PatternMapRuneSeventeenth
	PatternMapRuneEighteenth
	PatternMapRuneNineteenth
	PatternMapRuneTwentieth
	PatternMapRuneTwentyFirst
	PatternMapRuneTwentySecond
	PatternMapRuneTwentyThird
	PatternMapRuneTwentyFourth
	PatternMapRuneTwentyFifth
	PatternMapRuneTwentySixth
	PatternMapRuneTwentySeventh
	PatternMapRuneTwentyEighth
	PatternMapRuneTwentyNinth
	PatternMapRuneThirtieth
	PatternMapRuneThirtyFirst
	PatternMapRuneThirtySecond
	PatternMapRuneThirtyThird
	PatternMapRuneThirtyFourth
	PatternMapRuneThirtyFifth
	PatternMapRuneThirtySixth
	PatternMapRuneThirtySeventh
	PatternMapRuneThirtyEighth
	PatternMapRuneThirtyNinth
	PatternMapRuneFortieth
	PatternMapRuneFortyFirst
	PatternMapRuneFortySecond
	PatternMapRuneFortyThird
	PatternMapRuneFortyFourth
	PatternMapRuneFortyFifth
	PatternMapRuneFortySixth
	PatternMapRuneFortySeventh
	PatternMapRuneFortyEighth
	PatternMapRuneFortyNinth
	PatternMapRuneFiftieth
	PatternMapRuneFiftyFirst
	PatternMapRuneFiftySecond
	PatternMapRuneFiftyThird
	PatternMapRuneFiftyFourth
	PatternMapRuneFiftyFifth
	PatternMapRuneFiftySixth
	PatternMapRuneFiftySeventh
	PatternMapRuneFiftyEighth
	PatternMapRuneFiftyNinth
	PatternMapRuneSixtieth
	PatternMapRuneSixtyFirst
	PatternMapRuneSixtySecond
	PatternMapRuneSixtyThird
	PatternMapRuneSixtyFourth
	PatternMapRuneSixtyFifth
	PatternMapRuneSixtySixth
	PatternMapRuneSixtySeventh
	PatternMapRuneSixtyEighth
	PatternMapRuneSixtyNinth
	PatternMapRuneSeventieth
	PatternMapRuneSeventyFirst
	PatternMapRuneSeventySecond
	PatternMapRuneSeventyThird
	PatternMapRuneSeventyFourth
	PatternMapRuneSeventyFifth
	PatternMapRuneSeventySixth
	PatternMapRuneSeventySeventh
	PatternMapRuneSeventyEighth
	PatternMapRuneSeventyNinth
	PatternMapRuneEightieth
	PatternMapRuneEightyFirst
	PatternMapRuneEightySecond
	PatternMapRuneEightyThird
	PatternMapRuneEightyFourth
	PatternMapRuneEightyFifth
	PatternMapRuneEightySixth
	PatternMapRuneEightySeventh
	PatternMapRuneEightyEighth
	PatternMapRuneEightyNinth
	PatternMapRuneNinetieth
	PatternMapRuneNinetyFirst
	PatternMapRuneNinetySecond
	PatternMapRuneNinetyThird
	PatternMapRuneNinetyFourth
	PatternMapRuneNinetyFifth
	PatternMapRuneNinetySixth
	PatternMapRuneNinetySeventh
	PatternMapRuneNinetyEighth
	PatternMapRuneNinetyNinth
	PatternMapRuneOneHundredth
	PatternMapRuneOneHundredFirst
	PatternMapRuneOneHundredSecond
	PatternMapRuneOneHundredThird
	PatternMapRuneOneHundredFourth
	PatternMapRuneOneHundredFifth
	PatternMapRuneOneHundredSixth
	PatternMapRuneOneHundredSeventh
	PatternMapRuneOneHundredEighth
	PatternMapRuneOneHundredNinth
	PatternMapRuneOneHundredTenth
	PatternMapRuneOneHundredEleventh
	PatternMapRuneOneHundredTwelfth
	PatternMapRuneOneHundredThirteenth
	PatternMapRuneOneHundredFourteenth
	PatternMapRuneOneHundredFifteenth
	PatternMapRuneOneHundredSixteenth
	PatternMapRuneOneHundredSeventeenth
	PatternMapRuneOneHundredEighteenth
	PatternMapRuneOneHundredNineteenth
	PatternMapRuneOneHundredTwentieth
	PatternMapRuneOneHundredTwentyFirst
	PatternMapRuneOneHundredTwentySecond
	PatternMapRuneOneHundredTwentyThird
	PatternMapRuneOneHundredTwentyFourth
	PatternMapRuneOneHundredTwentyFifth
	PatternMapRuneOneHundredTwentySixth
	PatternMapRuneOneHundredTwentySeventh
	PatternMapRuneOneHundredTwentyEighth
	PatternMapRuneOneHundredTwentyNinth
	PatternMapRuneOneHundredThirtieth
	PatternMapRuneOneHundredThirtyFirst
	PatternMapRuneOneHundredThirtySecond
	PatternMapRuneOneHundredThirtyThird
	PatternMapRuneOneHundredThirtyFourth
	PatternMapRuneOneHundredThirtyFifth
	PatternMapRuneOneHundredThirtySixth
	PatternMapRuneOneHundredThirtySeventh
	PatternMapRuneOneHundredThirtyEighth
	PatternMapRuneOneHundredThirtyNinth
	PatternMapRuneOneHundredFortieth
	PatternMapRuneOneHundredFortyFirst
	PatternMapRuneOneHundredFortySecond
	PatternMapRuneOneHundredFortyThird
	PatternMapRuneOneHundredFortyFourth
	PatternMapRuneOneHundredFortyFifth
	PatternMapRuneOneHundredFortySixth
	PatternMapRuneOneHundredFortySeventh
	PatternMapRuneOneHundredFortyEighth
	PatternMapRuneOneHundredFortyNinth
	PatternMapRuneOneHundredFiftieth
	PatternMapRuneOneHundredFiftyFirst
	PatternMapRuneOneHundredFiftySecond
	PatternMapRuneOneHundredFiftyThird
	PatternMapRuneOneHundredFiftyFourth
	PatternMapRuneOneHundredFiftyFifth
	PatternMapRuneOneHundredFiftySixth
	PatternMapRuneOneHundredFiftySeventh
	PatternMapRuneOneHundredFiftyEighth
	PatternMapRuneOneHundredFiftyNinth
	PatternMapRuneOneHundredSixtieth
	PatternMapRuneOneHundredSixtyFirst
	PatternMapRuneOneHundredSixtySecond
	PatternMapRuneOneHundredSixtyThird
	PatternMapRuneOneHundredSixtyFourth
	PatternMapRuneOneHundredSixtyFifth
	PatternMapRuneOneHundredSixtySixth
	PatternMapRuneOneHundredSixtySeventh
	PatternMapRuneOneHundredSixtyEighth
	PatternMapRuneOneHundredSixtyNinth
	PatternMapRuneOneHundredSeventieth
	PatternMapRuneOneHundredSeventyFirst
	PatternMapRuneOneHundredSeventySecond
	PatternMapRuneOneHundredSeventyThird
	PatternMapRuneOneHundredSeventyFourth
	PatternMapRuneOneHundredSeventyFifth
	PatternMapRuneOneHundredSeventySixth
	PatternMapRuneOneHundredSeventySeventh
	PatternMapRuneOneHundredSeventyEighth
	PatternMapRuneOneHundredSeventyNinth
	PatternMapRuneOneHundredEightieth
	PatternMapRuneOneHundredEightyFirst
	PatternMapRuneOneHundredEightySecond
	PatternMapRuneOneHundredEightyThird
	PatternMapRuneOneHundredEightyFourth
	PatternMapRuneOneHundredEightyFifth
	PatternMapRuneOneHundredEightySixth
	PatternMapRuneOneHundredEightySeventh
	PatternMapRuneOneHundredEightyEighth
	PatternMapRuneOneHundredEightyNinth
	PatternMapRuneOneHundredNinetieth
	PatternMapRuneOneHundredNinetyFirst
	PatternMapRuneOneHundredNinetySecond
	PatternMapRuneOneHundredNinetyThird
	PatternMapRuneOneHundredNinetyFourth
	PatternMapRuneOneHundredNinetyFifth
	PatternMapRuneOneHundredNinetySixth
	PatternMapRuneOneHundredNinetySeventh
	PatternMapRuneOneHundredNinetyEighth
	PatternMapRuneOneHundredNinetyNinth
	PatternMapRuneTwoHundredth
)