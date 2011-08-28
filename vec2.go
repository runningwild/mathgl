package mathgl

import "fmt"

// 2 dimensional vector.
type Vec2 struct {
	x, y float32
}

// Fills the vector with the given float32
func (v *Vec2) Fill(x, y float32) {
	v.x = x
	v.y = y
}

// Returns the length  as float32
func (v *Vec2) Length() float32 {
	return Fsqrt32(Fsqr32(v.x) + Fsqr32(v.y))
}

// Returns the length as square as float32
func (v *Vec2) LengthSq() float32 {
	return Fsqr32(v.x) + Fsqr32(v.y)
}

// Normalize the vector
func (v *Vec2) Normalize() {
	var l float32 = 1.0 / v.Length()
	v.x *= l
	v.y *= l
}

// Adds the given Vec2 with the vector
func (v *Vec2) Add(x *Vec2) {
	v.x += x.x
	v.y += x.y
}

// Returns the cosine of the angle between the vectors as float32
func (v *Vec2) Dot(x *Vec2) float32 {
	return v.x*x.x + v.y*x.y
}

// Subtracts the given Vec2 from the vector
func (v *Vec2) Subtract(x *Vec2) {
	v.x -= x.x
	v.y -= x.y
}

// Transforms the Vec2 by a given Mat3
func (v *Vec2) Transform(m *Mat3) {
	var t Vec2
	t.x = v.x
	t.y = v.y

	v.x = t.x*m[0] + t.y*m[3] + m[6]
	v.y = t.x*m[1] + t.y*m[4] + m[7]
}

// Scales the vector with the given float32.
func (v *Vec2) Scale(s float32) {
	v.x *= s
	v.y *= s
}

// Assigns the given Vec2 to the Vec2
func (v *Vec2) Assign(x *Vec2) {
	if v == x {
		return
	}

	v.x = x.x
	v.y = x.y
}

// Returns true if the vectors are approximately equal in value
func (v *Vec2) AreEqual(x *Vec2) bool {
	return ((v.x < x.x+epsilon && v.x > x.x-epsilon) &&
		(v.y < x.y+epsilon && v.y > x.y-epsilon))
}

// Sets all the elements of Vec2 to zero.
func (v *Vec2) Zero() {
	v.x = 0.0
	v.y = 0.0
}

func (v *Vec2) String() string {
	return fmt.Sprintf("Vec2(%f, %f)", v.x, v.y)
}
