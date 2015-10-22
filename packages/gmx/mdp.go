package gmx

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

/* ==================================================================
	MDP struct
/ ================================================================== */

type MDP struct {
	defines []string

	// run
	integrator Integrator
	nsteps     int
	dt         float64
	comm_mode  CommMode
	nstcomm    int
	comm_grps  []string

	// em
	emtol  float64
	emstep float64

	// ld
	bd_fric float64
	ld_seed int

	// log
	nstxout              int
	nstvout              int
	nstfout              int
	nstlog               int
	nstcalcenergy        int
	nstenergy            int
	nstxout_compressed   int
	compressed_precision int

	// neighbor
	cutoff_scheme CutOffScheme
	nstype        NsType
	pbc           Pbc
	nstlist       int
	rlist         float64

	// coulomb
	coulombtype CoulombType
	rcoulomb    float64

	// vdw
	vdwtype  VdwType
	rvdw     float64
	dispcorr DispCorr

	// temperature
	tcoupl  TCouple
	tc_grps []string
	tau_t   float64
	ref_t   float64

	// pressure
	pcoupl           PCouple
	pcoupltype       PCoupleType
	tau_p            float64
	compressibility  float64
	ref_p            float64
	refcoord_scaling RefCoordScaling

	// constraint
	constraint           Constraint
	constraint_algorithm ConstAlg

	// pull
	pull               PullMode
	pull_geometry      PullGeometry
	pull_dim           [3]bool
	pull_r0            float64
	pull_r1            float64
	pull_start         bool
	pull_nstxout       int
	pull_nstfout       int
	pull_ngroups       int
	pull_ncoords       int
	pull_group1_name   string
	pull_group2_name   string
	pull_coord1_groups []int
	pull_coord1_rate   float64
}

/* ==================================================================
	Accessors and convertors
/ ================================================================== */

/*******************************/
/* initializer */
/*******************************/

func NewMDP() *MDP {
	return &MDP{}
}

/*******************************/
/* define */
/*******************************/

func convertDefines(m *MDP) string {
	if len(m.defines) == 0 {
		return "; nothing set"
	}

	return strings.Join(m.defines, " ")
}

/*******************************/
/* integrator */
/*******************************/

type Integrator int16

const (
	INTEGRATOR_NOTSET Integrator = 1 << iota
	INTEGRATOR_MD
	INTEGRATOR_STEEP
	INTEGRATOR_CG
)

func (m *MDP) SetIntegrator(n Integrator, nsteps int, dt float64) {
	m.integrator = n
	m.nsteps = nsteps
	m.dt = dt
}

func (m *MDP) Integrator() (Integrator, int, float64) {
	return m.integrator, m.nsteps, m.dt
}

func convertIntegrator(m *MDP) string {
	if m.integrator&INTEGRATOR_NOTSET != 0 {
		return "; integrator not set -- using defaults"
	}

	int_map := map[Integrator]string{
		INTEGRATOR_MD:    "md",
		INTEGRATOR_STEEP: "steep",
		INTEGRATOR_CG:    "cg",
	}

	// check we have the integrator
	if _, found := int_map[m.integrator]; !found {
		return "; << error : unknown integrator >> "
	}

	s := []string{
		makeMDPEntry("integrator", fmt.Sprintf("%s", int_map[m.integrator])),
		makeMDPEntry("nsteps", fmt.Sprintf("%d", m.nsteps)),
		makeMDPEntry("dt", fmt.Sprintf("%.4f", m.dt)),
		makeMDPEntry("emtol", fmt.Sprintf("%.1f", m.emtol)),
		makeMDPEntry("emstep", fmt.Sprintf("%.1f", m.emstep)),
	}

	return strings.Join(s, "\n")
}

/*******************************/
/* comm */
/*******************************/

type CommMode int8

const (
	COMM_MODE_NOTSET  CommMode = 1 << iota
	COMM_MODE_LINEAR           // COM translation
	COMM_MODE_ANGULAR          // COM translation and rotation
	COMM_MODE_NONE             // no restriction
)

func (m *MDP) SetCOMMode(mode CommMode, nstcomm int, comm_grps []string) {
	m.comm_mode = mode
	m.nstcomm = nstcomm
	m.comm_grps = comm_grps
}

func (m *MDP) COMMode() (CommMode, int, []string) {
	return m.comm_mode, m.nstcomm, m.comm_grps
}

func (m *MDP) convertCOMMode() string {
	if m.comm_mode&COMM_MODE_NOTSET != 0 {
		return "; comm_mode not set -- using defaults"
	}

	com_map := map[CommMode]string{
		COMM_MODE_LINEAR:  "linear",
		COMM_MODE_ANGULAR: "angular",
		COMM_MODE_NONE:    "none",
	}

	s := []string{
		makeMDPEntry("comm-mode", com_map[m.comm_mode]),
		makeMDPEntry("nstcomm", fmt.Sprintf("%d", m.nstcomm)),
		makeMDPEntry("comm-grps", strings.Join(m.comm_grps, " ")),
	}

	return strings.Join(s, "\n")
}

/*******************************/
/* log */
/*******************************/

func (m *MDP) SetLogFrequency(nstxout, nstvout, nstfout, nstlog, nstcalcenergy, nstenergy, nxtc, nxtc_prec int) {
	m.nstxout = nstxout
	m.nstvout = nstvout
	m.nstfout = nstfout
	m.nstcalcenergy = nstcalcenergy
	m.nstenergy = nstenergy
	m.nstxout_compressed = nxtc
	m.compressed_precision = nxtc_prec

}

func convertLog(m *MDP) string {
	s := []string{
		makeMDPEntry("nstlog", fmt.Sprintf("%d", m.nstlog)),
		makeMDPEntry("nstxout", fmt.Sprintf("%d", m.nstxout)),
		makeMDPEntry("nstvout", fmt.Sprintf("%d", m.nstfout)),
		makeMDPEntry("nstfout", fmt.Sprintf("%d", m.nstvout)),
		makeMDPEntry("nstenergy", fmt.Sprintf("%d", m.nstenergy)),
		makeMDPEntry("nstcalcenergy", fmt.Sprintf("%d", m.nstcalcenergy)),
		makeMDPEntry("nstxout-compressed", fmt.Sprintf("%d", m.nstxout_compressed)),
	}

	return strings.Join(s, "\n")
}

/*******************************/
/* neighbor */
/*******************************/

type CutOffScheme int8

const (
	CUTOFF_SCHEME_NOTSET CutOffScheme = 1 << iota
	CUTOFF_SCHEME_VERLET
	CUTOFF_SCHEME_GROUP
)

type NsType int8

const (
	NSTYPE_NOTSET NsType = 1 << iota
	NSTYPE_GRID
	NSTYPE_SIMPLE
)

type Pbc int8

const (
	PBC_NOTSET Pbc = 1 << iota
	PBC_XYZ
	PBC_NO
	PBC_XY
)

func (m *MDP) SetNeighborSearch(cuttoff CutOffScheme, nstype NsType, pbc Pbc, nstlist int, rlist float64) {
	m.cutoff_scheme = cuttoff
	m.nstype = nstype
	m.pbc = pbc
	m.nstlist = nstlist
	m.rlist = rlist
}

func (m *MDP) NeighborSearch() {

}

func convertNeighborSearch(m *MDP) string {
	if m.nstype&NSTYPE_NOTSET != 0 {
		return "; nstype not set -- using defaults"
	}

	cuttoff_map := map[CutOffScheme]string{
		CUTOFF_SCHEME_NOTSET: "",
		CUTOFF_SCHEME_VERLET: "verlet",
		CUTOFF_SCHEME_GROUP:  "group",
	}

	nstype_map := map[NsType]string{
		NSTYPE_NOTSET: "",
		NSTYPE_SIMPLE: "simple",
		NSTYPE_GRID:   "grid",
	}

	pbc_map := map[Pbc]string{
		PBC_NOTSET: "",
		PBC_NO:     "no",
		PBC_XYZ:    "xyz",
		PBC_XY:     "xy",
	}

	s := []string{
		makeMDPEntry("cuttoff-scheme", cuttoff_map[m.cutoff_scheme]),
		makeMDPEntry("ns-type", nstype_map[m.nstype]),
		makeMDPEntry("pbc", pbc_map[m.pbc]),
		makeMDPEntry("nstlist", fmt.Sprintf("%d", m.nstlist)),
		makeMDPEntry("rlist", fmt.Sprintf("%.2f", m.rlist)),
	}

	return strings.Join(s, "\n")
}

/*******************************/
/* coulomb */
/*******************************/

type CoulombType int16

const (
	COULOMBTYPE_NOTSET CoulombType = 1 << iota
	COULOMBTYPE_CUTOFF
	COULOMBTYPE_PME
)

func (m *MDP) SetElectrostatics() {

}

func (m *MDP) Electrostatics() {

}

func convertElectrostatics(m *MDP) string {
	if m.coulombtype&COULOMBTYPE_NOTSET != 0 {
		return "; coulombtype not set -- using defaults"
	}

	return ""
}

/*******************************/
/* vdw */
/*******************************/

type VdwType int8

const (
	VDWTYPE_NOTSET VdwType = 1 << iota
	VDWTYPE_CUTOFF
	VDWTYPE_PME
	VDWTYPE_SWITCH
)

type DispCorr int8

const (
	DISPCORR_NOTSET DispCorr = 1 << iota
	DISPCORR_NO
	DISPCORR_ENERPRES
	DISPCORR_ENER
)

func (m *MDP) SetVdw() {

}

func (m *MDP) Vdw() {

}

func convertVdw(m *MDP) string {
	if m.vdwtype&VDWTYPE_NOTSET != 0 {
		return "; vdwtype not set -- using defaults"
	}
	return ""
}

/*******************************/
/* tcoupl */
/*******************************/

type TCouple int16

const (
	TCOUPL_NOTSET TCouple = 1 << iota
	TCOUPL_NO
	TCOUPL_BERENDSEN
	TCOUPL_NOSE_HOOVER
	TCOUPL_ANDERSEN
	TCOUPL_ANDERSEN_MASSIVE
	TCOUPL_V_RESCALE
)

func (m *MDP) SetTCoupling() {

}

func (m *MDP) TCoupling() {

}

func convertTCoupling(m *MDP) string {
	return ""
}

/*******************************/
/* pcoupl */
/*******************************/

type PCouple int8

const (
	PCOUPL_NOTSET PCouple = 1 << iota
	PCOUPL_NO
	PCOUPL_BERENDSEN
	PCOUPL_PARINELLO_RAHMAN
)

/* pcoupl type */

type PCoupleType int8

const (
	PCOUPLTYPE_NOTSET PCoupleType = 1 << iota
	PCOUPLTYPE_ISOTROPIC
	PCOUPLTYPE_SEMIISOTROPIC
)

type RefCoordScaling int8

const (
	REFCOORD_SCALING_NOTSET RefCoordScaling = 1 << iota
	REFCOORD_SCALING_NO
	REFCOORD_SCALING_ALL
	REFCOORD_SCALING_COM
)

func (m *MDP) SetPCoupling() {

}

func (m *MDP) PCoupling() {

}

func convertPCoupling(m *MDP) string {
	return ""
}

/*******************************/
/* velocity */
/*******************************/

func (m *MDP) SetVelocityGeneration(v bool) {

}

func (m *MDP) VelocityGeneration() {

}

func convertVelocityGeneration(m *MDP) string {
	return ""
}

/*******************************/
/* constraint */
/*******************************/

type Constraint int8

const (
	CONSTRAINT_NOTSET Constraint = 1 << iota
	CONSTRAINT_NONE
	CONSTRAINT_H_BONDS
	CONSTRAINT_ALL_BONDS
	CONSTRAINT_H_ANGLE
	CONSTRAINT_ALL_ANGLES
)

type ConstAlg int8

const (
	CONSTALG_NOTSET ConstAlg = 1 << iota
	CONSTALG_LINCS
	CONSTALG_SHAKE
)

func (m *MDP) SetConstraints() {

}

func (m *MDP) Constraints() {

}

func convertConstraints(m *MDP) string {
	return ""
}

/*******************************/
/* pull */
/*******************************/

type PullMode int8

const (
	PULLMODE_NOTSET PullMode = 1 << iota
	PULLMODE_UMBRELLA
	PULLMODE_CONSTRAINT
	PULLMODE_CONSTANT_FORCE
)

type PullGeometry int8

const (
	PULL_GEOMETRY_NOTSET PullGeometry = 1 << iota
	PULL_GEOMETRY_DISTANCE
	PULL_GEOMETRY_DIRECTION
	PULL_GEOMETRY_DIRECTION_PERIODIC
	PULL_GEOMETRY_CYLINDER
)

func (m *MDP) SetPullConfig() {

}

func (m *MDP) PullConfig() {

}

func convertPullConfig(m *MDP) string {
	return ""
}

/*******************************/
/* helpers */
/*******************************/

func makeMDPEntry(key, value string) string {
	return fmt.Sprintf("%20s = %s", key, value)
}

/* ==================================================================
	MDP string generation
/ ================================================================== */

func (m *MDP) String() (string, error) {
	mdp_template := `
; defines
{{ convertDefines . }}	

; md
{{ convertIntegrator . }}

; log
{{ convertLog     . }}

; neighbor searching
{{ convertNeighborSearch . }}

; electrostatics
{{ convertElectrostatics . }}

; vdw
{{ convertVdw . }}

; tcoupling
{{ convertTCoupling . }}

; pcoupling
{{ convertPCoupling . }}

; generate velocity
{{ convertVelocityGeneration . }}

; pull
{{ convertPullConfig . }}
`

	funcMap := template.FuncMap{
		"convertDefines":            convertDefines,
		"convertLog":                convertLog,
		"convertNeighborSearch":     convertNeighborSearch,
		"convertElectrostatics":     convertElectrostatics,
		"convertVdw":                convertVdw,
		"convertTCoupling":          convertTCoupling,
		"convertPCoupling":          convertPCoupling,
		"convertVelocityGeneration": convertVelocityGeneration,
		"convertPullConfig":         convertPullConfig,
	}

	tmpl, err := template.New("rtp").Funcs(funcMap).Parse(mdp_template)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, m)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

/* ==================================================================
	MDP reader
/ ================================================================== */

func ReadMDPFile(fname string) (*MDP, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, nil
	}
	defer file.Close()

	return readMDP(file)

}

func ReadMDPString(str string) (*MDP, error) {
	return readMDP(strings.NewReader(str))
}

func readMDP(reader io.ReaderAt) (*MDP, error) {
	return nil, nil
}
