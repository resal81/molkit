package fragment

import (
	"strings"
)

type Fragment struct {
	name  string
	atoms []string // "H" "N" "O"
	bonds []string // "CA N"
}

type FragmentDatabase struct {
	fragments []*Fragment
}

func NewFragmentDatabase() *FragmentDatabase {
	return &FragmentDatabase{}
}

// The default database is instantiated using the data file
var DefaultFragmentDatabase = NewFragmentDatabase()

func init() {
	for _, line := range strings.Split(fragmentData, "\n") {
		line = stripComment(line)

		if strings.TrimSpace(line) == "" {
			continue
		}
	}
}

// stripComment removes Gromacs-style comments
func stripComment(line string) string {
	if strings.Contains(line, ";") {
		i := strings.Index(line, ";")
		return line[0:i]
	}
	return line
}

// getBracketField returns the field inside a Gromacs-style field
func getBracketField(line string) string {
	if strings.Contains(line, "[") && strings.Contains(line, "]") {
		i := strings.Index(line, "[")
		j := strings.Index(line, "]")

		out := line[i+1 : j-1]
		return strings.TrimSpace(out)
	}

	return line
}
