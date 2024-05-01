package pkg

import (
	"io/fs"
	"path/filepath"
)

func WalkDir(path string, fn fs.WalkDirFunc) error {
	path = filepath.Clean(path)
	err := filepath.WalkDir(path, fn)
	return err
}
