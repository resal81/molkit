package blocks

import (
	"fmt"
	"github.com/resal81/molkit/utils"
)

var (
	systemHash = utils.NewComponentHash()
)

/*
	System
*/

type System struct {
	id        int64
	Atoms     []*Atom
	Bonds     []*Bond
	Angles    []*Angle
	Dihedrals []*Dihedral
	Impropers []*Improper
	Links     []*Link
}

func NewSystem() *System {
	sys := &System{}
	sys.id = systemHash.Add(sys)
	return sys
}

func (s *System) Id() int64 {
	return s.id
}

func (s *System) String() string {
	return fmt.Sprintf("<system with %d atoms>", len(s.Atoms))
}
