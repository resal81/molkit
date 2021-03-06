package chm

import (
	"testing"

	"github.com/resal81/molkit/blocks"
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

func tempAtomTypeKey(at string) blocks.HashKey {
	atype := blocks.NewAtomType(at, blocks.AT_NULL)
	return blocks.GenerateHashKey(atype, blocks.HK_MODE_NORMAL)
}

func tempBondTypeKey(at1, at2 string) blocks.HashKey {
	btype := blocks.NewBondType(at1, at2, blocks.BT_NULL)
	return blocks.GenerateHashKey(btype, blocks.HK_MODE_NORMAL)
}

func tempAngleTypeKey(at1, at2, at3 string) blocks.HashKey {
	ntype := blocks.NewAngleType(at1, at2, at3, blocks.NT_NULL)
	return blocks.GenerateHashKey(ntype, blocks.HK_MODE_NORMAL)
}

func tempDihedralTypeKey(at1, at2, at3, at4 string) blocks.HashKey {
	dtype := blocks.NewDihedralType(at1, at2, at3, at4, blocks.DT_NULL)
	return blocks.GenerateHashKey(dtype, blocks.HK_MODE_NORMAL)
}

func tempImproperTypeKey(at1, at2, at3, at4 string) blocks.HashKey {
	dtype := blocks.NewImproperType(at1, at2, at3, at4, blocks.IT_NULL)
	return blocks.GenerateHashKey(dtype, blocks.HK_MODE_NORMAL)
}

func tempPairTypeKey(at1, at2 string) blocks.HashKey {
	ptype := blocks.NewPairType(at1, at2, blocks.PT_NULL)
	return blocks.GenerateHashKey(ptype, blocks.HK_MODE_NORMAL)
}

func tempCMapTypeKey(at1, at2, at3, at4, at5, at6, at7, at8 string) blocks.HashKey {
	ctype := blocks.NewCMapType(at1, at2, at3, at4, at5, at6, at7, at8, blocks.CT_NULL)
	return blocks.GenerateHashKey(ctype, blocks.HK_MODE_NORMAL)
}

func TestAtomTypes(t *testing.T) {
	s := `
    ATOMS
    MASS    52 C     12.01100 ! carbonyl C, peptide backbone
    MASS    81 OB    15.99900 ! carbonyl oxygen in acetic acid

    NONBONDED nbxmod  5 atom cdiel fshift vatom vdistance vfswitch -
    C      0.000000  -0.110000     2.000000 ! ALLOW   PEP POL ARO
    OB     0.000000  -0.120000     1.700000   0.000000  -0.120000     1.400000 ! ALLOW   PEP POL ARO
    `
	var vals = []struct {
		label                        string
		mass, lje, ljd, lje14, ljd14 float64
	}{
		{"C", 12.011, -0.11, 2.0, 0.0, 0.0},
		{"OB", 15.999, -0.12, 1.7, -0.12, 1.4},
	}

	frc, err := ReadPRMString(s)
	if err != nil {
		t.Fatal("could not read atom prm string => ", err)
	}

	ats := frc.AtomTypes()
	if v := len(ats); v != 2 {
		t.Errorf("# of atomtypes is wrong => %d, expected %d", v, 2)
	}

	for _, el := range vals {
		at := frc.AtomType(tempAtomTypeKey(el.label))
		if at == nil {
			t.Fatalf("atom type was not found for => %q", el.label)
		}

		if v := at.Mass(); v != el.mass {
			t.Errorf("mass is not right => %f, expected %f", v, el.mass)
		}
		if v := at.LJEnergy(); v != el.lje {
			t.Errorf("lje is not right => %f, expected %f", v, el.lje)
		}
		if v := at.LJDistance(); v != el.ljd {
			t.Errorf("ljd is not right => %f, expected %f", v, el.ljd)
		}

		if at.HasLJEnergy14Set() {
			if v := at.LJEnergy14(); v != el.lje14 {
				t.Errorf("lje14 is not right => %f, expected %f", v, el.lje14)
			}
			if v := at.LJDistance14(); v != el.ljd14 {
				t.Errorf("ljd14 is not right => %f, expected %f", v, el.ljd14)
			}
		}

	}
}

func TestBondTypes(t *testing.T) {
	s := `
    BONDS
    NH2   CT1   240.00      1.455  ! From LSN NH2-CT2
    CA   CA    305.000     1.3750 ! ALLOW   ARO
    `
	var vals = []struct {
		at1, at2 string
		kb, b0   float64
	}{
		{"NH2", "CT1", 240, 1.455},
		{"CA", "CA", 305, 1.375},
	}

	frc, err := ReadPRMString(s)
	if err != nil {
		t.Fatal("could not read bond prm string => ", err)
	}

	bts := frc.BondTypes()
	if v := len(bts); v != 2 {
		t.Errorf("# of bondtypes is wrong => %d, expected %d", v, 2)
	}

	for _, el := range vals {
		bt := frc.BondType(tempBondTypeKey(el.at1, el.at2))
		if bt == nil {
			t.Errorf("bond was not found for => %s_%s", el.at1, el.at2)
		}

		if v := bt.HarmonicConstant(); v != el.kb {
			t.Errorf("wrong kb => %f, expected %f", v, el.kb)
		}

		if v := bt.HarmonicDistance(); v != el.b0 {
			t.Errorf("wrong b0 => %f, expected %f", v, el.b0)
		}

	}

}

func TestAngleTyeps(t *testing.T) {
	s := `
    ANGLES
    !atom types     Ktheta    Theta0   Kub     S0
    CT3  CT1  CD    52.000    108.00              ! Ala cter
    NH2  CT1  HB    38.000    109.50   50.00   2.1400 ! From LSN NH2-CT2-HA
    `
	var vals = []struct {
		at1, at2, at3       string
		kt, theta, kub, r13 float64
	}{
		{"CT3", "CT1", "CD", 52, 108, 0, 0},
		{"NH2", "CT1", "HB", 38, 109.5, 50, 2.14},
	}

	frc, err := ReadPRMString(s)
	if err != nil {
		t.Fatal("could not read angle prm string => ", err)
	}

	angs := frc.AngleTypes()
	if v := len(angs); v != 2 {
		t.Errorf("# of angletypes is wrong => %d, expected %d", v, 2)
	}

	for _, el := range vals {
		at := frc.AngleType(tempAngleTypeKey(el.at1, el.at2, el.at3))
		if at == nil {
			t.Errorf("angle was not found for => %s_%s_%s", el.at1, el.at2, el.at3)
		}

		if v := at.Theta(); v != el.theta {
			t.Errorf("wrong theta => %f, expected %f", v, el.theta)
		}
		if v := at.ThetaConstant(); v != el.kt {
			t.Errorf("wrong kt => %f, expected %f", v, el.kt)
		}
		if at.HasR13Set() {
			if v := at.R13(); v != el.r13 {
				t.Errorf("wrong r13 => %f, expected %f", v, el.r13)
			}
			if v := at.UBConstant(); v != el.kub {
				t.Errorf("wrong kub => %f, expected %f", v, el.kub)
			}
		}
	}
}

func TestDihedralTyeps(t *testing.T) {
	s := `
    DIHEDRALS
    !atom types             Kchi    n   delta
    !Neutral N terminus
    NH2  CT1  C    O        0.0000  1     0.00
    NH2  CT1  C    NH1      5.0000  2     9.00
    X    CT3  OS   X       -0.1000  3     5.00 ! ALLOW   PEP POL
    `
	var vals = []struct {
		at1, at2, at3, at4 string
		kp, phi            float64
		mult               int
	}{
		{"NH2", "CT1", "C", "O", 0, 0, 1},
		{"NH2", "CT1", "C", "NH1", 5, 9, 2},
		{"X", "CT3", "OS", "X", -0.1, 5, 3},
	}

	frc, err := ReadPRMString(s)
	if err != nil {
		t.Fatal("could not read dihedral prm string => ", err)
	}

	dhs := frc.DihedralTypes()
	if v := len(dhs); v != 3 {
		t.Errorf("# of dihedraltypes is wrong => %d, expected %d", v, 3)
	}

	for _, el := range vals {
		dt := frc.DihedralType(tempDihedralTypeKey(el.at1, el.at2, el.at3, el.at4))
		if dt == nil {
			t.Errorf("dihedral was not found for => %s_%s_%s_%s", el.at1, el.at2, el.at3, el.at4)
		}

		if v := dt.Phi(); v != el.phi {
			t.Errorf("wrong phi => %f, expected %f", v, el.phi)
		}
		if v := dt.PhiConstant(); v != el.kp {
			t.Errorf("wrong kp => %f, expected %f", v, el.kp)
		}
		if v := dt.Multiplicity(); v != el.mult {
			t.Errorf("wrong mult => %f, expected %f", v, el.mult)
		}
	}
}

func TestImproperTyeps(t *testing.T) {
	s := `
    IMPROPER
    !atom types           Kpsi                   psi0
    HE2  HE2  CE2  CE2     3.0            0      5.00   !
    HR3  CPH1 NR2  CPH1    0.5000         0      1.0000 ! ALLOW ARO
    NC2  X    X    C      40.0000         0      0.0000 ! ALLOW   PEP POL ARO
    `

	var vals = []struct {
		at1, at2, at3, at4 string
		kp, psi            float64
	}{
		{"HE2", "HE2", "CE2", "CE2", 3, 5},
		{"HR3", "CPH1", "NR2", "CPH1", 0.5, 1},
		{"NC2", "X", "X", "C", 40, 0},
	}

	frc, err := ReadPRMString(s)
	if err != nil {
		t.Fatal("could not read improper prm string => ", err)
	}

	imps := frc.ImproperTypes()
	if v := len(imps); v != 3 {
		t.Errorf("# of impropertypes is wrong => %d, expected %d", v, 3)
	}

	for _, el := range vals {
		it := frc.ImproperType(tempImproperTypeKey(el.at1, el.at2, el.at3, el.at4))
		if it == nil {
			t.Errorf("improper was not found for => %s_%s_%s_%s", el.at1, el.at2, el.at3, el.at4)
		}

		if v := it.Psi(); v != el.psi {
			t.Errorf("wrong psi => %f, expected %f", v, el.psi)
		}
		if v := it.PsiConstant(); v != el.kp {
			t.Errorf("wrong kpsi => %f, expected %f", v, el.kp)
		}
	}

}

func TestNonBondedTyeps(t *testing.T) {
	s := `
    NBFIX
    SOD    OC       -0.075020   3.190 ! For prot
    LIT    OC2D2    -0.375020   5.190 ! For carb 
    `

	var vals = []struct {
		at1, at2 string
		ljd, lje float64
	}{
		{"SOD", "OC", 3.190, -0.075020},
		{"LIT", "OC2D2", 5.190, -0.375020},
	}

	frc, err := ReadPRMString(s)
	if err != nil {
		t.Fatal("could not read nonbonded prm string => ", err)
	}

	pts := frc.NonBondedTypes()
	if v := len(pts); v != 2 {
		t.Errorf("# of nonbondedtypes is wrong => %d, expected %d", v, 2)
	}

	for _, el := range vals {
		pt := frc.NonBondedType(tempPairTypeKey(el.at1, el.at2))
		if pt == nil {
			t.Errorf("pair was not found for => %s_%s", el.at1, el.at2)
		}

		if v := pt.LJDistance(); v != el.ljd {
			t.Errorf("wrong ljd => %f, expected %f", v, el.ljd)
		}
		if v := pt.LJEnergy(); v != el.lje {
			t.Errorf("wrong lje => %f, expected %f", v, el.lje)
		}
	}

}

func TestCMapTyeps(t *testing.T) {
	var vals = []struct {
		at1, at2, at3, at4, at5, at6, at7, at8 string
		vf, vl                                 float64
	}{
		{"C", "NH1", "CT1", "C", "NH1", "CT1", "C", "NH1", 0.126790, -2.231020},
		{"C", "NH1", "CT1", "C", "NH1", "CT1", "C", "N", 0.136790, -2.331020},
	}

	frc, err := ReadPRMString(cmap_string)
	if err != nil {
		t.Fatal("could not read cmpa prm string => ", err)
	}

	cmaps := frc.CMapTypes()
	if v := len(cmaps); v != 2 {
		t.Errorf("# of cmaptypes is wrong => %d, expected %d", v, 2)
	}

	for _, el := range vals {
		k := tempCMapTypeKey(el.at1, el.at2, el.at3, el.at4, el.at5, el.at6, el.at7, el.at8)
		c := frc.CMapType(k)
		if c == nil {
			t.Error("cmap was not found")
		}

		if v := len(c.Values()); v != 24*24 {
			t.Errorf("# of values in cmap is not right => %d, expected %d", v, 24*24)
		}
		if v := c.Values()[0]; v != el.vf {
			t.Errorf("cmpa[0] is wrong => %f, expected %f", v, el.vf)
		}
		if v := c.Values()[24*24-1]; v != el.vl {
			t.Errorf("cmpa[-1] is wrong => %f, expected %f", v, el.vl)
		}
	}

}

var cmap_string string = `
CMAP
C    NH1  CT1  C    NH1  CT1  C    NH1   24

!-180
0.126790 0.768700 0.971260 1.250970 2.121010
2.695430 2.064440 1.764790 0.755870 -0.713470
0.976130 -2.475520 -5.455650 -5.096450 -5.305850
-3.975630 -3.088580 -2.784200 -2.677120 -2.646060
-2.335350 -2.010440 -1.608040 -0.482250

!-165
-0.802290 1.377090 1.577020 1.872290 2.398990
2.461630 2.333840 1.904070 1.061460 0.518400
-0.116320 -3.575440 -5.284480 -5.160310 -4.196010
-3.276210 -2.715340 -1.806200 -1.101780 -1.210320
-1.008810 -0.637100 -1.603360 -1.776870

!-150
-0.634810 1.156210 1.624350 2.047200 2.653910
2.691410 2.296420 1.960450 1.324930 2.038290
-1.151510 -3.148610 -4.058280 -4.531850 -3.796370
-2.572090 -1.727250 -0.961410 -0.282910 -0.479120
-1.039340 -1.618060 -1.725460 -1.376360

!-135
0.214000 1.521370 1.977440 2.377950 2.929470
2.893410 2.435810 2.162970 1.761500 1.190090
-1.218610 -2.108900 -2.976100 -3.405340 -2.768440
-1.836030 -0.957950 0.021790 -0.032760 -0.665880
-1.321170 -1.212320 -0.893170 -0.897040

!-120
0.873950 1.959160 2.508990 2.841100 3.698960
3.309330 2.614300 2.481720 2.694660 1.082440
-0.398320 -1.761800 -2.945110 -3.294690 -2.308300
-0.855480 -0.087320 0.439040 0.691880 -0.586330
-1.027210 -0.976640 -0.467580 0.104020

!-105
1.767380 2.286650 2.818030 3.065500 3.370620
3.397440 2.730310 2.878790 2.542010 1.545240
-0.092150 -1.694440 -2.812310 -2.802430 -1.856360
-0.306240 -0.122440 0.444680 0.810150 -0.058630
-0.270290 -0.178830 0.202360 0.493810

!-90
1.456010 2.743180 2.589450 3.046230 3.451510
3.319160 3.052900 3.873720 2.420650 0.949100
0.008370 -1.382980 -2.138930 -2.087380 -1.268300
-0.494370 0.267580 0.908250 0.537520 0.306260
0.069540 0.097460 0.263060 0.603220

!-75
1.396790 3.349090 2.180920 2.942960 3.814070
3.675800 3.555310 3.887290 2.101260 -0.190940
-0.732240 -1.382040 -0.673880 -0.817390 -0.826980
-0.111800 0.053710 0.296400 0.692240 0.428960
-0.036100 -0.033820 -0.194300 0.400210

!-60
0.246650 1.229980 1.716960 3.168570 4.208190
4.366860 4.251080 3.348110 0.997540 -1.287540
-1.179900 -0.684300 -0.853660 -1.158760 -0.347550
0.114810 0.242800 0.322420 0.370140 -0.374950
-0.676940 -1.323430 -1.366650 -0.218770

!-45
-1.196730 0.078060 2.347410 4.211350 5.376000
5.364940 4.355200 2.436510 0.408470 -0.590840
-0.435960 -0.501210 -0.822230 -0.607210 0.057910
0.246580 -0.070570 0.379430 0.247770 -0.571680
-1.282910 -1.715770 -1.839820 -1.987110

!-30
-1.174720 1.067030 4.180460 6.741610 6.070770
4.781470 2.758340 1.295810 0.571150 -0.196480
0.251860 -0.732140 1.289360 1.497590 1.890550
2.198490 0.169290 0.534000 0.331780 -1.276320
-2.550070 -3.312150 -3.136670 -2.642260

!-15
0.293590 5.588070 3.732620 3.217620 3.272450
2.492320 1.563700 1.356760 0.831410 0.630170
1.591970 0.821920 0.486070 0.715760 0.996020
1.591580 -0.367400 0.181770 -0.613920 -2.267900
-3.516460 -3.597700 -3.043340 -1.765020

!0
2.832310 0.787990 0.323280 0.479230 0.628600
0.976330 1.238750 1.671950 1.645480 2.520340
1.606970 0.776350 0.119780 0.070390 0.121170
-1.569230 -1.213010 -1.846360 -2.744510 -3.792530
-3.934880 -3.615930 -2.675750 -0.924170

!15
-0.778340 -1.912680 -2.052140 -1.846280 -1.047430
0.183400 1.682950 2.223500 1.358370 2.448660
1.436920 0.678570 -0.237060 -0.535320 -0.790380
-2.182580 -3.251140 -4.195110 -4.269270 -3.908210
-3.455620 -2.773970 1.755370 0.313410

!30
-2.963810 -3.483730 -3.517080 -2.724860 -1.405510
0.336200 1.428450 1.394630 0.970370 2.462720
1.522430 0.553620 -0.407380 -1.482950 -3.613920
-4.159810 -4.945580 -4.784040 -3.764540 -2.959140
-1.963850 -1.071260 -1.599580 -2.445320

!45
-4.029070 -3.932660 -3.558480 -2.513980 -1.037320
0.362000 0.814380 0.754110 0.502370 1.903420
0.770220 -0.416420 -3.286310 -3.875270 -4.907800
-5.704430 -5.645660 -4.396040 -2.865450 -2.368170
-2.860490 -3.416560 -3.666490 -3.859070

!60
-3.338270 -2.960220 -2.311700 -1.272890 -0.246470
0.722610 0.668070 0.438130 2.395330 1.632470
-2.041450 -3.218100 -3.915080 -4.852510 -5.696500
-6.314370 -5.683690 -4.170620 -3.141000 -3.508820
-3.756430 -3.640810 -3.640430 -3.550690

!75
-2.244860 -1.632100 -1.000640 -0.170440 0.526440
0.823710 0.517140 -0.013120 -0.370910 -1.213720
-2.305650 -3.420580 -4.484960 -5.693140 -6.199150
-6.253870 -5.211310 -4.174380 -3.685150 -4.151360
-4.161970 -3.725150 -3.715310 -2.606760

!90
-1.720840 -1.177830 -0.428430 0.277730 0.807900
0.803260 0.482510 -0.336900 -0.786270 -1.774070
-2.793220 -3.828560 -5.211800 -6.636850 -6.989940
-6.108800 -5.452410 -3.911450 -4.321000 -4.587240
-4.102610 -3.772820 -3.157300 -2.648390

!105
-1.850640 -1.092420 -0.445020 0.128490 1.005520
0.884820 0.485850 -0.218470 -0.857670 -1.682330
-3.014400 -4.481110 -6.053510 -6.865400 -6.871130
-5.728240 -3.912230 -4.802110 -5.034640 -4.715990
-4.601080 -4.086220 -3.274630 -2.410940

!120
-1.969230 -1.116650 -0.540250 -0.150330 0.763520
1.038890 0.758480 0.313530 -0.333050 -1.872770
-3.366270 -5.008260 -6.124810 -7.034830 -6.724320
-3.700200 -4.510620 -5.185650 -5.361620 -4.847490
-4.444320 -4.004260 -3.415720 -2.751230

!135
-2.111250 -1.168960 -0.322790 -0.006920 0.316660
1.086270 0.939170 0.625340 -0.166360 -1.830310
-3.469470 -4.946030 -6.112560 -1.915580 -4.047310
-4.996740 -4.996730 -4.842690 -4.886620 -4.300540
-4.494620 -4.442210 -4.163570 -3.183510

!150
-1.757590 -0.403620 0.023920 0.362390 0.634520
1.264920 1.361360 0.948420 -0.073680 -1.483560
-3.152820 1.835120 -1.762860 -5.093660 -5.744830
-5.390070 -4.783930 -4.190630 -4.115420 -4.042280
-4.125570 -4.028550 -4.026100 -2.937910

!165
-0.810590 -0.071500 0.378890 0.543310 1.277880
1.641310 1.698840 1.519950 0.631950 -1.088670
-2.736530 -0.735240 -4.563830 -6.408350 -5.889450
-5.141750 -4.194970 -3.666490 -3.843450 -3.818830
-3.826180 -3.596820 -2.994790 -2.231020

! alanine before proline map
C    NH1  CT1  C    NH1  CT1  C    N     24

!-180
0.136790 0.768700 0.971260 1.250970 2.121010
2.695430 2.064440 1.764790 0.755870 -0.713470
0.976130 -2.475520 -5.455650 -5.096450 -5.305850
-3.975630 -3.088580 -2.784200 -2.677120 -2.646060
-2.335350 -2.010440 -1.608040 -0.482250

!-165
-0.802290 1.377090 1.577020 1.872290 2.398990
2.461630 2.333840 1.904070 1.061460 0.518400
-0.116320 -3.575440 -5.284480 -5.160310 -4.196010
-3.276210 -2.715340 -1.806200 -1.101780 -1.210320
-1.008810 -0.637100 -1.603360 -1.776870

!-150
-0.634810 1.156210 1.624350 2.047200 2.653910
2.691410 2.296420 1.960450 1.324930 2.038290
-1.151510 -3.148610 -4.058280 -4.531850 -3.796370
-2.572090 -1.727250 -0.961410 -0.282910 -0.479120
-1.039340 -1.618060 -1.725460 -1.376360

!-135
0.214000 1.521370 1.977440 2.377950 2.929470
2.893410 2.435810 2.162970 1.761500 1.190090
-1.218610 -2.108900 -2.976100 -3.405340 -2.768440
-1.836030 -0.957950 0.021790 -0.032760 -0.665880
-1.321170 -1.212320 -0.893170 -0.897040

!-120
0.873950 1.959160 2.508990 2.841100 3.698960
3.309330 2.614300 2.481720 2.694660 1.082440
-0.398320 -1.761800 -2.945110 -3.294690 -2.308300
-0.855480 -0.087320 0.439040 0.691880 -0.586330
-1.027210 -0.976640 -0.467580 0.104020

!-105
1.767380 2.286650 2.818030 3.065500 3.370620
3.397440 2.730310 2.878790 2.542010 1.545240
-0.092150 -1.694440 -2.812310 -2.802430 -1.856360
-0.306240 -0.122440 0.444680 0.810150 -0.058630
-0.270290 -0.178830 0.202360 0.493810

!-90
1.456010 2.743180 2.589450 3.046230 3.451510
3.319160 3.052900 3.873720 2.420650 0.949100
0.008370 -1.382980 -2.138930 -2.087380 -1.268300
-0.494370 0.267580 0.908250 0.537520 0.306260
0.069540 0.097460 0.263060 0.603220

!-75
1.396790 3.349090 2.180920 2.942960 3.814070
3.675800 3.555310 3.887290 2.101260 -0.190940
-0.732240 -1.382040 -0.673880 -0.817390 -0.826980
-0.111800 0.053710 0.296400 0.692240 0.428960
-0.036100 -0.033820 -0.194300 0.400210

!-60
0.246650 1.229980 1.716960 3.168570 4.208190
4.366860 4.251080 3.348110 0.997540 -1.287540
-1.179900 -0.684300 -0.853660 -1.158760 -0.347550
0.114810 0.242800 0.322420 0.370140 -0.374950
-0.676940 -1.323430 -1.366650 -0.218770

!-45
-1.196730 0.078060 2.347410 4.211350 5.376000
5.364940 4.355200 2.436510 0.408470 -0.590840
-0.435960 -0.501210 -0.822230 -0.607210 0.057910
0.246580 -0.070570 0.379430 0.247770 -0.571680
-1.282910 -1.715770 -1.839820 -1.987110

!-30
-1.174720 1.067030 4.180460 6.741610 6.070770
4.781470 2.758340 1.295810 0.571150 -0.196480
0.251860 -0.732140 1.289360 1.497590 1.890550
2.198490 0.169290 0.534000 0.331780 -1.276320
-2.550070 -3.312150 -3.136670 -2.642260

!-15
0.293590 5.588070 3.732620 3.217620 3.272450
2.492320 1.563700 1.356760 0.831410 0.630170
1.591970 0.821920 0.486070 0.715760 0.996020
1.591580 -0.367400 0.181770 -0.613920 -2.267900
-3.516460 -3.597700 -3.043340 -1.765020

!0
2.832310 0.787990 0.323280 0.479230 0.628600
0.976330 1.238750 1.671950 1.645480 2.520340
1.606970 0.776350 0.119780 0.070390 0.121170
-1.569230 -1.213010 -1.846360 -2.744510 -3.792530
-3.934880 -3.615930 -2.675750 -0.924170

!15
-0.778340 -1.912680 -2.052140 -1.846280 -1.047430
0.183400 1.682950 2.223500 1.358370 2.448660
1.436920 0.678570 -0.237060 -0.535320 -0.790380
-2.182580 -3.251140 -4.195110 -4.269270 -3.908210
-3.455620 -2.773970 1.755370 0.313410

!30
-2.963810 -3.483730 -3.517080 -2.724860 -1.405510
0.336200 1.428450 1.394630 0.970370 2.462720
1.522430 0.553620 -0.407380 -1.482950 -3.613920
-4.159810 -4.945580 -4.784040 -3.764540 -2.959140
-1.963850 -1.071260 -1.599580 -2.445320

!45
-4.029070 -3.932660 -3.558480 -2.513980 -1.037320
0.362000 0.814380 0.754110 0.502370 1.903420
0.770220 -0.416420 -3.286310 -3.875270 -4.907800
-5.704430 -5.645660 -4.396040 -2.865450 -2.368170
-2.860490 -3.416560 -3.666490 -3.859070

!60
-3.338270 -2.960220 -2.311700 -1.272890 -0.246470
0.722610 0.668070 0.438130 2.395330 1.632470
-2.041450 -3.218100 -3.915080 -4.852510 -5.696500
-6.314370 -5.683690 -4.170620 -3.141000 -3.508820
-3.756430 -3.640810 -3.640430 -3.550690

!75
-2.244860 -1.632100 -1.000640 -0.170440 0.526440
0.823710 0.517140 -0.013120 -0.370910 -1.213720
-2.305650 -3.420580 -4.484960 -5.693140 -6.199150
-6.253870 -5.211310 -4.174380 -3.685150 -4.151360
-4.161970 -3.725150 -3.715310 -2.606760

!90
-1.720840 -1.177830 -0.428430 0.277730 0.807900
0.803260 0.482510 -0.336900 -0.786270 -1.774070
-2.793220 -3.828560 -5.211800 -6.636850 -6.989940
-6.108800 -5.452410 -3.911450 -4.321000 -4.587240
-4.102610 -3.772820 -3.157300 -2.648390

!105
-1.850640 -1.092420 -0.445020 0.128490 1.005520
0.884820 0.485850 -0.218470 -0.857670 -1.682330
-3.014400 -4.481110 -6.053510 -6.865400 -6.871130
-5.728240 -3.912230 -4.802110 -5.034640 -4.715990
-4.601080 -4.086220 -3.274630 -2.410940

!120
-1.969230 -1.116650 -0.540250 -0.150330 0.763520
1.038890 0.758480 0.313530 -0.333050 -1.872770
-3.366270 -5.008260 -6.124810 -7.034830 -6.724320
-3.700200 -4.510620 -5.185650 -5.361620 -4.847490
-4.444320 -4.004260 -3.415720 -2.751230

!135
-2.111250 -1.168960 -0.322790 -0.006920 0.316660
1.086270 0.939170 0.625340 -0.166360 -1.830310
-3.469470 -4.946030 -6.112560 -1.915580 -4.047310
-4.996740 -4.996730 -4.842690 -4.886620 -4.300540
-4.494620 -4.442210 -4.163570 -3.183510

!150
-1.757590 -0.403620 0.023920 0.362390 0.634520
1.264920 1.361360 0.948420 -0.073680 -1.483560
-3.152820 1.835120 -1.762860 -5.093660 -5.744830
-5.390070 -4.783930 -4.190630 -4.115420 -4.042280
-4.125570 -4.028550 -4.026100 -2.937910

!165
-0.810590 -0.071500 0.378890 0.543310 1.277880
1.641310 1.698840 1.519950 0.631950 -1.088670
-2.736530 -0.735240 -4.563830 -6.408350 -5.889450
-5.141750 -4.194970 -3.666490 -3.843450 -3.818830
-3.826180 -3.596820 -2.994790 -2.331020

`
