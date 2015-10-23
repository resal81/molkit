package element

import (
	"strconv"
	"strings"
)

type Element struct {
	symbol string
	number uint8
	mass   float64
}

func NewElement(symbol string, number uint8, mass float64) *Element {
	return &Element{
		symbol: symbol,
		number: number,
		mass:   mass,
	}
}

var ElementsDatabase = map[string]*Element{}

func init() {
	for _, line := range strings.Split(elementData, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)

		number, err := strconv.ParseInt(fields[0], 10, 8)
		if err != nil {
			panic("element.go: cannot parse int")
		}

		symbol := fields[1]

		mass, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			panic("element.go: cannot parse float")
		}

		el := NewElement(symbol, uint8(number), mass)
		ElementsDatabase[symbol] = el
	}
}
