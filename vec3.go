package mathgl

import "fmt"

// 2 dimensional vector.
type Vec3 struct {
	X, Y, Z float32
}

// Fills the vector with the given float32
func (v *Vec3) Fill(x, y, z float32) {
	v.X = x
	v.Y = y
	v.Z = z
}

// Returns the length  as float32
func (v *Vec3) Length() float32 {
	return Fsqrt32(Fsqr32(v.X) + Fsqr32(v.Y) + Fsqr32(v.Z))
}

// Returns the length as square as float32
func (v *Vec3) LengthSq() float32 {
	return Fsqr32(v.X) + Fsqr32(v.Y) + Fsqr32(v.Z)
}

// Normalize the vector
func (v *Vec3) Normalize() {
	var l float32 = 1.0 / v.Length()
	v.X *= l
	v.Y *= l
	v.Z *= l
}

// Adds the given Vec3 with the vector
func (v *Vec3) Add(x *Vec3) {
	v.X += x.x
	v.Y += x.Y
	v.Z += x.Z
}

// Returns the cosine of the angle between the vectors as float32
func (v *Vec3) Dot(x *Vec3) float32 {
	return v.X*x.X + v.Y*x.Y + v.Z*x.Z
}

// Saves the Vec3 perpendicular to the given Vec3
func (v *Vec3) Cross(x *Vec3) {
	var t Vec3
	t.X = v.X
	t.Y = v.Y
	t.Z = v.Z

	v.X = (t.Y * x.Z) - (t.Z * x.Y)
	v.Y = (t.Z * x.X) - (t.X * x.Z)
	v.Z = (t.X * x.Y) - (t.Y * x.X)
}

// Subtracts the given Vec3 from the vector
func (v *Vec3) Subtract(x *Vec3) {
	v.X -= x.X
	v.Y -= x.Y
	v.Z -= x.Z
}

// Transforms the Vec3 by a given Mat4
func (v *Vec3) Transform(m *Mat4) {
	var t Vec3
	t.X = v.X
	t.Y = v.Y
	t.Z = v.Z

	v.X = t.X*m[0] + t.Y*m[4] + t.Z*m[8] + m[12]
	v.Y = t.X*m[1] + t.Y*m[5] + t.Z*m[9] + m[13]
	v.Z = t.X*m[2] + t.Y*m[6] + t.Z*m[10] + m[14]
}

// Transforms the Vec3 by a given Mat4 inversely
func (v *Vec3) InverseTransform(m *Mat4) {
	var t Vec3
	t.X = v.X - m[12]
	t.Y = v.Y - m[13]
	t.Z = v.Z - m[14]

	v.X = t.X*m[0] + t.Y*m[1] + t.Z*m[2]
	v.Y = t.X*m[4] + t.Y*m[5] + t.Z*m[6]
	v.Z = t.X*m[8] + t.Y*m[9] + t.Z*m[10]
}

// Transform a texture Vec3 with the given Mat4 matrix
func (v *Vec3) TransformCoord(m *Mat4) {
	var t Vec4
	t.Fill(v.X, v.Y, v.Z, 1.0)

	t.Transform(m)

	v.X = t.X / t.w
	v.Y = t.Y / t.w
	v.Z = t.Z / t.w
}

// Transform a normal Vec3 with the given Mat4 matrix. Omits the translation, only scaling + rotating
func (v *Vec3) TransformNormal(m *Mat4) {
	var t Vec3
	t.X = v.X
	t.Y = v.Y
	t.Z = v.Z

	v.X = t.X*m[0] + t.Y*m[4] + t.Z*m[8]
	v.Y = t.X*m[1] + t.Y*m[5] + t.Z*m[9]
	v.Z = t.X*m[2] + t.Y*m[6] + t.Z*m[10]
}

// Transforms a normal Vec3 with the given Mat4 matrix inversely. Omits the translation, only scaling + rotating
func (v *Vec3) InverseTransformNormal(m *Mat4) {
	var t Vec3
	t.X = v.X
	t.Y = v.Y
	t.Z = v.Z

	v.X = t.X*m[0] + t.Y*m[1] + t.Z*m[2]
	v.Y = t.X*m[4] + t.Y*m[5] + t.Z*m[6]
	v.Z = t.X*m[8] + t.Y*m[9] + t.Z*m[10]
}

// Scales a vector to the given length s in float32.
func (v *Vec3) Scale(s float32) {
	v.X *= s
	v.Y *= s
	v.Z *= s
}

// Returns true if the vectors are approximately equal in value
func (v *Vec3) AreEqual(x *Vec3) bool {
	return ((v.X < x.X+epsilon && v.X > x.X-epsilon) &&
		(v.Y < x.Y+epsilon && v.Y > x.Y-epsilon) &&
		(v.Z < x.Z+epsilon && v.Z > x.Z-epsilon))
}

// Assigns the given Vec3 to the Vec3
func (v *Vec3) Assign(x *Vec3) {
	if v == x {
		return
	}

	v.X = x.X
	v.Y = x.Y
	v.Z = x.Z
}

// Sets all the elements of Vec3 to zero
func (v *Vec3) Zero() {
	v.X = 0.0
	v.Y = 0.0
	v.Z = 0.0
}

func (v *Vec3) String() string {
	return fmt.Sprintf("Vec3(%f, %f, %f)", v.X, v.Y, v.Z)
}
