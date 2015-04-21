package utils

import (
	"github.com/resal81/molkit/blocks"
	"math"
)

func CalcDistance_Atoms(a1, a2 *blocks.Atom) float64 {
	c1 := a1.Coords()[0]
	c2 := a2.Coords()[0]

	r2 := math.Pow(c1[0]-c2[0], 2) + math.Pow(c1[1]-c2[1], 2) + math.Pow(c1[2]-c2[2], 2)
	return math.Sqrt(r2)
}
