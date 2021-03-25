package challenge_01

import (
	"math"
)

type Cylinder struct {
	Height float64
	Radius float64
}

func (cylinder Cylinder) SurfaceArea() float64 {
	return 2.0 * math.Pi * cylinder.Radius * (cylinder.Radius + cylinder.Height)
}
