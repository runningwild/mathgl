package mathgl

import "unsafe"

// Using the fast inverse root with better magic number
func Fsqrt32Go(x float32) float32 {
    const t float32 = 1.5
    var x2 float32 = x * 0.5
    var y  float32 = x
    i := *(*uint32)(unsafe.Pointer(&x))
    i = 0x5f375a86 - ( i >> 1 )
    y = *(*float32)(unsafe.Pointer(&i))
    y = y * ( t - ( x2 * y * y ) )
    return x * y
}

func Fsqrt32GoC(f float32, r *float32) {
	*r = Fsqrt32Go(f)
}
