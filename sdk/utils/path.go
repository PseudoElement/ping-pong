package utils

import (
	"os"
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

	exePath, _ := os.Executable()
	splitted := strings.Split(exePath, "/")
	splitted = splitted[:len(splitted)-1]
	rootPath := strings.Join(splitted, "/")

	return rootPath + "/" + relativePath
}
