package blocks

import (
	"testing"
)

func TestConnection(t *testing.T) {
	var conns = []struct {
		f1name, f2name, a1name, a2name string
	}{
		{"ALA", "LYS", "C", "N"},
		{"CYS", "ARG", "S", "O"},
	}

	var i int
	ln := NewLink()

	for _, el := range conns {
		f1 := NewFragment(el.f1name)
		f2 := NewFragment(el.f2name)

		a1 := NewAtom(el.a1name)
		a2 := NewAtom(el.a2name)

		a1.SetFragment(f1)
		a2.SetFragment(f2)

		c := NewConnection(a1, a2)

		if v := c.Fragment1().Name(); v != el.f1name {
			t.Errorf("wrong f1 => %q, expected %q", v, el.f1name)
		}
		if v := c.Fragment2().Name(); v != el.f2name {
			t.Errorf("wrong f2 => %q, expected %q", v, el.f2name)
		}
		if v := c.Atom1().Name(); v != el.a1name {
			t.Errorf("wrong a1 => %q, expected %q", v, el.a1name)
		}
		if v := c.Atom2().Name(); v != el.a2name {
			t.Errorf("wrong a2 => %q, expected %q", v, el.a2name)
		}

		ln.AddConnection(c)
		i++

		if v := len(ln.Connections()); v != i {
			t.Errorf("wrong number of connections => %q, expected %q", v, i)
		}

	}

}
