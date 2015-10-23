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

func (gm *Geom) MinMax() ([3]float64, [3]float64) {
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

	return [3]float64{minx, miny, minz}, [3]float64{maxx, maxy, maxz}
}

func (gm *Geom) Info() string {
	tmpl := template.Must(template.New("info").Parse(infoTemplate))

	min, max := gm.MinMax()
	dim := [3]float64{
		max[0] - min[0],
		max[1] - min[1],
		max[2] - min[2],
	}

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
		len(gm.Atoms()[0].Coords()),
		gm.CurrentFrame(),
		gm.CoG(),
		min,
		max,
		dim,
	}

	var b bytes.Buffer
	err := tmpl.Execute(&b, data)
	if err != nil {
		return "Could not generate info: " + err.Error()
	}

	return b.String()
}

const infoTemplate = `
************** Geometry Group ***************
{{.NAtoms}} atoms
{{.NFrames}} frames (current frame: {{.CurrFrame}})

CoG: {{.CoG | printf "%6.1f"}}
Min: {{.Min | printf "%6.1f"}}
Max: {{.Max | printf "%6.1f"}}
Dim: {{.Dim | printf "%6.1f"}}

*********************************************
`
