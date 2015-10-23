package gmx

import (
	"math"
	"testing"
)

const DELTA = 1e-8

func TestMDPIntegrator(t *testing.T) {

	var vals = []struct {
		integrator Integrator
		nsteps     int
		dt         float64
	}{
		{INTEGRATOR_MD, 1000, 0.002},
	}

	for _, el := range vals {
		m := NewMDP()
		m.SetIntegrator(el.integrator, el.nsteps, el.dt)

		i, n, d := m.Integrator()

		if i&el.integrator == 0 {
			t.Errorf("wrong integrator => %v, expected %v", i, el.integrator)
		}

		if n != el.nsteps {
			t.Errorf("wrong nsteps => %d, expected %d", n, el.nsteps)
		}

		if math.Abs(d-el.dt) > DELTA {
			t.Errorf("wrong dt => %f, expected %f", d, el.dt)
		}
	}

}
