package internal

import (
	"os"
	"path/filepath"
)

// GetFiles returns list of files from path
func GetFiles(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, file)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
