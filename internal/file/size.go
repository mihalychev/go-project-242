package file

import (
	"fmt"
	"os"
)

func GetSize(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if fileInfo.IsDir() {
		size, err := directoryFilesSize(path)
		if err != nil {
			return 0, err
		}

		return size, nil
	}

	return fileInfo.Size(), nil
}

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

func directoryFilesSize(path string) (int64, error) {
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

		if fileInfo.IsDir() {
			continue
		}

		sum += fileInfo.Size()
	}

	return sum, nil
}
