package pkg

import (
	"fmt"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"log"
	"sort"
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

func Stats(email, username, path string) {
	repoList := GetExistingReposFromStoreFile(path)
	CommitCount := make(map[int]int)

	for i := 0; i < DaysInSixMonths; i++ {
		CommitCount[i] = 0
	}

	for i := range repoList {
		CountCommits(repoList[i], email, username, func(day int) {
			CommitCount[day]++
		})
	}

	sortedMapToSlice := ConvertMapIntoSortedSlice(CommitCount)
	cols := buildColumns(sortedMapToSlice, CommitCount)
	PrintCells(cols)
}

func CountCommits(repoPath, email, username string, fn func(day int)) {
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
		if c.Author.Name != username && c.Author.Email != email {
			return nil
		}
		commitTime := c.Author.When
		duration := int(time.Since(commitTime).Seconds())
		limitSeconds := DaysInSixMonths * HoursInDay * MinutesInHour * SecondsInMinute
		if limitSeconds >= duration {
			day := duration / (SecondsInMinute * MinutesInHour * HoursInDay)
			fn(day)
		}

		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func ConvertMapIntoSortedSlice(m map[int]int) []int {
	var key []int
	for k, _ := range m {
		key = append(key, k)
	}
	sort.Ints(key)
	return key
}

func buildColumns(key []int, commitMap map[int]int) map[int]column {
	cols := make(map[int]column)
	col := column{}
	for _, k := range key {
		week := k / 7
		days := k % 7
		if days == 0 {
			col = column{}
		}
		col = append(col, commitMap[k])
		if days == 6 {
			cols[week] = col
		}
	}

	return cols
}

func PrintCells(cols map[int]column) {
	PrintMonths()
	for j := 6; j >= 0; j-- {
		for i := WeeksInLastSixMonths + 1; i >= 0; i-- {
			if i == WeeksInLastSixMonths+1 {
				PrintDayCol(j)
			}
			if col, ok := cols[i]; ok {
				//special case today
				if i == 0 && j == CalcOffset()-1 {
					printCell(col[j], true)
					continue
				} else {
					if len(col) > j {
						printCell(col[j], false)
						continue
					}
				}
			}
			printCell(0, false)
		}
		fmt.Printf("\n")
	}
}

func PrintMonths() {
	week := GetBeginningOfDay(time.Now()).Add(-(DaysInSixMonths * time.Hour * 24))
	month := week.Month()
	fmt.Printf("         ")
	for {
		if week.Month() != month {
			fmt.Printf("%s ", week.Month().String()[:3])
			month = week.Month()
		} else {
			fmt.Printf("    ")
		}

		week = week.Add(7 * time.Hour * 24)
		if week.After(time.Now()) {
			break
		}
	}
	fmt.Printf("\n")
}

func PrintDayCol(day int) {
	out := "     "
	switch day {
	case 1:
		out = " Mon "
	case 3:
		out = " Wed "
	case 5:
		out = " Fri "
	}

	fmt.Printf(out)
}

func printCell(val int, today bool) {
	escape := "\033[0;37;30m"
	switch {
	case val > 0 && val < 5:
		escape = "\033[1;30;47m"
	case val >= 5 && val < 10:
		escape = "\033[1;30;43m"
	case val >= 10:
		escape = "\033[1;30;42m"
	}

	if today {
		escape = "\033[1;37;45m"
	}

	if val == 0 {
		fmt.Printf(escape + "  - " + "\033[0m")
		return
	}

	str := "  %d "
	switch {
	case val >= 10:
		str = " %d "
	case val >= 100:
		str = "%d "
	}

	fmt.Printf(escape+str+"\033[0m", val)
}
