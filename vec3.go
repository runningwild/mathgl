package mathgl

import "fmt"

// 2 dimensional vector.
type Vec3 struct {
	x, y, z float32
}

// Fills the vector with the given float32
func (v *Vec3) Fill(x, y, z float32) {
	v.x = x
	v.y = y
	v.z = z
}

// Returns the length  as float32
func (v *Vec3) Length() float32 {
	return Fsqrt32(Fsqr32(v.x) + Fsqr32(v.y) + Fsqr32(v.z))
}

// Returns the length as square as float32
func (v *Vec3) LengthSq() float32 {
	return Fsqr32(v.x) + Fsqr32(v.y) + Fsqr32(v.z)
}

// Normalize the vector
func (v *Vec3) Normalize() {
	var l float32 = 1.0 / v.Length()
	v.x *= l
	v.y *= l
	v.z *= l
}

// Adds the given Vec3 with the vector
func (v *Vec3) Add(x *Vec3) {
	v.x += x.x
	v.y += x.y
	v.z += x.z
}

// Returns the cosine of the angle between the vectors as float32
func (v *Vec3) Dot(x *Vec3) float32 {
	return v.x*x.x + v.y*x.y + v.z*x.z
}

// Saves the Vec3 perpendicular to the given Vec3
func (v *Vec3) Cross(x *Vec3) {
	var t Vec3
	t.x = v.x
	t.y = v.y
	t.z = v.z

	v.x = (t.y * x.z) - (t.z * x.y)
	v.y = (t.z * x.x) - (t.x * x.z)
	v.z = (t.x * x.y) - (t.y * x.x)
}

// Subtracts the given Vec3 from the vector
func (v *Vec3) Subtract(x *Vec3) {
	v.x -= x.x
	v.y -= x.y
	v.z -= x.z
}

// Transforms the Vec3 by a given Mat4
func (v *Vec3) Transform(m *Mat4) {
	var t Vec3
	t.x = v.x
	t.y = v.y
	t.z = v.z

	v.x = t.x*m[0] + t.y*m[4] + t.z*m[8] + m[12]
	v.y = t.x*m[1] + t.y*m[5] + t.z*m[9] + m[13]
	v.z = t.x*m[2] + t.y*m[6] + t.z*m[10] + m[14]
}

// Transforms the Vec3 by a given Mat4 inversely
func (v *Vec3) InverseTransform(m *Mat4) {
	var t Vec3
	t.x = v.x - m[12]
	t.y = v.y - m[13]
	t.z = v.z - m[14]

	v.x = t.x*m[0] + t.y*m[1] + t.z*m[2]
	v.y = t.x*m[4] + t.y*m[5] + t.z*m[6]
	v.z = t.x*m[8] + t.y*m[9] + t.z*m[10]
}

// Transform a texture Vec3 with the given Mat4 matrix
func (v *Vec3) TransformCoord(m *Mat4) {
	var t Vec4
	t.Fill(v.x, v.y, v.z, 1.0)

	t.Transform(m)

	v.x = t.x / t.w
	v.y = t.y / t.w
	v.z = t.z / t.w
}

// Transform a normal Vec3 with the given Mat4 matrix. Omits the translation, only scaling + rotating
func (v *Vec3) TransformNormal(m *Mat4) {
	var t Vec3
	t.x = v.x
	t.y = v.y
	t.z = v.z

	v.x = t.x*m[0] + t.y*m[4] + t.z*m[8]
	v.y = t.x*m[1] + t.y*m[5] + t.z*m[9]
	v.z = t.x*m[2] + t.y*m[6] + t.z*m[10]
}

// Transforms a normal Vec3 with the given Mat4 matrix inversely. Omits the translation, only scaling + rotating
func (v *Vec3) InverseTransformNormal(m *Mat4) {
	var t Vec3
	t.x = v.x
	t.y = v.y
	t.z = v.z

	v.x = t.x*m[0] + t.y*m[1] + t.z*m[2]
	v.y = t.x*m[4] + t.y*m[5] + t.z*m[6]
	v.z = t.x*m[8] + t.y*m[9] + t.z*m[10]
}

// Scales a vector to the given length s in float32. Does normalize the vector beforehand!
func (v *Vec3) Scale(s float32) {
	v.Normalize()
	v.x *= s
	v.y *= s
	v.z *= s
}

// Returns true if the vectors are approximately equal in value
func (v *Vec3) AreEqual(x *Vec3) bool {
	return ((v.x < x.x+epsilon && v.x > x.x-epsilon) &&
		(v.y < x.y+epsilon && v.y > x.y-epsilon) &&
		(v.z < x.z+epsilon && v.z > x.z-epsilon))
}

// Assigns the given Vec3 to the Vec3
func (v *Vec3) Assign(x *Vec3) {
	if v == x {
		return
	}

	v.x = x.x
	v.y = x.y
	v.z = x.z
}

// Sets all the elements of Vec3 to zero
func (v *Vec3) Zero() {
	v.x = 0.0
	v.y = 0.0
	v.z = 0.0
}

func (v *Vec3) String() string {
	return fmt.Sprintf("Vec3(%f, %f, %f)", v.x, v.y, v.z)
}
