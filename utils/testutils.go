package utils

import (
	"testing"
)

func AssertNil(t *testing.T, err error, msg string) {
	if err != nil {
		t.Fatalf("%s : '%s'", msg, err)
	}
}

func CheckEqInt8(t *testing.T, v1, v2 int8, msg string) {
	if v1 != v2 {
		t.Errorf("%s : %d != %d", msg, v1, v2)
	}
}

func CheckEqInt(t *testing.T, i1, i2 int, msg string) {
	if i1 != i2 {
		t.Errorf("%s : %d != %d", msg, i1, i2)
	}
}

func CheckEqFloat64(t *testing.T, v1, v2 float64, msg string) {
	if v1 != v2 {
		t.Errorf("%s : %f != %f", msg, v1, v2)
	}
}

func CheckEqString(t *testing.T, v1, v2, msg string) {
	if v1 != v2 {
		t.Errorf("%s : %s != %s", msg, v1, v2)
	}
}

func CheckTrue(t *testing.T, cond bool, msg string) {
	if !cond {
		t.Errorf("%s", msg)
	}
}
