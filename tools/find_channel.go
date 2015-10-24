package main

import (
	"fmt"
	// "github.com/gonum/matrix/mat64"
	"github.com/resal81/molkit/blocks/selection"
	"github.com/resal81/molkit/molio/pdb"
)

func main() {
	st, err := pdb.ReadPdbFile("../molio/pdb/testdata/vdac.pdb")
	if err != nil {
		panic(err)
	}

	sdf := selection.NewSelDef()
	sdf.Keyword |= selection.All
	sel := selection.NewSelection(st, sdf)

	fmt.Print(sel.Info())
	// atoms := st.Atoms()
	// gm := geom.NewGeom(atoms)

	// fmt.Println("Before centring:")
	// fmt.Print(gm.Info())

	// gm.CenterCoG()
	// fmt.Println("After centring:")
	// fmt.Print(gm.Info())

	// gr := geom.NewGridInt8(gm, 2, 0.2)
	// fmt.Println(gr.Info())

	// // mat := gm.Matrix()
	// // fmt.Println(mat64.Formatted(mat, mat64.Prefix(" "), mat64.Excerpt(3)))

	// sig, v := gm.PCA()
	// fmt.Println(sig)
	// fmt.Println(v)

}
