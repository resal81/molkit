package blocks

import (
	"testing"
)

func TestGMXSetup(t *testing.T) {
	var gs = []struct {
		nbf      int
		combrule int
		genpairs bool
		flj      float64
		fqq      float64
	}{
		{1, 2, true, 0.5, 0.8},
	}

	for _, el := range gs {
		g := NewGMXSetup(el.nbf, el.combrule, el.genpairs, el.flj, el.fqq)

		if v := g.NbFunc(); v != el.nbf {
			t.Errorf("wrong nbFunc => %q, expected %q", v, el.nbf)
		}
		if v := g.CombinationRule(); v != el.combrule {
			t.Errorf("wrong combRule => %q, expected %q", v, el.combrule)
		}
		if v := g.GeneratePairs(); v != el.genpairs {
			t.Errorf("wrong genPairs => %q, expected %q", v, el.genpairs)
		}
		if v := g.FudgeLJ(); v != el.flj {
			t.Errorf("wrong fudgeLJ => %q, expected %q", v, el.flj)
		}
		if v := g.FudgeQQ(); v != el.fqq {
			t.Errorf("wrong fudgeQQ => %q, expected %q", v, el.fqq)
		}
	}

}

func TestForceField(t *testing.T) {

}
