package element

import (
	"testing"
)

func TestElement(t *testing.T) {
	if len(ElementsDatabase) == 0 {
		t.Fatal("Element database is empty")
	}

	C := ElementsDatabase["C"]
	if C.number != 6 || C.mass != 12.0107 {
		t.Errorf("Wrong element data for carbon: %d, %f", C.number, C.mass)
	}
}
