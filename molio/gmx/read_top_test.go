package gmx

import (
	"testing"
)

func TestReadTop(t *testing.T) {
	_, err := ReadTOPFile("../../testdata/tmp.top")
	if err != nil {
		t.Errorf("could not read top file: %s", err)
	}
}
