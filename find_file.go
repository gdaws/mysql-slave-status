package main

import (
	"path"
	"os"
)

func FindFile(filename string, searchPaths []string) []string {

	found := make([]string, 0, len(searchPaths))

	for _, basePath := range searchPaths {

		filePath := path.Join(basePath, filename)

		_, err := os.Stat(filePath)

		if err == nil {
			found = append(found, filePath)
		}
	}

	return found
}
