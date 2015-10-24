package selection

import (
	"bytes"
	"github.com/gonum/matrix/mat64"
	"github.com/resal81/molkit/blocks/atom"
	"github.com/resal81/molkit/blocks/geom"
	"github.com/resal81/molkit/blocks/structure"
	"text/template"
)

// ***********************************************************************//
// Selection Keywords
// ***********************************************************************//

type selKeyword int64

const (
	All selKeyword = 1 << iota
	Protein
	Nucleic
	Lipid
	Water
	Ion
	Calpha
	Backbone
)

// ***********************************************************************//
// Selection Definition
// ***********************************************************************//

type SelDef struct {
	Keyword        selKeyword
	ResidueNames   []string
	ResidueSerials []int
	AtomNames      []string
	AtomSerials    []int
	ChainNames     []string
}

func NewSelDef() *SelDef {
	return &SelDef{
		ResidueNames:   make([]string, 0),
		ResidueSerials: make([]int, 0),
		AtomNames:      make([]string, 0),
		AtomSerials:    make([]int, 0),
		ChainNames:     make([]string, 0),
	}
}

// ***********************************************************************//
// Selection
// ***********************************************************************//

type Selection struct {
	atoms     []*atom.Atom
	currFrame int
}

func NewSelection(st *structure.Structure, sdf *SelDef) *Selection {
	sl := Selection{}

	if sdf.Keyword&All != 0 {
		sl.atoms = st.Atoms()
		return &sl
	}

	panic("Not implemented yet")
}

// NAtoms returns the number of atoms in the selection.
func (sel *Selection) NAtoms() int {
	return len(sel.Atoms())
}

// Atoms returns the underlying slice of atoms for the selection.
func (sel *Selection) Atoms() []*atom.Atom {
	return sel.atoms
}

// NFrames returns the number of frames.
func (sel *Selection) NFrames() int {
	if len(sel.Atoms()) > 0 {
		return len(sel.Atoms()[0].Coords())
	}
	return 0
}

// CurrentFrame returns the currently set frame number.
func (sel *Selection) CurrentFrame() int {
	return sel.currFrame
}

// SetCurrentFrame sets the current frame number.
func (sel *Selection) SetCurrentFrame(frame int) {
	// TODO check that frame is within limits
	sel.currFrame = frame
}

// CurrentMatrix creates a new N x 3 matrix of cooridnates for the current frame
func (sel *Selection) CurrentMatrix() *mat64.Dense {
	crds := make([]float64, sel.NAtoms()*3)
	frame := sel.CurrentFrame()

	for j, at := range sel.Atoms() {
		crd := at.CoordAtFrame(frame)
		crds[j*3] = crd[0]
		crds[j*3+1] = crd[1]
		crds[j*3+2] = crd[2]
	}

	return mat64.NewDense(sel.NAtoms(), 3, crds)
}

// UpdateUsingMatrix uses the provided matrix to update the coordinates of the current frame
func (sel *Selection) UpdateUsingMatrix(mat *mat64.Dense) {
	m, _ := mat.Dims()
	if m != sel.NAtoms() {
		panic("Not same number of coordinate between selection and matrix")
	}

	atoms := sel.Atoms()
	frame := sel.CurrentFrame()
	for i := 0; i < m; i++ {
		atoms[i].SetCoordAtFrame(frame, []float64{mat.At(i, 0), mat.At(i, 1), mat.At(i, 2)})
	}
}

// MinMax returns the minimum and maximum corners of the bounding box
func (sel *Selection) MinMaxDim() ([]float64, []float64, []float64) {
	mat := sel.CurrentMatrix()
	min, max := geom.MatrixColMinMax(mat)
	dim := geom.SliceSubtract(max, min)

	return min, max, dim
}

// CoG returns center of selection.
func (sel *Selection) CoG() []float64 {
	mat := sel.CurrentMatrix()
	return geom.MatrixColMean(mat)
}

// CenterCoG centers the coordinate for the current frame at the origin.
func (sel *Selection) CenterCoG() {
	cog := sel.CoG()
	frame := sel.CurrentFrame()
	for _, at := range sel.Atoms() {
		crd := at.CoordAtFrame(frame)
		crd = []float64{crd[0] - cog[0], crd[1] - cog[1], crd[2] - cog[2]}
		at.SetCoordAtFrame(frame, crd)
	}

}

// func (sel *Selection) CovMatrix() *mat64.Dense {
// 	mat := sel.Matrix()

// 	out := mat64.NewDense(3, 3, make([]float64, 9))
// 	out.Mul(mat.T(), mat)
// 	out.Scale(1.0/float64(sel.NAtoms()-1), out)

// 	return out
// }

// SVD returns (U, S, V) that saticifies `sel.Matrix() == U * S * V.T`.
// U is m x n, S is n x n, and v is n x n.
// func (sel *Selection) PCA() ([]float64, *mat64.Dense) {
// 	cov := sel.CovMatrix()
// 	svd := mat64.SVD(cov, 1e-6, 0.001, true, true)
// 	return svd.Sigma, svd.V
// }

// Info returns some info about this selection group
func (sel *Selection) Info() string {
	tmpl := template.Must(template.New("info").Parse(selInfoTemplate))

	min, max, dim := sel.MinMaxDim()
	data := struct {
		NAtoms    int
		NFrames   int
		CurrFrame int
		CoG       []float64
		Min       []float64
		Max       []float64
		Dim       []float64
	}{
		sel.NAtoms(),
		sel.NFrames(),
		sel.CurrentFrame(),
		sel.CoG(),
		min,
		max,
		dim,
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, data); err != nil {
		return "Could not generate geom info: " + err.Error()
	}

	return b.String()
}

const selInfoTemplate = `
** Selection *******************************
{{.NAtoms}} atoms, {{.NFrames}} frames (current : {{.CurrFrame}})

CoG: {{.CoG | printf "%6.1f"}}
Min: {{.Min | printf "%6.1f"}}
Max: {{.Max | printf "%6.1f"}}
Dim: {{.Dim | printf "%6.1f"}}
--------------------------------------------

`
