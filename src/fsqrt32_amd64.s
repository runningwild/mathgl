// func Fsqrt32(x float32) float32
TEXT ·Fsqrt32(SB),7,$0
	SQRTSS x+0(FP), X0
	MOVSD X0, r+8(FP)
	RET
