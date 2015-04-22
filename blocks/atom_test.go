package blocks

import (
	"testing"
)

func TestNewAtom(t *testing.T) {
	a1 := NewAtom()
	a2 := NewAtom()

	if a1.Id() != 1 {
		t.Errorf("a1.Id() != 1, instead is %d", a1.Id())
	}

	if a2.Id() != 2 {
		t.Errorf("a2.Id() != 2, instead is %d", a2.Id())
	}
}
