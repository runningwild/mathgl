package mathgl

// 4 dimensional vector.
type Vec4 struct {
	x, y, z, w float32
}

// Fills the vector with the given float32
func (v *Vec4) Fill(x, y, z, w float32) {}

func (v *Vec4) Transform(m *Mat4) {}
