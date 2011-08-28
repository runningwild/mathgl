// MathGL is a simple 3D math library written in Go which should help writing OpenGL code.
package mathgl

// 3x3 Matrix type. Column major.
type Mat3 [9]float32

// Sets the matrix to a 3x3 identity matrix.
func (m *Mat3) Identity() {
	m[0] = 1
	m[1] = 0
	m[2] = 0

	m[3] = 0
	m[4] = 1
	m[5] = 0

	m[6] = 0
	m[7] = 0
	m[8] = 1
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

	*m = adjugate
}

// Inverse the matrix with the given determinant in float32. Returns true if the inverse could be build.
func (m *Mat3) Inverse(determinate float32) bool {
	var detInv float32

	if determinate == 0.0 {
		return false
	}

	detInv = 1.0 / determinate
	m.Adjugate()
	m.ScalarMultiply(detInv)

	return true
}


// Returns true if the matrix is a identity matrix.
func (m *Mat3) IsIdentity() bool {
	var identity Mat3
	identity.Identity()
	if m.AreEqual(&identity) {
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
	*m = tmp
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

	*m = out
}

// Multiplies the matrix with a given scalar in float32.
func (m *Mat3) ScalarMultiply(factor float32) {
	for i := range m {
		m[i] *= factor
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

// Set the matrix to a scaling matrix, which scale with given x,y floats32
func (m *Mat3) Scaling(x, y float32) {
	m.Identity()
	m[0] = x
	m[4] = y
}


// Set the matrix to a translation matrix, which translates with given x,y floats32
func (m *Mat3) Translation(x, y float32) {
	m.Identity()
	m[6] = x
	m[7] = y
}

// Set the matrix to a matrix that rotates around the x-axis
func (m *Mat3) RotationX(radians float32) {
	m[0] = 1.0
	m[1] = 0.0
	m[2] = 0.0

	m[3] = 0.0
	m[4] = Fcos32(radians)
	m[5] = Fsin32(radians)

	m[6] = 0.0
	m[7] = -Fsin32(radians)
	m[8] = Fcos32(radians)
}

// Set the matrix to a matrix that rotates around the y-axis
func (m *Mat3) RotationY(radians float32) {
	m[0] = Fcos32(radians)
	m[1] = 0.0
	m[2] = -Fsin32(radians)

	m[3] = 0.0
	m[4] = 1.0
	m[5] = 0.0

	m[6] = Fsin32(radians)
	m[7] = 0.0
	m[8] = Fcos32(radians)
}

// Set the matrix to a matrix that rotates around the z-axis
func (m *Mat3) RotationZ(radians float32) {
	m[0] = Fcos32(radians)
	m[1] = Fsin32(radians)
	m[2] = 0.0

	m[3] = -Fsin32(radians)
	m[4] = Fcos32(radians)
	m[5] = 0.0

	m[6] = 0.0
	m[7] = 0.0
	m[8] = 1.0
}

// Sets the matrix to a matrix that rotates with the help of the given quaternion
func (m *Mat3) RotationQuaternion(pIn *Quaternion) {
	m[0] = 1.0 - 2.0*(pIn.y*pIn.y+pIn.z*pIn.z)
	m[1] = 2.0 * (pIn.x*pIn.y - pIn.w*pIn.z)
	m[2] = 2.0 * (pIn.x*pIn.z + pIn.w*pIn.y)

	m[3] = 2.0 * (pIn.x*pIn.y + pIn.w*pIn.z)
	m[4] = 1.0 - 2.0*(pIn.x*pIn.x+pIn.z*pIn.z)
	m[5] = 2.0 * (pIn.y*pIn.z - pIn.w*pIn.x)

	m[6] = 2.0 * (pIn.x*pIn.z - pIn.w*pIn.y)
	m[7] = 2.0 * (pIn.y*pIn.z + pIn.w*pIn.x)
	m[8] = 1.0 - 2.0*(pIn.x*pIn.x+pIn.y*pIn.y)
}

// Sets the matrix to a matrix that rotates with the help of the given vector Vec3 and angle float32
func (m *Mat3) RotationAxisAngle(axis Vec3, radians float32) {
	rcos := Fcos32(radians)
	rsin := Fsin32(radians)

	axis.Normalize()

	m[0] = rcos + axis.x*axis.x*(1-rcos)
	m[1] = axis.z*rsin + axis.y*axis.x*(1-rcos)
	m[2] = -axis.y*rsin + axis.z*axis.x*(1-rcos)

	m[3] = -axis.z*rsin + axis.x*axis.y*(1-rcos)
	m[4] = rcos + axis.y*axis.y*(1-rcos)
	m[5] = axis.x*rsin + axis.z*axis.y*(1-rcos)

	m[6] = axis.y*rsin + axis.x*axis.z*(1-rcos)
	m[7] = -axis.x*rsin + axis.y*axis.z*(1-rcos)
	m[8] = rcos + axis.z*axis.z*(1-rcos)
}
