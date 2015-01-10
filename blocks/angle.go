package blocks

import (
	"sync/atomic"
)

var (
	angleid_counter int64 = 0
)

type Angle struct {
	id int64

	atom1, atom2, atom3 *Atom
}

func NewAngle() *Angle {
	id := atomic.AddInt64(&angleid_counter, 1)
	return &Angle{
		id: id,
	}
}

func (a *Angle) Id() int64 {
	return a.id
}
