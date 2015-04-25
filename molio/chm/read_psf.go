package chm

import (
	//"bufio"
	//"errors"
	//"fmt"
	"io"
	"os"
	"strings"

	//"github.com/resal81/molkit/ff"
	"github.com/resal81/molkit/blocks"
)

type psfFileFormat int32

const (
	psf_format_XPLOR psfFileFormat = 1 << iota
)

/**********************************************************
* ReadPSFFile
**********************************************************/

func ReadPSFFile(fname string) (*blocks.System, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, nil
	}

	return readpsf(file)
}

/**********************************************************
* ReadPSFString
**********************************************************/

func ReadPSFString(s string) (*blocks.System, error) {

	return readpsf(strings.NewReader(s))
}

/**********************************************************
* readpsf
**********************************************************/

type psfLevel int64

type resData struct {
	resname string
	resnumb int
	segname string
}

func readpsf(reader io.Reader) (*blocks.System, error) {

	return nil, nil
	/*
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

		tmp_atoms_map := make(map[int64]*blocks.Atom, 0) // atom.Id() : *blocks.Atom
		tmp_resdata_map := make(map[int64]*resData, 0)   // atom.Id() : *resData

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
				a, rd, err := psfParseAtomsLine(line)
				if err != nil {
					return nil, fmt.Errorf("in line: {'%s'} - reason: {'%s'}", line, err)
				}

				tmp_atoms_map[a.Id()] = a
				tmp_resdata_map[a.Id()] = rd

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
	*/
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

/**********************************************************
* Line parsers
**********************************************************/

/*
func psfParseAtomsLine(line string) (*blocks.Atom, *resData, error) {
	fields := strings.Fields(line)
	if len(fields) != 9 {
		return nil, nil, errors.New("bad length in atoms line")
	}

	var aser, resnumb int
	var chg, mass float64
	var aname, atype, resname, segname, tmp string

	n, err := fmt.Sscanf(line,
		"%d %s %d %s %s %s %f %f %s",
		&aser, &segname, &resnumb, &resname, &aname, &atype, &chg, &mass, &tmp)
	if n != 9 {
		return nil, nil, errors.New("# of scanned fields is not 9")
	}
	if err != nil {
		return nil, nil, err
	}

	a := blocks.NewAtom()
	a.Name = aname
	a.Serial = aser

	at := blocks.NewAtomType()
	at.Label = atype
	at.Mass = mass
	at.ParCharge = chg

	at.Setting |= blocks.AT_HAS_MASS_SET
	at.Setting |= blocks.AT_HAS_PAR_CHARGE_SET

	a.Type = at

	rd := resData{
		resname: resname,
		resnumb: resnumb,
		segname: segname,
	}

	return a, &rd, nil
}

*/
