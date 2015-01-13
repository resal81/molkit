package chm

import (
	"testing"
)

func TestPRMRead(t *testing.T) {
	// fname := "../../testdata/chm_prm/par_all22_prot.prm"
	fnames := []string{
		"../../testdata/chm_prm/par_all22_prot.prm",
		"../../testdata/chm_prm/par_all35_ethers.prm",
		"../../testdata/chm_prm/par_all36_prot.prm",
		"../../testdata/chm_prm/par_all36_na.prm",
		"../../testdata/chm_prm/par_all36_lipid.prm",
		"../../testdata/chm_prm/par_all36_carb.prm",
		"../../testdata/chm_prm/par_all36_cgenff.prm",
	}
	_, err := ReadPRMFiles(fnames...)
	if err != nil {
		t.Fatalf("could not read prm file: %s", err)
	}
}
