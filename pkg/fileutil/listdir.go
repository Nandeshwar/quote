package fileutil

import (
	"os"
	"path/filepath"
)

func ListDir(dirName string) ([]string, error) {
	var files []string

	root := dirName
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
