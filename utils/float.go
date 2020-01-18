package utils

import "math"

func FloatRoundDown(float float64, scale int) float64 {
	scalePow := math.Pow10(scale)
	return math.Floor(float*scalePow) / scalePow
}
func FloatRoundUp(float float64, scale int) float64 {
	scalePow := math.Pow10(scale)
	return math.Ceil(float*scalePow) / scalePow
}
func FloatRoundNearest(float float64, scale int) float64 {
	scalePow := math.Pow10(scale)
	return math.Round(float*scalePow) / scalePow
}
