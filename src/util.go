/**
 * mathgl: util.go
 * User: Nils Hasenbanck
 * Date: 21.08.11
 * Time: 13:35
 */
package mathgl

import "unsafe"

const (
	PI         float32 = 3.141592
	PIover180  float32 = 0.017453
	PIunder180 float32 = 57.295779
	epsilon    float32 = 1.0 / 64.0
)

func Fsqr32(s float32) float32 {
	return s * s
}
func Fdeg2rad32(degrees float32) float32 {
	return degrees * PIover180
}

func Frad2deg32(radians float32) float32 {
	return radians * PIunder180
}

func Fmin32(lhs float32, rhs float32) float32 {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func Fmax32(lhs float32, rhs float32) float32 {
	if lhs > rhs {
		return lhs
	}
	return rhs
}

func Falmostequal32(lhs float32, rhs float32) bool {
	return (lhs+epsilon > rhs && lhs-epsilon < rhs)
}

// TODO: ASM me with SSE
// Using the fast inverse root with better magic number
func Fsqrt32(x float32) float32 {
    const t float32 = 1.5
    var x2 float32 = x * 0.5
    var y  float32 = x
    i := *(*uint32)(unsafe.Pointer(&x))
    i = 0x5f375a86 - ( i >> 1 )
    y = *(*float32)(unsafe.Pointer(&i))
    y = y * ( t - ( x2 * y * y ) )
    return x * y
}
