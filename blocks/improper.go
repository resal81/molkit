package blocks

import (
	"sync/atomic"
)

var (
	improperid_counter int64 = 0
)

type Improper struct {
	id int64

	atom1, atom2, atom3, atom4 *Atom
}

func NewImproper() *Improper {
	id := atomic.AddInt64(&improperid_counter, 1)
	return &Improper{
		id: id,
	}
}

func (m *Improper) Id() int64 {
	return m.id
}
