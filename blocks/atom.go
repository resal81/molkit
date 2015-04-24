package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	atmHash = utils.NewComponentHash()
)

/**********************************************************
* AtomType
**********************************************************/

type ATSetting int64

const (
	AT_TYPE_CHM_1 ATSetting = 1 << iota // Normal CHARMM ATOM
	AT_TYPE_GMX_1                       // Normal GROMACS ATOM
	AT_HAS_PROTONS_SET
	AT_HAS_MASS_SET
	AT_HAS_LJ_DISTANCE_SET
	AT_HAS_LJ_ENERGY_SET
	AT_HAS_LJ_DISTANCE_14_SET
	AT_HAS_LJ_ENERGY_14_SET
	AT_HAS_CHARGE_SET
	AT_HAS_PARTIAL_CHARGE_SET
	AT_HAS_RADIUS_SET
)

type AtomType struct {
	label      string
	protons    int
	mass       float64
	ljDist     float64
	ljEnergy   float64
	ljDist14   float64
	ljEnergy14 float64
	charge     int
	parCharge  float64
	radius     float64
	setting    ATSetting
}

/* new atomtype */

func NewAtomType(label string) *AtomType {
	return &AtomType{
		label: label,
	}
}

/* label */

func (at *AtomType) Label() string {
	return at.label
}

/* protons */

func (at *AtomType) SetProtons(v int) {
	at.setting |= AT_HAS_PROTONS_SET
	at.protons = v
}

func (at *AtomType) HasProtonsSet() bool {
	return at.setting&AT_HAS_PROTONS_SET != 0
}

func (at *AtomType) Protons() int {
	return at.protons
}

/* mass */

func (at *AtomType) SetMass(v float64) {
	at.setting |= AT_HAS_MASS_SET
	at.mass = v
}

func (at *AtomType) HasMassSet() bool {
	return at.setting&AT_HAS_MASS_SET != 0
}

func (at *AtomType) Mass() float64 {
	return at.mass
}

/* ljDist */

func (at *AtomType) SetLJDistance(v float64) {
	at.setting |= AT_HAS_LJ_DISTANCE_SET
	at.ljDist = v
}

func (at *AtomType) HasLJDistanceSet() bool {
	return at.setting&AT_HAS_LJ_DISTANCE_SET != 0
}

func (at *AtomType) LJDistance() float64 {
	return at.ljDist
}

/* ljDist14 */

func (at *AtomType) SetLJDistance14(v float64) {
	at.setting |= AT_HAS_LJ_DISTANCE_14_SET
	at.ljDist14 = v
}

func (at *AtomType) HasLJDistance14Set() bool {
	return at.setting&AT_HAS_LJ_DISTANCE_14_SET != 0
}

func (at *AtomType) LJDistance14() float64 {
	return at.ljDist14
}

/* ljEnergy */

func (at *AtomType) SetLJEnergy(v float64) {
	at.setting |= AT_HAS_LJ_ENERGY_SET
	at.ljEnergy = v
}

func (at *AtomType) HasLJEnergySet() bool {
	return at.setting&AT_HAS_LJ_ENERGY_SET != 0
}

func (at *AtomType) LJEnergy() float64 {
	return at.ljEnergy
}

/* ljEnergy14 */

func (at *AtomType) SetLJEnergy14(v float64) {
	at.setting |= AT_HAS_LJ_ENERGY_14_SET
	at.ljEnergy14 = v
}

func (at *AtomType) HasLJEnergy14Set() bool {
	return at.setting&AT_HAS_LJ_ENERGY_14_SET != 0
}

func (at *AtomType) LJEnergy14() float64 {
	return at.ljEnergy14
}

/* charge */

func (at *AtomType) SetCharge(v int) {
	at.setting |= AT_HAS_CHARGE_SET
	at.charge = v
}

func (at *AtomType) HasChargeSet() bool {
	return at.setting&AT_HAS_CHARGE_SET != 0
}

func (at *AtomType) Charge() int {
	return at.charge
}

/* par charge */

func (at *AtomType) SetPartialCharge(v float64) {
	at.setting |= AT_HAS_PARTIAL_CHARGE_SET
	at.parCharge = v
}

func (at *AtomType) HasPartialChargeSet() bool {
	return at.setting&AT_HAS_PARTIAL_CHARGE_SET != 0
}

func (at *AtomType) PartialCharge() float64 {
	return at.parCharge
}

/* radius */

func (at *AtomType) SetRadius(v float64) {
	at.setting |= AT_HAS_RADIUS_SET
	at.radius = v
}

func (at *AtomType) HasRadiusSet() bool {
	return at.setting&AT_HAS_RADIUS_SET != 0
}

func (at *AtomType) Radius() float64 {
	return at.radius
}

/* setting */

func (at *AtomType) Setting() ATSetting {
	return at.setting
}

/**********************************************************
* Atom
**********************************************************/

type Atom struct {
	id        int64
	name      string
	serial    int64
	bFactor   float64
	occupancy float64
	altLoc    string
	isHetero  bool
	coords    [][3]float64
	tipe      *AtomType
	fragment  *Fragment
	bonds     []*Bond
}

/* new atom */

func NewAtom(name string) *Atom {
	at := &Atom{
		name: name,
	}
	id := atmHash.Add(at)
	at.id = id
	return at
}

/* id */

func (a *Atom) Id() int64 {
	return a.id
}

/* name */

func (a *Atom) Name() string {
	return a.name
}

/* serial */

func (a *Atom) SetSerial(s int64) {
	a.serial = s
}

func (a *Atom) Serial() int64 {
	return a.serial
}

/* bfactor */

func (a *Atom) SetBFactor(v float64) {
	a.bFactor = v
}

func (a *Atom) BFactor() float64 {
	return a.bFactor
}

/* occupancy */

func (a *Atom) SetOccupancy(v float64) {
	a.occupancy = v
}

func (a *Atom) Occupancy() float64 {
	return a.occupancy
}

/* altloc */

func (a *Atom) SetAltLoc(v string) {
	a.altLoc = v
}

func (a *Atom) AltLoc() string {
	return a.altLoc
}

/* ishetero */

func (a *Atom) SetHetero(v bool) {
	a.isHetero = v
}

func (a *Atom) IsHetero() bool {
	return a.isHetero
}

/* coords */
func (a *Atom) AddCoord(c [3]float64) {
	a.coords = append(a.coords, c)
}

func (a *Atom) Coords() [][3]float64 {
	return a.coords
}

/* type */

func (a *Atom) SetType(v *AtomType) {
	a.tipe = v
}

func (a *Atom) Type() *AtomType {
	return a.tipe
}

/* fragment */

func (a *Atom) SetFragment(v *Fragment) {
	a.fragment = v
}

func (a *Atom) Fragment() *Fragment {
	return a.fragment
}

/* bond */

func (a *Atom) AddBond(b *Bond) {
	a.bonds = append(a.bonds, b)
}

func (a *Atom) Bonds() []*Bond {
	return a.bonds
}
