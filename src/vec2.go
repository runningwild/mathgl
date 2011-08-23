package mathgl

type Vec2 struct {
    x float32
    y float32
}

func (v *Vec2) Fill(x float32, y float32) *Vec2 {
    v.x = x
    v.y = y
}

func (v *Vec2) Length() float32 {
    return Fsqrt32(Fsqr32(v.x) + Fsqr32(v.y))
}

func (v *Vec2) Lengthsq() float32 {
    return Fsqr32(v.x) + Fsqr32(v.y)
}

func (v *Vec2) Normalize() {
	var l float32 = 1.0 / v.Length()
	v.x *= l
	v.y *= l
}

func (v *Vec2) Add(x *Vec2) {
	v.x += x.x
	v.y += x.y
}

func (v *Vec2) Dot(x *Vec2) float32 {
    return v.x * x.x + v.y * x.y
}

func (v *Vec2) Subtract(x *Vec2) {
	v.x -= x.x
	v.y -= x.y
}

func (v *Vec2) Transform(m *Mat3) {
    var t Vec2
    // TODO: Remove this code when we have copy
    t.x = v.x
    t.y = v.y

    v.x = t.x * m[0] + t.y * m[3] + m[6]
    v.y = t.x * m[1] + t.y * m[4] + m[7]
}

func (v *Vec2) Scale(s float32) {
	v.x = v.x * s
	v.y = v.y * s
}

func (v *Vec2) Areequal(x *Vec2) bool {
    return ((v.x < x.x + epsilon && v.x > x.x - epsilon) &&
	    (v.y < x.y + epsilon && v.y > x.y - epsilon))
}
