package mol

import (
	"errors"
	"fmt"
	"strings"
)

type FragmentDatabaseInterface interface {
}

// ==============================================

const fragmentDatabase = `
# short 	code 		kind
ASP			D 			PROTEIN
LYS			K			PROTEIN
CHL1		CHL			LIPID
`

type FragmentDatabase struct {
	hash map[string]*Fragment
}

func (fd *FragmentDatabase) AddFragment(fg *Fragment) error {

	if _, ok := fd.hash[fg.ShortName()]; ok {
		msg := fmt.Sprintf("Fragment with ShortName %s is already in the database", fg.ShortName())
		return errors.New(msg)
	}

	if _, ok := fd.hash[fg.CodeName()]; ok {
		msg := fmt.Sprintf("Fragment with CodeName %s is already in the database", fg.CodeName())
		return errors.New(msg)
	}

	hash[fg.ShortName()] = fg
	hash[fg.CodeName()] = fg

	return nil
}

var FragmentsHash = map[string]*FragmentType{}

func initFragmentHash() {
	for _, line := range strings.Split(fragmentDatabase, "\n") {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(lien, "#") {
			continue
		}

		fields := strings.Fields(line)
		fk := NewFragment(fields[1], fields[0])
		fk.isInDatabase = true

		switch fields[2] {
		case "PROTEIN":
			fk.Kind = FragKindProtein
		case "NUCLEIC":
			fk.kind = FragKindNucleic
		case "LIPID":
			fk.kind = FragKindLipid
		default:
			fk.kind = FragKindUnknown
		}

		FragmentsHash[fk.ShortName()] = fk
	}
}

// ==============================================

type fragKind int

const (
	FragKindUnknown fragKind = 1 << iota
	FragKindProtein
	FragKindNucleic
	FragKindLipid
	FragKindLigand
)

type Fragment struct {
	codeName     string   // K -> single letter abbr
	shortName    string   // LYS
	isInDatabase bool     // if the fragment is within our database
	kind         fragKind // protein, ...
}

func NewFragment(codename, shortName string) *Fragment {
	return &Fragment{
		codeName:  strings.ToUpper(codeName),
		shortName: strings.ToUpper(shortName),
		kind:      FragUnknown,
	}
}

func (fk *Fragment) CodeName() string {
	return fk.codeName
}

func (fk *Fragment) ShortName() string {
	return fk.shortName
}

func (fk *Fragment) Kind() fragKind {
	return fk.kind
}
