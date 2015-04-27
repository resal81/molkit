package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strings"

	"bytes"
	"github.com/resal81/molkit/blocks"
	"github.com/resal81/molkit/molio/chm"
	"text/template"
)

func ConvertCHMFragmentToGMX(frag *blocks.Fragment) (string, error) {
	str := `

[ {{.Name}} ]
  [ atoms ]
{{ range $index, $atom := .Atoms }}{{ $atom.Name | printf "%8s" }}{{ $atom.Type.Label | printf "%8s" }}{{ $atom.Type.PartialCharge | printf "%8.3f" }}{{ $index | printf "%3d\n" }}{{ end }}

  [ bonds ]
{{ range $bond := .Bonds }}{{ $bond.Atom1.Name | printf "%8s"}}{{ $bond.Atom2.Name | printf "%8s\n"}}{{end}}

  [ impropers ]
{{ range $imp := .Impropers }}{{ $imp.Atom1.Name | printf "%8s"}}{{ $imp.Atom2.Name | printf "%8s"}}{{ $imp.Atom3.Name | printf "%8s"}}{{ $imp.Atom4.Name | printf "%8s\n"}}{{end}}

  [ cmap ]
{{ range $cm := .CMaps }}{{ $cm.Atom1.Name | printf "%8s"}}{{ $cm.Atom2.Name | printf "%8s"}}{{ $cm.Atom3.Name | printf "%8s"}}{{ $cm.Atom4.Name | printf "%8s"}}{{ $cm.Atom8.Name | printf "%8s\n"}}{{end}}

`
	tmpl, err := template.New("rtp").Parse(str)
	if err != nil {
		return "", err
	}

	//
	type resSum struct {
		Name      string
		Atoms     []*blocks.Atom
		Bonds     []*blocks.Bond
		Impropers []*blocks.Improper
		CMaps     []*blocks.CMap
	}

	res := resSum{}
	res.Name = frag.Name()
	res.Atoms = frag.Atoms()
	res.Bonds = frag.Bonds()
	res.Impropers = frag.Impropers()
	res.CMaps = frag.CMaps()

	if frag.HasLinkerNext() {
		ln := frag.LinkerNext()
		if v := ln.Bond(); v != nil {
			res.Bonds = append(res.Bonds, v)
		}
		if v := ln.Improper(); v != nil {
			res.Impropers = append(res.Impropers, v)
		}
	}

	if frag.HasLinkerPrev() {
		ln := frag.LinkerPrev()
		if v := ln.Bond(); v != nil {
			res.Bonds = append(res.Bonds, v)
		}
		if v := ln.Improper(); v != nil {
			res.Impropers = append(res.Impropers, v)
		}
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, res)
	if err != nil {
		return "", err
	}

	return b.String(), nil

}

func BondTypeToString(bt *blocks.BondType) string {
	s := fmt.Sprintf("%8s %8s %5d %12.8f %12.2f\n", bt.AType1(), bt.AType2(), 1, bt.HarmonicDistance(), bt.HarmonicConstant())
	return s
}

func AngleTypeToString(at *blocks.AngleType) string {
	s := fmt.Sprintf("%8s %8s %8s %5d %12.6f %12.6f %12.8f %12.2f\n", at.AType1(), at.AType2(), at.AType3(), 5, at.Theta(), at.ThetaConstant(), at.R13(), at.UBConstant())
	return s
}

func DihedralTypeToString(dt *blocks.DihedralType) string {
	s := fmt.Sprintf("%8s %8s %8s %8s %5d %12.6f %12.6f %5d\n", dt.AType1(), dt.AType2(), dt.AType3(), dt.AType4(), 9, dt.Phi(), dt.PhiConstant(), dt.Multiplicity())
	return s
}

func ImproperTypeToString(it *blocks.ImproperType) string {
	s := fmt.Sprintf("%8s %8s %8s %8s %5d %12.6f %12.6f\n", it.AType1(), it.AType2(), it.AType3(), it.AType4(), 2, it.Psi(), it.PsiConstant())
	return s
}

func CMapTypeToString(ct *blocks.CMapType) string {
	s := fmt.Sprintf("%s %s %s %s %s %d %d %d", ct.AType1(), ct.AType2(), ct.AType3(), ct.AType4(), ct.AType8(), 1, 24, 24)

	for i, v := range ct.Values() {
		if i%10 == 0 {
			s += "\\\n"
		}
		s += fmt.Sprintf(" %10.8f", v)
	}
	s += "\n\n"
	return s
}

func AtomTypeToString(at *blocks.AtomType) string {
	// ;type atnum         mass   charge ptype           sigma  epsilon
	//  ALG1    13    26.981540    0.000  A  0.356359487256  2.71960

	s := fmt.Sprintf("%8s %3d %12.6f %6.3f %3s %14.12f %12.9f\n",
		at.Label(), at.Protons(), at.Mass(), at.PartialCharge(), "A", at.LJDistance(), at.LJEnergy())

	return s
}

func OneFourTypeToString(pt *blocks.PairType) string {
	// [ pairtypes ]
	// CP1     C  1  0.347450500075  0.138767581228

	s := fmt.Sprintf("%8s %8s %2d %14.12f %14.12f\n", pt.AType1(), pt.AType2(), 1, pt.LJDistance14(), pt.LJEnergy14())
	return s
}

func NonBondedTypeToString(nt *blocks.PairType) string {
	// ; NBFIX  rmin=<charmm_rmin>/2^(1/6), eps=4.184*<charmm_eps>
	//;name   type1  type2  1  sigma   epsilon
	//[ nonbond_params ]
	//SOD   CLA  1  0.332394311738  0.350933000000

	s := fmt.Sprintf("%8s %8s %2d %14.12f %14.12f\n", nt.AType1(), nt.AType2(), 1, nt.LJDistance(), nt.LJEnergy())
	return s
}

func ConvertCHMParToGMX(ff *blocks.ForceField) (string, error) {
	str := `
[ bondtypes ]
;      i        j  func           b0           kb
{{ range $bt := .BondTypes}}{{ BondTypeToString $bt}}{{ end }}

[ angletypes ]
;      i        j        k  func       theta0       ktheta          ub0          kub
{{ range $at := .AngleTypes}}{{ AngleTypeToString $at }}{{ end }}

[ dihedraltypes ]
;      i        j        k        l  func         phi0         kphi  mult
{{ range $dt := .DihedralTypes}}{{ DihedralTypeToString $dt }}{{ end }}

[ dihedraltypes ]
; 'improper' dihedrals 
;      i        j        k        l  func         phi0         kphi
{{ range $it := .ImproperTypes}}{{ ImproperTypeToString $it}}{{ end }} 

[ cmaptypes ]
{{ range $ct := .CMapTypes}}{{ CMapTypeToString $ct}}{{ end }}

[ atomtypes ]
{{ range $at := .AtomTypes}}{{ AtomTypeToString $at}}{{ end }}

[ pairtypes ]
{{ range $ot := .OneFourTypes}}{{ OneFourTypeToString $ot}}{{ end }}

[ nonbond_params ]
{{ range $nt := .NonBondedTypes}}{{ NonBondedTypeToString $nt}}{{ end }}

`
	funcMap := template.FuncMap{
		"BondTypeToString":      BondTypeToString,
		"AngleTypeToString":     AngleTypeToString,
		"DihedralTypeToString":  DihedralTypeToString,
		"ImproperTypeToString":  ImproperTypeToString,
		"CMapTypeToString":      CMapTypeToString,
		"AtomTypeToString":      AtomTypeToString,
		"OneFourTypeToString":   OneFourTypeToString,
		"NonBondedTypeToString": NonBondedTypeToString,
	}

	tmpl, err := template.New("bonding").Funcs(funcMap).Parse(str)
	if err != nil {
		return "", err
	}

	type ffSum struct {
		BondTypes      []*blocks.BondType
		AngleTypes     []*blocks.AngleType
		DihedralTypes  []*blocks.DihedralType
		ImproperTypes  []*blocks.ImproperType
		CMapTypes      []*blocks.CMapType
		OneFourTypes   []*blocks.PairType
		NonBondedTypes []*blocks.PairType
		AtomTypes      []*blocks.AtomType
	}

	convFF := ffSum{}

	// Bond
	for _, bt1 := range ff.BondTypes() {
		bt2, err := bt1.ConvertTo(blocks.BT_TYPE_GMX_1)
		if err != nil {
			return "", err
		}
		convFF.BondTypes = append(convFF.BondTypes, bt2)
	}

	// Angle
	for _, at1 := range ff.AngleTypes() {
		at2, err := at1.ConvertTo(blocks.NT_TYPE_GMX_5)
		if err != nil {
			return "", err
		}
		convFF.AngleTypes = append(convFF.AngleTypes, at2)
	}

	// Dihedral
	for _, dt1 := range ff.DihedralTypes() {
		dt2, err := dt1.ConvertTo(blocks.DT_TYPE_GMX_9)
		if err != nil {
			return "", err
		}
		convFF.DihedralTypes = append(convFF.DihedralTypes, dt2)
	}

	// Improper
	for _, it1 := range ff.ImproperTypes() {
		it2, err := it1.ConvertTo(blocks.IT_TYPE_GMX_2)
		if err != nil {
			return "", err
		}
		convFF.ImproperTypes = append(convFF.ImproperTypes, it2)
	}

	// CMap
	for _, ct1 := range ff.CMapTypes() {
		ct2, err := ct1.ConvertTo(blocks.CT_TYPE_GMX_1)
		if err != nil {
			return "", err
		}
		convFF.CMapTypes = append(convFF.CMapTypes, ct2)
	}

	// Atom
	for _, at1 := range ff.AtomTypes() {
		at2, err := at1.ConvertTo(blocks.AT_TYPE_GMX_1)
		if err != nil {
			return "", err
		}
		convFF.AtomTypes = append(convFF.AtomTypes, at2)
	}

	// NonBonded
	for _, nt1 := range ff.NonBondedTypes() {
		nt2, err := nt1.ConvertTo(blocks.PT_TYPE_GMX_1)
		if err != nil {
			return "", err
		}
		convFF.NonBondedTypes = append(convFF.NonBondedTypes, nt2)
	}

	// OneFour
	ats := []*blocks.AtomType{}
	for _, at := range ff.AtomTypes() {
		ats = append(ats, at)
	}

	mix_e := func(x, y float64) float64 {
		return math.Pow(x*y, 0.5)
	}

	mix_d := func(x, y float64) float64 {
		return (x + y) / 2.0
	}

	for i, a1 := range ats {
		for _, a2 := range ats[i:] {
			np1 := blocks.NewPairType(a1.Label(), a2.Label(), blocks.PT_TYPE_CHM_1)

			var e, d float64

			switch {

			case a1.HasLJEnergy14Set() && a1.HasLJDistance14Set():

				switch {

				case a2.HasLJEnergy14Set() && a2.HasLJDistance14Set():
					// both have lj14 set

					e = mix_e(a1.LJEnergy14(), a2.LJEnergy14())
					d = mix_d(a1.LJDistance14(), a2.LJDistance14())

				default:
					// only a1 has lj14 set
					e = mix_e(a1.LJEnergy14(), a2.LJEnergy())
					d = mix_d(a1.LJDistance14(), a2.LJDistance())

				}

			case a2.HasLJEnergy14Set() && a2.HasLJDistance14Set():
				// only a2 has lj14 set
				e = mix_e(a1.LJEnergy(), a2.LJEnergy14())
				d = mix_d(a1.LJDistance(), a2.LJDistance14())

			default:
				// neither a1 or a2 have lj14 set
				// don't generate by default
				continue
			}

			np1.SetLJEnergy14(e)
			np1.SetLJDistance14(d)

			np2, err := np1.ConvertTo(blocks.PT_TYPE_GMX_1)
			if err != nil {
				return "", err
			}
			convFF.OneFourTypes = append(convFF.OneFourTypes, np2)
		}
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, convFF)
	if err != nil {
		return "", err
	}

	return b.String(), nil

}

func main() {
	// for multiple files, separate them by ,

	var topFile string
	var parFile string
	//var resName string

	flag.StringVar(&topFile, "top", "", "CHARMM topology file. Separate multiple by ','.")
	flag.StringVar(&parFile, "par", "", "CHARMM parameter file. Separate multiple by ','.")
	//flag.StringVar(&resName, "res", "", "Target RESI name")
	flag.Parse()

	//if resName == "" {
	//fmt.Println("Please provide a residue name. See -h for more info.")
	//return
	//}

	if topFile != "" {

		topFiles := strings.Split(topFile, ",")

		top, err := chm.ReadTOPFiles(topFiles...)
		if err != nil {
			log.Fatalln(err)
		}

		for _, v := range top.Fragments() {
			rtp, err := ConvertCHMFragmentToGMX(v)
			if err != nil {
				log.Fatalln(err)
			}
			_ = rtp
			//fmt.Println(rtp)
		}

		//residue
		//res := top.Fragment(blocks.HashKey(resName))
		//if res == nil {
		//fmt.Printf("residue not found => %s\n", resName)
		//return
		//}
	}

	if parFile != "" {

		parFiles := strings.Split(parFile, ",")

		// prm files
		ff, err := chm.ReadPRMFiles(parFiles...)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(len(ff.AtomTypes()))

		// bonding
		bonding, err := ConvertCHMParToGMX(ff)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(bonding)

	}

}
