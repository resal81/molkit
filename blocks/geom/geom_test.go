package geom

import (
	"testing"
)

func TestGeom(t *testing.T) {

}

func TestSliceSubtract(t *testing.T) {
	sl1 := []float64{1, 2, 3}
	sl2 := []float64{5, 6, 7}

	out := SliceSubtract(sl2, sl1)
	if out[0] != 4 || out[1] != 4 || out[2] != 4 {
		t.Errorf("Wrong slice subtractin: expected `[4, 4, 4]`, got `%v`", out)
	}
}
