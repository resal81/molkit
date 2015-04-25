package chm

import (
	"testing"

	//"github.com/resal81/molkit/blocks"
)

func TestTOPRead(t *testing.T) {
	fnames := []string{
		"../../testdata/chm_top/top_all22_prot.rtf",
		"../../testdata/chm_top/top_all35_ethers.rtf",
		"../../testdata/chm_top/top_all36_prot.rtf",
		"../../testdata/chm_top/top_all36_na.rtf",
		"../../testdata/chm_top/top_all36_lipid.rtf",
		"../../testdata/chm_top/top_all36_carb.rtf",
		"../../testdata/chm_top/top_all36_cgenff.rtf",
	}
	_, err := ReadTOPFiles(fnames...)
	if err != nil {
		t.Errorf("could not read prm file -> %s", err)
	}
}
