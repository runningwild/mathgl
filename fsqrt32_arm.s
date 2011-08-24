// func Sqrt(x float32) float32
TEXT ·Fsqrt32(SB),7,$0
	MOVF   x+0(FP),F0
	SQRTF  F0,F0
	MOVF  F0,r+4(FP)
	RET
