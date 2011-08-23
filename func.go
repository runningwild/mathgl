// MathGL is a simple 3D math library written in Go which should help writing OpenGL code.
package mathgl

// Some constants which are often in 3D math.
const (
	PI         float32 = 3.141592
	PIover180  float32 = 0.017453
	PIunder180 float32 = 57.295779
	epsilon    float32 = 1.0 / 64.0
)

// The Square of a given float32.
func Fsqr32(s float32) float32 {
	return s * s
}

// Returns the radius value from a given degree value given in float32.
func Fdeg2rad32(degrees float32) float32 {
	return degrees * PIover180
}

// Returns the degrees value from a given radius value given in float32.
func Frad2deg32(radians float32) float32 {
	return radians * PIunder180
}

// Returns the smallest value of two given float32 values.
func Fmin32(lhs float32, rhs float32) float32 {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

// Returns biggest value of two given float32 values.
func Fmax32(lhs float32, rhs float32) float32 {
	if lhs > rhs {
		return lhs
	}
	return rhs
}

// Returns true if two float32 are almost the same (the threshold is epsilon = 1/64).
func FalmostEqual32(lhs float32, rhs float32) bool {
	return (lhs+epsilon > rhs && lhs-epsilon < rhs)
}

