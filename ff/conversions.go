package ff

import (
	"math"
)

/**********************************************************
* Helpers
**********************************************************/

func convertSigma(sig float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return sig
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// self.charmm['param']['ljl'] * 2 * 0.1 / (2**(1.0/6.0))
			return sig * 2 * 0.1 / math.Pow(2.0, 1.0/6.0)
		default:
			panic("not implemented")
		}

	default:
		panic("not implemented")
	}
}

func convertEpsilon(eps float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return eps
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// abs(self.charmm['param']['lje']) * 4.184
			return math.Abs(eps) * 4.184
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}

}

func convertHarmonicConstant(kb float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return kb
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// self.charmm['param']['kb'] * 2 * 4.184 * (1.0 / 0.01)   # nm^2
			return kb * 2 * 4.184 * 100
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}

func convertHarmonicDistance(b0 float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return b0
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// self.charmm['param']['b0'] * 0.1
			return b0 * 0.1
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}

func convertThetaConstant(kt float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return kt
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// self.charmm['param']['ktetha'] * 2 * 4.184
			return kt * 2 * 4.184
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}

func convertTheta(theta float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return theta
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// self.charmm['param']['tetha0']
			return theta
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}

func convertUBConstant(kub float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return kub
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// self.charmm['param']['kub'] * 2 * 4.184 * 10 * 10
			return kub * 2 * 4.184 * 100
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}

func convertR13(r13 float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return r13
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// self.charmm['param']['s0'] * 0.1
			return r13 * 0.1
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}

func convertPhiConstant(kphi float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return kphi
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// dih['kchi'] * 4.184
			return kphi * 4.184
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}

func convertPhi(phi float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return phi
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// dih['delta']
			return phi
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}

func convertMutl(mult int8, from, to ffTypes) int8 {
	if from&to != 0 {
		return mult
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// dih['n']
			return mult
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}

func convertPsiConstant(kpsi float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return kpsi
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// imp['kpsi'] * 2 * 4.184
			return kpsi * 2 * 4.184
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}

func convertPsi(psi float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return psi
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			// imp['psi0']
			return psi
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}

func convertCMap(cm float64, from, to ffTypes) float64 {
	if from&to != 0 {
		return cm
	}

	switch from {
	case FF_CHARMM:
		switch to {
		case FF_GROMACS:
			return cm
		default:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}
}