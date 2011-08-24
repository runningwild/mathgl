// func Fsqrt32(x float32) float32
TEXT ·Fsqrt32(SB),7,$0
	FMOVF x+0(FP), F0
	FSQRT
	FMOVFP F0, r+4(FP)
	RET
