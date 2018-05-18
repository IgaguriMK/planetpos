package vec

import "math"

type Vec struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func (v Vec) GetX() float64 { return v.X }
func (v Vec) GetY() float64 { return v.Y }
func (v Vec) GetZ() float64 { return v.Z }

var (
	ZeroVec = Vec{0, 0, 0}
	OneX    = Vec{1, 0, 0}
	OneY    = Vec{0, 1, 0}
	OneZ    = Vec{0, 0, 0}
	One     = Vec{1, 1, 1}
)

type Vec3 interface {
	GetX() float64
	GetY() float64
	GetZ() float64
}

func Abs(v Vec3) float64 {
	return math.Sqrt(v.GetX()*v.GetX() + v.GetY()*v.GetY() + v.GetZ()*v.GetZ())
}

func Add(v, u Vec3) Vec {
	return Vec{
		X: v.GetX() + u.GetX(),
		Y: v.GetY() + u.GetY(),
		Z: v.GetZ() + u.GetZ(),
	}
}

func Diff(v, u Vec3) Vec {
	return Vec{
		X: v.GetX() - u.GetX(),
		Y: v.GetY() - u.GetY(),
		Z: v.GetZ() - u.GetZ(),
	}
}

func Scalar(v Vec3, k float64) Vec {
	return Vec{
		X: k * v.GetX(),
		Y: k * v.GetY(),
		Z: k * v.GetZ(),
	}
}

func Unit(v Vec3) Vec {
	if v.GetX() == 0 && v.GetY() == 0 && v.GetZ() == 0 {
		return ZeroVec
	}

	a := Abs(v)

	return Vec{
		X: v.GetX() / a,
		Y: v.GetY() / a,
		Z: v.GetZ() / a,
	}
}

func Dot(v, u Vec3) float64 {
	return v.GetX()*u.GetX() + v.GetY()*u.GetY() + v.GetZ()*u.GetZ()
}

func Dist(v, u Vec3) float64 {
	return Abs(Diff(u, v))
}

func Near(v, u Vec3, dist float64) bool {
	d := Diff(v, u)

	return d.X*d.X+d.Y*d.Y+d.Z*d.Z < dist*dist
}
