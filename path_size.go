package code

import (
	"code/src/file"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := file.GetSize(path, all, recursive)
	if err != nil {
		return "", err
	}

	formattedSize, err := file.FormatSize(size, human)
	if err != nil {
		return "", err
	}

	return formattedSize, nil
}
