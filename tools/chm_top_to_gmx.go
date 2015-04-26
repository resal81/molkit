package main

import (
	"flag"
	"fmt"
	"log"
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
	s := fmt.Sprintf("")

	return s
}

func OneFourTypeToString(at *blocks.PairType) string {
	s := fmt.Sprintf("")

	return s
}

func NonBondedTypeToString(at *blocks.PairType) string {
	s := fmt.Sprintf("")

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
{{ range ot := .OneFourTypes}}{{ OneFourTypeToString $ot}}{{ end }}

[ nonbond_params ]
{{ range nt := .NonBondedType}}{{ NonBondedTypeToString $nt}}{{ end }}

`
	funcMap := template.FuncMap{
		"BondTypeToString":     BondTypeToString,
		"AngleTypeToString":    AngleTypeToString,
		"DihedralTypeToString": DihedralTypeToString,
		"ImproperTypeToString": ImproperTypeToString,
		"CMapTypeToString":     CMapTypeToString,
	}

	tmpl, err := template.New("bonding").Funcs(funcMap).Parse(str)
	if err != nil {
		return "", err
	}

	type ffSum struct {
		BondTypes     []*blocks.BondType
		AngleTypes    []*blocks.AngleType
		DihedralTypes []*blocks.DihedralType
		ImproperTypes []*blocks.ImproperType
		CMapTypes     []*blocks.CMapType
	}

	convFF := ffSum{}
	for _, bt1 := range ff.BondTypes() {
		bt2, err := bt1.ConvertTo(blocks.BT_TYPE_GMX_1)
		if err != nil {
			return "", err
		}
		convFF.BondTypes = append(convFF.BondTypes, bt2)
	}

	for _, at1 := range ff.AngleTypes() {
		at2, err := at1.ConvertTo(blocks.NT_TYPE_GMX_5)
		if err != nil {
			return "", err
		}
		convFF.AngleTypes = append(convFF.AngleTypes, at2)
	}

	for _, dt1 := range ff.DihedralTypes() {
		dt2, err := dt1.ConvertTo(blocks.DT_TYPE_GMX_9)
		if err != nil {
			return "", err
		}
		convFF.DihedralTypes = append(convFF.DihedralTypes, dt2)
	}

	for _, it1 := range ff.ImproperTypes() {
		it2, err := it1.ConvertTo(blocks.IT_TYPE_GMX_2)
		if err != nil {
			return "", err
		}
		convFF.ImproperTypes = append(convFF.ImproperTypes, it2)
	}

	for _, ct1 := range ff.CMapTypes() {
		ct2, err := ct1.ConvertTo(blocks.CT_TYPE_GMX_1)
		if err != nil {
			return "", err
		}
		convFF.CMapTypes = append(convFF.CMapTypes, ct2)
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
			log.Println(rtp)
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

		// bonding
		bonding, err := ConvertCHMParToGMX(ff)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(bonding)

	}

}
