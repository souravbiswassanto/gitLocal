package pkg

import (
	"fmt"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"log"
	"time"
)

func Test() {
	repo, err := git.PlainOpen("/home/saurov/go/src/github.com/souravbiswassanto/api-bookserver")
	if err != nil {
		log.Fatalf("%v", err)
	}
	a, err := repo.Branch("main")
	fmt.Println(a.Name, a.Remote)
	b, err := repo.Remotes()
	for _, r := range b {
		fmt.Println(r.String(), " config", r.Config().Name, "urls", r.Config().URLs)
	}
	// git local branches
	refs, err := repo.Branches()
	err = refs.ForEach(func(reference *plumbing.Reference) error {
		fmt.Println(reference.Name(), reference.Target(), reference.String())
		return nil
	})

	head, err := repo.Head()
	iterator, err := repo.Log(&git.LogOptions{From: head.Hash()})
	if err != nil {
		log.Fatalf("err = %v", err)
	}
	err = iterator.ForEach(func(commit *object.Commit) error {
		fmt.Println(commit)
		return nil
	})

}

func Stats(email string, path string) {
	repoList := GetExistingReposFromStoreFile(path)
	CommitCount := make(map[int]int)
	for i := 0; i <= DaysInSixMonths; i++ {
		CommitCount[i] = 0
	}
	for i := range repoList {
		CountCommits(repoList[i], func(day int) {
			CommitCount[day]++
		})
	}
	for key, val := range CommitCount {
		fmt.Println(key, val)
	}
}

func CountCommits(repoPath string, fn func(day int)) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalln(err)
	}
	head, err := repo.Head()
	if err != nil {
		log.Fatalln(err)
	}
	commitIterator, err := repo.Log(&git.LogOptions{
		From: head.Hash(),
	})
	if err != nil {
		log.Fatalln(err)
	}
	err = commitIterator.ForEach(func(c *object.Commit) error {
		commitTime := c.Author.When
		duration := int(time.Since(commitTime).Seconds())
		limitSeconds := DaysInSixMonths * HoursInDay * MinutesInHour * SecondsInMinute

		if limitSeconds >= duration {
			day := duration / (SecondsInMinute * MinutesInHour * HoursInDay)
			fn(day)
		}

		return nil
	})
}
