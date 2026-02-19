package file

import (
	"fmt"
	"os"
	"strings"
)

func FormatSize(bytes int64, human bool) string {
	size := float64(bytes)
	if size < 1024 || !human {
		return fmt.Sprintf("%.0fB", size)
	}

	for _, unit := range []string{"KB", "MB", "GB", "TB", "PB", "EB"} {
		size /= 1024
		if size < 1024 {
			return fmt.Sprintf("%.1f%s", size, unit)
		}
	}
	return fmt.Sprintf("%.1fEB", size)
}

func GetSize(path string, all bool) (int64, error) {
	if isHidden(path) && !all {
		return 0, nil
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if fileInfo.IsDir() {
		size, err := directoryFilesSize(path, all)
		if err != nil {
			return 0, err
		}

		return size, nil
	}

	return fileInfo.Size(), nil
}

func isHidden(path string) bool {
	pathParts := strings.Split(path, "/")
	return pathParts[len(pathParts) - 1][0] == '.'
}

func directoryFilesSize(path string, all bool) (int64, error) {
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

		if fileInfo.IsDir() || (isHidden(fileInfo.Name()) && !all) {
			continue
		}

		sum += fileInfo.Size()
	}

	return sum, nil
}
