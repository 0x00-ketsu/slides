package utils

import (
	"os"
)

func IsFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}
