package main

import (
	"fmt"
	"math"
)

var step float64 = 2 * math.Pi / 10000

func main() {
	eccentricity := 0.854

	max := OrbitAngle(2*math.Pi, eccentricity)

	fmt.Println("r\ta\ts")
	last := 0.0
	for r := 0.001; r < 1; r += 0.001 {
		a := OrbitAngleInverse(max*r, eccentricity)
		fmt.Printf("%f\t%f\t%f\n", r, a, a-last)
		last = a
	}
}

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
