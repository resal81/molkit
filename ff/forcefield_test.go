package ff

import (
	"testing"
)

func TestNewForceField(t *testing.T) {
	ff := NewForceField(FF_GROMACS)
	if ff.Type()&FF_GROMACS == 0 {
		t.Error("wrong params db type")
	}
}
