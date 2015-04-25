package chm

import (
	"testing"
)

func TestReadPSF(t *testing.T) {
	_, err := ReadPSFFile("../../testdata/chm_psf/glic.psf")
	if err != nil {
		t.Fatalf("cound not parse the psf file")
	}
}
