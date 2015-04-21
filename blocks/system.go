package blocks

import (
	"fmt"
	"github.com/resal81/molkit/utils"
)

var (
	systemHash = utils.NewComponentHash()
)

type System struct {
	id int64

	Polymers []*Polymer
}

func NewSystem() *System {
	sys := &System{}
	sys.id = systemHash.Add(sys)
	return sys
}

func (s *System) deletePolymer(p1 *Polymer) {
	for i, p2 := range s.Polymers {
		if p1.Id() == p2.Id() {
			s.Polymers = append(s.Polymers[:i], s.Polymers[i+1:]...)
			return
		}
	}
}

func (s *System) Id() int64 {
	return s.id
}

func (s *System) Atoms() []*Atom {
	n := 0
	for _, poly := range s.Polymers {
		n += len(poly.Atoms())
	}

	var i int = 0
	out := make([]*Atom, n, n)
	for _, poly := range s.Polymers {
		for _, atm := range poly.Atoms() {
			out[i] = atm
			i++
		}
	}
	return out
}

func (s *System) Fragments() []*Fragment {
	n := 0
	for _, poly := range s.Polymers {
		n += len(poly.Fragments)
	}

	var i int = 0
	out := make([]*Fragment, n)
	for _, poly := range s.Polymers {
		for _, frag := range poly.Fragments {
			out[i] = frag
			i++
		}
	}
	return out
}

func (s *System) String() string {
	return fmt.Sprintf("<system with %d atoms>", len(s.Atoms()))
}
