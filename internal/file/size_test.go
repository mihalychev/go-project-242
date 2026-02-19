package file

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetSizeTableDriven(t *testing.T) {
	var tests = []struct {
		name string
		path string
		all bool
		recursive bool
		expected int64
	}{
		{ "GetSize for file", "../../testdata/test_cat.png", false, false, 43689 },
		{ "GetSize for hidden file", "../../testdata/test_size_dir/.hidden_elephant.png", false, false, 0 },
		{ "GetSize for hidden file with all flag", "../../testdata/test_size_dir/.hidden_elephant.png", true, false, 678429 },

		{ "GetSize for directory", "../../testdata/test_size_dir", false, false, 167353 + 10717 },
		{ "GetSize for directory with all flag", "../../testdata/test_size_dir", true, false, 167353 + 10717 + 678429 },
		{ "GetSize for hidden directory", "../../testdata/.hidden_test_size_dir", false, false, 0 },
		{ "GetSize for hidden directory with all flag", "../../testdata/.hidden_test_size_dir", true, false, 167353 },
		{ "GetSize for directory with recursive flag", "../../testdata/test_size_dir", false, true, 167353 + 10717 * 2 },
		{ "GetSize for directory with all and recursive flags", "../../testdata/test_size_dir", true, true, 167353 + 10717 * 2 + 678429 },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetSize(tt.path, tt.all, tt.recursive)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestFormatSizeNotHuman(t *testing.T) {
	expected := "43689B"

	size, err := GetSize("../../testdata/test_cat.png", false, false)
	assert.NoError(t, err)

	res := FormatSize(size, false)
	assert.Equal(t, res, expected)
}

func TestFormatSizeHuman(t *testing.T) {
	expected := "42.7KB" // 43689 / 1024 = 42.665...

	size, err := GetSize("../../testdata/test_cat.png", false, false)
	assert.NoError(t, err)

	res := FormatSize(size, true)
	assert.Equal(t, res, expected)
}
