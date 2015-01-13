package chm

import (
	"testing"
)

func TestPRMRead(t *testing.T) {
	fname := "../../testdata/chm_prm/par_all22_prot.prm"
	_, err := ReadPRMFiles(fname)
	if err != nil {
		t.Fatalf("could not read prm file: %s", err)
	}
}
