package sample

//go:generate go run ./.. -type=TwoTypeFirst,TwoTypeSecond -trimprefix=@me -nametransform=snake_case_lower -fromstringgenfn -marshaljson -outputtransform=snake_case_lower

type TwoTypeFirst uint8

const (
	TwoTypeFirstFirst TwoTypeFirst = iota * 2
	TwoTypeFirstSecond
	TwoTypeFirstThird
	TwoTypeFirstFourth
	TwoTypeFirstFifth
	TwoTypeFirstSixth
	TwoTypeFirstSeventh
	TwoTypeFirstEighth
	TwoTypeFirstNinth
	TwoTypeFirstTenth
)

type TwoTypeSecond uint8

const (
	TwoTypeSecondFirst TwoTypeSecond = iota
	TwoTypeSecondSecond
	TwoTypeSecondThird
	TwoTypeSecondFourth
	TwoTypeSecondFifth
	TwoTypeSecondSixth
	TwoTypeSecondSeventh
	TwoTypeSecondEighth
	TwoTypeSecondNinth
	TwoTypeSecondTenth
)
