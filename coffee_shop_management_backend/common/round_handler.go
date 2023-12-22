package common

import "math"

const maxDecimalNumber = 3

func roundToDecimal(num float32, decimalPlaces int) float32 {
	shift := math.Pow10(decimalPlaces)
	return float32(math.Round(float64(num)*shift)) / float32(shift)
}

func RoundToInt(num float32) int {
	return int(roundToDecimal(num, 0))
}

func CustomRound(num *float32) {
	roundedNum := roundToDecimal(*num, maxDecimalNumber)
	*num = roundedNum
}
