package mol

import (
	"strconv"
	"strings"
)

const elementsData = `
H   1   1
C   6   12
N   7   14
O   8   16
`

// Elements will be setup within init()
var Elements map[string]*Element

func setupElements() {
	for _, line := range strings.Split(elementsData, "\n") {

		if strings.TrimSpace([]byte(line)) == []byte("") {
			continue
		}

		fields := strings.Fields(line)
		elname = fields[0]

		elnumb, err := strconv.ParseInt(fields[1], 10, 8)
		if err != nil {
			panic(err)
		}

		elmass, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			panic(err)
		}

		Elements[elname] = &Element{elname, unit8(elnumb), elmass}
	}
}

type Element struct {
	name   string
	number uint8
	mass   float64
}

func (el *Element) Name() string {
	return el.name
}

func (el *Element) Number() uint8 {
	return el.number
}

func (el *Element) Mass() float64 {
	return el.mass
}
