package geom

import (
	"github.com/gonum/matrix/mat64"
)

// MatrixColMean calculates the mean of each column.
func MatrixColMean(mat *mat64.Dense) []float64 {
	nrows, ncols := mat.Dims()
	out := make([]float64, ncols)

	for i := 0; i < nrows; i += 1 {
		for j := 0; j < ncols; j += 1 {
			out[j] += mat.At(i, j)
		}
	}

	for j := 0; j < ncols; j += 1 {
		out[j] /= float64(nrows)
	}

	return out
}

// MatrixColMinMax calculates the minimum and maximum of each column.
func MatrixColMinMax(mat *mat64.Dense) ([]float64, []float64) {
	nrows, ncols := mat.Dims()
	min := make([]float64, ncols)
	max := make([]float64, ncols)

	for i := 0; i < nrows; i += 1 {
		for j := 0; j < ncols; j += 1 {
			el := mat.At(i, j)
			if el < min[j] {
				min[j] = el
			}

			if el > max[j] {
				max[j] = el
			}
		}
	}

	return min, max
}

//
func SliceSubtract(A, B []float64) []float64 {
	if len(A) != len(B) {
		panic("The length of slices are not the same")
	}
	out := make([]float64, len(A))
	for i := range A {
		out[i] = A[i] - B[i]
	}

	return out
}
