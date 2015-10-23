package blocks

/**********************************************************
* Connection
**********************************************************/

// A Connection is made of two atoms.
type Connection struct {
	frag1 *Fragment
	frag2 *Fragment
	atom1 *Atom
	atom2 *Atom
}

func NewConnection(atom1, atom2 *Atom) *Connection {
	return &Connection{
		frag1: atom1.Fragment(),
		frag2: atom2.Fragment(),
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

// A Link can have one or more Connection.
type Link struct {
	conns []*Connection
}

func NewLink() *Link {
	return &Link{}
}

func (ln *Link) AddConnection(c *Connection) {
	ln.conns = append(ln.conns, c)
}

func (ln *Link) Connections() []*Connection {
	return ln.conns
}

/**********************************************************
* Linker
**********************************************************/

type Linker struct {
	frag1    *Fragment
	frag2    *Fragment
	bond     *Bond
	improper *Improper
}

func NewLinker() *Linker {
	return &Linker{}
}

func (l *Linker) SetFragment1(frag1 *Fragment) {
	l.frag1 = frag1
}

func (l *Linker) SetFragment2(frag2 *Fragment) {
	l.frag2 = frag2
}

func (l *Linker) Fragment1() *Fragment {
	return l.frag1
}

func (l *Linker) Fragment2() *Fragment {
	return l.frag2
}

func (l *Linker) SetBond(b *Bond) {
	l.bond = b
}

func (l *Linker) Bond() *Bond {
	return l.bond
}

func (l *Linker) SetImproper(m *Improper) {
	l.improper = m
}

func (l *Linker) Improper() *Improper {
	return l.improper
}
