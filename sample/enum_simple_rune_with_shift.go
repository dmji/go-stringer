package sample

//go:generate go run ./../. -type=SimpleRuneWithShift -trimprefix=SimpleRuneWithShift -linecomment=true -nametransform=snake_case_lower -outputtransform=snake_case_lower -extraconstsnamesuffix=_i18n_ID -extraconstsvaluetransform=snake_case_lower -extraconstsvaluesuffix=Presentation

type SimpleRuneWithShift uint8

const (
	SimpleRuneWithShiftFirst SimpleRuneWithShift = iota + 25
	SimpleRuneWithShiftSecond
	SimpleRuneWithShiftThird
	SimpleRuneWithShiftFourth
	SimpleRuneWithShiftFifth
	SimpleRuneWithShiftSixth
	SimpleRuneWithShiftSeventh
	SimpleRuneWithShiftEighth
	SimpleRuneWithShiftNinth
	SimpleRuneWithShiftTenth
	SimpleRuneWithShiftEleventh
	SimpleRuneWithShiftTwelfth
	SimpleRuneWithShiftThirteenth
	SimpleRuneWithShiftFourteenth
	SimpleRuneWithShiftFifteenth
	SimpleRuneWithShiftSixteenth
	SimpleRuneWithShiftSeventeenth
	SimpleRuneWithShiftEighteenth
	SimpleRuneWithShiftNineteenth
	SimpleRuneWithShiftTwentieth
	SimpleRuneWithShiftTwentyFirst
	SimpleRuneWithShiftTwentySecond
	SimpleRuneWithShiftTwentyThird
	SimpleRuneWithShiftTwentyFourth
	SimpleRuneWithShiftTwentyFifth
	SimpleRuneWithShiftTwentySixth
	SimpleRuneWithShiftTwentySeventh
	SimpleRuneWithShiftTwentyEighth
	SimpleRuneWithShiftTwentyNinth
	SimpleRuneWithShiftThirtieth
	SimpleRuneWithShiftThirtyFirst
	SimpleRuneWithShiftThirtySecond
	SimpleRuneWithShiftThirtyThird
	SimpleRuneWithShiftThirtyFourth
	SimpleRuneWithShiftThirtyFifth
	SimpleRuneWithShiftThirtySixth
	SimpleRuneWithShiftThirtySeventh
	SimpleRuneWithShiftThirtyEighth
	SimpleRuneWithShiftThirtyNinth
	SimpleRuneWithShiftFortieth
	SimpleRuneWithShiftFortyFirst
	SimpleRuneWithShiftFortySecond
	SimpleRuneWithShiftFortyThird
	SimpleRuneWithShiftFortyFourth
	SimpleRuneWithShiftFortyFifth
	SimpleRuneWithShiftFortySixth
	SimpleRuneWithShiftFortySeventh
	SimpleRuneWithShiftFortyEighth
	SimpleRuneWithShiftFortyNinth
	SimpleRuneWithShiftFiftieth
	SimpleRuneWithShiftFiftyFirst
	SimpleRuneWithShiftFiftySecond
	SimpleRuneWithShiftFiftyThird
	SimpleRuneWithShiftFiftyFourth
	SimpleRuneWithShiftFiftyFifth
	SimpleRuneWithShiftFiftySixth
	SimpleRuneWithShiftFiftySeventh
	SimpleRuneWithShiftFiftyEighth
	SimpleRuneWithShiftFiftyNinth
	SimpleRuneWithShiftSixtieth
	SimpleRuneWithShiftSixtyFirst
	SimpleRuneWithShiftSixtySecond
	SimpleRuneWithShiftSixtyThird
	SimpleRuneWithShiftSixtyFourth
	SimpleRuneWithShiftSixtyFifth
	SimpleRuneWithShiftSixtySixth
	SimpleRuneWithShiftSixtySeventh
	SimpleRuneWithShiftSixtyEighth
	SimpleRuneWithShiftSixtyNinth
	SimpleRuneWithShiftSeventieth
	SimpleRuneWithShiftSeventyFirst
	SimpleRuneWithShiftSeventySecond
	SimpleRuneWithShiftSeventyThird
	SimpleRuneWithShiftSeventyFourth
	SimpleRuneWithShiftSeventyFifth
	SimpleRuneWithShiftSeventySixth
	SimpleRuneWithShiftSeventySeventh
	SimpleRuneWithShiftSeventyEighth
	SimpleRuneWithShiftSeventyNinth
	SimpleRuneWithShiftEightieth
	SimpleRuneWithShiftEightyFirst
	SimpleRuneWithShiftEightySecond
	SimpleRuneWithShiftEightyThird
	SimpleRuneWithShiftEightyFourth
	SimpleRuneWithShiftEightyFifth
	SimpleRuneWithShiftEightySixth
	SimpleRuneWithShiftEightySeventh
	SimpleRuneWithShiftEightyEighth
	SimpleRuneWithShiftEightyNinth
	SimpleRuneWithShiftNinetieth
	SimpleRuneWithShiftNinetyFirst
	SimpleRuneWithShiftNinetySecond
	SimpleRuneWithShiftNinetyThird
	SimpleRuneWithShiftNinetyFourth
	SimpleRuneWithShiftNinetyFifth
	SimpleRuneWithShiftNinetySixth
	SimpleRuneWithShiftNinetySeventh
	SimpleRuneWithShiftNinetyEighth
	SimpleRuneWithShiftNinetyNinth
	SimpleRuneWithShiftOneHundredth
	SimpleRuneWithShiftOneHundredFirst
	SimpleRuneWithShiftOneHundredSecond
	SimpleRuneWithShiftOneHundredThird
	SimpleRuneWithShiftOneHundredFourth
	SimpleRuneWithShiftOneHundredFifth
	SimpleRuneWithShiftOneHundredSixth
	SimpleRuneWithShiftOneHundredSeventh
	SimpleRuneWithShiftOneHundredEighth
	SimpleRuneWithShiftOneHundredNinth
	SimpleRuneWithShiftOneHundredTenth
	SimpleRuneWithShiftOneHundredEleventh
	SimpleRuneWithShiftOneHundredTwelfth
	SimpleRuneWithShiftOneHundredThirteenth
	SimpleRuneWithShiftOneHundredFourteenth
	SimpleRuneWithShiftOneHundredFifteenth
	SimpleRuneWithShiftOneHundredSixteenth
	SimpleRuneWithShiftOneHundredSeventeenth
	SimpleRuneWithShiftOneHundredEighteenth
	SimpleRuneWithShiftOneHundredNineteenth
	SimpleRuneWithShiftOneHundredTwentieth
	SimpleRuneWithShiftOneHundredTwentyFirst
	SimpleRuneWithShiftOneHundredTwentySecond
	SimpleRuneWithShiftOneHundredTwentyThird
	SimpleRuneWithShiftOneHundredTwentyFourth
	SimpleRuneWithShiftOneHundredTwentyFifth
	SimpleRuneWithShiftOneHundredTwentySixth
	SimpleRuneWithShiftOneHundredTwentySeventh
	SimpleRuneWithShiftOneHundredTwentyEighth
	SimpleRuneWithShiftOneHundredTwentyNinth
	SimpleRuneWithShiftOneHundredThirtieth
	SimpleRuneWithShiftOneHundredThirtyFirst
	SimpleRuneWithShiftOneHundredThirtySecond
	SimpleRuneWithShiftOneHundredThirtyThird
	SimpleRuneWithShiftOneHundredThirtyFourth
	SimpleRuneWithShiftOneHundredThirtyFifth
	SimpleRuneWithShiftOneHundredThirtySixth
	SimpleRuneWithShiftOneHundredThirtySeventh
	SimpleRuneWithShiftOneHundredThirtyEighth
	SimpleRuneWithShiftOneHundredThirtyNinth
	SimpleRuneWithShiftOneHundredFortieth
	SimpleRuneWithShiftOneHundredFortyFirst
	SimpleRuneWithShiftOneHundredFortySecond
	SimpleRuneWithShiftOneHundredFortyThird
	SimpleRuneWithShiftOneHundredFortyFourth
	SimpleRuneWithShiftOneHundredFortyFifth
	SimpleRuneWithShiftOneHundredFortySixth
	SimpleRuneWithShiftOneHundredFortySeventh
	SimpleRuneWithShiftOneHundredFortyEighth
	SimpleRuneWithShiftOneHundredFortyNinth
	SimpleRuneWithShiftOneHundredFiftieth
	SimpleRuneWithShiftOneHundredFiftyFirst
	SimpleRuneWithShiftOneHundredFiftySecond
	SimpleRuneWithShiftOneHundredFiftyThird
	SimpleRuneWithShiftOneHundredFiftyFourth
	SimpleRuneWithShiftOneHundredFiftyFifth
	SimpleRuneWithShiftOneHundredFiftySixth
	SimpleRuneWithShiftOneHundredFiftySeventh
	SimpleRuneWithShiftOneHundredFiftyEighth
	SimpleRuneWithShiftOneHundredFiftyNinth
	SimpleRuneWithShiftOneHundredSixtieth
	SimpleRuneWithShiftOneHundredSixtyFirst
	SimpleRuneWithShiftOneHundredSixtySecond
	SimpleRuneWithShiftOneHundredSixtyThird
	SimpleRuneWithShiftOneHundredSixtyFourth
	SimpleRuneWithShiftOneHundredSixtyFifth
	SimpleRuneWithShiftOneHundredSixtySixth
	SimpleRuneWithShiftOneHundredSixtySeventh
	SimpleRuneWithShiftOneHundredSixtyEighth
	SimpleRuneWithShiftOneHundredSixtyNinth
	SimpleRuneWithShiftOneHundredSeventieth
	SimpleRuneWithShiftOneHundredSeventyFirst
	SimpleRuneWithShiftOneHundredSeventySecond
	SimpleRuneWithShiftOneHundredSeventyThird
	SimpleRuneWithShiftOneHundredSeventyFourth
	SimpleRuneWithShiftOneHundredSeventyFifth
	SimpleRuneWithShiftOneHundredSeventySixth
	SimpleRuneWithShiftOneHundredSeventySeventh
	SimpleRuneWithShiftOneHundredSeventyEighth
	SimpleRuneWithShiftOneHundredSeventyNinth
	SimpleRuneWithShiftOneHundredEightieth
	SimpleRuneWithShiftOneHundredEightyFirst
	SimpleRuneWithShiftOneHundredEightySecond
	SimpleRuneWithShiftOneHundredEightyThird
	SimpleRuneWithShiftOneHundredEightyFourth
	SimpleRuneWithShiftOneHundredEightyFifth
	SimpleRuneWithShiftOneHundredEightySixth
	SimpleRuneWithShiftOneHundredEightySeventh
	SimpleRuneWithShiftOneHundredEightyEighth
	SimpleRuneWithShiftOneHundredEightyNinth
	SimpleRuneWithShiftOneHundredNinetieth
	SimpleRuneWithShiftOneHundredNinetyFirst
	SimpleRuneWithShiftOneHundredNinetySecond
	SimpleRuneWithShiftOneHundredNinetyThird
	SimpleRuneWithShiftOneHundredNinetyFourth
	SimpleRuneWithShiftOneHundredNinetyFifth
	SimpleRuneWithShiftOneHundredNinetySixth
	SimpleRuneWithShiftOneHundredNinetySeventh
	SimpleRuneWithShiftOneHundredNinetyEighth
	SimpleRuneWithShiftOneHundredNinetyNinth
	SimpleRuneWithShiftTwoHundredth
)
