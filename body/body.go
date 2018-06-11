package body

import (
	"math"
	"math/rand"

	"github.com/IgaguriMK/planetpos/model"
)

const (
	DaySeconds = 24 * 60 * 60
	AU         = 149597870700.0
	LS         = 299792458.0
	AU_per_LS  = AU / LS
)

type Body struct {
	IsReference    bool
	Name           string
	SemiMajorAxis  float64
	Eccentricity   float64
	ArgOfPeriapsis float64
	Inclination    float64
	Period         float64
}

func FromJsonModel(b model.Body, refBody string) *Body {
	return &Body{
		IsReference:    b.Name == refBody,
		Name:           b.Name,
		SemiMajorAxis:  b.SemiMajorAxis * AU_per_LS,
		Eccentricity:   b.OrbitalEccentricity,
		ArgOfPeriapsis: degToRad(b.ArgOfPeriapsis),
		Inclination:    degToRad(b.OrbitalInclination),
		Period:         float64(b.OrbitalPeriod * DaySeconds),
	}
}

func degToRad(deg float64) float64 {
	return math.Pi * deg / 180
}

func (b *Body) CalcPos(v Variable, t int64) pmath.Vec3 {
	// 軌道上の角座標の角度θ(angle)
	meanRot := v.MeanAnomary / (2 * math.Pi)
	tRot := float64(t) / b.Period

	rotRatio := math.Mod(tRot+meanRot, 1.0)
	maxOmega := OrbitAngle(2*math.Pi, b.Eccentricity)

	angle := OrbitAngleInverse(rotRatio*maxOmega, b.Eccentricity)

	// 軌道面上の位置 P_plane
	l := b.SemiMajorAxis * (1 - b.Eccentricity*b.Eccentricity) / 2
	r := l / (1 + b.Eccentricity*math.Cos(angle+b.ArgOfPeriapsis))

	pPlane := pmath.Vec3{
		r * math.Cos(angle),
		0,
		r * math.Sin(angle),
	}

	// 軌道傾斜角i
	pIncl := pmath.Apply(
		pmath.RotX(b.Inclination),
		pPlane,
	)

	// 昇交点赤経Ω
	pSpace := pmath.Apply(
		pmath.RotY(v.AscendingNode),
		pIncl,
	)

	return pSpace
}

var step float64 = 2 * math.Pi / 10000

func OrbitAngle(angle, eccentricity float64) float64 {
	if step == 0.0 {
		panic("step is zero.")
	}

	var sum float64 = 0
	var rad float64

	for rad = 0.0; rad < angle; rad += step {
		sum += step * orbitAngleInner(rad, eccentricity)
	}

	return sum
}

func OrbitAngleInverse(value, eccentricity float64) float64 {
	if step == 0.0 {
		panic("step is zero.")
	}

	var sum float64 = 0.0
	var angle float64 = 0.0

	for angle < 2*math.Pi {
		sum += step * orbitAngleInner(angle, eccentricity)
		if sum > value {
			return angle
		}
		angle += step
	}

	return 2 * math.Pi
}

func orbitAngleInner(angle, eccentricity float64) float64 {
	return 1 / math.Pow(1+eccentricity*math.Cos(angle), 2)
}

type Variable struct {
	AscendingNode float64
	MeanAnomary   float64
	Error         float64
	Evaluated     bool
}

func RandomMeanAnomary(count int) []Variable {
	res := make([]Variable, 0, count)

	for i := 0; i < count; i++ {
		res = append(
			res,
			Variable{
				AscendingNode: 0,
				MeanAnomary:   2 * math.Pi * rand.Float64(),
			},
		)
	}
}

func RandomVariables(count int) []Variable {
	res := make([]Variable, 0, count)

	for i := 0; i < count; i++ {
		res = append(
			res,
			Variable{
				AscendingNode: 2 * math.Pi * rand.Float64(),
				MeanAnomary:   2 * math.Pi * rand.Float64(),
			},
		)
	}
}
