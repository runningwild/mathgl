package mathgl

import (
    "fmt"
    "testing"
)

type squareTest struct {
	in, out float32
}

var squareTests = []squareTest{
	squareTest{1.5, 1.224744871},
	squareTest{5.5, 2.34520788},
	squareTest{10.25, 3.201562119},
}

// We test if the sqrt function has more then 4% error
func TestSqrtFloat32(t *testing.T) {
	var error, result float32
	for _, st := range squareTests {
		result = Fsqrt32(st.in)
		error = ( Fmax32(st.out,result) / Fmin32(st.out,result) ) - 1
		if error > 0.04 {
		message := fmt.Sprintf("the error is too big: srqrt(%f) with error %f\n", st.in, error)
		t.Errorf(message)
		}
	}
}
