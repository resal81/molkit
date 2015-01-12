package ff

import (
	"testing"
)

func TestNewForceField(t *testing.T) {
	ff := NewForceField(FF_SOURCE_GROMACS)
	if ff.Kind()&FF_SOURCE_GROMACS == 0 {
		t.Error("wrong params db type")
	}
}
