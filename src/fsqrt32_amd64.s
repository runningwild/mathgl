// func FSqrt32(x float32) float32
TEXT ·FSqrt32(SB),7,$0
	SQRTSS x+0(FP), X0      // 0(FP) is the first argument, x is just a name
	MOVSS X0, r+8(FP)       // 8(FP) is the return argument, r is just a name
	RET
