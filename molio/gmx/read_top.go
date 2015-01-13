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

const (
	errCouldNotBeParsed int32 = 1 << iota
	errBadFunctionType
	errBadNumbFields
	errCouldNotFindTopAtom
)

var errMessages map[int32]string = map[int32]string{
	errCouldNotBeParsed:    "'%s' : line could not be parsed",
	errBadFunctionType:     "'%s' : bad fn type",
	errBadNumbFields:       "'%s' : wrong number of fields",
	errCouldNotFindTopAtom: "'%s' : could not find *TopAtom",
}

var (
	ErrDefaults  = errors.New("[defaults] line could not be parsed")
	ErrAtomTypes = errors.New("[atomtypes] line could not be parsed")
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
				forcefield.SetPropsGMX(defaults)

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
				bn, err := parseBonds(line, curr_topmol)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}

				curr_topmol.AddTopBond(bn)

			case G_PAIRS:
				pr, err := parsePairs(line, curr_topmol)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}

				curr_topmol.AddTopPair(pr)

			case G_ANGLES:
				ag, err := parseAngles(line, curr_topmol)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}

				curr_topmol.AddTopAngle(ag)

			case G_DIHEDRALS:
				dh, err := parseDihedral(line, curr_topmol)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, nil, err
				}

				if dh.Kind()&ff.FF_DIHEDRAL_TYPE_1 != 0 || dh.Kind()&ff.FF_DIHEDRAL_TYPE_9 != 0 {
					curr_topmol.AddTopDihedral(dh)
				} else if dh.Kind()&ff.FF_DIHEDRAL_TYPE_2 != 0 {
					curr_topmol.AddTopImproper(dh)
				} else {
					panic("should not reach here")
				}

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

//
func genError(group string, errType int32) error {
	if msg, ok := errMessages[errType]; ok {
		return fmt.Errorf(msg, group)
	} else {
		return fmt.Errorf("'%s' : unknown error", group)
	}
}

//
func checkLine(s string, exp_lens []int, exp_fns []int8, fn_index int) (nfields int, out_fn int8, err error) {
	fields := strings.Fields(s)

	// check length
	len_ok := false
	for _, l := range exp_lens {
		if len(fields) == l {
			nfields = l
			len_ok = true
			break
		}
	}

	if !len_ok {
		return 0, 0, errors.New("bad length")
	}

	// check fn
	if fn_index > len(fields)-1 {
		return 0, 0, errors.New("bad index")
	}
	curr_fn, e := strconv.ParseInt(fields[fn_index], 10, 8)
	if e != nil {
		return 0, 0, errors.New("could not extract fn")
	}

	fn_ok := false
	for _, f := range exp_fns {
		if int8(curr_fn) == f {
			out_fn = f
			fn_ok = true
			break
		}
	}

	if !fn_ok {
		return 0, 0, errors.New("bad fn")
	}

	return nfields, out_fn, nil

}

// parses [defaults]
func parseDefaults(s string) (*ff.GMXProps, error) {
	// ; nbfunc	comb-rule	gen-pairs	fudgeLJ	fudgeQQ

	var nbfunc, combrule int8
	var fudgeLJ, fudgeQQ float64
	var genpairs string

	n, err := fmt.Sscanf(s, "%d %d %s %f %f", &nbfunc, &combrule, &genpairs, &fudgeLJ, &fudgeQQ)
	if err != nil || n != 5 {
		return nil, ErrDefaults
	}

	gd := ff.NewGMXProps(nbfunc, combrule, genpairs == "yes", fudgeLJ, fudgeQQ)
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
		return nil, genError("[atomtypes]", errBadNumbFields)
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
		return nil, nil, genError("[atoms]", errBadNumbFields)
	}

	var name, atype, resname string
	var chg, mass float64
	var ser, cgnr, resnumb int64

	n, err := fmt.Sscanf(s, "%d %s %d %s %s %d %f %f", &ser, &atype, &resnumb, &resname, &name, &cgnr, &chg, &mass)
	if n != 8 || err != nil {
		return nil, nil, genError("[atoms]", errCouldNotBeParsed)
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
		return nil, genError("[nonbonded]", errCouldNotBeParsed)
	}

	switch fn {
	case 1:
		nbt := ff.NewNonBondedType(at1, at2, ff.FF_NON_BONDED_TYPE_1)
		nbt.SetSigma(sig)
		nbt.SetEpsilon(eps)
		return nbt, nil
	default:
		return nil, genError("[nonbonded]", errBadFunctionType)
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
		return nil, genError("[pairtypes]", errCouldNotBeParsed)
	}

	switch fn {
	case 1:
		pt := ff.NewPairType(at1, at2, ff.FF_PAIR_TYPE_1)
		pt.SetSigma14(sig14)
		pt.SetEpsilon14(eps14)
		return pt, nil
	default:
		return nil, genError("[pairtypes]", errBadFunctionType)
	}

}

// parses [pairs]
func parsePairs(s string, topPol *ff.TopPolymer) (*ff.TopPair, error) {
	// ; ai    aj  funct   c6  c12

	// check the number of fields
	nfields, fn, err := checkLine(s, []int{3, 5}, []int8{1}, 2)
	if err != nil {
		return nil, genError("[pairs]", errCouldNotBeParsed)
	}

	var ai, aj int64
	var tmp int8
	var sig14, eps14 float64

	switch nfields {
	case 3:
		n, err := fmt.Sscanf(s, "%d %d %d", &ai, &aj, &tmp)
		if n != 3 || err != nil {
			return nil, genError("[pairs]", errCouldNotBeParsed)
		}

	case 5:
		n, err := fmt.Sscanf(s, "%d %d %d %f %f", &ai, &aj, &fn, &sig14, &eps14)
		if n != 5 || err != nil {
			return nil, genError("[pairs]", errCouldNotBeParsed)
		}
	}

	// find the *AtomType
	a1 := topPol.AtomBySerial(ai)
	a2 := topPol.AtomBySerial(aj)

	if a1 == nil || a2 == nil {
		return nil, genError("[pairs]", errCouldNotFindTopAtom)
	}

	if fn == 1 {
		// create pair
		p := ff.NewTopPair(a1, a2, ff.FF_PAIR_TYPE_1)

		// if we have custom parameters, create a corresponding *PairType
		if nfields == 5 {
			pt := ff.NewPairType(a1.AtomType(), a2.AtomType(), ff.FF_PAIR_TYPE_1)
			p.SetCustomPairType(pt)
		}

		return p, nil

	} else {

		return nil, genError("[pairs]", errBadFunctionType)
	}

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
		return nil, genError("[bondtypes]", errCouldNotBeParsed)
	}

	switch fn {
	case 1:
		bt := ff.NewBondType(at1, at2, ff.FF_BOND_TYPE_1)
		bt.SetHarmonicConstant(kr)
		bt.SetHarmonicDistance(r0)
		return bt, nil
	default:
		return nil, genError("[bondtypes]", errBadFunctionType)
	}
	return nil, nil
}

// parses [bonds]
func parseBonds(s string, topPol *ff.TopPolymer) (*ff.TopBond, error) {
	// ; ai    aj  funct   b0  Kb

	// check the number of fields
	nfields, fn, err := checkLine(s, []int{3, 5}, []int8{1}, 2)
	if err != nil {
		return nil, genError("[pairs]", errCouldNotBeParsed)
	}

	// parse the fields
	// if there are 5 fields, the fields[3] and fields[4] are b0 and kb
	var ai, aj int64
	var tmp int8
	var b0, kb float64

	switch nfields {
	case 3:
		n, err := fmt.Sscanf(s, "%d %d %d", &ai, &aj, &tmp)
		if n != 3 || err != nil {
			return nil, genError("[bonds]", errCouldNotBeParsed)
		}

	case 5:
		n, err := fmt.Sscanf(s, "%d %d %d %f %f", &ai, &aj, &tmp, &b0, &kb)
		if n != 5 || err != nil {
			return nil, genError("[bonds]", errCouldNotBeParsed)
		}
	}

	// find the *AtomType
	a1 := topPol.AtomBySerial(ai)
	a2 := topPol.AtomBySerial(aj)

	if a1 == nil || a2 == nil {
		return nil, genError("[bonds]", errCouldNotFindTopAtom)
	}

	if fn == 1 {
		// create bond
		b := ff.NewTopBond(a1, a2, ff.FF_BOND_TYPE_1)

		// if we have custom parameters, create a corresponding *BondType
		if nfields == 5 {
			bt := ff.NewBondType(a1.AtomType(), a2.AtomType(), ff.FF_BOND_TYPE_1)
			b.SetCustomBondType(bt)
		}

		return b, nil

	} else {
		return nil, genError("[bonds]", errBadFunctionType)
	}

}

// parses [angletypes]
func parseAngleTypes(s string) (*ff.AngleType, error) {
	//; i	j	k	func	th0	cth	S0	Kub

	_, fn, err := checkLine(s, []int{6, 8}, []int8{1, 5}, 3)
	if err != nil {
		return nil, genError("[angletypes]", errCouldNotBeParsed)
	}

	var at1, at2, at3 string
	var tmp int8
	var thet, kt, r13, kub float64

	var at *ff.AngleType

	switch fn {
	case 1:
		n, err := fmt.Sscanf(s, "%s %s %s %d %f %f", &at1, &at2, &at3, &tmp, &thet, &kt)
		if n != 6 || err != nil {
			return nil, genError("[angletypes]", errCouldNotBeParsed)
		}
		at = ff.NewAngleType(at1, at2, at3, ff.FF_ANGLE_TYPE_1)

	case 5:
		n, err := fmt.Sscanf(s, "%s %s %s %d %f %f %f %f", &at1, &at2, &at3, &tmp, &thet, &kt, &r13, &kub)
		if n != 8 || err != nil {
			return nil, genError("[angletypes]", errCouldNotBeParsed)
		}
		at = ff.NewAngleType(at1, at2, at3, ff.FF_ANGLE_TYPE_5)
	}

	at.SetThetaConstant(kt)
	at.SetTheta(thet)

	if fn == 5 {
		at.SetR13(r13)
		at.SetUBConstant(kub)
	}

	return at, nil

}

// parses [angles]
func parseAngles(s string, topPol *ff.TopPolymer) (*ff.TopAngle, error) {
	// ; ai    aj  ak  funct   th0 cth S0  Kub

	nfields, fn, err := checkLine(s, []int{4, 6, 8}, []int8{1, 5}, 3)
	if err != nil {
		return nil, genError("[angles]", errCouldNotBeParsed)
	}

	var ai, aj, ak int64
	var tmp int8
	var thet, kt, r13, kub float64

	switch fn {
	case 1:
		if nfields == 4 {
			n, err := fmt.Sscanf(s, "%d %d %d %d", &ai, &aj, &ak, &tmp)
			if n != 4 || err != nil {
				return nil, genError("[angles]", errCouldNotBeParsed)
			}

		} else if nfields == 6 {
			n, err := fmt.Sscanf(s, "%d %d %d %d %f %f", &ai, &aj, &ak, &tmp, &thet, &kt)
			if n != 6 || err != nil {
				return nil, genError("[angles]", errCouldNotBeParsed)
			}

		} else {
			return nil, genError("[angles]", errBadNumbFields)
		}

	case 5:
		if nfields == 4 {
			n, err := fmt.Sscanf(s, "%d %d %d %d", &ai, &aj, &ak, &tmp)
			if n != 4 || err != nil {
				return nil, genError("[angles]", errCouldNotBeParsed)
			}

		} else if nfields == 8 {
			n, err := fmt.Sscanf(s, "%d %d %d %d %f %f %f %f", &ai, &aj, &ak, &tmp, &thet, &kt, &r13, &kub)
			if n != 8 || err != nil {
				return nil, genError("[angles]", errCouldNotBeParsed)
			}

		} else {
			return nil, genError("[angles]", errBadNumbFields)
		}

	}

	a1 := topPol.AtomBySerial(ai)
	a2 := topPol.AtomBySerial(aj)
	a3 := topPol.AtomBySerial(ak)

	if a1 == nil || a2 == nil || a3 == nil {
		return nil, genError("[angles]", errCouldNotFindTopAtom)
	}

	var tg *ff.TopAngle
	switch fn {
	case 1:
		tg = ff.NewTopAngle(a1, a2, a3, ff.FF_ANGLE_TYPE_1)
		switch nfields {
		case 6:
			at := ff.NewAngleType(a1.AtomType(), a2.AtomType(), a3.AtomType(), ff.FF_ANGLE_TYPE_1)
			at.SetThetaConstant(kt)
			at.SetTheta(thet)

			tg.SetCustomAngleType(at)
		}
	case 5:
		tg = ff.NewTopAngle(a1, a2, a3, ff.FF_ANGLE_TYPE_5)
		switch nfields {
		case 8:
			at := ff.NewAngleType(a1.AtomType(), a2.AtomType(), a3.AtomType(), ff.FF_ANGLE_TYPE_5)
			at.SetThetaConstant(kt)
			at.SetTheta(thet)
			at.SetUBConstant(kub)
			at.SetR13(r13)

			tg.SetCustomAngleType(at)
		}
	}

	return tg, nil
}

// parses [dihedraltypes]
func parseDihedralTypes(s string) (*ff.DihedralType, error) {
	// ; i	j	k	l	func	phi0	cp	mult

	// find the function type
	_, fn, err := checkLine(s, []int{7, 8}, []int8{1, 2, 9}, 4)
	if err != nil {
		return nil, genError("[dihedraltypes]", errCouldNotBeParsed)
	}

	var at1, at2, at3, at4 string
	var tmp, mult int8
	var phi, psi, kphi, kpsi float64

	switch fn {
	case 1, 9:
		n, err := fmt.Sscanf(s, "%s %s %s %s %d %f %f %d", &at1, &at2, &at3, &at4, &tmp, &phi, &kphi, &mult)
		if n != 8 || err != nil {
			return nil, genError("[dihedraltypes]", errCouldNotBeParsed)
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
			return nil, genError("[dihedraltypes]", errCouldNotBeParsed)
		}
		dt := ff.NewDihedralType(at1, at2, at3, at4, ff.FF_DIHEDRAL_TYPE_2)
		dt.SetPsiConstant(kpsi)
		dt.SetPsi(psi)
		return dt, nil
	}

	return nil, genError("[dihedraltypes]", errCouldNotBeParsed)
}

// parses [dihedrals]
func parseDihedral(s string, topPol *ff.TopPolymer) (*ff.TopDihedral, error) {
	// ; ai    aj  ak  al  funct   phi0    cp  mult
	// ; ai    aj  ak  al  funct   q0  cq

	// find the function type
	nfields, fn, err := checkLine(s, []int{5, 7, 8}, []int8{1, 2, 9}, 4)
	if err != nil {
		return nil, genError("[dihedrals]", errCouldNotBeParsed)
	}

	var ai, aj, ak, am int64
	var tmp, mult int8
	var phi, psi, kphi, kpsi float64

	switch fn {
	case 1, 9:
		if nfields == 5 {
			n, err := fmt.Sscanf(s, "%d %d %d %d %d", &ai, &aj, &ak, &am, &tmp)
			if n != 5 || err != nil {
				return nil, genError("[dihedrals]", errCouldNotBeParsed)
			}
		} else if nfields == 8 {
			n, err := fmt.Sscanf(s, "%d %d %d %d %d %f %f %d", &ai, &aj, &ak, &am, &tmp, &phi, &kphi, &mult)
			if n != 8 || err != nil {
				return nil, genError("[dihedrals]", errCouldNotBeParsed)
			}
		}

	case 2:
		if nfields == 5 {
			n, err := fmt.Sscanf(s, "%d %d %d %d %d", &ai, &aj, &ak, &am, &tmp)
			if n != 5 || err != nil {
				return nil, genError("[dihedrals]", errCouldNotBeParsed)
			}
		} else if nfields == 7 {
			n, err := fmt.Sscanf(s, "%d %d %d %d %d %f %f %d", &ai, &aj, &ak, &am, &tmp, &psi, &kpsi)
			if n != 7 || err != nil {
				return nil, genError("[dihedrals]", errCouldNotBeParsed)
			}
		}
	}

	a1 := topPol.AtomBySerial(ai)
	a2 := topPol.AtomBySerial(aj)
	a3 := topPol.AtomBySerial(ak)
	a4 := topPol.AtomBySerial(am)

	if a1 == nil || a2 == nil || a3 == nil || a4 == nil {
		return nil, genError("[dihedrals]", errCouldNotFindTopAtom)
	}

	var dh *ff.TopDihedral

	switch fn {
	case 1:
		dh = ff.NewTopDihedral(a1, a2, a3, a4, ff.FF_DIHEDRAL_TYPE_1)
		if nfields == 8 {
			dt := ff.NewDihedralType(a1.AtomType(), a2.AtomType(), a3.AtomType(), a4.AtomType(), ff.FF_DIHEDRAL_TYPE_1)
			dt.SetPhiConstant(kphi)
			dt.SetPhi(phi)
			dt.SetMult(mult)

			dh.SetCustomDihedralType(dt)
		}
	case 9:
		dh = ff.NewTopDihedral(a1, a2, a3, a4, ff.FF_DIHEDRAL_TYPE_9)
		if nfields == 8 {
			dt := ff.NewDihedralType(a1.AtomType(), a2.AtomType(), a3.AtomType(), a4.AtomType(), ff.FF_DIHEDRAL_TYPE_9)
			dt.SetPhiConstant(kphi)
			dt.SetPhi(phi)
			dt.SetMult(mult)

			dh.SetCustomDihedralType(dt)
		}
	case 2:
		dh = ff.NewTopDihedral(a1, a2, a3, a4, ff.FF_DIHEDRAL_TYPE_2)
		if nfields == 7 {
			dt := ff.NewDihedralType(a1.AtomType(), a2.AtomType(), a3.AtomType(), a4.AtomType(), ff.FF_DIHEDRAL_TYPE_2)
			dt.SetPsiConstant(kpsi)
			dt.SetPsi(psi)

			dh.SetCustomDihedralType(dt)
		}
	}

	return dh, nil
}

// parses [moleculetypes]
func parseMoleculeTypes(s string) (*ff.TopPolymer, error) {
	// ; name	nrexcl

	fields := strings.Fields(s)
	if len(fields) != 2 {
		return nil, genError("[moleculetype]", errBadNumbFields)
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
