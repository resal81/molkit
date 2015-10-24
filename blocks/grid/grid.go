package geom

import (
	"bytes"
	"text/template"
)

type GridInt8 struct {
	gpoints [3]int
	gmin    [3]float64
	gmax    [3]float64
	spacing float64
	grid    [][][]int8
}

// NewGridInt8 generates a 3d grid that cover the specified Geom.
// `padding` is the additional pad at each direction, and `spacing` is the
// distance between grid points.
func NewGridInt8(gm *Geom, padding float64, spacing float64) *GridInt8 {
	gr := GridInt8{
		spacing: spacing,
	}

	gmin, gmax, gpoints := findGridSize(gm, padding, spacing)
	gr.gmin = gmin
	gr.gmax = gmax
	gr.gpoints = gpoints

	gr.grid = make([][][]int8, gr.gpoints[0])
	for i := range gr.grid {
		gr.grid[i] = make([][]int8, gr.gpoints[1])

		for j := range gr.grid[i] {
			gr.grid[i][j] = make([]int8, gr.gpoints[2])
		}
	}

	return &gr
}

// findGridSize is a helper method that finds min, max and points of the grid.
func findGridSize(gm *Geom, padding float64, spacing float64) ([3]float64, [3]float64, [3]int) {
	min, max, _ := gm.MinMaxDim()

	gmin := [3]float64{
		min[0] - padding,
		min[1] - padding,
		min[2] - padding,
	}

	gmax := [3]float64{
		max[0] + padding,
		max[1] + padding,
		max[2] + padding,
	}

	gpoints := [3]int{
		int((gmax[0] - gmin[0]) / spacing),
		int((gmax[1] - gmin[1]) / spacing),
		int((gmax[2] - gmin[2]) / spacing),
	}

	return gmin, gmax, gpoints
}

func (gr *GridInt8) Grid() [][][]int8 {
	return gr.grid
}

func (gr *GridInt8) Min() [3]float64 {
	return gr.gmin
}

func (gr *GridInt8) Max() [3]float64 {
	return gr.gmax
}

func (gr *GridInt8) Points() [3]int {
	return gr.gpoints
}

func (gr *GridInt8) coordToIndex(crd []float64) []int {
	nx := int((crd[0] - gr.gmin[0]) / gr.spacing)
	ny := int((crd[1] - gr.gmin[1]) / gr.spacing)
	nz := int((crd[2] - gr.gmin[2]) / gr.spacing)

	return []int{nx, ny, nz}
}

func (gr *GridInt8) indexToCoord(ind []int) []float64 {
	x := float64(ind[0])*gr.spacing + gr.gmin[0]
	y := float64(ind[1])*gr.spacing + gr.gmin[1]
	z := float64(ind[2])*gr.spacing + gr.gmin[2]

	return []float64{x, y, z}
}

func (gr *GridInt8) SetValue(crd []float64, val int8) {
	ind := gr.coordToIndex(crd)
	gr.Grid()[ind[0]][ind[1]][ind[2]] = val
}

func (gr *GridInt8) Info() string {
	tmpl := template.Must(template.New("info").Parse(gridInfoTemplate))

	ntot := gr.Points()[0] * gr.Points()[1] * gr.Points()[2]
	mem := ntot / 1e6

	data := struct {
		Nx, Ny, Nz, Ntot, Mem int
	}{
		gr.Points()[0],
		gr.Points()[1],
		gr.Points()[2],
		ntot,
		mem,
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, data); err != nil {
		return "Could not generate grid info: " + err.Error()
	}

	return b.String()
}

const gridInfoTemplate = `
************** Grid ***************
Grid (int8)  -> {{.Nx}} x {{.Ny}} x {{.Nz}} 
Total points -> {{.Ntot}}
Memory       -> {{.Mem}} MB
-----------------------------------

`
