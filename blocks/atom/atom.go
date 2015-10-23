package atom

type Atom struct {
	name   string
	serial uint64
}

func NewAtom(name string, serial uint64) *Atom {
	if name == "" {
		panic("Atom: name cannot be empty")
	}

	return &Atom{
		name:   name,
		serial: serial,
	}
}

func (a *Atom) Name() string {
	return a.name
}

func (a *Atom) Serial() uint64 {
	return a.serial
}
