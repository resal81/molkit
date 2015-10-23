package fragment

import (
	"testing"
)

func TestStripComment(t *testing.T) {
	var data = []struct {
		lineIn  string
		lineOut string
	}{
		{"; a full line comment", ""},
		{";", ""},
		{"[ bond ] ; some info ", "[ bond ] "},
	}

	for _, d := range data {
		if out := stripComment(d.lineIn); out != d.lineOut {
			t.Errorf("Wrong stripComment result: got '%s', expected '%s'", out, d.lineOut)
		}
	}
}

func TestGetBracketField(t *testing.T) {
	var data = []struct {
		line   string
		result string
	}{
		{" [ bond ] ; with comment and extra [ bracket ] ", "bond"},
		{" no bracket ", " no bracket "},
	}

	for _, d := range data {
		if out := getBracketField(d.line); out != d.result {
			t.Errorf("Wrong getBracketField result: got '%s', expected '%s'", out, d.result)
		}
	}
}
