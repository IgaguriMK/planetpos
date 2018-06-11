package pmath

import (
	"fmt"
	"math"
)

type Vec3 [3]float64

func (v Vec3) String() string {
	return fmt.Sprintf("[%.3f, %.3f, %.3f]", v[0], v[1], v[2])
}

func Diff(v, u Vec3) Vec3 {
	var r Vec3
	r[0] = v[0] - u[0]
	r[1] = v[1] - u[1]
	r[2] = v[2] - u[2]
	return r
}

func Dist(v, u Vec3) float64 {
	d := Diff(v, u)
	return math.Sqrt(d[0]*d[0] + d[1]*d[1] + d[2]*d[2])
}

type Mat3 [3][3]float64

func (m Mat3) String() string {
	return fmt.Sprintf(
		"[%.3f, %.3f, %.3f; %.3f, %.3f, %.3f; %.3f, %.3f, %.3f]",
		m[0][0], m[0][1], m[0][2],
		m[1][0], m[1][1], m[1][2],
		m[2][0], m[2][1], m[2][2],
	)
}

func RotX(angle float64) Mat3 {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return Mat3{
		[3]float64{1, 0, 0},
		[3]float64{0, c, -s},
		[3]float64{0, s, c},
	}
}

func RotZ(angle float64) Mat3 {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return Mat3{
		[3]float64{c, 0, s},
		[3]float64{0, 1, 0},
		[3]float64{-s, 0, c},
	}
}

func Apply(m Mat3, v Vec3) Vec3 {
	return Vec3{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2],
		m[2][0]*v[0] + m[2][1]*v[1] + m[2][2]*v[2],
	}
}
