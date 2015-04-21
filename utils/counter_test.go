package utils

import (
	"testing"
)

func TestComponentHash(t *testing.T) {

	c := NewComponentHash()

	if c.Length() != 0 {
		t.Fatalf("Length of a new Component hash must be zero. It is %d", c.Length())
	}

	i1 := c.Add(20)
	i2 := c.Add(40)

	if c.Length() != 2 {
		t.Fatalf("Length of hash should be two. It is %d", c.Length())
	}

	el1, err1 := c.Get(i1)
	el2, err2 := c.Get(i2)

	if err1 != nil {
		t.Fatalf(err1.Error())
	}

	if err2 != nil {
		t.Fatalf(err2.Error())
	}

	if el1.(int) != 20 {
		t.Fatalf("First element should be 20. It is %v", el1)
	}
	if el2.(int) != 40 {
		t.Fatalf("Second element should be 40. It is %v", el1)
	}

}
