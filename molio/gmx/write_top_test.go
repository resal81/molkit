package gmx

import (
	"testing"
)

func TestWriteTopToString(t *testing.T) {
	s, err := WriteTopToString(nil, nil)
	if err != nil {
		t.Fatalf("%s", err)
	}

	if s != "hello world" {
		t.Errorf("s is %s", s)
	}
}
