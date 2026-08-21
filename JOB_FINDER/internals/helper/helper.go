package helper

import (
	"os"
)

func Helper(a string) string {
	return ""
}

func CheckDirectory(path string) string {
	files, err := os.ReadDir(path)

	if err != nil {
		return "empty"
	}

	if len(files) == 0 {
		return "empty"
	}
	return "full"
}
