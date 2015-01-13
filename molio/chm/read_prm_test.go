package chm

import (
	"testing"

	"github.com/resal81/molkit/ff"
	"github.com/resal81/molkit/utils"
)

func TestPRMRead(t *testing.T) {
	fnames := []string{
		"../../testdata/chm_prm/par_all22_prot.prm",
		"../../testdata/chm_prm/par_all35_ethers.prm",
		"../../testdata/chm_prm/par_all36_prot.prm",
		"../../testdata/chm_prm/par_all36_na.prm",
		"../../testdata/chm_prm/par_all36_lipid.prm",
		"../../testdata/chm_prm/par_all36_carb.prm",
		"../../testdata/chm_prm/par_all36_cgenff.prm",
		"../../testdata/chm_prm/wat_ion.prm",
	}
	_, err := ReadPRMFiles(fnames...)
	if err != nil {
		t.Fatalf("could not read prm file -> %s", err)
	}
}

func TestAtomTypes(t *testing.T) {
	s := `
    ATOMS
    MASS    52 C     12.01100 ! carbonyl C, peptide backbone
    MASS    57 CPH1  12.01100 ! his CG and CD2 carbons
    MASS    66 CS    12.01100 ! thiolate carbon
    MASS    67 CE1   12.01100 ! for alkene; RHC=CR
    MASS    81 OB    15.99900 ! carbonyl oxygen in acetic acid

    NONBONDED nbxmod  5 atom cdiel fshift vatom vdistance vfswitch -
    C      0.000000  -0.110000     2.000000 ! ALLOW   PEP POL ARO
    CE1    0.000000  -0.068000     2.090000 ! 
    CPH1   0.000000  -0.050000     1.800000 ! ALLOW ARO
    CS     0.000000  -0.110000     2.200000 ! ALLOW SUL
    OB     0.000000  -0.120000     1.700000   0.000000  -0.120000     1.400000 ! ALLOW   PEP POL ARO
    `

	frc, err := ReadPRMString(s)
	utils.AssertNotNil(t, err, "could not read prm string")

	ats := frc.AtomTypes()
	utils.CheckEqInt(t, len(ats), 5, "AtomTypes()")
	utils.CheckEqFloat64(t, ats[0].Mass(), 12.01100, "mass is not right")
	utils.CheckEqFloat64(t, ats[1].LJDist(ff.FF_CHARMM), 2.09, "rmin/2 is not right")
	utils.CheckEqFloat64(t, ats[2].LJEnergy(ff.FF_CHARMM), -0.05, "epsilon is not right")
	utils.CheckEqFloat64(t, ats[4].LJDist14(ff.FF_CHARMM), 1.4, "rmin/2 for 1-4 interaction is not right")
	utils.CheckEqFloat64(t, ats[4].LJEnergy14(ff.FF_CHARMM), -0.12, "epsilon for 1-4 interaction is not right")

	utils.CheckTrue(t, ats[0].HasLJDistSet(), "ats[0] should have LJDist")
	utils.CheckTrue(t, ats[0].HasLJEnergySet(), "ats[0] should have LJEnergy")
	utils.CheckTrue(t, ats[0].HasMassSet(), "ats[0] should have Mass")
	utils.CheckTrue(t, !ats[0].HasLJDist14Set(), "ats[0] should not have LJDist14")
	utils.CheckTrue(t, !ats[0].HasLJEnergy14Set(), "ats[0] should not have LJEnergy14")
	utils.CheckTrue(t, !ats[0].HasChargeSet(), "ats[0] should not have Charge")
	utils.CheckTrue(t, !ats[0].HasProtonsSet(), "ats[0] should not have Protons")

	utils.CheckTrue(t, ats[4].HasLJDistSet(), "ats[4] should have LJDist")
	utils.CheckTrue(t, ats[4].HasLJEnergySet(), "ats[4] should have LJEnergy")
	utils.CheckTrue(t, ats[4].HasLJDist14Set(), "ats[4] should have LJDist14")
	utils.CheckTrue(t, ats[4].HasLJEnergy14Set(), "ats[4] should have LJEnergy14")

}

func TestBondTypes(t *testing.T) {
	s := `
    BONDS
    NH2   CT1   240.00      1.455  ! From LSN NH2-CT2
    CA   CA    305.000     1.3750 ! ALLOW   ARO
    `
	frc, err := ReadPRMString(s)
	utils.AssertNotNil(t, err, "could not read prm string")

	bts := frc.BondTypes()
	utils.CheckEqInt(t, len(bts), 2, "the length of BondTypes is not right")
	utils.CheckEqFloat64(t, bts[0].HarmonicConstant(ff.FF_CHARMM), 240.0, "bts[0] kb is not right")
	utils.CheckEqFloat64(t, bts[1].HarmonicDistance(ff.FF_CHARMM), 1.375, "bts[0] b0 is not right")

	utils.CheckTrue(t, bts[0].HasHarmonicConstantSet(), "ats[0] should have kb set")
	utils.CheckTrue(t, bts[0].HasHarmonicDistanceSet(), "ats[0] should have b0 set")

}
