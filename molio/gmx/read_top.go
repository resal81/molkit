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

	"github.com/resal81/molkit/blocks"
	"github.com/resal81/molkit/ff"
)

var (
	ErrDefaults  = errors.New("[defaults] could not be parsed")
	ErrAtomTypes = errors.New("[atomtypes] could not be parsed")
)

type topologyLevel int8
type topologyGroup int64

func ReadTOPFile(fname string) (*blocks.System, *ff.ForceField, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	return readtop(file)

}

func ReadTOPString(s string) (*blocks.System, *ff.ForceField, error) {
	reader := strings.NewReader(s)
	return readtop(reader)
}

func readtop(reader io.Reader) (*blocks.System, *ff.ForceField, error) {

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

	// var curr_topmol *ff.TopPolymer
	forcefield := ff.NewForceField(ff.FF_GROMACS)

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

				if dt.Setting()&ff.DHT_TYPE_1 != 0 || dt.Setting()&ff.DHT_TYPE_9 != 0 {
					forcefield.AddDihedralType(dt)

				} else if dt.Setting()&ff.DHT_TYPE_2 != 0 {
					forcefield.AddImproperType(dt)

				} else {
					return nil, nil, errors.New("dihedral type is not valid")
				}

			}

		}

		if level == L_MOLTYPES {
			switch group {
			case G_ATOMS:
			}
		}

		if level == L_SYSTEM {
			switch group {
			case G_SYSTEM:
				continue
			}
		}

	}

	return nil, forcefield, nil

}

//
func cleanLine(s string) string {
	i := strings.Index(s, ";")
	if i != -1 {
		s = s[:i]
	}

	s = strings.TrimSpace(s)
	return s
}

// Extract whats inside brackets: [ ... ]
func getGroup(s string) string {
	s = strings.TrimLeft(s, "[")
	s = strings.TrimRight(s, "]")
	s = strings.TrimSpace(s)
	return s
}

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
		nbt := ff.NewNonBondedType(at1, at2, ff.NBT_TYPE_1)
		nbt.SetSigma(sig)
		nbt.SetEpsilon(eps)
		return nbt, nil
	default:
		return nil, errors.New("nonbonded function type is not 1")
	}

}

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
		pt := ff.NewPairType(at1, at2, ff.PT_TYPE_1)
		pt.SetSigma14(sig14)
		pt.SetEpsilon14(eps14)
		return pt, nil
	default:
		return nil, errors.New("pairtype function type is not 1")
	}

}

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
		bt := ff.NewBondType(at1, at2, ff.BT_TYPE_1)
		bt.SetHarmonicConstant(kr)
		bt.SetHarmonicDistance(r0)
		return bt, nil
	default:
		return nil, errors.New("bondtype function type is not 1")
	}
	return nil, nil
}

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
		at := ff.NewAngleType(at1, at2, at3, ff.ANG_TYPE_1)
		at.SetThetaConstant(kt)
		at.SetTheta(thet)
		return at, nil
	case 5:
		n, err := fmt.Sscanf(s, "%s %s %s %d %f %f %f %f", &at1, &at2, &at3, &tmp, &thet, &kt, &r13, &kub)
		if n != 8 || err != nil {
			return nil, errors.New("could not parse angletype")
		}
		at := ff.NewAngleType(at1, at2, at3, ff.ANG_TYPE_5)
		at.SetThetaConstant(kt)
		at.SetTheta(thet)
		at.SetR13(r13)
		at.SetUBConstant(kub)
		return at, nil
	default:
		return nil, errors.New("angletype function type is not 1 or 5")
	}
}

func parseDihedralTypes(s string) (*ff.DihedralType, error) {
	// ; i	j	k	l	func	phi0	cp	mult
	fields := strings.Fields(s)
	fn, err := strconv.ParseInt(fields[4], 10, 8)
	if err != nil {
		return nil, errors.New("could not determine function type")
	}

	var at1, at2, at3, at4 string
	var tmp, mult int8
	var phi, psi, kphi, kpsi float64

	switch fn {
	case 1:
		n, err := fmt.Sscanf(s, "%s %s %s %s %d %f %f %d", &at1, &at2, &at3, &at4, &tmp, &phi, &kphi, &mult)
		if n != 8 || err != nil {
			return nil, errors.New("could not parse dihedraltype")
		}
		dt := ff.NewDihedralType(at1, at2, at3, at4, ff.DHT_TYPE_1)
		dt.SetPhi(phi)
		dt.SetPhiConstant(kphi)
		dt.SetMult(mult)
		return dt, nil

	case 9:
		n, err := fmt.Sscanf(s, "%s %s %s %s %d %f %f %d", &at1, &at2, &at3, &at4, &tmp, &phi, &kphi, &mult)
		if n != 8 || err != nil {
			return nil, errors.New("could not parse dihedraltype")
		}
		dt := ff.NewDihedralType(at1, at2, at3, at4, ff.DHT_TYPE_9)
		dt.SetPhi(phi)
		dt.SetPhiConstant(kphi)
		dt.SetMult(mult)
		return dt, nil

	case 2:
		n, err := fmt.Sscanf(s, "%s %s %s %s %d %f %f", &at1, &at2, &at3, &at4, &tmp, &psi, &kpsi)
		if n != 7 || err != nil {
			return nil, errors.New("could not parse dihedraltype")
		}
		dt := ff.NewDihedralType(at1, at2, at3, at4, ff.DHT_TYPE_2)
		dt.SetPsiConstant(kpsi)
		dt.SetPsi(psi)
		return dt, nil
	default:
		return nil, errors.New("dihedraltype function type is not 1, 2 or 9")
	}
}

//
