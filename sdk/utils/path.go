package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type AppPath string

func (path *AppPath) Valid() bool {
	return strings.HasPrefix(string(*path), "../")
}

func (path *AppPath) Normalize() string {
	relativePath, ok := strings.CutPrefix(string(*path), "../")
	if !ok {
		panic("[AppPath_Normalize] invalid path: " + *path)
	}

	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	rootPath := filepath.Dir(exePath)
	fullPath := rootPath + "/" + relativePath

	log.Println("[AppPath_Normalize] fullPath:", fullPath)

	return fullPath
}
