package atomtype

import (
	"testing"
)

func TestAtomType(t *testing.T) {
	at := NewAtomType("CA")
	if at.PqrCharge() != 0.0 {
		t.Error("PqrCharge is not zero")
	}
}
