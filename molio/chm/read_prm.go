package chm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/resal81/molkit/blocks"
)

/**********************************************************
* ReadPRMFile
**********************************************************/

// Parses one or multiple CHARMM prm files.
func ReadPRMFiles(fnames ...string) (*blocks.ForceField, error) {

	ff := blocks.NewForceField(blocks.FF_TYPE_CHM)

	for _, fname := range fnames {
		file, err := os.Open(fname)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		err = readprm(file, ff)
		if err != nil {
			return nil, err
		}
	}

	return ff, nil

}

/**********************************************************
* ReadPRMString
**********************************************************/

// Parses a prm string (e.g. contents of a file).
func ReadPRMString(s string) (*blocks.ForceField, error) {
	frc := blocks.NewForceField(blocks.FF_TYPE_CHM)
	reader := strings.NewReader(s)
	err := readprm(reader, frc)
	return frc, err
}

/**********************************************************
* readprm
**********************************************************/

type prmLevel int64

func readprm(reader io.Reader, frc *blocks.ForceField) error {

	const (
		L_MASS prmLevel = 1 << iota
		L_BONDS
		L_ANGLES
		L_DIHEDRALS
		L_IMPROPERS
		L_NONBONDED
		L_NBFIX
		L_CMAP
		L_IGNORE
	)

	var lvl prmLevel
	massDB := map[string]float64{}

	cmap_header := ""                       // first line in a cmap block; has 8 atomtypes followed by a number
	var cmap_str_vals []string = []string{} // the numerical values for the cmap, in string format

	// read file
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		// clean out comments and flanking whitespace
		line = cleanPRMLine(line)
		if line == "" {
			continue
		}

		// check for END keyword
		if strings.ToUpper(line) == "END" {
			break
		}

		// all lines must be longer than 4
		if len(line) < 4 {
			panic("line length less that 5 " + line)
		}

		switch strings.ToUpper(line[:4]) {
		case "ATOM":
			lvl = L_MASS
			continue
		case "BOND":
			lvl = L_BONDS
			continue
		case "ANGL":
			lvl = L_ANGLES
			continue
		case "DIHE":
			lvl = L_DIHEDRALS
			continue
		case "IMPR":
			lvl = L_IMPROPERS
			continue
		case "CMAP":
			lvl = L_CMAP
			continue
		case "NONB":
			lvl = L_NONBONDED
			continue
		case "NBFI":
			lvl = L_NBFIX
			continue
		case "CUTN":
			lvl = L_IGNORE
			continue
		case "HBON":
			lvl = L_IGNORE
			continue
		}

		switch lvl {
		case L_MASS:
			name, mass, err := prmParseMassLine(line)
			if err != nil {
				return fmt.Errorf("in line: {'%s'} - reason: {'%s'}", line, err)
			}
			massDB[name] = mass

		case L_BONDS:
			bt, err := prmParseBondType(line)
			if err != nil {
				return fmt.Errorf("in line: {'%s'} - reason: {'%s'}", line, err)
			}
			frc.AddBondType(bt)

		case L_ANGLES:
			at, err := prmParseAngleType(line)
			if err != nil {
				return fmt.Errorf("in line: {'%s'} - reason: {'%s'}", line, err)
			}
			frc.AddAngleType(at)

		case L_DIHEDRALS:
			dh, err := prmParseDihedralType(line)
			if err != nil {
				return fmt.Errorf("in line: {'%s'} - reason: {'%s'}", line, err)
			}
			frc.AddDihedralType(dh)

		case L_IMPROPERS:
			im, err := prmParseImproperType(line)
			if err != nil {
				return fmt.Errorf("in line: {'%s'} - reason: {'%s'}", line, err)
			}
			frc.AddImproperType(im)

		case L_CMAP:
			if cmap_header == "" {
				// this is the header line
				cmap_header = line
				cmap_str_vals = []string{}
			} else {
				cmap_str_vals = append(cmap_str_vals, strings.Fields(line)...)
				if len(cmap_str_vals) == 24*24 {
					cm, err := prmParseCMapType(24, 24, strings.Fields(cmap_header), cmap_str_vals)
					if err != nil {
						return fmt.Errorf("in line: {'%s'} - reason: {'%s'}", line, err)
					}
					frc.AddCMapType(cm)
					cmap_header = ""
				}
			}

		case L_NONBONDED:
			at, err := prmParseNonBonded(line)
			if err != nil {
				return fmt.Errorf("in line: {'%s'} - reason: {'%s'}", line, err)
			}
			if m, ok := massDB[at.Label()]; ok {
				at.SetMass(m)
				frc.AddAtomType(at)
			} else {
				return fmt.Errorf("could not find mass for atom type: %s", at.Label)
			}

		case L_NBFIX:
			nb, err := prmParseNBFixType(line)
			if err != nil {
				return fmt.Errorf("in line: {'%s'} - reason: {'%s'}", line, err)
			}
			frc.AddNonBondedType(nb)

		case L_IGNORE:
			continue
		}

	}

	return nil
}

/**********************************************************
* Helpers
**********************************************************/

// removes comments plus leading and tailing spaces
func cleanPRMLine(s string) string {
	i := strings.Index(s, "!")
	if i != -1 {
		s = s[:i]
	}

	j := strings.Index(s, "*")
	if j != -1 {
		s = s[:j]
	}

	s = strings.TrimSpace(s)
	return s
}

//
func prmCheckLineFields(s string, exp_lens []int, header string) (nfields int, err error) {

	fields := strings.Fields(s)
	for _, l := range exp_lens {
		if len(fields) == l {
			return l, nil
		}
	}

	return 0, fmt.Errorf("%s : bad length", header)
}

/**********************************************************
* Line parsers
**********************************************************/

//
func prmParseMassLine(s string) (string, float64, error) {
	/*
		ATOMS
		MASS    41 H      1.00800 ! polar H
		MASS    42 HC     1.00800 ! N-ter H
	*/

	if strings.HasPrefix(s, "MASS") {
		fields := strings.Fields(s)
		if len(fields) != 4 {
			return "", 0, fmt.Errorf("bad length in MASS line")
		}
		name := fields[2]
		m, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return "", 0, err
		}

		return name, m, nil
	} else {
		return "", 0, fmt.Errorf("bad length in MASS line")
	}
}

//
func prmParseBondType(s string) (*blocks.BondType, error) {
	/*

		BONDS
		!V(bond) = Kb(b - b0)**2
		!Kb: kcal/mole/A**2
		!b0: A
		!atom type Kb          b0
		NH2   CT1   240.00      1.455  ! From LSN NH2-CT2
		CA   CA    305.000     1.3750 ! ALLOW   ARO
	*/

	// atype1 atype2  Kb  b0
	if _, err := prmCheckLineFields(s, []int{4}, "BONDS"); err != nil {
		return nil, err
	}

	var at1, at2 string
	var kb, b0 float64

	n, err := fmt.Sscanf(s, "%s %s %f %f", &at1, &at2, &kb, &b0)
	if n != 4 || err != nil {
		return nil, fmt.Errorf("error paring BONDS line")
	}

	bt := blocks.NewBondType(at1, at2, blocks.BT_TYPE_CHM_1)
	bt.SetHarmonicDistance(b0)
	bt.SetHarmonicConstant(kb)

	return bt, nil

}

//
func prmParseAngleType(s string) (*blocks.AngleType, error) {
	/*

		ANGLES
		!V(angle) = Ktheta(Theta - Theta0)**2
		!V(Urey-Bradley) = Kub(S - S0)**2
		!Ktheta: kcal/mole/rad**2
		!Theta0: degrees
		!Kub: kcal/mole/A**2 (Urey-Bradley)
		!S0: A
		!atom types     Ktheta    Theta0   Kub     S0
		NH2  CT1  CT2   67.700    110.00                  ! From LSN NH2-CT2-CT2
		NH2  CT1  HB    38.000    109.50   50.00   2.1400 ! From LSN NH2-CT2-HA

	*/

	// atyp1 atype2 atype3     Ktheta    Theta0   Kub     S0
	nfields, err := prmCheckLineFields(s, []int{5, 7}, "ANGLES")
	if err != nil {
		return nil, err
	}

	var at1, at2, at3 string
	var kt, theta, kub, r13 float64

	var angt *blocks.AngleType

	switch nfields {
	case 5:
		n, err := fmt.Sscanf(s, "%s %s %s %f %f", &at1, &at2, &at3, &kt, &theta)
		if n != 5 || err != nil {
			return nil, fmt.Errorf("could not parse angletype - 5")
		}
		angt = blocks.NewAngleType(at1, at2, at3, blocks.NT_TYPE_CHM_1)
		angt.SetTheta(theta)
		angt.SetThetaConstant(kt)
		angt.SetR13(0)
		angt.SetUBConstant(0)
	case 7:
		n, err := fmt.Sscanf(s, "%s %s %s %f %f %f %f", &at1, &at2, &at3, &kt, &theta, &kub, &r13)
		if n != 7 || err != nil {
			return nil, fmt.Errorf("could not parse angletype - 7")
		}
		angt = blocks.NewAngleType(at1, at2, at3, blocks.NT_TYPE_CHM_1)
		angt.SetTheta(theta)
		angt.SetThetaConstant(kt)
		angt.SetR13(r13)
		angt.SetUBConstant(kub)
	}

	return angt, nil
}

//
func prmParseDihedralType(s string) (*blocks.DihedralType, error) {
	/*
		DIHEDRALS
		!V(dihedral) = Kchi(1 + cos(n(chi) - delta))
		!Kchi: kcal/mole
		!n: multiplicity
		!delta: degrees
		!atom types             Kchi    n   delta
		!Neutral N terminus
		H    NH2  CT1  CT3      0.110   3     0.00  ! From LSN HC-NH2-CT2-CT2
		C    CT1  NH1  C        0.2000  1   180.00  ! ALLOW PEP
	*/

	// atype1 atype2 atype3  atype4 Kchi    n   delta
	_, err := prmCheckLineFields(s, []int{7}, "DIHEDRALS")
	if err != nil {
		return nil, err
	}

	var at1, at2, at3, at4 string
	var mult int
	var kphi, phi float64

	n, err := fmt.Sscanf(s, "%s %s %s %s %f %d %f", &at1, &at2, &at3, &at4, &kphi, &mult, &phi)
	if n != 7 || err != nil {
		return nil, fmt.Errorf("could not parse dihedraltype")
	}

	dh := blocks.NewDihedralType(at1, at2, at3, at4, blocks.DT_TYPE_CHM_1)

	dh.SetPhi(phi)
	dh.SetPhiConstant(kphi)
	dh.SetMultiplicity(mult)

	return dh, nil
}

//
func prmParseImproperType(s string) (*blocks.ImproperType, error) {
	/*
		IMPROPER
		!V(improper) = Kpsi(psi - psi0)**2
		!Kpsi: kcal/mole/rad**2
		!psi0: degrees
		!note that the second column of numbers (0) is ignored
		!atom types           Kpsi                   psi0
		HE2  HE2  CE2  CE2     3.0            0      0.00   !
		HR1  NR1  NR2  CPH2    0.5000         0      0.0000 ! ALLOW ARO
	*/

	// atype1 atype2 atype3  atype4  Kpsi ign psi0
	_, err := prmCheckLineFields(s, []int{7}, "IMPROPER")
	if err != nil {
		return nil, err
	}

	var at1, at2, at3, at4, tmp string
	var kpsi, psi float64

	n, err := fmt.Sscanf(s, "%s %s %s %s %f %s %f", &at1, &at2, &at3, &at4, &kpsi, &tmp, &psi)
	if n != 7 || err != nil {
		return nil, fmt.Errorf("could not parse dihedraltype")
	}

	it := blocks.NewImproperType(at1, at2, at3, at4, blocks.IT_TYPE_CHM_1)
	it.SetPsi(psi)
	it.SetPsiConstant(kpsi)

	return it, nil
}

//
func prmParseNonBonded(s string) (*blocks.AtomType, error) {

	/*
		NONBONDED nbxmod  5 atom cdiel fshift vatom vdistance vfswitch -
		cutnb 14.0 ctofnb 12.0 ctonnb 10.0 eps 1.0 e14fac 1.0 wmin 1.5
		!V(Lennard-Jones) = Eps,i,j[(Rmin,i,j/ri,j)**12 - 2(Rmin,i,j/ri,j)**6]
		!epsilon: kcal/mole, Eps,i,j = sqrt(eps,i * eps,j)
		!Rmin/2: A, Rmin,i,j = Rmin/2,i + Rmin/2,j
		!atom  ignored    epsilon      Rmin/2   ignored   eps,1-4       Rmin/2,1-4
		CE2    0.000000  -0.064000     2.080000 !
		CP1    0.000000  -0.020000     2.275000   0.000000  -0.010000     1.900000 ! ALLOW   ALI

	*/

	// atom  ignored    epsilon      Rmin/2   ignored   eps,1-4       Rmin/2,1-4

	nfields, err := prmCheckLineFields(s, []int{4, 7}, "NONBONDED")
	if err != nil {
		return nil, err
	}

	var at1, tmp1, tmp2 string
	var dist, en, dist14, en14 float64

	if nfields == 4 {
		n, err := fmt.Sscanf(s, "%s %s %f %f", &at1, &tmp1, &en, &dist)
		if n != 4 {
			return nil, fmt.Errorf("NONBONDED line doesn't have 4 fields")
		}
		if err != nil {
			return nil, err
		}
	} else if nfields == 7 {
		n, err := fmt.Sscanf(s, "%s %s %f %f %s %f %f", &at1, &tmp1, &en, &dist, &tmp2, &en14, &dist14)
		if n != 7 || err != nil {
			return nil, fmt.Errorf("NONBONDED line doesn't have 7 fields")
		}
		if err != nil {
			return nil, err
		}
	}

	atm := blocks.NewAtomType(at1, blocks.AT_TYPE_CHM_1)
	atm.SetLJDistance(dist)
	atm.SetLJEnergy(en)

	if nfields == 7 {
		atm.SetLJDistance14(dist14)
		atm.SetLJEnergy14(en14)
	}

	return atm, nil

}

func prmParseNBFixType(s string) (*blocks.PairType, error) {
	/*
		NBFIX
		SOD    OC       -0.075020   3.190 ! For prot carboxylate groups
		SOD    OCL      -0.075020   3.190 ! For lipid carboxylate groups
	*/
	// atype1  atype2   Emin    Rmin
	if _, err := prmCheckLineFields(s, []int{4}, "NBFIX"); err != nil {
		return nil, err
	}

	var at1, at2 string
	var dist, en float64

	n, err := fmt.Sscanf(s, "%s %s %f %f", &at1, &at2, &en, &dist)
	if n != 4 || err != nil {
		return nil, fmt.Errorf("could not parse nbfix")
	}

	nb := blocks.NewPairType(at1, at2, blocks.PT_TYPE_CHM_1)
	nb.SetLJDistance(dist)
	nb.SetLJEnergy(en)

	return nb, nil
}

//
func prmParseCMapType(nx, ny int, atypes []string, vals []string) (*blocks.CMapType, error) {
	/*
		CMAP
		! 2D grid correction data.  The following surfaces are the correction
		! to the CHARMM22 phi, psi alanine, proline and glycine dipeptide surfaces.
		! Use of CMAP requires generation with the topology file containing the
		! CMAP specifications along with version 31 or later of CHARMM.  Note that
		! use of "skip CMAP" yields the charmm22 energy surfaces.
		! alanine map
		C    NH1  CT1  C    NH1  CT1  C    NH1   24

		!-180
		0.126790 0.768700 0.971260 1.250970 2.121010
		2.695430 2.064440 1.764790 0.755870 -0.713470
		0.976130 -2.475520 -5.455650 -5.096450 -5.305850
		-3.975630 -3.088580 -2.784200 -2.677120 -2.646060
		-2.335350 -2.010440 -1.608040 -0.482250

		!-165
		-0.802290 1.377090 1.577020 1.872290 2.398990
		2.461630 2.333840 1.904070 1.061460 0.518400
		-0.116320 -3.575440 -5.284480 -5.160310 -4.196010
		-3.276210 -2.715340 -1.806200 -1.101780 -1.210320
		-1.008810 -0.637100 -1.603360 -1.776870
	*/

	//[C NH1 CT1 C NH1 CT1 C N 24]

	if nx*ny != len(vals) {
		return nil, fmt.Errorf("nx and ny are %d and %d, but len(vals) is %d", nx, ny, len(vals))
	}

	vals_f := make([]float64, len(vals))
	for i, v := range vals {
		fv, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}

		vals_f[i] = fv
	}
	cm := blocks.NewCMapType(
		atypes[0],
		atypes[1],
		atypes[2],
		atypes[3],
		atypes[4],
		atypes[5],
		atypes[6],
		atypes[7],
		blocks.CT_TYPE_CHM_1,
	)
	cm.SetValues(vals_f)

	return cm, nil
}
