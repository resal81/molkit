package blocks

/**********************************************************
* Connection
**********************************************************/

type Connection struct {
	frag1 *Fragment
	frag2 *Fragment
	atom1 *Atom
	atom2 *Atom
}

func NewConnection(atom1, atom2 *Atom) *Connection {
	return &Connection{
		frag1: atom1.Fragment,
		frag2: atom2.Fragment,
		atom1: atom1,
		atom2: atom2,
	}
}

func (c *Connection) Fragment1() *Fragment {
	return c.frag1
}

func (c *Connection) Fragment2() *Fragment {
	return c.frag2
}

func (c *Connection) Atom1() *Atom {
	return c.atom1
}

func (c *Connection) Atom2() *Atom {
	return c.atom2
}

/**********************************************************
* Link
**********************************************************/

type Link struct {
	conns []*Connection
}

func NewLink() *Link {
	return &Link{}
}

func (ln *Link) AddConnection(c *Connection) {
	ln.conns = append(ln.conns, c)
}
