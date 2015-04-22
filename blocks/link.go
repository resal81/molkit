package blocks

/*
	Link
*/

type Link struct {
	Fragment1   *Fragment
	Fragment2   *Fragment
	Connections [][2]int64
}
