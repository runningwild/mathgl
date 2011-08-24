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

// We test if the sqrt function has more then 1% error
func TestFsqrt32(t *testing.T) {
	var error, result float32
	for _, st := range squareTests {
		result = Fsqrt32(st.in)
		error = (Fmax32(st.out, result) / Fmin32(st.out, result)) - 1
		if error > 0.01 {
			message := fmt.Sprintf("The error is too big: srqrt(%f) with error %f\n", st.in, error)
			t.Errorf(message)
		}
	}
}

func TestFsin32(t *testing.T) {
	if Fcos32(Fdeg2rad32(360)) != 1.0 {
		message := fmt.Sprintf("cos(360.0) is not 1, it is %f)\n", Fcos32(Fdeg2rad32(360)))
		t.Errorf(message)
	}
	if Fsin32(Fdeg2rad32(90)) != 1.0 {
		message := fmt.Sprintf("sin(90.0) is not 1, it is %f)\n", Fsin32(Fdeg2rad32(90)))
		t.Errorf(message)
	}
}

func TestVec2(t *testing.T) {
	v2 := new(Vec2)
	if v2.x != 0 || v2.y != 0 {
		t.Errorf("Initialized Vec2 is not zero\n")
	}
	v2.Fill(4.0, 3.0)
	length := v2.Length()
	if length != 5.0 {
		message := fmt.Sprintf("Length of Vec2(4.0,3.0) should be 5.0 but is %f\n", length)
		t.Errorf(message)
	}
	v2.Normalize()
	length = v2.Length()
	if length != 1.0 {
		message := fmt.Sprintf("Length of Vec2(%f,%f) should be 1.0 but is %f\n", v2.x, v2.y, length)
		t.Errorf(message)
	}
	v2.Add(v2)
	if v2.x != 1.6 || v2.y != 1.2 {
		message := fmt.Sprintf("Vec2(%f,%f) should be Vec2(1.6,1.2) after add function\n", v2.x, v2.y)
		t.Errorf(message)
	}
	dot := v2.Dot(v2)
	if dot != 4.0 {
		message := fmt.Sprintf("Dot product of Vec2(%f,%f) should be 4.0 but is %f \n", v2.x, v2.y, dot)
		t.Errorf(message)
	}
	sub := new(Vec2)
	sub.Fill(0.6, 0.2)
	v2.Subtract(sub)
	if v2.x != 1.0 || v2.y != 1.0 {
		message := fmt.Sprintf("Vector should be Vec(1.0,1.0) but is Vec2(%f,%f) after subtraction\n", v2.x, v2.y)
		t.Errorf(message)
	}
	identity := MakeIdentityMat3()
	v2.Transform(identity)
	if v2.x != 1.0 || v2.y != 1.0 {
		message := fmt.Sprintf("Vector should be Vec(1.0,1.0) but is Vec2(%f,%f) after transformation with identity matrix\n", v2.x, v2.y)
		t.Errorf(message)
	}
	v2.Scale(5.0)
	if v2.x != 5.0 || v2.y != 5.0 {
		message := fmt.Sprintf("Vector should be Vec(5.0,5.0) but is Vec2(%f,%f) after scale with scalar 5.0\n", v2.x, v2.y)
		t.Errorf(message)
	}
	if !v2.AreEqual(v2) {
		message := fmt.Sprintf("Vector is not equal with himelf. We screwed up badly!", v2.x, v2.y)
		t.Errorf(message)
	}

}
