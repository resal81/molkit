package blocks

/**********************************************************
* Link
**********************************************************/

type Link struct {
	conns [][2]*Atom
}

func NewLink() *Link {
	return &Link{}
}

func (ln *Link) Connect(a1, a2 *Atom) {
	ln.conns = append(ln.conns, [2]*Atom{a1, a2})
}
