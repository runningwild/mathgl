package mathgl

type Quaternion struct {
	X, Y, Z, W float32
}

func (q *Quaternion) RotationMatrix(rotation *Mat3) {}
func (q *Quaternion) QuaternionToAxisAngle() (*Vec3, float32) {
	return nil, 0.0
}
