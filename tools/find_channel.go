package main

import (
	"fmt"
	"github.com/resal81/molkit/blocks/geom"
	"github.com/resal81/molkit/molio/pdb"
)

func main() {
	st, err := pdb.ReadPdbFile("../molio/pdb/testdata/vdac.pdb")
	if err != nil {
		panic(err)
	}

	atoms := st.Atoms()
	gm := geom.NewGeom(atoms)
	fmt.Print(gm.Info())

}
