package simpleround

import "math"

// Round rounds number to nearest whole integer
func Round(f float64) float64 {
	return math.Floor(f + .5)
}

// PrecisionRound rounds to given number of decimal places
func PrecisionRound(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return Round(f*shift) / shift
}

// RoundDollars wraps PrecisionRound to round to 2 decimals
func RoundDollars(f float64) float64 {
	return PrecisionRound(f, 2)
}
