package mathgl

import "fmt"

// 4 dimensional vector.
type Vec4 struct {
	X, Y, Z, W float32
}

// Fills the vector with the given float32
func (v *Vec4) Fill(x, y, z, w float32) {
	v.X = x
	v.Y = y
	v.Z = z
	v.W = w
}

// Returns the length  as float32
func (v *Vec4) Length() float32 {
	return Fsqrt32(Fsqr32(v.X) + Fsqr32(v.Y) + Fsqr32(v.Z) + Fsqr32(v.W))
}

// Returns the length as square as float32
func (v *Vec4) LengthSq() float32 {
	return Fsqr32(v.X) + Fsqr32(v.Y) + Fsqr32(v.Z) + Fsqr32(v.W)
}

// Normalize the vector
func (v *Vec4) Normalize() {
	var l float32 = 1.0 / v.Length()
	v.X *= l
	v.Y *= l
	v.Z *= l
	v.W *= l
}

// Adds the given Vec4 with the vector
func (v *Vec4) Add(x *Vec4) {
	v.X += x.X
	v.Y += x.Y
	v.Z += x.Z
	v.W += x.W
}

// Returns the cosine of the angle between the vectors as float32
func (v *Vec4) Dot(x *Vec4) float32 {
	// Todo should this better be a Vec3 Dot product?
	return v.X*x.X + v.Y*x.Y + v.Z*x.Z + v.W*x.W
}

// Saves the Vec4 perpendicular to the given Vec4 (Attention: This is a Vec3 Cross Product in a homogeneous 4D environment!)
func (v *Vec4) Cross(x *Vec4) {
	var t Vec4
	t.Assign(v)

	v.X = (t.Y * x.Z) - (t.Z * x.Y)
	v.Y = (t.Z * x.X) - (t.X * x.Z)
	v.Z = (t.X * x.Y) - (t.Y * x.X)
}

// Subtracts the given Vec4 from the vector
func (v *Vec4) Subtract(x *Vec4) {
	v.X -= x.X
	v.Y -= x.Y
	v.Z -= x.Z
	v.W -= v.W
}

// Transforms the Vec4 by a given Mat4
func (v *Vec4) Transform(m *Mat4) {
	var t Vec4
	t.Assign(v)

	v.X = t.X*m[0] + t.Y*m[4] + t.Z*m[8] + t.W*m[12]
	v.Y = t.X*m[1] + t.Y*m[5] + t.Z*m[9] + t.W*m[13]
	v.Z = t.X*m[2] + t.Y*m[6] + t.Z*m[10] + t.W*m[14]
	v.W = t.X*m[3] + t.Y*m[7] + t.Z*m[11] + t.W*m[15]
}

/// Loops through an input slice transforming each Vec4 by the given Mat3
func (v *Vec3) TransformArray(x []Vec4, m *Mat4) {
	for _, item := range x {
		// TODO: We should test this
		item.Transform(m)
	}
}


// Scales a vector to the given length s in float32.
func (v *Vec4) Scale(s float32) {
	v.X *= s
	v.Y *= s
	v.Z *= s
	v.W *= s
}

// Returns true if the vectors are approximately equal in value
func (v *Vec4) AreEqual(x *Vec4) bool {
	return ((v.X < x.X+epsilon && v.X > x.X-epsilon) &&
		(v.Y < x.Y+epsilon && v.Y > x.Y-epsilon) &&
		(v.Z < x.Z+epsilon && v.Z > x.Z-epsilon) &&
		(v.W < x.W+epsilon && v.W > x.W-epsilon))
}

// Assigns the given Vec4 to the Vec4
func (v *Vec4) Assign(x *Vec4) {
	if v == x {
		return
	}

	v.X = x.X
	v.Y = x.Y
	v.Z = x.Z
	v.W = x.W
}

// Sets all the elements of Vec4 to zero
func (v *Vec4) Zero() {
	v.X = 0.0
	v.Y = 0.0
	v.Z = 0.0
	v.W = 0.0
}

func (v *Vec4) String() string {
	return fmt.Sprintf("Vec4(%f, %f, %f, %f)", v.X, v.Y, v.Z, v.W)
}
