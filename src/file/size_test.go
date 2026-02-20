package file

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

const (
	testDataDir = "../../testdata/dir1/"

	// Test files sizes in bytes
	dogSize      = 167353
	elephantSize = 678429
)

func TestGetSizeForBlankPath(t *testing.T) {
	_, err := GetSize("", false, false)
	assert.Error(t, err)
}

func TestGetSizeTableDriven(t *testing.T) {
	var tests = []struct {
		name      string
		path      string
		all       bool
		recursive bool
		expected  int64
	}{
		{"GetSize for file", dataDir("dog.png"), false, false, dogSize},
		{"GetSize for hidden file", dataDir(".hidden_elephant.png"), false, false, 0},
		{"GetSize for hidden file with all flag", dataDir(".hidden_elephant.png"), true, false, elephantSize},
		{"GetSize for for file in hidden directory", dataDir(".hidden_nested/dog.png"), false, false, 0},
		{"GetSize for for file in hidden directory with all flag", dataDir(".hidden_nested/dog.png"), true, false, dogSize},

		{"GetSize for directory", testDataDir, false, false, dogSize},
		{"GetSize for directory with all flag", testDataDir, true, false, dogSize + elephantSize},
		{"GetSize for hidden directory", dataDir(".hidden_nested"), false, false, 0},
		{"GetSize for hidden directory with all flag", dataDir(".hidden_nested"), true, false, dogSize},
		{"GetSize for directory with recursive flag", testDataDir, false, true, dogSize * 2},
		{"GetSize for directory with all and recursive flags", testDataDir, true, true, dogSize*3 + elephantSize},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetSize(tt.path, tt.all, tt.recursive)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestFormatSizeForNegativeSize(t *testing.T) {
	_, err := FormatSize(-1, false)
	assert.Error(t, err)
}

func TestFormatSizeTableDriven(t *testing.T) {
	var tests = []struct {
		name     string
		human    bool
		expected string
	}{
		{"FormatSize with human flag", true, "163.4KB"}, // 167353 / 1024 = 163.43...
		{"FormatSize without human flag", false, "167353B"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size, err := GetSize(dataDir("dog.png"), false, false)
			assert.NoError(t, err)

			res, err := FormatSize(size, tt.human)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func dataDir(path string) string {
	return filepath.Join(testDataDir, path)
}
