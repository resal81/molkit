package blocks

type TopolsDB struct {
	exclusions  []*TopExclusion
	rest_pos    []*TopPositionRestraint
	rest_dist   []*TopDistanceRestraint
	rest_ang    []*TopAngleRestraint
	rest_dih    []*TopDihedralRestraint
	rest_orient []*TopOrientationRestraint
	settle      *TopSettle
}
