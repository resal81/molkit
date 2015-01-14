package chm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/resal81/molkit/ff"
)

type psfFileFormat int32

const (
	psf_format_XPLOR psfFileFormat = 1 << iota
)

/**********************************************************
* ReadPSFFile
**********************************************************/

func ReadPSFFile(fname string) (*ff.TopSystem, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, nil
	}

	return readpsf(file)
}

/**********************************************************
* ReadPSFString
**********************************************************/

func ReadPSFString(s string) (*ff.TopSystem, error) {

	return readpsf(strings.NewReader(s))
}

/**********************************************************
* readpsf
**********************************************************/

type psfLevel int64

func readpsf(reader io.Reader) (*ff.TopSystem, error) {

	const (
		L_TITLE psfLevel = 1 << iota
		L_ATOMS
		L_BONDS
		L_ANGLES
		L_DIHEDRALS
		L_IMPROPERS
		L_CMPAS
		L_DONORS
		L_ACCEPTORS
		L_NNB
		L_NS2
		L_NUMLPH
		L_MOLNT
		L_NGRP
	)

	var format psfFileFormat
	var lvl psfLevel

	// var curr_pol *ff.TopPolymer
	// pol_list := make([]*ff.TopPolymer, 1)
	// topSys := ff.NewTopSystem()

	var lineno int64

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		lineno++

		if lineno == 1 {
			if strings.HasPrefix(line, "PSF") {
				// TODO check for other formats
				format = psf_format_XPLOR
			} else {
				return nil, fmt.Errorf("first line doesn't start with PSF")
			}
			continue
		}

		line = cleanPSFLine(line, format)

		if line == "" {
			continue
		}

		if strings.Index(line, "!") != -1 {
			switch {
			case strings.Index(line, "NTITLE") != -1:
				lvl = L_TITLE
			case strings.Index(line, "NATOM") != -1:
				lvl = L_ATOMS
			case strings.Index(line, "NBOND") != -1:
				lvl = L_BONDS
			case strings.Index(line, "NTHETA") != -1:
				lvl = L_ANGLES
			case strings.Index(line, "NPHI") != -1:
				lvl = L_DIHEDRALS
			case strings.Index(line, "NIMPHI") != -1:
				lvl = L_IMPROPERS
			case strings.Index(line, "NDON") != -1:
				lvl = L_DONORS
			case strings.Index(line, "NACC") != -1:
				lvl = L_ACCEPTORS
			case strings.Index(line, "NNB") != -1:
				lvl = L_NNB
			case strings.Index(line, "NST2") != -1:
				lvl = L_NS2
			case strings.Index(line, "MOLNT") != -1:
				lvl = L_MOLNT
			case strings.Index(line, "NUMLPH") != -1:
				lvl = L_NUMLPH
			case strings.Index(line, "NCRTERM") != -1:
				lvl = L_CMPAS
			case strings.Index(line, "NGRP") != -1:
				lvl = L_NGRP
			default:
				return nil, fmt.Errorf("unknown keyword: in line '%s", line)
			}

			continue
		}

		switch lvl {
		case L_TITLE:
		case L_ATOMS:
			_, err := psfParseAtomsLine(line)
			if err != nil {
				return psfLogError("in line: {'%s'} - reason: {'%s'}", line, err)
			}

		case L_BONDS:
		case L_ANGLES:
		case L_DIHEDRALS:
		case L_IMPROPERS:
		case L_CMPAS:
		case L_DONORS:
		case L_ACCEPTORS:
		case L_NNB:
		case L_NS2:
		case L_NUMLPH:
		case L_MOLNT:
		case L_NGRP:
		default:
			return nil, errors.New("unknow psf level")
		}
	}

	return nil, nil
}

/**********************************************************
* Helpers
**********************************************************/

//
func cleanPSFLine(line string, format psfFileFormat) string {
	if format&psf_format_XPLOR != 0 {
	}

	return line
}

//
func psfLogError(format string, args ...interface{}) (*ff.TopSystem, error) {
	return nil, fmt.Errorf(format, args...)
}

/**********************************************************
* Line parsers
**********************************************************/

func psfParseAtomsLine(line string) (*ff.TopAtom, error) {
	return nil, nil
}
