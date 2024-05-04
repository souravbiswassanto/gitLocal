package pkg

import "log"

func ShowLocalGitContrib(email string, paths []string) {
	Scan(paths)
	repoStoreFileName, err := GetRepoStoreFileName()
	if err != nil {
		log.Fatalln(err)
	}
	Stats(email, repoStoreFileName)
}

func Scan(paths []string) {
	folders, err := GetFolders(paths)
	if err != nil {
		log.Fatalf("err = %v", err)
	}
	repoStoreFileName, err := GetRepoStoreFileName()
	if err != nil {
		log.Fatalf("err = %v", err)
	}
	err = CreateRepoStoreFile(repoStoreFileName)
	if err != nil {
		log.Fatalf("err = %v", err)
	}
	existingReposFromStoreFile := GetExistingReposFromStoreFile(repoStoreFileName)
	mergedRepoSlice := MergeSlice(folders, existingReposFromStoreFile)
	WriteReposToFile(repoStoreFileName, mergedRepoSlice)
}
