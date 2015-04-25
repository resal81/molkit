package blocks

/*
	CMapType
*/

type CMSetting int64

const (
	CT_NULL CMSetting = 1 << iota
	CT_TYPE_CHM_1
	CT_TYPE_GMX_1
)

type CMapType struct {
	aType1  string
	aType2  string
	aType3  string
	aType4  string
	aType5  string
	aType6  string
	aType7  string
	aType8  string
	values  []float64
	setting CMSetting
}

func NewCMapType(at1, at2, at3, at4, at5, at6, at7, at8 string, t CMSetting) *CMapType {
	return &CMapType{
		aType1:  at1,
		aType2:  at2,
		aType3:  at3,
		aType4:  at4,
		aType5:  at5,
		aType6:  at6,
		aType7:  at7,
		aType8:  at8,
		setting: t,
	}
}

func (ct *CMapType) AType1() string {
	return ct.aType1
}

func (ct *CMapType) AType2() string {
	return ct.aType2
}

func (ct *CMapType) AType3() string {
	return ct.aType3
}

func (ct *CMapType) AType4() string {
	return ct.aType4
}

func (ct *CMapType) AType5() string {
	return ct.aType5
}

func (ct *CMapType) AType6() string {
	return ct.aType6
}

func (ct *CMapType) AType7() string {
	return ct.aType7
}

func (ct *CMapType) AType8() string {
	return ct.aType8
}

func (ct *CMapType) Setting() CMSetting {
	return ct.setting
}

func (ct *CMapType) SetValues(vs []float64) {
	ct.values = vs
}

func (ct *CMapType) Values() []float64 {
	return ct.values
}

/*
	CMap
*/

type CMap struct {
	atom1 *Atom
	atom2 *Atom
	atom3 *Atom
	atom4 *Atom
	atom5 *Atom
	atom6 *Atom
	atom7 *Atom
	atom8 *Atom
	tipe  *CMapType
}

func NewCMap(a1, a2, a3, a4, a5, a6, a7, a8 *Atom) *CMap {
	return &CMap{
		atom1: a1,
		atom2: a2,
		atom3: a3,
		atom4: a4,
		atom5: a5,
		atom6: a6,
		atom7: a7,
		atom8: a8,
	}
}

func (c *CMap) Atom1() *Atom {
	return c.atom1
}

func (c *CMap) Atom2() *Atom {
	return c.atom2
}

func (c *CMap) Atom3() *Atom {
	return c.atom3
}

func (c *CMap) Atom4() *Atom {
	return c.atom4
}

func (c *CMap) Atom5() *Atom {
	return c.atom5
}

func (c *CMap) Atom6() *Atom {
	return c.atom6
}

func (c *CMap) Atom7() *Atom {
	return c.atom7
}

func (c *CMap) Atom8() *Atom {
	return c.atom8
}

func (c *CMap) SetType(ct *CMapType) {
	c.tipe = ct
}

func (c *CMap) Type() *CMapType {
	return c.tipe
}
