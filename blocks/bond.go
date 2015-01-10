package blocks

import (
	"sync/atomic"
)

var (
	bondid_counter int64 = 0
)

type Bond struct {
	id int64

	atom1 *Atom
	atom2 *Atom
}

func NewBond() *Bond {
	id := atomic.AddInt64(&bondid_counter, 1)
	return &Bond{
		id: id,
	}
}

func (b *Bond) Id() int64 {
	return b.id
}
