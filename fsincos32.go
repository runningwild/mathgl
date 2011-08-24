package mathgl

import "math"

func Fcos32(x float32) float32 {
	return float32(math.Cos(float64(x)))
}
func Fsin32(x float32) float32 {
	return float32(math.Sin(float64(x)))
}
