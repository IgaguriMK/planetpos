package model

type Bodies struct {
	Name    string `json:"name"`
	RefBody string `json:"refBody"`
	Bodies  []Body `json:"bodies"`
}

type Body struct {
	Name                string  `json:"name"`
	ArgOfPeriapsis      float64 `json:"argOfPeriapsis"`
	OrbitalEccentricity float64 `json:"orbitalEccentricity"`
	OrbitalInclination  float64 `json:"orbitalInclination"`
	OrbitalPeriod       float64 `json:"orbitalPeriod"`
	SemiMajorAxis       float64 `json:"semiMajorAxis"`
	Offset              int64   `json:"offset"`
	Radius              float64 `json:"radius"`
	Type                string  `json:"type"`
}
