package mathgl

import "fmt"

// 2 dimensional vector.
type Vec2 struct {
	X, Y float32
}

// Fills the vector with the given float32
func (v *Vec2) Fill(x, y float32) {
	v.X = x
	v.Y = y
}

// Returns the length  as float32
func (v *Vec2) Length() float32 {
	return Fsqrt32(Fsqr32(v.X) + Fsqr32(v.Y))
}

// Returns the length as square as float32
func (v *Vec2) LengthSq() float32 {
	return Fsqr32(v.X) + Fsqr32(v.Y)
}

// Normalize the vector
func (v *Vec2) Normalize() {
	var l float32 = 1.0 / v.Length()
	v.X *= l
	v.Y *= l
}

func (v *Vec2) Cross() {
	v.X, v.Y = -v.Y, v.X
}

// Adds the given Vec2 with the vector
func (v *Vec2) Add(x *Vec2) {
	v.X += x.X
	v.Y += x.Y
}

// Returns the cosine of the angle between the vectors as float32
func (v *Vec2) Dot(x *Vec2) float32 {
	return v.X*x.X + v.Y*x.Y
}

// Subtracts the given Vec2 from the vector
func (v *Vec2) Subtract(x *Vec2) {
	v.X -= x.X
	v.Y -= x.Y
}

// Transforms the Vec2 by a given Mat3
func (v *Vec2) Transform(m *Mat3) {
	var t Vec2
	t.X = v.X
	t.Y = v.Y

	v.X = t.X*m[0] + t.Y*m[3] + m[6]
	v.Y = t.X*m[1] + t.Y*m[4] + m[7]
}

// Scales the vector with the given float32.
func (v *Vec2) Scale(s float32) {
	v.X *= s
	v.Y *= s
}

// Assigns the given Vec2 to the Vec2
func (v *Vec2) Assign(x *Vec2) {
	if v == x {
		return
	}

	v.X = x.X
	v.Y = x.Y
}

// Returns true if the vectors are approximately equal in value
func (v *Vec2) AreEqual(x *Vec2) bool {
	return ((v.X < x.X+epsilon && v.X > x.X-epsilon) &&
		(v.Y < x.Y+epsilon && v.Y > x.Y-epsilon))
}

// Sets all the elements of Vec2 to zero.
func (v *Vec2) Zero() {
	v.X = 0.0
	v.Y = 0.0
}

func (v *Vec2) String() string {
	return fmt.Sprintf("Vec2(%f, %f)", v.X, v.Y)
}
