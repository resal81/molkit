package geom

import (
	"bytes"
	"github.com/resal81/molkit/blocks/atom"
	"text/template"
)

type Geom struct {
	atoms     []*atom.Atom
	currFrame int
}

func NewGeom(atoms []*atom.Atom) *Geom {
	return &Geom{
		atoms: atoms,
	}
}

func (gm *Geom) Atoms() []*atom.Atom {
	return gm.atoms
}

func (gm *Geom) SetCurrentFrame(frame int) {
	gm.currFrame = frame
}

func (gm *Geom) CurrentFrame() int {
	return gm.currFrame
}

func (gm *Geom) NFrames() int {
	if len(gm.Atoms()) > 0 {
		return len(gm.Atoms()[0].Coords())
	}
	return 0
}

// CoG returns center of geometry
func (gm *Geom) CoG() [3]float64 {
	var cx, cy, cz float64

	for _, at := range gm.Atoms() {
		coord := at.CoordAtFrame(gm.currFrame)
		cx += coord[0]
		cy += coord[1]
		cz += coord[2]
	}

	n := float64(len(gm.Atoms()))
	return [3]float64{cx / n, cy / n, cz / n}
}

// MinMax returns the minimum and maximum corners of the bounding box
func (gm *Geom) MinMaxDim() ([3]float64, [3]float64, [3]float64) {
	var minx, miny, minz float64
	var maxx, maxy, maxz float64

	for _, at := range gm.Atoms() {
		coord := at.CoordAtFrame(gm.currFrame)

		if coord[0] < minx {
			minx = coord[0]
		} else if coord[0] > maxx {
			maxx = coord[0]
		}

		if coord[1] < miny {
			miny = coord[1]
		} else if coord[1] > maxy {
			maxy = coord[1]
		}

		if coord[2] < minz {
			minz = coord[2]
		} else if coord[2] > maxz {
			maxz = coord[2]
		}
	}

	min := [3]float64{minx, miny, minz}
	max := [3]float64{maxx, maxy, maxz}
	dim := [3]float64{
		maxx - minx,
		maxy - miny,
		maxz - minz,
	}

	return min, max, dim
}

// CenterCoG centers the frame at origin
func (gm *Geom) CenterCoG() {
	i := gm.CurrentFrame()
	cog := gm.CoG()
	for _, at := range gm.Atoms() {
		crd := at.CoordAtFrame(i)
		crd = []float64{crd[0] - cog[0], crd[1] - cog[1], crd[2] - cog[2]}
		at.SetCoordAtFrame(i, crd)
	}

}

//
func (gm *Geom) CenterCoGForAllFrames() {
	for i := 0; i < gm.NFrames(); i++ {
		gm.SetCurrentFrame(i)
		gm.CenterCoG()
	}
}

//
func (gm *Geom) FindCavities() {
}

// Info returns some info about this geometry group
func (gm *Geom) Info() string {
	tmpl := template.Must(template.New("info").Parse(geomInfoTemplate))

	min, max, dim := gm.MinMaxDim()
	data := struct {
		NAtoms    int
		NFrames   int
		CurrFrame int
		CoG       [3]float64
		Min       [3]float64
		Max       [3]float64
		Dim       [3]float64
	}{
		len(gm.Atoms()),
		gm.NFrames(),
		gm.CurrentFrame(),
		gm.CoG(),
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

const geomInfoTemplate = `
************** Geometry Group ***************
{{.NAtoms}} atoms, {{.NFrames}} frames (current : {{.CurrFrame}})

CoG: {{.CoG | printf "%6.1f"}}
Min: {{.Min | printf "%6.1f"}}
Max: {{.Max | printf "%6.1f"}}
Dim: {{.Dim | printf "%6.1f"}}

`
