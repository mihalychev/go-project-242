package code

import (
	"code/src/file"
	"errors"
)

func GetPathSize(path string, all, human, recursive bool) (string, error) {
	if path == "" {
		return "", errors.New("path argument must be passed")
	}

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
