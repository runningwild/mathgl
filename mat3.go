/*
   MathGL is a simple 3D math library which should help writing OpenGL code.
*/
package mathgl

// 3x3 Matrix type.
type Mat3 [9]float32

// Returns a 3x3 identity matrix.
func MakeIdentityMat3() *Mat3 {
	// This code won't trigger the GC
	m := Mat3{1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0}
	return &m
}

// Fills the matrix with the given float32.
func (m *Mat3) Fill(content float32) {
	for i := range m {
		m[i] = content
	}
}

// Returns the calculated determinant from the matrix as float32.
func (m *Mat3) Determinant() float32 {
	var determinant float32

	// We use the rule of sarrus to get the determinant
	determinant = m[0]*m[4]*m[8] + m[1]*m[5]*m[6] + m[2]*m[3]*m[7]
	determinant -= m[2]*m[4]*m[6] + m[0]*m[5]*m[7] + m[1]*m[3]*m[8]

	return determinant
}

// Adjugates the matrix.
func (m *Mat3) Adjugate() {
	var adjugate Mat3

	//  the transpose of its cofactor matrix
	adjugate[0] = m[4]*m[8] - m[5]*m[7]
	adjugate[1] = m[2]*m[7] - m[1]*m[8]
	adjugate[2] = m[1]*m[5] - m[2]*m[4]
	adjugate[3] = m[5]*m[6] - m[3]*m[8]
	adjugate[4] = m[0]*m[8] - m[2]*m[6]
	adjugate[5] = m[2]*m[3] - m[0]*m[5]
	adjugate[6] = m[3]*m[7] - m[4]*m[6]
	adjugate[7] = m[1]*m[6] - m[0]*m[7]
	adjugate[8] = m[0]*m[4] - m[1]*m[3]

	// This will trigger the GC
	m = &adjugate
}

// Inverse the matrix with the given determinant in float32.
func (m *Mat3) Inverse(determinate float32) {
	var detInv float32

	if determinate == 0.0 {
		panic("Division through ZERO at calculating the Inverse!")
	}

	detInv = 1.0 / determinate
	m.Adjugate()
	m.ScalarMultiply(detInv)
}


// Returns true if the matrix is a identity matrix.
func (m *Mat3) IsIdentity() bool {
	identity := MakeIdentityMat3()
	if m.AreEqual(identity) {
		return true
	}
	return false
}

// Transpose the matrix
func (m *Mat3) Transpose() {
	var tmp Mat3
	for z := 0; z < 3; z++ {
		for x := 0; x < 3; x++ {
			tmp[(z*3)+x] = m[(x*3)+z]
		}
	}
	// This will trigger the GC
	m = &tmp
}

// Multiplies the matrix with a given Mat3 matrix
func (m *Mat3) Multiply(in *Mat3) {
	var out Mat3

	out[0] = m[0]*in[0] + m[3]*in[1] + m[6]*in[2]
	out[1] = m[1]*in[0] + m[4]*in[1] + m[7]*in[2]
	out[2] = m[2]*in[0] + m[5]*in[1] + m[8]*in[2]

	out[3] = m[0]*in[3] + m[3]*in[4] + m[6]*in[5]
	out[4] = m[1]*in[3] + m[4]*in[4] + m[7]*in[5]
	out[5] = m[2]*in[3] + m[5]*in[4] + m[8]*in[5]

	out[6] = m[0]*in[6] + m[3]*in[7] + m[6]*in[8]
	out[7] = m[1]*in[6] + m[4]*in[7] + m[7]*in[8]
	out[8] = m[2]*in[6] + m[5]*in[7] + m[8]*in[8]

	// This will trigger the GC
	m = &out
}

// Multiplies the matrix with a given scalar in float32.
func (m *Mat3) ScalarMultiply(factor float32) {
	for i := range m {
		m[i] = m[i] * factor
	}
}

// Assigns the values of the input matrix
func (m *Mat3) Assign(input *Mat3) {
	for i, x := range input {
		m[i] = x
	}
}

// Returns true if the 2 matrices are equal (approximately)
func (m *Mat3) AreEqual(candidate *Mat3) bool {
	for i, x := range candidate {
		if !(m[i]+epsilon > x &&
			m[i]-epsilon < x) {
			return false
		}
	}
	return true
}

// Returns a scaling matrix, which scale with given x,y floats32
func MakeScalingMat3(x, y float32) *Mat3 {
	out := MakeIdentityMat3()
	out[0] = x
	out[4] = y

	return out
}


// Returns a translation matrix, which translates with given x,y floats32
func MakeTranslationMat3(x, y float32) *Mat3 {
	out := MakeIdentityMat3()
	out[6] = x
	out[7] = y

	return out
}

// Returns a matrix that rotates around the x-axis
func MakeRotationXMat3(radians float32) *Mat3 {
	var m mat3

	m[0] = 1.0
	m[1] = 0.0
	m[2] = 0.0

	m[3] = 0.0
	m[4] = Fcos32(radians)
	m[5] = Fsin32(radians)

	m[6] = 0.0
	m[7] = -Fsin32(radians)
	m[8] = Fcos32(radians)

	return &m
}

// Returns a matrix that rotates around the y-axis
func MakeRotationYMat3(radians float32) *Mat3 {
	var m mat3

	m[0] = Fcos32(radians)
	m[1] = 0.0
	m[2] = -Fsin32(radians)

	m[3] = 0.0
	m[4] = 1.0
	m[5] = 0.0

	m[6] = Fsin32(radians)
	m[7] = 0.0
	m[8] = Fcos32(radians)

	return &m
}

// Returns a matrix that rotates around the z-axis
func MakeRotationYMat3(radians float32) *Mat3 {
	var m mat3

	m[0] = Fcos32(radians)
	m[1] = -Fsin32(radians)
	m[2] = 0.0

	m[3] = Fsin32(radians)
	m[4] = Fcos32(radians)
	m[5] = 0.0

	m[6] = 0.0
	m[7] = 0.0
	m[8] = 1.0

	return &m
}
