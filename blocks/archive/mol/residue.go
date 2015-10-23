package mol

// ==============================================

type Residue struct {
	serial   int
	fragment *Fragment
	Atoms    []*Atom
}

func NewResidue(serial int) *Residue {
	return &Residue{
		serial: serial,
		Atoms:  make([]*Atom),
	}
}
