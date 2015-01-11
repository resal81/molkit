package ff

import (
	"testing"
)

func TestNewParmasDB(t *testing.T) {
	pdb := NewParamsDB(P_DB_GROMACS)
	if pdb.Type()&P_DB_GROMACS == 0 {
		t.Error("wrong params db type")
	}
}
