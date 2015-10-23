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

	fmt.Println("Before centring:")
	fmt.Print(gm.Info())

	gm.CenterCoG()
	fmt.Println("After centring:")
	fmt.Print(gm.Info())

	gr := geom.NewGridInt8(gm, 2, 0.2)
	fmt.Println(gr.Info())

}
