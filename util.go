package mathgl

import (
  "fmt"
)

// Polys need to be defined in clock-wise order
type Poly []Vec2

func (p *Poly) Clip(s *Seg2) {
  var start int
  for start = 0; start < len(*p); start++ {
    if s.Right(&(*p)[start]) {
      break
    }
  }
  if start == len(*p) {
    *p = (*p)[0:0]
    return
  }
  clip1, clip2 := -1, -1
  var isect1, isect2 Vec2
  for _i := range (*p) {
    prev := (_i + start) % len(*p)
    i := (_i + start + 1) % len(*p)
    if clip1 == -1 {
      if s.Left(&(*p)[i]) {
        clip1 = i
        seg := Seg2{(*p)[i], (*p)[prev]}
        isect1 = seg.Isect(s)
      }
    } else {
      if !s.Left(&(*p)[i]) {
        clip2 = i
        seg := Seg2{(*p)[i], (*p)[prev]}
        isect2 = seg.Isect(s)
        break
      }
    }
  }
  if clip2 == -1 {
    return
  }
  fmt.Printf("Isect at %d: %v, %d: %v\n", clip1, isect1, clip2, isect2)
  var clipper Poly
  clipper = append(clipper, isect1)
  clipper = append(clipper, isect2)
  for _i := range *p {
    i := (_i + clip2) % len(*p)
    if i == clip1 { break }
    clipper = append(clipper, (*p)[i])
  }
  *p = clipper
}



type Seg2 struct {
  A, B Vec2
}

func (a Seg2) Ray() Vec2 {
  var v Vec2
  v.Assign(&a.B)
  v.Subtract(&a.A)
  return v
}

// Returns a Vec2 indicating the intersection point of the lines passing
// through segments a and b
func (u Seg2) Isect(v *Seg2) Vec2 {
  vy := v.B.Y - v.A.Y
  vx := v.A.X - v.B.X
  n := (v.A.X - u.A.X) * vy + (v.A.Y - u.A.Y) * vx
  d := (u.B.X - u.A.X) * vy + (u.B.Y - u.A.Y) * vx
  f := n/d
  return Vec2{ u.A.X + (u.B.X - u.A.X) * f, u.A.Y + (u.B.Y - u.A.Y) * f}
}
func (a Seg2) DistFromOrigin() float32 {
  a_ray := a.Ray()
  a_ray.Cross()
  r := (Seg2{a_ray, Vec2{0,0}}).Isect(&a)
  return r.Length()
}

// Returns true iff u lies to the left of a
func (a Seg2) Left(u *Vec2) bool {
  v := a.Ray()
  v.Cross()
  var v2 Vec2
  v2.Assign(u)
  v2.Subtract(&a.A)
  return v.Dot(&v2) > 0
}

// Returns true iff u lies to the left of a
func (a Seg2) Right(u *Vec2) bool {
  v := a.Ray()
  v.Cross()
  var v2 Vec2
  v2.Assign(u)
  v2.Subtract(&a.A)
  return v.Dot(&v2) < 0
}


