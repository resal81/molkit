package gmx

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
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

func ReadTOPFile(fname string) (*blocks.System, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return readtop(file)

}

func readtop(reader io.Reader) (*blocks.System, error) {

	// settings

	// main levels
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
	paramsDB := ff.NewParamsDB(ff.P_DB_GROMACS)

	// read file line by line
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || line == "" {
			continue
		}

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
				return nil, fmt.Errorf("unknown group in the top file: %s", line)
			}

			continue
		}

		if level == L_PARAMS {
			switch group {
			case G_DEFAULTS:
				defaults, err := parseDefaults(line)
				if err != nil {
					log.Printf("error in line: '%s' \n", line)
					return nil, err
				}
				paramsDB.SetGMXDefaults(defaults)
			case G_ATOMTYPES:

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

	return nil, nil

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
	var mass, chg, sig, eps, sig14, eps14 float32

	n, err := fmt.Sscanf(s, "%s %d %f %f %s %f %f %f %f", &name, &prot, &mass, &chg, &pt, &sig, &eps, &sig14, &eps14)
	if err != nil {
		return nil, err
	}
	if n != 7 || n != 9 {
		return nil, ErrAtomTypes
	}

	at := ff.NewAtomType(name, ff.AT_GMX_ATOMTYPE)
	at.SetProtons(prot)
	at.SetMass(mass)
	at.SetCharge(chg)
	at.SetSigma(sig)
	at.SetEpsilon(eps)

	if n == 9 {
		at.SetSigma14(sig14)
		at.SetEpsilon14(eps14)
	}

	return at, nil
}
