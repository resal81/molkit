package atom

import (
	"testing"
)

func TestNewAtom(t *testing.T) {

	a := NewAtom("CA", 1)
	if a.Name() != "CA" || a.Serial() != 1 {
		t.Error("wrong atom info")
	}

}
