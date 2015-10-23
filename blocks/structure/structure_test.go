package structure

import (
	"testing"
)

func TestStructure(t *testing.T) {
	st := NewStructure()
	if len(st.Chains()) != 0 {
		t.Error("The Structure.Chains() return non-zero")
	}
}
