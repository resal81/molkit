package mol

type Chain struct {
	chainName string
	Fragments []*Fragment
}

func NewChain(name string) *Chain {
	return &Chain{
		chainName: name,
		Fragments: make([]*Fragment),
	}
}
