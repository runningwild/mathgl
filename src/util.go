package mathgl

const (
	PI         float32 = 3.141592
	PIover180  float32 = 0.017453
	PIunder180 float32 = 57.295779
	epsilon    float32 = 1.0 / 64.0
)

func FSqr32(s float32) float32 {
	return s * s
}
func FDegToRad32(degrees float32) float32 {
	return degrees * PIover180
}

func FRadToDeg32(radians float32) float32 {
	return radians * PIunder180
}

func FMin32(lhs float32, rhs float32) float32 {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func FMax32(lhs float32, rhs float32) float32 {
	if lhs > rhs {
		return lhs
	}
	return rhs
}

func FAlmostEqual32(lhs float32, rhs float32) bool {
	return (lhs+epsilon > rhs && lhs-epsilon < rhs)
}

