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

func FormatSize(bytes int64) string {
	if bytes <= 1024 {
		return fmt.Sprintf("%dB", bytes)
	}

	kilobytes := bytes / 1024
	if kilobytes <= 1024 {
		return fmt.Sprintf("%dKB", kilobytes)
	}

	megabytes := kilobytes / 1024
	if megabytes <= 1024 {
		return fmt.Sprintf("%dMB", megabytes)
	}

	gigabytes := kilobytes / 1024
	return fmt.Sprintf("%dGB", gigabytes)
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
