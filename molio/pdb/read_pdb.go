package pdb

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/resal81/molkit/blocks"
)

func ReadPDBFile(fname string) (*blocks.System, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return readpdb(file)

}

func readpdb(reader io.Reader) (*blocks.System, error) {

	// setting
	var first_model bool = true
	var line_index int64 = 0
	var prev_rname string = "--"
	var prev_rserial int64 = -1
	var prev_chain string = "--"

	lineIndexToAtomMap := map[int64]*blocks.Atom{}
	chainLetterToChainMap := map[string]*blocks.Polymer{}

	// initial values
	system := blocks.NewSystem()
	var chain *blocks.Polymer
	var residue *blocks.Fragment

	// read file line by line
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "ENDMDL") {
			first_model = false
			line_index = 0
		}

		if strings.HasPrefix(line, "ATOM") || strings.HasPrefix(line, "HETATM") {

			if first_model {
				atm, extra, err := parseAtomLineFull(line)
				if err != nil {
					return nil, err
				}

				// check for new chain
				if extra.chain != prev_chain {
					if ch, ok := chainLetterToChainMap[extra.chain]; ok {
						chain = ch
					} else {
						// fmt.Println("new chain", extra.chain)
						chain = blocks.NewPolymer()
						system.AddPolymer(chain)
						chainLetterToChainMap[extra.chain] = chain
					}

					// update the prev_chain
					prev_chain = extra.chain
				}

				// check for new residue
				if extra.rname != prev_rname || extra.rserial != prev_rserial {
					// fmt.Println("new residue", extra.rname, extra.rserial)

					residue = blocks.NewFragment()
					residue.SetName(extra.rname)
					residue.SetSerial(extra.rserial)
					chain.AddFragment(residue)

					// update the prev residue
					prev_rname = extra.rname
					prev_rserial = extra.rserial
				}

				residue.AddAtom(atm)
				// fmt.Println(len(residue.Atoms()))

				// store a pointer to atom for a multi-model pdb file
				lineIndexToAtomMap[line_index] = atm
				line_index++

			} else {
				// add the coordinate to a previously known atom
				coord, err := parseAtomLineCoord(line)
				if err != nil {
					return nil, err
				}

				if atm, ok := lineIndexToAtomMap[line_index]; ok {
					atm.AddCoord(coord[0], coord[1], coord[2])
					line_index++
				} else {
					return nil, fmt.Errorf("could not find atom matching lineIndexAtomMap[%d]", line_index)
				}

			} // first_model

		} // HETATOM or ATOM

	} // for loop

	return system, nil
}

type extraAtomData struct {
	rserial int64
	rname   string
	chain   string
}

func parseAtomLineFull(line string) (*blocks.Atom, *extraAtomData, error) {
	if len(line) < 67 {
		return nil, nil, fmt.Errorf("atom line too short: %s", line)
	}
	/*
	   %6s     0:6     flag
	   %5d     6:11    atom number
	   space   11:12
	   %4s     12:16   atom name
	   %1s     16:17   altloc
	   %4s     17:21   res name
	   %1s     21:22   chain
	   %4d     22:26   res number
	   space   26:30
	   %8.3f   30:38   x
	   %8.3f   38:46   y
	   %8.3f   46:54   z
	   %6.2f   54:60   occ
	   %6.2f   60:66   bf

	   HETATM    1  C   ACE A   0      37.266  12.061  15.716  1.00 82.39           C
	   ssssssDDDDD-ssssSssssSdddd----ffffffffFFFFFFFFffffffffFFFFFFffffff
	   01234567890123456789012345678901234567890123456789012345678901234567890123456789
	   0         1         2         3         4         5         6
	*/

	flag := strings.TrimSpace(line[0:6])
	aname := strings.TrimSpace(line[12:16])
	altloc := strings.TrimSpace(line[16:17])
	rname := strings.TrimSpace(line[17:21])
	chain := strings.TrimSpace(line[21:22])

	aserial, err_aserial := strconv.ParseInt(strings.TrimSpace(line[6:11]), 10, 64)
	if err_aserial != nil {
		aserial, err_aserial = strconv.ParseInt(strings.TrimSpace(line[6:11]), 16, 64)
	}

	rserial, err_rserial := strconv.ParseInt(strings.TrimSpace(line[22:26]), 10, 64)
	x, err_x := strconv.ParseFloat(strings.TrimSpace(line[30:38]), 64)
	y, err_y := strconv.ParseFloat(strings.TrimSpace(line[38:46]), 64)
	z, err_z := strconv.ParseFloat(strings.TrimSpace(line[46:54]), 64)
	occ, err_occ := strconv.ParseFloat(strings.TrimSpace(line[54:60]), 32)
	bf, err_bf := strconv.ParseFloat(strings.TrimSpace(line[60:66]), 32)

	if err_aserial != nil || err_rserial != nil {
		return nil, nil, fmt.Errorf("problem atom or residue serial in line: %s\n", line)
	}

	if err_x != nil || err_y != nil || err_z != nil {
		return nil, nil, fmt.Errorf("problem atom coordinates in line: %s\n", line)
	}

	if err_occ != nil || err_bf != nil {
		return nil, nil, fmt.Errorf("problem atom occupancy or bfactor in line: %s\n", line)
	}

	a := blocks.NewAtom()
	if flag == "HETATM" {
		a.SetPropIsHetero(true)
	}
	a.SetName(aname)
	a.SetSerial(aserial)
	a.SetPropAltloc(altloc)
	a.SetPropOccupancy(float32(occ))
	a.SetPropBFactor(float32(bf))
	a.AddCoord(x, y, z)

	// err := molio.AtomicNumberForAtom(a)
	// if err != nil {
	// 	return nil, nil, err
	// }

	return a, &extraAtomData{rserial: rserial, rname: rname, chain: chain}, nil

}

func parseAtomLineCoord(line string) ([3]float64, error) {
	x, err_x := strconv.ParseFloat(strings.TrimSpace(line[30:38]), 64)
	y, err_y := strconv.ParseFloat(strings.TrimSpace(line[38:46]), 64)
	z, err_z := strconv.ParseFloat(strings.TrimSpace(line[46:54]), 64)

	if err_x != nil || err_y != nil || err_z != nil {
		return [3]float64{0, 0, 0}, fmt.Errorf("problem atom coordinates in line: %s\n", line)
	}

	return [3]float64{x, y, z}, nil
}
