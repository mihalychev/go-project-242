package file

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Test files sizes in bytes
const (
	CatSize      = 43689
	DogSize      = 167353
	MonkeySize   = 10717
	ElephantSize = 678429
)

func TestGetSizeForBlankPath(t *testing.T) {
	_, err := GetSize("", false, false)
	assert.Error(t, err)
}

func TestGetSizeTableDriven(t *testing.T) {
	var tests = []struct {
		name string
		path string
		all bool
		recursive bool
		expected int64
	}{
		{ "GetSize for file", "../../testdata/cat.png", false, false, CatSize },
		{ "GetSize for hidden file", "../../testdata/test_size_dir/.hidden_elephant.png", false, false, 0 },
		{ "GetSize for hidden file with all flag", "../../testdata/test_size_dir/.hidden_elephant.png", true, false, ElephantSize },

		{ "GetSize for directory", "../../testdata/test_size_dir", false, false, DogSize + MonkeySize },
		{ "GetSize for directory with all flag", "../../testdata/test_size_dir", true, false, ElephantSize + DogSize + MonkeySize },
		{ "GetSize for hidden directory", "../../testdata/.hidden_test_size_dir", false, false, 0 },
		{ "GetSize for hidden directory with all flag", "../../testdata/.hidden_test_size_dir", true, false, DogSize },
		{ "GetSize for directory with recursive flag", "../../testdata/test_size_dir", false, true, MonkeySize * 2 + DogSize },
		{ "GetSize for directory with all and recursive flags", "../../testdata/test_size_dir", true, true, MonkeySize * 2 + ElephantSize + DogSize },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetSize(tt.path, tt.all, tt.recursive)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestGetSizeForNegativeSize(t *testing.T) {
	_, err := FormatSize(-1, false)
	assert.Error(t, err)
}

func TestFormatSizeTableDriven(t *testing.T) {
	var tests = []struct {
		name string
		human bool
		expected string
	}{
		{ "FormatSize with human flag", true, "42.7KB" }, // 43689 / 1024 = 42.665...
		{ "FormatSize without human flag", false, "43689B" },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size, err := GetSize("../../testdata/cat.png", false, false)
			assert.NoError(t, err)

			res, err := FormatSize(size, tt.human)
			assert.NoError(t, err)
			assert.Equal(t, res, tt.expected)
		})
	}
}
