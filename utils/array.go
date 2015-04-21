package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Make2DArray_f64(N int) [][]float64 {
	out := make([][]float64, N)
	for i := 0; i < N; i++ {
		out[i] = make([]float64, N)
	}

	return out
}

func Write2DArray_f64(fname string, arr [][]float64) error {
	str := ""
	n1 := len(arr)
	for i := 0; i < n1; i++ {

		n2 := len(arr[i])
		line := make([]string, n2)

		for j := 0; j < n2; j++ {
			line[j] = fmt.Sprintf("%.2f ", arr[i][j])
		}

		str += strings.Join(line, " ")
		str += "\n"
	}

	err := ioutil.WriteFile(fname, []byte(str), 0644)
	return err
}
