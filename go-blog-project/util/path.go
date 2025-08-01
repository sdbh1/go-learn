package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindProjectRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// 向上查找 go.mod 文件来确定项目根目录
	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("go.mod not found")
		}
		currentDir = parentDir
	}
}
