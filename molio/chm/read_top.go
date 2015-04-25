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
func ReadTOPFiles(fnames ...string) (*blocks.ForceField, error) {

	frc := blocks.NewForceField(blocks.FF_TYPE_CHM)

	for _, fname := range fnames {
		file, err := os.Open(fname)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		err = readtop(file, frc)
		if err != nil {
			return nil, err
		}
	}

	return frc, nil

}

/**********************************************************
* ReadTOPString
**********************************************************/

// Parses a prm string (e.g. contents of a file).
func ReadTOPString(s string) (*blocks.ForceField, error) {

	frc := blocks.NewForceField(blocks.FF_TYPE_CHM)

	reader := strings.NewReader(s)

	err := readtop(reader, frc)

	return frc, err
}

/**********************************************************
* readtop
**********************************************************/

func readtop(reader io.Reader, frc *blocks.ForceField) error {

	var fg *blocks.Fragment                // a pointer to hold reference to the current fragment
	name_atom := map[string]*blocks.Atom{} // a map of atomname:*Atom for the current fragment

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		line = cleanTOPLine(line)

		if line == "" {
			continue
		}

		line = strings.ToUpper(line) // make all letters uppercase

		// break out of for loop if END keyword is found
		if line == "END" {
			break
		}

		switch {

		case strings.HasPrefix(line, "PRES"):
			// PRES is for caps
			// should ignore it before reaching RESI
			fg = nil

		case strings.HasPrefix(line, "RESI"):
			// RESI line can have 2 or 3 fields. If 3, the last element is the total charge
			// currently ingnoring total charge

			fields := strings.Fields(line)

			fg = blocks.NewFragment(fields[1])
			name_atom = map[string]*blocks.Atom{}
			frc.AddFragment(fg)

		case strings.HasPrefix(line, "GROU"):
			// GROUP line - ignore it

		case strings.HasPrefix(line, "ATOM"):

			// check for a valid fragment
			if fg == nil {
				continue
			}

			fields := strings.Fields(line)

			if len(fields) != 4 {
				return fmt.Errorf("ATOM line is not right => %s", line)
			}

			// new atom
			a := blocks.NewAtom(fields[1])

			// new atom type
			at := blocks.NewAtomType(fields[2], blocks.AT_TYPE_CHM_1)

			// parse partial charge
			ch, err := strconv.ParseFloat(fields[3], 64)
			if err != nil {
				return fmt.Errorf("could not parse charge in this line => %s", line)
			}

			at.SetPartialCharge(ch)

			// connect both
			a.SetType(at)

			// add to the current fragment and atom map
			fg.AddAtom(a)

			// store it in the name map
			name_atom[a.Name()] = a

		case strings.HasPrefix(line, "BOND") || strings.HasPrefix(line, "DOUB") || strings.HasPrefix(line, "TRIP"):

			// check for a valid fragment
			if fg == nil {
				continue
			}

			fields := strings.Fields(line)

			if (len(fields)-1)%2 != 0 {
				return fmt.Errorf("BOND line, bad number of entries => %s", line)
			}

			for i := 1; i < len(fields)-1; i += 2 {
				n1 := fields[i]
				n2 := fields[i+1]

				// check for connection specifiers
				if strings.HasPrefix(n1, "+") || strings.HasPrefix(n1, "-") {
					return fmt.Errorf("invalid connection specifier => %s", line)

				}

				// linker next
				if strings.HasPrefix(n2, "+") {
					if !fg.HasLinkerNext() {
						ln := blocks.NewLinker()
						fg.SetLinkerNext(ln)
					}

					a1 := name_atom[n1]
					a2 := blocks.NewAtom(n2[1:])
					b := blocks.NewBond(a1, a2)
					ln := fg.LinkerNext()
					ln.SetBond(b)

					continue
				}

				// linker previous
				if strings.HasPrefix(n2, "-") {
					if !fg.HasLinkerNext() {
						ln := blocks.NewLinker()
						fg.SetLinkerPrev(ln)
					}

					a1 := name_atom[n1]
					a2 := blocks.NewAtom(n2[1:])
					b := blocks.NewBond(a1, a2)
					ln := fg.LinkerPrev()
					ln.SetBond(b)

					continue
				}

				// find atoms and create a bond
				a1 := name_atom[n1]
				a2 := name_atom[n2]

				if a1 == nil || a2 == nil {
					return fmt.Errorf("BOND line, one or both atoms were not found => fragment: %s, line: %s", fg.Name(), line)
				}

				b := blocks.NewBond(a1, a2)
				fg.AddBond(b)
			}

		case strings.HasPrefix(line, "IMPR"):

			// check that fragment is valid
			if fg == nil {
				continue
			}

			fields := strings.Fields(line)

			if (len(fields)-1)%4 != 0 {
				return fmt.Errorf("IMPROPER line, bad number of entries => %s", line)
			}

			for i := 1; i < len(fields)-1; i += 4 {
				n1 := fields[i]
				n2 := fields[i+1]
				n3 := fields[i+2]
				n4 := fields[i+3]

				// check for connection specifiers
				var tmp string
				tmp = strings.Join([]string{n1, n2, n3, n4}, "_")

				if strings.Contains(tmp, "+") && strings.Contains(tmp, "-") {
					msg := fmt.Sprintf("both '+' and '-' specifiers were found in the improper line => %s (%s)", fg.Name(), line)
					return fmt.Errorf(msg)
				}

				// linker prev
				if strings.Contains(tmp, "-") {

					// check we have only one '-'
					if strings.Count(tmp, "-") != 1 {
						msg := fmt.Sprintf("more than one '-' specifier in the improper line => %s (%s)", fg.Name(), line)
						return fmt.Errorf(msg)
					}

					// check we have linker prev
					if !fg.HasLinkerPrev() {
						ln := blocks.NewLinker()
						fg.SetLinkerPrev(ln)
					}

					a1 := name_atom[n1]
					a2 := blocks.NewAtom(n2[1:])
					a3 := name_atom[n3]
					a4 := name_atom[n4]

					im := blocks.NewImproper(a1, a2, a3, a4)
					ln := fg.LinkerPrev()
					ln.SetImproper(im)

					continue
				}

				// linker next
				if strings.Contains(tmp, "+") {

					// check we have only one '+'
					if strings.Count(tmp, "+") != 1 {
						msg := fmt.Sprintf("more than one '+' specifier in the improper line => %s (%s)", fg.Name(), line)
						return fmt.Errorf(msg)
					}

					// check we have linker next
					if !fg.HasLinkerNext() {
						ln := blocks.NewLinker()
						fg.SetLinkerNext(ln)
					}

					a1 := name_atom[n1]
					a2 := name_atom[n2]
					a3 := blocks.NewAtom(n3[1:])
					a4 := name_atom[n4]

					im := blocks.NewImproper(a1, a2, a3, a4)
					ln := fg.LinkerNext()
					ln.SetImproper(im)

					continue
				}

				// normal case
				a1 := name_atom[n1]
				a2 := name_atom[n2]
				a3 := name_atom[n3]
				a4 := name_atom[n4]

				if a1 == nil || a2 == nil || a3 == nil || a4 == nil {
					return fmt.Errorf("not all atoms for the improper were found => %s (%s)", fg.Name(), line)
				}
				im := blocks.NewImproper(a1, a2, a3, a4)
				fg.AddImproper(im)
			}

		case strings.HasPrefix(line, "CMAP"):
			if fg == nil {
				continue
			}

			fields := strings.Fields(line)

			if (len(fields)-1)%8 != 0 {
				return fmt.Errorf("number of entries is not right for CMAP => %s (%s)", fg.Name(), line)
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

				// check for connection specifiers
				var tmp string
				tmp = strings.Join([]string{n1, n2, n3, n4, n5, n6, n7, n8}, "_")

				if strings.Contains(tmp, "+") || strings.Contains(tmp, "-") {
					if strings.Count(tmp, "+") != 1 || strings.Count(tmp, "-") != 1 {
						msg := fmt.Sprintf("more than one '+' and '-' specifiers were found for the cmap => %s (%s)", fg.Name(), line)
						return fmt.Errorf(msg)
					}

					// TODO read CMAP for connections

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
					return fmt.Errorf("not all atoms for the cmap were found => %s (%s)", fg.Name(), line)
				}
				cm := blocks.NewCMap(a1, a2, a3, a4, a5, a6, a7, a8)
				fg.AddCMap(cm)
			}

		case strings.HasPrefix(line, "MASS"):
		case strings.HasPrefix(line, "DECL"):
		case strings.HasPrefix(line, "DEFA"):
		case strings.HasPrefix(line, "AUTO"):
		case strings.HasPrefix(line, "DELE"):
		case strings.HasPrefix(line, "DONO"):
		case strings.HasPrefix(line, "ACCE"):
		case strings.HasPrefix(line, "IC"):
		case strings.HasPrefix(line, "PATC"):
		case strings.HasPrefix(line, "BILD"):

		case strings.HasPrefix(line, "3") && len(strings.Fields(line)) <= 2:
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
