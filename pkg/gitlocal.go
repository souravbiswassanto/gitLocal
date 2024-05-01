package pkg

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func ShowLocalGitContrib(email string, paths []string) {
	var folders []string
	for _, path := range paths {
		err := WalkDir(path, func(dir string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(dir, "vendor") {
				return filepath.SkipDir
			}
			
			if d.IsDir() && strings.HasSuffix(dir, ".git") {
				folders = append(folders, strings.TrimSuffix(dir, "/.git"))
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}
