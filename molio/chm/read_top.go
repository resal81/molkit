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
* ReadTOPFile
**********************************************************/

// Parses one or multiple CHARMM prm files.
func ReadTOPFiles(fnames ...string) ([]*blocks.Fragment, error) {

	fgs := []*blocks.Fragment{}

	for _, fname := range fnames {
		file, err := os.Open(fname)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		err = readtop(file, fgs)
		if err != nil {
			return nil, err
		}
	}

	return fgs, nil

}

/**********************************************************
* ReadTOPString
**********************************************************/

// Parses a prm string (e.g. contents of a file).
func ReadTOPString(s string) ([]*blocks.Fragment, error) {
	fgs := []*blocks.Fragment{}
	reader := strings.NewReader(s)
	err := readtop(reader, fgs)
	return fgs, err
}

/**********************************************************
* readtop
**********************************************************/

func readtop(reader io.Reader, fgs []*blocks.Fragment) error {

	var fg *blocks.Fragment
	name_atom := map[string]*blocks.Atom{}

	// read file
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		line = cleanTOPLine(line)
		if line == "" {
			continue
		}

		// check for END keyword
		if strings.ToUpper(line) == "END" {
			break
		}

		if len(line) < 4 {
			continue
		}

		head := strings.ToUpper(line[:4])

		switch {
		case strings.HasPrefix(head, "MASS"):

		case strings.HasPrefix(head, "PRES"):
			fg = nil

		case strings.HasPrefix(head, "RESI"):
			if fg != nil {
				fgs = append(fgs, fg)
			}

			fields := strings.Fields(line)
			// residue line can have 2 or 3 fields

			fg = blocks.NewFragment(strings.ToUpper(fields[1]))
			name_atom = map[string]*blocks.Atom{}

			// TODO add total charge

		case strings.HasPrefix(head, "GROU"):

		case strings.HasPrefix(head, "ATOM"):
			if fg == nil {
				continue
			}

			fields := strings.Fields(line)
			if len(fields) != 4 {
				return fmt.Errorf("ATOM line is not right => %s", line)
			}

			// new atom
			a := blocks.NewAtom(strings.ToUpper(fields[1]))

			// new atom type
			at := blocks.NewAtomType(strings.ToUpper(fields[2]), blocks.AT_TYPE_CHM_1)
			ch, err := strconv.ParseFloat(fields[3], 64)
			if err != nil {
				return fmt.Errorf("could not parse charge in this line => %s", line)
			}
			at.SetPartialCharge(ch)

			// connect both
			a.SetType(at)

			// add to the current fragment and atom map
			fg.AddAtom(a)
			name_atom[a.Name()] = a

		case strings.HasPrefix(head, "BOND") || strings.HasPrefix(head, "DOUB") || strings.HasPrefix(head, "TRIP"):
			if fg == nil {
				continue
			}
			fields := strings.Fields(line)
			if (len(fields)-1)%2 != 0 {
				return fmt.Errorf("uneven number of atoms for bond formation => %s", line)
			}
			for i := 1; i < len(fields)-1; i += 2 {
				n1 := strings.ToUpper(fields[i])
				n2 := strings.ToUpper(fields[i+1])

				// skip bonds between residues for now - TODO
				if isConnectionAtom(n1) || isConnectionAtom(n2) {
					continue
				}

				a1 := name_atom[n1]
				a2 := name_atom[n2]
				if a1 == nil || a2 == nil {
					return fmt.Errorf("one or both atoms were not found for bond => fragment: %s, line: %s", fg.Name(), line)
				}

				b := blocks.NewBond(a1, a2)
				fg.AddBond(b)
			}

		case strings.HasPrefix(head, "IMPR"):
			if fg == nil {
				continue
			}
			fields := strings.Fields(line)
			if (len(fields)-1)%4 != 0 {
				return fmt.Errorf("uneven number of atoms for improper formation => %s", line)
			}
			for i := 1; i < len(fields)-1; i += 4 {
				n1 := strings.ToUpper(fields[i])
				n2 := strings.ToUpper(fields[i+1])
				n3 := strings.ToUpper(fields[i+2])
				n4 := strings.ToUpper(fields[i+3])

				if isConnectionAtom(n1) || isConnectionAtom(n2) || isConnectionAtom(n3) || isConnectionAtom(n4) {
					continue
				}
				a1 := name_atom[n1]
				a2 := name_atom[n2]
				a3 := name_atom[n3]
				a4 := name_atom[n4]

				if a1 == nil || a2 == nil || a3 == nil || a4 == nil {
					return fmt.Errorf("one or more atoms were not found for improper => fragment: %s, line: %s", fg.Name(), line)
				}
				im := blocks.NewImproper(a1, a2, a3, a4)
				fg.AddImproper(im)
			}

		case strings.HasPrefix(head, "CMAP"):
			if fg == nil {
				continue
			}
			fields := strings.Fields(line)
			if (len(fields)-1)%8 != 0 {
				return fmt.Errorf("uneven number of atoms for cmap formation => %s", line)
			}
			for i := 1; i < len(fields)-1; i += 8 {
				n1 := strings.ToUpper(fields[i])
				n2 := strings.ToUpper(fields[i+1])
				n3 := strings.ToUpper(fields[i+2])
				n4 := strings.ToUpper(fields[i+3])
				n5 := strings.ToUpper(fields[i+4])
				n6 := strings.ToUpper(fields[i+5])
				n7 := strings.ToUpper(fields[i+6])
				n8 := strings.ToUpper(fields[i+7])

				if isConnectionAtom(n1) || isConnectionAtom(n2) || isConnectionAtom(n3) || isConnectionAtom(n4) {
					continue
				}
				if isConnectionAtom(n5) || isConnectionAtom(n6) || isConnectionAtom(n7) || isConnectionAtom(n8) {
					continue
				}
				a1 := name_atom[n1]
				a2 := name_atom[n2]
				a3 := name_atom[n3]
				a4 := name_atom[n4]
				a5 := name_atom[n5]
				a6 := name_atom[n6]
				a7 := name_atom[n7]
				a8 := name_atom[n8]

				if a1 == nil || a2 == nil || a3 == nil || a4 == nil || a5 == nil || a6 == nil || a7 == nil || a8 == nil {
					return fmt.Errorf("one or more atoms were not found for improper => %s", line)
				}
				cm := blocks.NewCMap(a1, a2, a3, a4, a5, a6, a7, a8)
				fg.AddCMap(cm)
			}

		case strings.HasPrefix(head, "DECL"):
		case strings.HasPrefix(head, "DEFA"):
		case strings.HasPrefix(head, "AUTO"):
		case strings.HasPrefix(head, "DELE"):
		case strings.HasPrefix(head, "DONO"):
		case strings.HasPrefix(head, "ACCE"):
		case strings.HasPrefix(head, "IC"):
		case strings.HasPrefix(head, "PATC"):
		case strings.HasPrefix(head, "BILD"):
		case strings.HasPrefix(head, "3") && len(strings.Fields(line)) <= 2:
			// assuming version number - improve it TODO

		default:
			return fmt.Errorf("Unrecognized line => %s", line)
		}

	}

	return nil
}

/**********************************************************
* Helpers
**********************************************************/

// removes comments plus leading and tailing spaces
func cleanTOPLine(s string) string {
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

func isConnectionAtom(aname string) bool {
	if strings.HasPrefix(aname, "+") || strings.HasPrefix(aname, "-") {
		return true
	}
	return false
}
