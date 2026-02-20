package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDataDir = "testdata/dir2/"

func TestGetPathSizeForBlankPath(t *testing.T) {
	_, err := GetPathSize("", false, false, false)
	assert.Error(t, err)
}

func TestGetPathSizeForIncorrectPath(t *testing.T) {
	_, err := GetPathSize("incorrect", false, false, false)
	assert.Error(t, err)
}

func TestGetPathSizeTableDriven(t *testing.T) {
	var tests = []struct {
		name      string
		path      string
		all       bool
		human     bool
		recursive bool
		expected  string
	}{
		{"GetPathSize", testDataDir, false, false, false, "2B"},                // a.txt (2)
		{"GetPathSize all", testDataDir, true, false, false, "9B"},             // a.txt + .hidden.txt (2 + 7)
		{"GetPathSize recursive", testDataDir, false, false, true, "7B"},       // a.txt + deep.txt (2 + 5)
		{"GetPathSize all + recursive", testDataDir, true, false, true, "14B"}, // a.txt + .hidden.txt + deep.txt (2 + 7 + 5)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetPathSize(tt.path, tt.recursive, tt.human, tt.all)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, res)
		})
	}
}
