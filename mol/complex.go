package mol

type Complex struct {
	complexName string
	Chains      []*Chain
}

func NewComplex(name string) *Complex {
	return &Complex{
		complexName: name,
		Chains:      make([]*Chain),
	}
}
