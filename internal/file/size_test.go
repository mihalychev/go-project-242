package file

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetSizeForFile(t *testing.T) {
	var expected int64 = 43689

	res, err := GetSize("../../testdata/test_cat.png")
	assert.NoError(t, err)
	assert.Equal(t, res, expected)
}

func TestGetSizeForDirectory(t *testing.T) {
	var expected int64 = 167353 + 10717

	res, err := GetSize("../../testdata/test_size_dir")
	assert.NoError(t, err)
	assert.Equal(t, res, expected)
}

func TestFormatSizeNotHuman(t *testing.T) {
	expected := "43689B"

	size, err := GetSize("../../testdata/test_cat.png")
	assert.NoError(t, err)

	res := FormatSize(size, false)
	assert.Equal(t, res, expected)
}

func TestFormatSizeHuman(t *testing.T) {
	expected := "42.7KB" // 43689 / 1024 = 42.665...

	size, err := GetSize("../../testdata/test_cat.png")
	assert.NoError(t, err)

	res := FormatSize(size, true)
	assert.Equal(t, res, expected)
}
