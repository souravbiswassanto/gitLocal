package pkg

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func GetFolders(paths []string) ([]string, error) {
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
			return []string{}, err
		}

	}
	return folders, nil
}

func GetRepoStoreFileName() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("/home/%s/.gitLocal", strings.ToLower(usr.Name))
	return fileName, err
}
func CreateRepoStoreFile(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(filePath)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func GetExistingReposFromStoreFile(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("err = %v", err)
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		if err != io.EOF {
			log.Fatalf("err = %v", err)
		}
	}
	return lines
}

func existsInSlice(key string, In []string) bool {
	for _, str := range In {
		if str == key {
			return true
		}
	}
	return false
}

func MergeSlice(s1 []string, s2 []string) []string {
	var newString []string
	for _, str := range s1 {
		if !existsInSlice(str, s2) {
			newString = append(newString, str)
		}
	}
	newString = append(newString, s2...)
	return newString
}

func WriteReposToFile(filePath string, repos []string) {
	content := strings.Join(repos, "\n")
	err := os.WriteFile(filePath, []byte(content), 0755)
	if err != nil {
		log.Fatalf("err = %v", err)
	}
}
