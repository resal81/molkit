package main

import (
	"flag"
	"fmt"

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

func main() {
	var topFile string
	var resName string

	flag.StringVar(&topFile, "top", "", "CHARMM topology file")
	flag.StringVar(&resName, "res", "", "Target RESI name")
	flag.Parse()

	if topFile == "" {
		fmt.Println("Please provide a topology file. See -h for more info.")
		return
	}

	if resName == "" {
		fmt.Println("Please provide a residue name. See -h for more info.")
		return
	}

	ff, err := chm.ReadTOPFiles(topFile)
	if err != nil {
		fmt.Println("%s", err)
		return
	}

	res := ff.Fragment(blocks.HashKey(resName))
	if res == nil {
		fmt.Printf("residue not found => %s\n", resName)
		return
	}

	out, err := ConvertCHMFragmentToGMX(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(out)
}
