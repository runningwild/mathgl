package mathgl

import "fmt"

// 4 dimensional vector.
type Vec4 struct {
	x, y, z, w float32
}

// Fills the vector with the given float32
func (v *Vec4) Fill(x, y, z, w float32) {
	v.x = x
	v.y = y
	v.z = z
	v.w = w
}

// Returns the length  as float32
func (v *Vec4) Length() float32 {
	return Fsqrt32(Fsqr32(v.x) + Fsqr32(v.y) + Fsqr32(v.z) + Fsqr32(v.w))
}

// Returns the length as square as float32
func (v *Vec4) LengthSq() float32 {
	return Fsqr32(v.x) + Fsqr32(v.y) + Fsqr32(v.z) + Fsqr32(v.w)
}

// Normalize the vector
func (v *Vec4) Normalize() {
	var l float32 = 1.0 / v.Length()
	v.x *= l
	v.y *= l
	v.z *= l
	v.w *= l
}

// Adds the given Vec4 with the vector
func (v *Vec4) Add(x *Vec4) {
	v.x += x.x
	v.y += x.y
	v.z += x.z
	v.w += x.w
}

// Returns the cosine of the angle between the vectors as float32
func (v *Vec4) Dot(x *Vec4) float32 {
	// Todo should this better be a Vec3 Dot product?
	return v.x*x.x + v.y*x.y + v.z*x.z + v.w*x.w
}

// Saves the Vec4 perpendicular to the given Vec4 (Attention: This is a Vec3 Cross Product in a homogeneous 4D environment!)
func (v *Vec4) Cross(x *Vec4) {
	var t Vec4
	t.Assign(v)

	v.x = (t.y * x.z) - (t.z * x.y)
	v.y = (t.z * x.x) - (t.x * x.z)
	v.z = (t.x * x.y) - (t.y * x.x)
}

// Subtracts the given Vec4 from the vector
func (v *Vec4) Subtract(x *Vec4) {
	v.x -= x.x
	v.y -= x.y
	v.z -= x.z
	v.w -= v.w
}

// Transforms the Vec4 by a given Mat4
func (v *Vec4) Transform(m *Mat4) {
	var t Vec4
	t.Assign(v)

	v.x = t.x*m[0] + t.y*m[4] + t.z*m[8] + t.w*m[12]
	v.y = t.x*m[1] + t.y*m[5] + t.z*m[9] + t.w*m[13]
	v.z = t.x*m[2] + t.y*m[6] + t.z*m[10] + t.w*m[14]
	v.w = t.x*m[3] + t.y*m[7] + t.z*m[11] + t.w*m[15]
}

/// Loops through an input slice transforming each Vec4 by the given Mat3
func (v *Vec3) TransformArray(x []Vec4, m *Mat4) {
	for _, item := range x {
		// TODO: We should test this
		item.Transform(m)
	}
}


// Scales a vector to the given length s in float32. Does not normalize first, you should do that!
func (v *Vec4) Scale(s float32) {
	v.Normalize()
	v.x *= s
	v.y *= s
	v.z *= s
	v.w *= s
}

// Returns true if the vectors are approximately equal in value
func (v *Vec4) AreEqual(x *Vec4) bool {
	return ((v.x < x.x+epsilon && v.x > x.x-epsilon) &&
		(v.y < x.y+epsilon && v.y > x.y-epsilon) &&
		(v.z < x.z+epsilon && v.z > x.z-epsilon) &&
		(v.w < x.w+epsilon && v.w > x.w-epsilon))
}

// Assigns the given Vec4 to the Vec4
func (v *Vec4) Assign(x *Vec4) {
	if v == x {
		return
	}

	v.x = x.x
	v.y = x.y
	v.z = x.z
	v.w = x.w
}

// Sets all the elements of Vec4 to zero
func (v *Vec4) Zero() {
	v.x = 0.0
	v.y = 0.0
	v.z = 0.0
	v.w = 0.0
}

func (v *Vec4) String() string {
	return fmt.Sprintf("Vec4(%f, %f, %f, %f)", v.x, v.y, v.z, v.w)
}
