package mathgl

// 4x4 Matrix type. Column major.
type Mat4 [16]float32

// Sets the matrix to a 3x3 identity matrix.
func (m *Mat4) Identity() {
	m[0] = 1
	m[1] = 0
	m[2] = 0
	m[3] = 0

	m[4] = 0
	m[5] = 1
	m[6] = 0
	m[7] = 0

	m[8] = 0
	m[9] = 0
	m[10] = 1
	m[11] = 0

	m[12] = 0
	m[13] = 0
	m[14] = 0
	m[15] = 1
}

// Fills the matrix with the given float32.
func (m *Mat4) Fill(content float32) {
	for i := range m {
		m[i] = content
	}
}

// Returns the calculated determinant from the matrix as float32.
func (m *Mat4) Determinant() float32 {
	var determinant float32

	determinant = m[12]*m[9]*m[6]*m[3] - m[8]*m[13]*m[6]*m[3] - m[12]*m[5]*m[10]*m[3] + m[4]*m[13]*m[10]*m[3] +
		m[8]*m[5]*m[14]*m[3] - m[4]*m[9]*m[14]*m[3] - m[12]*m[9]*m[2]*m[7] + m[8]*m[13]*m[2]*m[7] +
		m[12]*m[1]*m[10]*m[7] - m[0]*m[13]*m[10]*m[7] - m[8]*m[1]*m[14]*m[7] + m[0]*m[9]*m[14]*m[7] +
		m[12]*m[5]*m[2]*m[11] - m[4]*m[13]*m[2]*m[11] - m[12]*m[1]*m[6]*m[11] + m[0]*m[13]*m[6]*m[11] +
		m[4]*m[1]*m[14]*m[11] - m[0]*m[5]*m[14]*m[11] - m[8]*m[5]*m[2]*m[15] + m[4]*m[9]*m[2]*m[15] +
		m[8]*m[1]*m[6]*m[15] - m[0]*m[9]*m[6]*m[15] - m[4]*m[1]*m[10]*m[15] + m[0]*m[5]*m[10]*m[15]
	return determinant
}

// Returns the item at the given row and column
func (m *Mat4) get(row, col int) float32 {
	return m[row+4*col]
}

// Sets the item at the given row and column
func (m *Mat4) set(row, col int, value float32) {
	m[row+4*col] = value
}

// Swaps the given items at the given locations
func (m *Mat4) swap(r1, c1, r2, c2 int) {
	tmp := m.get(r1, c1)
	m.set(r1, c1, m.get(r2, c2))
	m.set(r2, c2, tmp)
}


//Returns an upper and a lower triangular matrix which are L and R in the Gauss algorithm
func gaussj(a, b *Mat4) bool {
	var i, j, k, l, ll, icol, irow int
	var n, m int = 4, 4
	var big, dum, pivinv float32
	var indxc [4]int
	var indxr [4]int
	var ipiv [4]int

	for i = 0; i < n; i++ {
		big = 0.0
		for j = 0; j < n; j++ {
			if ipiv[j] != 1 {
				for k = 0; k < n; k++ {
					if ipiv[k] == 0 {
						if Fabs32(a.get(j, k)) >= big {
							big = Fabs32(a.get(j, k))
							irow = j
							icol = k
						}
					}
				}
			}
		}
		(ipiv[icol])++
		if irow != icol {
			for l = 0; l < n; l++ {
				a.swap(irow, l, icol, l)
			}
			for l = 0; l < m; l++ {
				b.swap(irow, l, icol, l)
			}
		}
		indxr[i] = irow
		indxc[i] = icol
		if a.get(icol, icol) == 0.0 {
			return false
		}
		pivinv = 1.0 / a.get(icol, icol)
		a.set(icol, icol, 1.0)
		for l = 0; l < n; l++ {
			a.set(icol, l, a.get(icol, l)*pivinv)
		}
		for l = 0; l < m; l++ {
			b.set(icol, l, b.get(icol, l)*pivinv)
		}

		for ll = 0; ll < n; ll++ {
			if ll != icol {
				dum = a.get(ll, icol)
				a.set(ll, icol, 0.0)
				for l = 0; l < n; l++ {
					a.set(ll, l, a.get(ll, l)-a.get(icol, l)*dum)
				}
				for l = 0; l < m; l++ {
					b.set(ll, l, a.get(ll, l)-b.get(icol, l)*dum)
				}
			}
		}
	}
	//    This is the end of the main loop over columns of the reduction. It only remains to unscram-
	//    ble the solution in view of the column interchanges. We do this by interchanging pairs of
	//    columns in the reverse order that the permutation was built up.
	for l = n - 1; l >= 0; l-- {
		if indxr[l] != indxc[l] {
			for k = 0; k < n; k++ {
				a.swap(k, indxr[l], k, indxc[l])
			}
		}
	}
	return true
}

// Inverse the matrix with the given determinant in float32. Returns true if the inverse could be build.
func (m *Mat4) Inverse(determinate float32) bool {
	var inv, tmp Mat4
	inv.Assign(m)
	tmp.Identity()

	if gaussj(&inv, &tmp) == false {
		return false
	}

	m.Assign(&inv)
	return true
}


// Returns true if the matrix is a identity matrix.
func (m *Mat4) IsIdentity() bool {
	var identity Mat4
	identity.Identity()
	if m.AreEqual(&identity) {
		return true
	}
	return false
}

// Transpose the matrix
func (m *Mat4) Transpose() {
	var tmp Mat4
	for z := 0; z < 4; z++ {
		for x := 0; x < 4; x++ {
			tmp[(z*4)+x] = m[(x*4)+z]
		}
	}
	*m = tmp
}

// Multiplies the matrix with a given Mat4 matrix
func (m *Mat4) Multiply(in *Mat4) {
	var out Mat4

	// TODO: Anybody want to write some SSE code for the AMD64?
	out[0] = m[0]*in[0] + m[4]*in[1] + m[8]*in[2] + m[12]*in[3]
	out[1] = m[1]*in[0] + m[5]*in[1] + m[9]*in[2] + m[13]*in[3]
	out[2] = m[2]*in[0] + m[6]*in[1] + m[10]*in[2] + m[14]*in[3]
	out[3] = m[3]*in[0] + m[7]*in[1] + m[11]*in[2] + m[15]*in[3]

	out[4] = m[0]*in[4] + m[4]*in[5] + m[8]*in[6] + m[12]*in[7]
	out[5] = m[1]*in[4] + m[5]*in[5] + m[9]*in[6] + m[13]*in[7]
	out[6] = m[2]*in[4] + m[6]*in[5] + m[10]*in[6] + m[14]*in[7]
	out[7] = m[3]*in[4] + m[7]*in[5] + m[11]*in[6] + m[15]*in[7]

	out[8] = m[0]*in[8] + m[4]*in[9] + m[8]*in[10] + m[12]*in[11]
	out[9] = m[1]*in[8] + m[5]*in[9] + m[9]*in[10] + m[13]*in[11]
	out[10] = m[2]*in[8] + m[6]*in[9] + m[10]*in[10] + m[14]*in[11]
	out[11] = m[3]*in[8] + m[7]*in[9] + m[11]*in[10] + m[15]*in[11]

	out[12] = m[0]*in[12] + m[4]*in[13] + m[8]*in[14] + m[12]*in[15]
	out[13] = m[1]*in[12] + m[5]*in[13] + m[9]*in[14] + m[13]*in[15]
	out[14] = m[2]*in[12] + m[6]*in[13] + m[10]*in[14] + m[14]*in[15]
	out[15] = m[3]*in[12] + m[7]*in[13] + m[11]*in[14] + m[15]*in[15]

	*m = out
}

// Multiplies the matrix with a given scalar in float32.
func (m *Mat4) ScalarMultiply(factor float32) {
	for i := range m {
		m[i] = m[i] * factor
	}
}

// Assigns the values of the input matrix
func (m *Mat4) Assign(input *Mat4) {
	for i, x := range input {
		m[i] = x
	}
}

// Returns true if the 2 matrices are equal (approximately)
func (m *Mat4) AreEqual(candidate *Mat4) bool {
	for i, x := range candidate {
		if !(m[i]+epsilon > x &&
			m[i]-epsilon < x) {
			return false
		}
	}
	return true
}

// Set the matrix to a scaling matrix, which scale with given x,y floats32
func (m *Mat4) Scaling(x, y, z float32) {
	m.Identity()
	m[0] = x
	m[5] = y
	m[10] = z
}


// Set the matrix to a translation matrix, which translates with given x,y floats32
func (m *Mat4) Translation(x, y, z float32) {
	m.Identity()
	m[12] = x
	m[13] = y
	m[14] = z
}

// Set the matrix to a matrix that rotates around the x-axis
func (m *Mat4) RotationX(radians float32) {
	m[0] = 1.0
	m[1] = 0.0
	m[2] = 0.0
	m[3] = 1.0

	m[4] = 0.0
	m[5] = Fcos32(radians)
	m[6] = Fsin32(radians)
	m[7] = 0.0

	m[8] = 0.0
	m[9] = -Fsin32(radians)
	m[10] = Fcos32(radians)
	m[11] = 0.0

	m[12] = 0.0
	m[13] = 0.0
	m[14] = 0.0
	m[15] = 1.0
}

// Set the matrix to a matrix that rotates around the y-axis
func (m *Mat4) RotationY(radians float32) {
	m[0] = Fcos32(radians)
	m[1] = 0.0
	m[2] = -Fsin32(radians)
	m[3] = 0.0

	m[4] = 0.0
	m[5] = 1.0
	m[6] = 0.0
	m[7] = 0.0

	m[8] = Fsin32(radians)
	m[9] = 0.0
	m[10] = Fcos32(radians)
	m[11] = 0.0

	m[12] = 0.0
	m[13] = 0.0
	m[14] = 0.0
	m[15] = 1.0
}

// Set the matrix to a matrix that rotates around the z-axis
func (m *Mat4) RotationZ(radians float32) {
	m[0] = Fcos32(radians)
	m[1] = Fsin32(radians)
	m[2] = 0.0
	m[3] = 0.0

	m[4] = -Fsin32(radians)
	m[5] = Fcos32(radians)
	m[6] = 0.0
	m[7] = 0.0

	m[8] = 0.0
	m[9] = 0.0
	m[10] = 1.0
	m[11] = 0.0

	m[12] = 0.0
	m[13] = 0.0
	m[14] = 0.0
	m[15] = 1.0
}

// Sets the matrix to a matrix that rotates with the help of the given quaternion
func (m *Mat4) RotationQuaternion(pIn *Quaternion) {
	m[0] = 1.0 - 2.0*(pIn.y*pIn.y+pIn.z*pIn.z)
	m[1] = 2.0 * (pIn.x*pIn.y - pIn.w*pIn.z)
	m[2] = 2.0 * (pIn.x*pIn.z + pIn.w*pIn.y)
	m[3] = 0.0

	m[4] = 2.0 * (pIn.x*pIn.y + pIn.w*pIn.z)
	m[5] = 1.0 - 2.0*(pIn.x*pIn.x+pIn.z*pIn.z)
	m[6] = 2.0 * (pIn.y*pIn.z - pIn.w*pIn.x)
	m[7] = 0.0

	m[8] = 2.0 * (pIn.x*pIn.z - pIn.w*pIn.y)
	m[9] = 2.0 * (pIn.y*pIn.z + pIn.w*pIn.x)
	m[10] = 1.0 - 2.0*(pIn.x*pIn.x+pIn.y*pIn.y)
	m[11] = 0.0

	m[12] = 0.0
	m[13] = 0.0
	m[14] = 0.0
	m[15] = 1.0
}

// Sets the matrix to a matrix that rotates with the help of the given vector Vec3 and angle float32
func (m *Mat4) RotationAxisAngle(axis Vec3, radians float32) {
	rcos := Fcos32(radians)
	rsin := Fsin32(radians)

	axis.Normalize()

	m[0] = rcos + axis.x*axis.x*(1-rcos)
	m[1] = axis.z*rsin + axis.y*axis.x*(1-rcos)
	m[2] = -axis.y*rsin + axis.z*axis.x*(1-rcos)
	m[3] = 0.0

	m[4] = -axis.z*rsin + axis.x*axis.y*(1-rcos)
	m[5] = rcos + axis.y*axis.y*(1-rcos)
	m[6] = axis.x*rsin + axis.z*axis.y*(1-rcos)
	m[7] = 0.0

	m[8] = axis.y*rsin + axis.x*axis.z*(1-rcos)
	m[9] = -axis.x*rsin + axis.y*axis.z*(1-rcos)
	m[10] = rcos + axis.z*axis.z*(1-rcos)
	m[11] = 0.0

	m[12] = 0.0
	m[13] = 0.0
	m[14] = 0.0
	m[15] = 1.0
}
