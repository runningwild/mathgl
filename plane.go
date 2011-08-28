package mathgl

type PlaneEnum int

const (
	PLANE_LEFT PlaneEnum = iota
	PLANE_RIGHT
	PLANE_BOTTOM
	PLANE_TOP
	PLANE_NEAR
	PLANE_FAR
)

type PointClassificationEnum int

const (
	POINT_INFRONT_OF_PLANE PointClassificationEnum = iota
	POINT_BEHIND_PLANE
	POINT_ON_PLANE
)

type Plane struct {
	A, B, C, D float32
}
