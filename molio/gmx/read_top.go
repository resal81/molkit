package gmx

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/resal81/molkit/ff"
)

var (
	ErrDefaults  = errors.New("[defaults] could not be parsed")
	ErrAtomTypes = errors.New("[atomtypes] could not be parsed")
)

type topologyLevel int8
type topologyGroup int64

/**********************************************************
* ReadTOPFile
**********************************************************/

func ReadTOPFile(fname string) (*ff.TopSystem, *ff.ForceField, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	return readtop(file)

}

/**********************************************************
* ReadTOPString
**********************************************************/

func ReadTOPString(s string) (*ff.TopSystem, *ff.ForceField, error) {
	reader := strings.NewReader(s)
	return readtop(reader)
}

/**********************************************************
* readtop
**********************************************************/

func readtop(reader io.Reader) (*ff.TopSystem, *ff.ForceField, error) {

	// the three main levels
	const (
		L_PARAMS topologyLevel = 1 << iota
		L_MOLTYPES
		L_SYSTEM
	)

	// groups
	const (
		G_DEFAULTS topologyGroup = 1 << iota
		G_ATOMTYPES
		G_BONDTYPES
		G_PAIRTYPES
		G_ANGLETYPES
		G_DIHEDRALTYPES
		G_CONSTRAINTTYPES
		G_NONBONDPARAMS

		G_MOLECULETYPE
		G_ATOMS
		G_BONDS
		G_PAIRS
		G_ANGLES
		G_DIHEDRALS
		G_EXCLUSIONS
		G_CONSTRAINTS
		G_SETTLES
		G_POSREST
		G_DISTREST
		G_ANGLEREST
		G_DIHEDRALREST
		G_ORIENTREST

		G_SYSTEM
		G_MOLECULES
	)

	var level topologyLevel
	var group topologyGroup

	var curr_topmol *ff.TopPolymer
	var curr_topres *ff.TopFragment
	topsys := ff.NewTopSystem()
	var prev_resname string = "_"
	var prev_resnumb int64 = -1

	forcefield := ff.NewForceField(ff.FF_SOURCE_GROMACS)

	// read file line by line
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		line = cleanLine(line)
		if line == "" {
			continue
		}

		// if encounters a "[ ... ]", find the group
		if strings.HasPrefix(line, "[") {
			line = getGroup(line)

			switch line {
			case "defaults":
				level = L_PARAMS
				group = G_DEFAULTS
			case "atomtypes":
				group = G_ATOMTYPES
			case "nonbond_params":
				group = G_NONBONDPARAMS
			case "bondtypes":
				group = G_BONDTYPES
			case "pairtypes":
				group = G_PAIRTYPES
			case "angletypes":
				group = G_ANGLETYPES
			case "dihedraltypes":
				group = G_DIHEDRALTYPES
			case "moleculetype":
				level = L_MOLTYPES
				group = G_MOLECULETYPE
			case "atoms":
				group = G_ATOMS
			case "bonds":
				group = G_BONDS
			case "pairs":
				group = G_PAIRS
			case "angles":
				group = G_ANGLES
			case "dihedrals":
				group = G_DIHEDRALS
			case "position_restraints":
				group = G_POSREST
			case "dihedral_restraints":
				group = G_DIHEDRALREST
			case "settles":
				group = G_SETTLES
			case "exclusions":
				group = G_EXCLUSIONS
			case "system":
				level = L_SYSTEM
				group = G_SYSTEM
			case "molecules":
				group = G_MOLECULES
			default:
				return nil, nil, fmt.Errorf("unknown group in the top file: %s", line)
			}

			continue
		}

		if level == L_PARAMS {
			switch group {

			case G_DEFAULTS:
				defaults, err := parseDefaults(line)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}
				forcefield.SetGMXDefaults(defaults)

			case G_ATOMTYPES:
				at, err := parseAtomTypes(line)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}
				forcefield.AddAtomType(at)

			case G_NONBONDPARAMS:
				nb, err := parseNonBondedTypes(line)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}
				forcefield.AddNonBondedType(nb)

			case G_PAIRTYPES:
				pt, err := parsePairTypes(line)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}
				forcefield.AddPairType(pt)

			case G_BONDTYPES:
				bt, err := parseBondTypes(line)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}
				forcefield.AddBondType(bt)

			case G_ANGLETYPES:
				ag, err := parseAngleTypes(line)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}
				forcefield.AddAngleType(ag)

			case G_DIHEDRALTYPES:
				dt, err := parseDihedralTypes(line)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}

				if dt.Kind()&ff.FF_DIHEDRAL_TYPE_1 != 0 || dt.Kind()&ff.FF_DIHEDRAL_TYPE_9 != 0 {
					forcefield.AddDihedralType(dt)

				} else if dt.Kind()&ff.FF_DIHEDRAL_TYPE_2 != 0 {
					forcefield.AddImproperType(dt)

				} else {
					return nil, nil, errors.New("dihedral type is not valid")
				}

			}

		}

		if level == L_MOLTYPES {
			switch group {

			case G_MOLECULETYPE:
				pol, err := parseMoleculeTypes(line)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}
				curr_topmol = pol
				topsys.RegisterTopPolymer(curr_topmol)

			case G_ATOMS:
				at, rd, err := parseAtoms(line)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}

				if rd.name != prev_resname || rd.serial != prev_resnumb {
					curr_topres = ff.NewTopFragment()
					curr_topres.SetName(rd.name)
					curr_topres.SetSerial(rd.serial)

					curr_topmol.AddTopFragment(curr_topres)

					prev_resname = rd.name
					prev_resnumb = rd.serial
				}

				curr_topres.AddTopAtom(at)
				curr_topmol.AddTopAtom(at)

			case G_BONDS:
			case G_PAIRS:
			case G_ANGLES:
			case G_DIHEDRALS:
			}
		}

		if level == L_SYSTEM {
			switch group {
			case G_SYSTEM:
				continue
			case G_MOLECULES:
			}
		}

	}

	return topsys, forcefield, nil

}

// removes comments plus leading and tailing spaces
func cleanLine(s string) string {
	i := strings.Index(s, ";")
	if i != -1 {
		s = s[:i]
	}

	s = strings.TrimSpace(s)
	return s
}

// Extract the group inside brackets: [ ... ]
func getGroup(s string) string {
	s = strings.TrimLeft(s, "[")
	s = strings.TrimRight(s, "]")
	s = strings.TrimSpace(s)
	return s
}

// parses [defaults]
func parseDefaults(s string) (*ff.GMXDefaults, error) {
	// ; nbfunc	comb-rule	gen-pairs	fudgeLJ	fudgeQQ

	var nbfunc, combrule int8
	var fudgeLJ, fudgeQQ float32
	var genpairs string

	n, err := fmt.Sscanf(s, "%d %d %s %f %f", &nbfunc, &combrule, &genpairs, &fudgeLJ, &fudgeQQ)
	if err != nil || n != 5 {
		return nil, ErrDefaults
	}

	gd := ff.NewGMXDefaults(nbfunc, combrule, genpairs == "yes", fudgeLJ, fudgeQQ)
	return gd, nil

}

// parses [atomstypes]
func parseAtomTypes(s string) (*ff.AtomType, error) {
	// ; name	at.num	mass	charge	ptype	sigma	epsilon	;	sigma_14	epsilon_14

	var name, pt string
	var prot int8
	var mass, chg, sig, eps, sig14, eps14 float64

	fields := strings.Fields(s)
	switch len(fields) {
	case 7:
		_, err := fmt.Sscanf(s, "%s %d %f %f %s %f %f", &name, &prot, &mass, &chg, &pt, &sig, &eps)
		if err != nil {
			return nil, err
		}
	case 9:
		_, err := fmt.Sscanf(s, "%s %d %f %f %s %f %f %f %f", &name, &prot, &mass, &chg, &pt, &sig, &eps, &sig14, &eps14)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("the number of fields in the [atomtypes] line is not right")
	}

	at := ff.NewAtomType(name)
	at.SetProtons(prot)
	at.SetMass(mass)
	at.SetCharge(chg)
	at.SetSigma(sig)
	at.SetEpsilon(eps)

	if len(fields) == 9 {
		at.SetSigma14(sig14)
		at.SetEpsilon14(eps14)
	}

	return at, nil
}

type resData struct {
	name   string
	serial int64
}

// parses [atoms]
func parseAtoms(s string) (*ff.TopAtom, *resData, error) {
	// ; nr	type	resnr	residu	atom	cgnr	charge	mass

	fields := strings.Fields(s)
	if len(fields) != 8 {
		return nil, nil, errors.New("number of fields in [atoms] in not 8")
	}

	var name, atype, resname string
	var chg, mass float64
	var ser, cgnr, resnumb int64

	n, err := fmt.Sscanf(s, "%d %s %d %s %s %d %f %f", &ser, &atype, &resnumb, &resname, &name, &cgnr, &chg, &mass)
	if n != 8 || err != nil {
		return nil, nil, errors.New("could not parse [atoms] line")
	}

	a := ff.NewTopAtom()
	a.SetName(name)
	a.SetSerial(ser)
	a.SetAtomType(atype)
	a.SetCharge(chg)
	a.SetMass(mass)
	a.SetCGNR(cgnr)

	return a, &resData{resname, resnumb}, nil
}

// parses [nonbonded-params]
func parseNonBondedTypes(s string) (*ff.NonBondedType, error) {
	// ; i	j	func	sigma	epsilon
	var at1, at2 string
	var fn int8
	var sig, eps float64
	fields := strings.Fields(s)

	n, err := fmt.Sscanf(s, "%s %s %d %f %f", &at1, &at2, &fn, &sig, &eps)
	if len(fields) != 5 || n != 5 || err != nil {
		return nil, errors.New("error parsing [nonbonded-params] line")
	}

	switch fn {
	case 1:
		nbt := ff.NewNonBondedType(at1, at2, ff.FF_NON_BONDED_TYPE_1)
		nbt.SetSigma(sig)
		nbt.SetEpsilon(eps)
		return nbt, nil
	default:
		return nil, errors.New("nonbonded function type is not 1")
	}

}

// parses [pairtypes]
func parsePairTypes(s string) (*ff.PairType, error) {
	// ; i	j	func	sigma1-4	epsilon1-4
	var at1, at2 string
	var fn int8
	var sig14, eps14 float64
	fields := strings.Fields(s)

	n, err := fmt.Sscanf(s, "%s %s %d %f %f", &at1, &at2, &fn, &sig14, &eps14)
	if len(fields) != 5 || n != 5 || err != nil {
		return nil, errors.New("error parsing [pairtypes] line")
	}

	switch fn {
	case 1:
		pt := ff.NewPairType(at1, at2, ff.FF_PAIR_TYPE_1)
		pt.SetSigma14(sig14)
		pt.SetEpsilon14(eps14)
		return pt, nil
	default:
		return nil, errors.New("pairtype function type is not 1")
	}

}

// parses [pairs]
func parsePairs(s string, topPol *ff.TopPolymer) (*ff.TopPair, error) {
	// ; ai    aj  funct   c6  c12

	return nil, nil
}

// parses [bondtypes]
func parseBondTypes(s string) (*ff.BondType, error) {
	// ; i	j	func	b0	Kb
	var at1, at2 string
	var fn int8
	var r0, kr float64
	fields := strings.Fields(s)

	n, err := fmt.Sscanf(s, "%s %s %d %f %f", &at1, &at2, &fn, &r0, &kr)
	if len(fields) != 5 || n != 5 || err != nil {
		return nil, errors.New("error parsing [bondtypes] line")
	}

	switch fn {
	case 1:
		bt := ff.NewBondType(at1, at2, ff.FF_BOND_TYPE_1)
		bt.SetHarmonicConstant(kr)
		bt.SetHarmonicDistance(r0)
		return bt, nil
	default:
		return nil, errors.New("bondtype function type is not 1")
	}
	return nil, nil
}

// parses [bonds]
func parseBonds(s string, topPol *ff.TopPolymer) (*ff.TopBond, error) {
	// ; ai    aj  funct   b0  Kb

	// check if we have the right number of fields (3 or 5)
	fields := strings.Fields(s)
	if len(fields) != 3 || len(fields) != 5 {
		return nil, errors.New("number of fields in [bonds] is not 3 or 5")
	}

	// parse the fields
	// if there are 5 fields, the fields[3] and fields[4] are b0 and kb

	var ai, aj int64
	var fn int8
	var b0, kb float64

	switch len(fields) {
	case 3:
		n, err := fmt.Sscanf(s, "%d %d %d", &ai, &aj, &fn)
		if n != 3 || err != nil {
			return nil, errors.New("problem parsing [bonds] line with 3 fields")
		}

	case 5:
		n, err := fmt.Sscanf(s, "%d %d %d %f %f", &ai, &aj, &fn, &b0, &kb)
		if n != 5 || err != nil {
			return nil, errors.New("problem parsing [bonds] line with 5 fields")
		}
	}

	// find the *AtomType
	a1 := topPol.AtomBySerial(ai)
	a2 := topPol.AtomBySerial(aj)

	if a1 == nil || a2 == nil {
		return nil, errors.New("could not find either ai or aj")
	}

	// check if we have a supported function type
	if fn != 1 {
		return nil, errors.New("[bonds] fn must be 1")
	}

	// create bond
	b := ff.NewTopBond(a1, a2, ff.FF_BOND_TYPE_1)

	// if we have custom parameters, create a corresponding *BondType
	if len(fields) == 5 {
		bt := ff.NewBondType(a1.AtomType(), a2.AtomType(), ff.FF_BOND_TYPE_1)
		b.SetCustomBondType(bt)
	}

	// we're good
	return b, nil

}

// parses [angletypes]
func parseAngleTypes(s string) (*ff.AngleType, error) {
	//; i	j	k	func	th0	cth	S0	Kub
	fields := strings.Fields(s)
	fn, err := strconv.ParseInt(fields[3], 10, 8)
	if err != nil {
		return nil, errors.New("could not determine function type")
	}

	var at1, at2, at3 string
	var tmp int8
	var thet, kt, r13, kub float64

	switch fn {
	case 1:
		n, err := fmt.Sscanf(s, "%s %s %s %d %f %f", &at1, &at2, &at3, &tmp, &thet, &kt)
		if n != 6 || err != nil {
			return nil, errors.New("could not parse angletype")
		}
		at := ff.NewAngleType(at1, at2, at3, ff.FF_ANGLE_TYPE_1)
		at.SetThetaConstant(kt)
		at.SetTheta(thet)
		return at, nil
	case 5:
		n, err := fmt.Sscanf(s, "%s %s %s %d %f %f %f %f", &at1, &at2, &at3, &tmp, &thet, &kt, &r13, &kub)
		if n != 8 || err != nil {
			return nil, errors.New("could not parse angletype")
		}
		at := ff.NewAngleType(at1, at2, at3, ff.FF_ANGLE_TYPE_5)
		at.SetThetaConstant(kt)
		at.SetTheta(thet)
		at.SetR13(r13)
		at.SetUBConstant(kub)
		return at, nil
	default:
		return nil, errors.New("angletype function type is not 1 or 5")
	}
}

// parses [angles]
func parseAngles(s string, topPol *ff.TopPolymer) (*ff.TopAngle, error) {
	// ; ai    aj  ak  funct   th0 cth S0  Kub

	return nil, nil
}

// parses [dihedraltypes]
func parseDihedralTypes(s string) (*ff.DihedralType, error) {
	// ; i	j	k	l	func	phi0	cp	mult

	fields := strings.Fields(s)

	// find the function type
	fn, err := strconv.ParseInt(fields[4], 10, 8)
	if err != nil {
		return nil, errors.New("could not determine function type")
	}

	var at1, at2, at3, at4 string
	var tmp, mult int8
	var phi, psi, kphi, kpsi float64

	switch fn {
	case 1, 9:
		n, err := fmt.Sscanf(s, "%s %s %s %s %d %f %f %d", &at1, &at2, &at3, &at4, &tmp, &phi, &kphi, &mult)
		if n != 8 || err != nil {
			return nil, errors.New("could not parse [dihedraltypes]")
		}

		var dt *ff.DihedralType
		if fn == 1 {
			dt = ff.NewDihedralType(at1, at2, at3, at4, ff.FF_DIHEDRAL_TYPE_1)
		} else if fn == 9 {
			dt = ff.NewDihedralType(at1, at2, at3, at4, ff.FF_DIHEDRAL_TYPE_9)
		}

		dt.SetPhi(phi)
		dt.SetPhiConstant(kphi)
		dt.SetMult(mult)
		return dt, nil

	case 2:
		n, err := fmt.Sscanf(s, "%s %s %s %s %d %f %f", &at1, &at2, &at3, &at4, &tmp, &psi, &kpsi)
		if n != 7 || err != nil {
			return nil, errors.New("could not parse [dihedraltypes]")
		}
		dt := ff.NewDihedralType(at1, at2, at3, at4, ff.FF_DIHEDRAL_TYPE_2)
		dt.SetPsiConstant(kpsi)
		dt.SetPsi(psi)
		return dt, nil
	default:
		return nil, errors.New("[dihedraltype] function type is not 1, 2 or 9")
	}
}

// parses [dihedrals]
func parseDihedral(s string, topPol *ff.TopPolymer) (*ff.TopDihedral, error) {
	// ; ai    aj  ak  al  funct   phi0    cp  mult
	// ; ai    aj  ak  al  funct   q0  cq

	return nil, nil
}

// parses [moleculetypes]
func parseMoleculeTypes(s string) (*ff.TopPolymer, error) {
	// ; name	nrexcl

	fields := strings.Fields(s)
	if len(fields) != 2 {
		return nil, errors.New("number of fields in [moleculetype] is not 2")
	}

	name := fields[0]
	nrexcl, err := strconv.ParseInt(fields[1], 10, 8)
	if err != nil {
		return nil, err
	}

	p := ff.NewTopPolymer()
	p.SetName(name)
	p.SetNrExcl(int8(nrexcl))

	return p, nil

}

//
