package chm

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/resal81/molkit/ff"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadPRMFiles(fnames ...string) (*ff.ForceField, error) {

	frc := ff.NewForceField(ff.FF_CHARMM)

	for _, fname := range fnames {
		file, err := os.Open(fname)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		err = readprm(file, frc)
		if err != nil {
			return nil, err
		}
	}

	return frc, nil

}

func ReadPRMString(s string) (*ff.ForceField, error) {
	frc := ff.NewForceField(ff.FF_CHARMM)
	reader := strings.NewReader(s)
	err := readprm(reader, frc)
	return frc, err
}

type prmLevel int64

func readprm(reader io.Reader, frc *ff.ForceField) error {

	const (
		L_ATOMS prmLevel = 1 << iota
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

	cmap_header := ""
	var cmap_str_vals []string = []string{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		line = cleanLine(line)
		if line == "" {
			continue
		}

		if strings.ToUpper(line) == "END" {
			break
		}

		if len(line) < 4 {
			panic("line length less that 5 " + line)
		}

		switch strings.ToUpper(line[:4]) {
		case "ATOM":
			lvl = L_ATOMS
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
		case "cutn":
			lvl = L_IGNORE
			continue
		case "HBON":
			lvl = L_IGNORE
			continue
		}

		switch lvl {
		case L_ATOMS:
			name, mass, err := parseAtomType(line)
			if err != nil {
				log.Printf("error in line: %s", line)
				return err
			}
			massDB[name] = mass

		case L_BONDS:
			bt, err := parseBondType(line)
			if err != nil {
				log.Printf("error in line: %s", line)
				return err
			}
			frc.AddBondType(bt)

		case L_ANGLES:
			at, err := parseAngleType(line)
			if err != nil {
				log.Printf("error in line: %s", line)
				return err
			}
			frc.AddAngleType(at)

		case L_DIHEDRALS:
			dh, err := parseDihedralType(line)
			if err != nil {
				log.Printf("error in line: %s", line)
				return err
			}
			frc.AddDihedralType(dh)

		case L_IMPROPERS:
			im, err := parseImproperType(line)
			if err != nil {
				log.Printf("error in line: %s", line)
				return err
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
					cm, err := parseCMapType(24, 24, strings.Fields(cmap_header), cmap_str_vals)
					if err != nil {
						log.Printf("error in line: %s", line)
						return err
					}
					frc.AddCMapType(cm)
					cmap_header = ""
				}
			}
		case L_NBFIX:
		case L_IGNORE:
			continue
		}

	}

	return nil
}

// removes comments plus leading and tailing spaces
func cleanLine(s string) string {
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
func checkLineFields(s string, exp_lens []int) (nfields int, err error) {
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
		return 0, errors.New("bad length")
	}

	return nfields, nil
}

func parseAtomType(s string) (string, float64, error) {
	if strings.HasPrefix(s, "MASS") {
		fields := strings.Fields(s)
		if len(fields) != 4 {
			return "", 0, errors.New("bad length in MASS line")
		}

		name := fields[2]
		m, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return "", 0, err
		}

		return name, m, nil

	} else {
		panic("ATOMS line without MASS prefix")
	}
}

//
func parseBondType(s string) (*ff.BondType, error) {

	// atype1 atype2  Kb  b0
	_, err := checkLineFields(s, []int{4})
	if err != nil {
		return nil, err
	}

	var at1, at2 string
	var kb, b0 float64

	n, err := fmt.Sscanf(s, "%s %s %f %f", &at1, &at2, &kb, &b0)
	if n != 4 || err != nil {
		return nil, errors.New("error paring BONDS line")
	}

	bt := ff.NewBondType(at1, at2, ff.FF_BOND_TYPE_1, ff.FF_CHARMM)
	bt.SetHarmonicConstant(kb)
	bt.SetHarmonicDistance(b0)

	return bt, nil

}

//
func parseAngleType(s string) (*ff.AngleType, error) {

	// atyp1 atype2 atype3     Ktheta    Theta0   Kub     S0
	nfields, err := checkLineFields(s, []int{5, 7})
	if err != nil {
		return nil, err
	}

	var at1, at2, at3 string
	var kt, theta, kub, r13 float64

	switch nfields {
	case 5:
		n, err := fmt.Sscanf(s, "%s %s %s %f %f", &at1, &at2, &at3, &kt, &theta)
		if n != 5 || err != nil {
			return nil, errors.New("could not parse angletype - 5")
		}
	case 7:
		n, err := fmt.Sscanf(s, "%s %s %s %f %f %f %f", &at1, &at2, &at3, &kt, &theta, &kub, &r13)
		if n != 7 || err != nil {
			return nil, errors.New("could not parse angletype - 7")
		}
	}

	at := ff.NewAngleType(at1, at2, at3, ff.FF_ANGLE_TYPE_5, ff.FF_CHARMM)
	at.SetThetaConstant(kt)
	at.SetTheta(theta)
	at.SetUBConstant(kub)
	at.SetR13(r13)

	return at, nil
}

//
func parseDihedralType(s string) (*ff.DihedralType, error) {

	// atype1 atype2 atype3  atype4 Kchi    n   delta
	_, err := checkLineFields(s, []int{7})
	if err != nil {
		return nil, err
	}

	var at1, at2, at3, at4 string
	var mult int8
	var kphi, phi float64

	n, err := fmt.Sscanf(s, "%s %s %s %s %f %d %f", &at1, &at2, &at3, &at4, &kphi, &mult, &phi)
	if n != 7 || err != nil {
		return nil, errors.New("could not parse dihedraltype")
	}

	dh := ff.NewDihedralType(at1, at2, at3, at4, ff.FF_DIHEDRAL_TYPE_9, ff.FF_CHARMM)
	dh.SetPhiConstant(kphi)
	dh.SetPhi(phi)
	dh.SetMult(mult)

	return dh, nil
}

//
func parseImproperType(s string) (*ff.ImproperType, error) {
	// atype1 atype2 atype3  atype4  Kpsi ign psi0
	_, err := checkLineFields(s, []int{7})
	if err != nil {
		return nil, err
	}

	var at1, at2, at3, at4, tmp string
	var kpsi, psi float64

	n, err := fmt.Sscanf(s, "%s %s %s %s %f %s %f", &at1, &at2, &at3, &at4, &kpsi, &tmp, &psi)
	if n != 7 || err != nil {
		return nil, errors.New("could not parse dihedraltype")
	}

	it := ff.NewImproperType(at1, at2, at3, at4, ff.FF_IMPROPER_TYPE_1, ff.FF_CHARMM)
	it.SetPsiConstant(kpsi)
	it.SetPsi(psi)

	return it, nil
}

//
func parseNonBondedType(s string) (*ff.AtomType, error) {
	return nil, nil
}

//
func parseCMapType(nx, ny int, atypes, vals []string) (*ff.CMapType, error) {

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

	cm := ff.NewCMapType(nx, ny, ff.FF_CHARMM)
	cm.SetValues(vals_f)
	return cm, nil
}

//
