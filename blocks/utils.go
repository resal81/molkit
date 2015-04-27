package blocks

import (
	"fmt"
	"regexp"
	"strings"
)

type atomInfo struct {
	protons int
	reg     *regexp.Regexp
}

var PROTONS_LIB = map[string]*atomInfo{
	"H":  &atomInfo{1, regexp.MustCompile(`^\d*H[A-Z0-9]*$`)},
	"LI": &atomInfo{3, regexp.MustCompile(`^LI[A-Z0-9]*$`)},
	"C":  &atomInfo{6, regexp.MustCompile(`^C[A-Z0-9]*$`)},
	"N":  &atomInfo{7, regexp.MustCompile(`^N[A-Z0-9]*$`)},
	"O":  &atomInfo{8, regexp.MustCompile(`^O[A-Z0-9]*$`)},
	"F":  &atomInfo{9, regexp.MustCompile(`^F[A-Z0-9]*$`)},
	"NA": &atomInfo{11, regexp.MustCompile(`^NA\d*$`)},
	"MG": &atomInfo{12, regexp.MustCompile(`^MG\d*$`)},
	"AL": &atomInfo{13, regexp.MustCompile(`^AL[A-Z0-9]*$`)},
	"P":  &atomInfo{15, regexp.MustCompile(`^P[A-Z0-9]*$`)},
	"S":  &atomInfo{16, regexp.MustCompile(`^S[A-Z0-9]*$`)},
	"CL": &atomInfo{17, regexp.MustCompile(`^CL\d*$`)},
	"K":  &atomInfo{19, regexp.MustCompile(`^K\d*$`)},
	"CA": &atomInfo{20, regexp.MustCompile(`^CA\d*$`)},
	"FE": &atomInfo{26, regexp.MustCompile(`^FE\d*$`)},
	"ZN": &atomInfo{30, regexp.MustCompile(`^ZN\d*$`)},
	"BR": &atomInfo{35, regexp.MustCompile(`^BR[A-Z0-9]*$`)},
	"RB": &atomInfo{37, regexp.MustCompile(`^RB[A-Z0-9]*$`)},
	"I":  &atomInfo{53, regexp.MustCompile(`^I[A-Z0-9]*$`)},
	"BA": &atomInfo{56, regexp.MustCompile(`^BA[A-Z0-9]*$`)},
}

func AtomNameToProtons(name1 string) (int, error) {

	name2 := strings.ToUpper(name1)

	// Calcium and C-alpha are both CA
	if name2 == "CA" {
		return 0, fmt.Errorf("CA is ambiguous - determine is based on context")
	}

	if name2 == "CL" || name2 == "NA" || name2 == "FE" {
		return PROTONS_LIB[name2].protons, nil
	}

	if name2 == "RUB" {
		return PROTONS_LIB["RB"].protons, nil
	}

	if name2 == "BAR" {
		return PROTONS_LIB["BA"].protons, nil
	}

	if len(name2) == 1 {
		if _, ok := PROTONS_LIB[name2]; !ok {
			return 0, fmt.Errorf("protons for atom name not found => %s", name1)
		}
		return PROTONS_LIB[name2].protons, nil
	}

	for _, v := range PROTONS_LIB {
		if v.reg.Match([]byte(name2)) {
			return v.protons, nil
		}
	}

	return 0, fmt.Errorf("protons not found for => %s", name1)

}
