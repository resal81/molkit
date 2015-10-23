package pdb

import (
	"bufio"
	"github.com/resal81/molkit/blocks/atom"
	"github.com/resal81/molkit/blocks/chain"
	"github.com/resal81/molkit/blocks/residue"
	"github.com/resal81/molkit/blocks/structure"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadPdbFile(fname string) (*structure.Structure, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return readPdb(file)
}

func ReadPdbString(str string) (*structure.Structure, error) {
	reader := strings.NewReader(str)
	return readPdb(reader)
}

func readPdb(reader io.Reader) (*structure.Structure, error) {

	st := structure.NewStructure()
	firstModel := true                       // specifies the first model in an NMR-stye PDB
	atomCounter := 0                         // specifies the index for the current atom. Is resetted per model
	atomsList := []*atom.Atom{}              // a temporary list of atoms
	prevResName := "NONE"                    //
	var prevResSerial int64 = 0              //
	prevChainName := "NONE"                  //
	var currResidue *residue.Residue         //
	var currChain *chain.Chain               //
	var chainMap = map[string]*chain.Chain{} // temporarily holds reference to the chains

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "ENDMDL") {
			firstModel = false
			atomCounter = 0
		}

		if strings.HasPrefix(line, "ATOM") || strings.HasPrefix(line, "HETATM") {
			if firstModel {
				var x, y, z, beta, occ float64
				var aname, rname, altloc, chname string
				var aserial, rserial int64
				var err error

				if aserial, err = ii(line[6:11]); err != nil {
					return nil, err
				}
				aname = strings.TrimSpace(line[12:16])
				altloc = strings.TrimSpace(line[16:17])
				rname = strings.TrimSpace(line[17:21])
				chname = strings.TrimSpace(line[21:22])

				if rserial, err = ii(line[22:26]); err != nil {
					return nil, err
				}

				if x, err = ff(line[30:38]); err != nil {
					return nil, err
				}
				if y, err = ff(line[38:46]); err != nil {
					return nil, err
				}
				if z, err = ff(line[46:54]); err != nil {
					return nil, err
				}

				if occ, err = ff(line[54:60]); err != nil {
					return nil, err
				}
				if beta, err = ff(line[60:66]); err != nil {
					return nil, err
				}

				at := atom.NewAtom(aname, aserial)
				at.SetPdbAltLoc(altloc)
				at.SetPdbOccupancy(occ)
				at.SetPdbBeta(beta)
				at.AddCoord([]float64{x, y, z})

				if chname != prevChainName {
					if ch, ok := chainMap[chname]; ok {
						currChain = ch
					} else {
						currChain = chain.NewChain(chname)
						chainMap[chname] = currChain
						st.AddChain(currChain)
					}

					prevChainName = chname
				}

				if rname != prevResName || rserial != prevResSerial {
					currResidue = residue.NewResidue(rname, rserial)
					prevResName = rname
					prevResSerial = rserial
					currChain.AddResidue(currResidue)
				}

				currResidue.AddAtom(at)
				atomsList = append(atomsList, at)

			} else {
				var x, y, z float64
				var err error

				if x, err = ff(line[30:38]); err != nil {
					return nil, err
				}
				if y, err = ff(line[38:46]); err != nil {
					return nil, err
				}
				if z, err = ff(line[46:54]); err != nil {
					return nil, err
				}

				at := atomsList[atomCounter]
				at.AddCoord([]float64{x, y, z})
			}

			atomCounter += 1
		}

	}

	return st, nil
}

// ff strips the space and converts it to float64
func ff(s string) (float64, error) {
	s = strings.TrimSpace(s)
	return strconv.ParseFloat(s, 64)
}

// ii strips the space and converts it to int
func ii(s string) (int64, error) {
	s = strings.TrimSpace(s)
	return strconv.ParseInt(s, 10, 64)
}
