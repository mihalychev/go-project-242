package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FormatSize(bytes int64, human bool) (string, error) {
	if bytes < 0 {
		return "", errors.New("incorrect size value")
	}

	if bytes < 1024 || !human {
		return fmt.Sprintf("%dB", bytes), nil
	}

	size := float64(bytes)
	for _, unit := range []string{"KB", "MB", "GB", "TB", "PB", "EB"} {
		size /= 1024
		if size < 1024 {
			return fmt.Sprintf("%.1f%s", size, unit), nil
		}
	}
	return fmt.Sprintf("%.1fEB", size), nil
}

func GetSize(path string, all, recursive bool) (int64, error) {
	if isHidden(path) && !all {
		return 0, nil
	}

	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if fileInfo.IsDir() {
		size, err := directoryFilesSize(path, all, recursive)
		if err != nil {
			return 0, err
		}

		return size, nil
	}

	return fileInfo.Size(), nil
}

func directoryFilesSize(path string, all, recursive bool) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var sum int64 = 0
	for _, v := range entries {
		fileInfo, err := v.Info()
		if err != nil {
			return 0, err
		}

		if isHidden(fileInfo.Name()) && !all {
			continue
		}

		if fileInfo.IsDir() {
			if !recursive {
				continue
			}

			nestedPath := filepath.Join(path, fileInfo.Name())
			nestedSize, err := directoryFilesSize(nestedPath, all, recursive)
			if err != nil {
				return 0, err
			}

			sum += nestedSize
		} else {
			sum += fileInfo.Size()
		}
	}

	return sum, nil
}

func isHidden(path string) bool {
	pathParts := strings.Split(path, "/")
	for _, v := range pathParts {
		if strings.HasPrefix(v, ".") && v != "." && v != ".." {
			return true
		}
	}

	return false
}
