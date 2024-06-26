package pp

import (
	"testing"
)

type colorTest struct {
	input  string
	color  uint16
	result string
}

var tests = []colorTest{
	{
		"blue on red",
		Blue | BackgroundRed,
		"\x1b[34m\x1b[41mblue on red\x1b[0m",
	},
	{
		"magenta on white",
		Magenta | BackgroundWhite,
		"\x1b[35m\x1b[47mmagenta on white\x1b[0m",
	},
	{
		"cyan",
		Cyan,
		"\x1b[36mcyan\x1b[0m",
	},
	{
		"default on red",
		BackgroundRed,
		"\x1b[41mdefault on red\x1b[0m",
	},
	{
		"default bold on yellow",
		Bold | BackgroundYellow,
		"\x1b[43m\x1b[1mdefault bold on yellow\x1b[0m",
	},
	{
		"bold",
		Bold,
		"\x1b[1mbold\x1b[0m",
	},
	{
		"no color at all",
		NoColor,
		"no color at all",
	},
}

func TestColorize(t *testing.T) {
	for _, test := range tests {
		if output := colorizeText(test.input, test.color); output != test.result {
			t.Errorf("Expected %q, got %q", test.result, output)
		}
	}
}
