package blocks

/*
	CMapType
*/

type CMSetting int64

const (
	CM_TYPE_CHM_1 CMSetting = 1 << iota
)

type CMapType struct {
	NX     int
	NY     int
	Values []string
	AType1 string
	AType2 string
	AType3 string
	AType4 string
	AType5 string
	AType6 string
	AType7 string
	AType8 string
}

/*
	CMap
*/

type CMap struct {
	Atom1 *Atom
	Type  *CMapType
}
