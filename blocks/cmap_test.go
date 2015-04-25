package blocks

import (
	"testing"
)

func TestCMapType(t *testing.T) {
	var cmts = []struct {
		at1, at2, at3, at4, at5, at6, at7, at8 string
		ct                                     CMSetting
	}{
		{"C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", CT_TYPE_CHM_1},
	}

	for _, el := range cmts {
		cm := NewCMapType(el.at1, el.at2, el.at3, el.at4, el.at5, el.at6, el.at7, el.at8, el.ct)

		if v := cm.Setting(); v&el.ct == 0 {
			t.Errorf("wrong cmap type => %q, expected %q", v, el.ct)
		}

		if v := cm.AType1(); v != el.at1 {
			t.Errorf("wrong at1 => %q, %q", v, el.at1)
		}
		if v := cm.AType2(); v != el.at2 {
			t.Errorf("wrong at1 => %q, %q", v, el.at2)
		}
		if v := cm.AType3(); v != el.at3 {
			t.Errorf("wrong at1 => %q, %q", v, el.at3)
		}
		if v := cm.AType4(); v != el.at4 {
			t.Errorf("wrong at1 => %q, %q", v, el.at4)
		}
		if v := cm.AType5(); v != el.at5 {
			t.Errorf("wrong at1 => %q, %q", v, el.at5)
		}
		if v := cm.AType6(); v != el.at6 {
			t.Errorf("wrong at1 => %q, %q", v, el.at6)
		}
		if v := cm.AType7(); v != el.at7 {
			t.Errorf("wrong at1 => %q, %q", v, el.at7)
		}
		if v := cm.AType8(); v != el.at8 {
			t.Errorf("wrong at1 => %q, %q", v, el.at8)
		}
	}
}

func TestCMap(t *testing.T) {
	var cms = []struct {
		an1, an2, an3, an4, an5, an6, an7, an8 string
	}{
		{"C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8"},
	}

	for _, el := range cms {
		a1 := NewAtom(el.an1)
		a2 := NewAtom(el.an2)
		a3 := NewAtom(el.an3)
		a4 := NewAtom(el.an4)
		a5 := NewAtom(el.an5)
		a6 := NewAtom(el.an6)
		a7 := NewAtom(el.an7)
		a8 := NewAtom(el.an8)
		c := NewCMap(a1, a2, a3, a4, a5, a6, a7, a8)

		if v := c.Atom1().Name(); v != el.an1 {
			t.Errorf("wrong an1 => %q, expected %q", v, el.an1)
		}
		if v := c.Atom2().Name(); v != el.an2 {
			t.Errorf("wrong an2 => %q, expected %q", v, el.an2)
		}
		if v := c.Atom3().Name(); v != el.an3 {
			t.Errorf("wrong an3 => %q, expected %q", v, el.an3)
		}
		if v := c.Atom4().Name(); v != el.an4 {
			t.Errorf("wrong an4 => %q, expected %q", v, el.an4)
		}
		if v := c.Atom5().Name(); v != el.an5 {
			t.Errorf("wrong an5 => %q, expected %q", v, el.an5)
		}
		if v := c.Atom6().Name(); v != el.an6 {
			t.Errorf("wrong an6 => %q, expected %q", v, el.an6)
		}
		if v := c.Atom7().Name(); v != el.an7 {
			t.Errorf("wrong an7 => %q, expected %q", v, el.an7)
		}
		if v := c.Atom8().Name(); v != el.an8 {
			t.Errorf("wrong an8 => %q, expected %q", v, el.an8)
		}
	}
}
