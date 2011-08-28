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

func TestFsincos32(t *testing.T) {
	cos := Fcos32(Fdeg2rad32(450))
	if !FalmostEqual32(cos, 0) {
		message := fmt.Sprintf("cos(450.0) is not 0, it is %e)\n", cos)
		t.Errorf(message)
	}
	cos = Fcos32(Fdeg2rad32(180))
	if !FalmostEqual32(cos, -1) {
		message := fmt.Sprintf("cos(180.0) is not -1, it is %e)\n", cos)
		t.Errorf(message)
	}
	cos = Fcos32(Fdeg2rad32(45))
	if !FalmostEqual32(cos, 0.7071) {
		message := fmt.Sprintf("cos(180.0) is not -1, it is %e)\n", cos)
		t.Errorf(message)
	}
	sin := Fsin32(Fdeg2rad32(540))
	if !FalmostEqual32(sin, 0) {
		message := fmt.Sprintf("sin(540.0) is not 0, it is %e)\n", sin)
		t.Errorf(message)
	}
	sin = Fsin32(Fdeg2rad32(45))
	if !FalmostEqual32(sin, 0.7071) {
		message := fmt.Sprintf("sin(540.0) is not 0, it is %e)\n", sin)
		t.Errorf(message)
	}
}

func TestVec2(t *testing.T) {
	v2 := new(Vec2)
	if v2.X != 0 || v2.Y != 0 {
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
		message := fmt.Sprintf("Length of Vec2(%f,%f) should be 1.0 but is %f\n", v2.X, v2.Y, length)
		t.Errorf(message)
	}
	v2.Add(v2)
	if v2.X != 1.6 || v2.Y != 1.2 {
		message := fmt.Sprintf("Vec2(%f,%f) should be Vec2(1.6,1.2) after add function\n", v2.X, v2.Y)
		t.Errorf(message)
	}
	dot := v2.Dot(v2)
	if dot != 4.0 {
		message := fmt.Sprintf("Dot product of Vec2(%f,%f) should be 4.0 but is %f \n", v2.X, v2.Y, dot)
		t.Errorf(message)
	}
	sub := new(Vec2)
	sub.Fill(0.6, 0.2)
	v2.Subtract(sub)
	if v2.X != 1.0 || v2.Y != 1.0 {
		message := fmt.Sprintf("Vector should be Vec(1.0,1.0) but is Vec2(%f,%f) after subtraction\n", v2.X, v2.Y)
		t.Errorf(message)
	}
	var identity Mat3
	identity.Identity()
	v2.Transform(&identity)
	if v2.X != 1.0 || v2.Y != 1.0 {
		message := fmt.Sprintf("Vector should be Vec(1.0,1.0) but is Vec2(%f,%f) after transformation with identity matrix\n", v2.X, v2.Y)
		t.Errorf(message)
	}
	v2.Scale(5.0)
	if v2.X != 5.0 || v2.Y != 5.0 {
		message := fmt.Sprintf("Vector should be Vec(5.0,5.0) but is Vec2(%f,%f) after scale with scalar 5.0\n", v2.X, v2.Y)
		t.Errorf(message)
	}
	if !v2.AreEqual(v2) {
		message := fmt.Sprintf("Vector is not equal with himelf. We screwed up badly!", v2.X, v2.Y)
		t.Errorf(message)
	}

}


func TestMat3(t *testing.T) {
	m := Mat3{5.0, 8.0, 1.0, 2.0, 9.0, 3.0, 4.0, 7.0, 4.0}
	n := m
	det := m.Determinant()
	if det != 85 {
		t.Errorf("Determinant is not 85! It is %f", det)
	}
	if !m.Inverse() {
		t.Errorf("Determinant was 0! Can't calculate inverse!")
	}
	m.Multiply(&n)
	if !m.IsIdentity() {
		t.Errorf("The Mat3 matrix is not a identity matrix after multiplying itself with its inverse.")
	}
}


func TestMat4(t *testing.T) {
	var m Mat4
	m.Identity()
	det := m.Determinant()
	if det != 1 {
		t.Errorf("Determinant of identity matrix is not 1! It is %f", det)
	}
	m = Mat4{1.0, 6.0, 2.0, 2.0, 8.0, 4.0, 2.0, 9.0, 7.0, 2.0, 4.0, 1.0, 10.0, 9.0, 5.0, 5.0}
	n := m
	if !m.Inverse() {
		t.Errorf("Determinant was 0! Can't calculate inverse!")
	}
	m.Multiply(&n)
	if !m.IsIdentity() {
		t.Errorf("The Mat4 matrix is not a identity matrix after multiplying itself with its inverse.")
	}
}
