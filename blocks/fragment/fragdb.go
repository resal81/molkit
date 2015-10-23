package fragment

import (
	"strconv"
	"strings"
)

// The default database is instantiated using the data file
var DefaultFragmentDatabase *FragmentDatabase

func init() {
	DefaultFragmentDatabase = ParseRtpContent(defaultFragmentData)
}

// ***********************************************************************//
// FragmentDatabase struct
// ***********************************************************************//

type FragmentDatabase struct {
	fragments map[string]*Fragment
}

func NewFragmentDatabase() *FragmentDatabase {
	return &FragmentDatabase{
		fragments: make(map[string]*Fragment),
	}
}

func (fd *FragmentDatabase) AddFragment(fg *Fragment) {
	fd.fragments[fg.name] = fg
}

func (fd *FragmentDatabase) Size() int {
	return len(fd.fragments)
}

func (fd *FragmentDatabase) GetFragmentByName(name string) *Fragment {
	return fd.fragments[name]
}

// ***********************************************************************//
// Internal atom representation
// ***********************************************************************//

type atom struct {
	name   string
	kind   string
	charge float64
}

// ***********************************************************************//
// Internal bond representation
// ***********************************************************************//

type bond struct {
	a1 string
	a2 string
}

// ***********************************************************************//
// Fragment struct
// ***********************************************************************//

type Fragment struct {
	name  string
	atoms []*atom
	bonds []*bond
}

func NewFragment(name string) *Fragment {
	return &Fragment{
		name:  name,
		atoms: make([]*atom, 0),
		bonds: make([]*bond, 0),
	}
}

func (fr *Fragment) addAtom(a *atom) {
	fr.atoms = append(fr.atoms, a)
}

func (fr *Fragment) addBond(b *bond) {
	fr.bonds = append(fr.bonds, b)
}

func (fr *Fragment) Atoms() []*atom {
	return fr.atoms
}

func (fr *Fragment) Bonds() []*bond {
	return fr.bonds
}

// ***********************************************************************//
// RTP parser
// ***********************************************************************//

func ParseRtpContent(content string) *FragmentDatabase {
	var fdb = NewFragmentDatabase()

	var currFrag *Fragment
	var currSection string

	for _, line := range strings.Split(content, "\n") {
		line = stripComment(line)

		if strings.TrimSpace(line) == "" {
			continue
		}

		switch {
		case strings.HasPrefix(line, "["):
			fragName := getBracketField(line)
			currFrag = NewFragment(fragName)
			fdb.AddFragment(currFrag)

		case strings.HasPrefix(line, " ["):
			field := getBracketField(line)
			switch field {
			case "atoms":
				currSection = "atoms"
			case "bonds":
				currSection = "bonds"
			default:
				currSection = "other"
			}

		default:
			switch currSection {
			case "atoms":
				fields := strings.Fields(line)
				aname := fields[0]
				akind := fields[1]

				charge, err := strconv.ParseFloat(fields[2], 64)
				if err != nil {
					panic("ParseRtpContent: could not parse float")
				}

				at := &atom{name: aname, kind: akind, charge: charge}
				currFrag.addAtom(at)

			case "bonds":
				fields := strings.Fields(line)
				b := &bond{a1: fields[0], a2: fields[1]}
				currFrag.addBond(b)

			default:
			}

		}
	}

	return fdb
}

// ***********************************************************************//
// Helpers
// ***********************************************************************//

// stripComment removes Gromacs-style comments
func stripComment(line string) string {
	if strings.Contains(line, ";") {
		i := strings.Index(line, ";")
		return line[0:i]
	}
	return line
}

// getBracketField returns the field inside a Gromacs-style field
func getBracketField(line string) string {
	if strings.Contains(line, "[") && strings.Contains(line, "]") {
		i := strings.Index(line, "[")
		j := strings.Index(line, "]")

		out := line[i+1 : j-1]
		return strings.TrimSpace(out)
	}

	return line
}
