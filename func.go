// MathGL is a simple 3D math library written in Go which should help writing OpenGL code.
package mathgl

import (
	"unsafe"
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

// Returns the absolute value from a float
func Fabs32(f float32) float32 {
	// Anybody knows a speed hack without type-conversion?
	i := *(*int32)(unsafe.Pointer(&f)) & 0x7fffffff
	return *(*float32)(unsafe.Pointer(&i))
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

// The following SIN/COS functions come from an forum thread from the user Nick:
// http://www.devmaster.net/forums/showthread.php?t=5784 provided 'as is'.
// This approximation is incredible fast (2x faster then the 64bit version
// with two casts)

// Returns the sin of a given float32 radiant
func Fsin32(x float32) float32 {
	x = x * (1.0 / PI)

	z := (x + 25165824.0)
	x = x - (z - 25165824.0)

	y := x - x*Fabs32(x)

	const Q float32 = 3.1
	const P float32 = 3.6

	return y * (Q + P*Fabs32(y))
}

// Returns the cos of a given float32 radiant
func Fcos32(x float32) float32 {
	const PIhalf float32 = PI / 2
	x = x + PIhalf // shift for cos
	x = x * (1.0 / PI)

	z := (x + 25165824.0)
	x = x - (z - 25165824.0)

	y := x - x*Fabs32(x)

	const Q float32 = 3.1
	const P float32 = 3.6

	return y * (Q + P*Fabs32(y))
}
