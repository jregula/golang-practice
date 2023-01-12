package main
import (
"fmt"
)
// search for open job seeker with salary less than 100

type jobSeeker struct {
	name string
	openToWork bool
	skills []string
	location string
	salary float64
}

func(j *jobSeeker) Search() string {
	if j.salary < 100 && j.openToWork == true {
		fmt.Println(j.name)
		return j.name
	}
	return "none"
}
type Searcher interface {
	Search() string
}

func main() {
	Ruzin := &jobSeeker{
		"Ruzin",
		true,
		[]string{},
		"UK",
		999,
	}

	Joshua := &jobSeeker{
		"Joshua",
		false,
		[]string{"everything", "more of everything"},
		"Spain",
		999,
	}
	jobSeekerPeople := []Searcher{
		Ruzin,
		Joshua,
	}

	for _,i := range jobSeekerPeople {
		i.Search()

	}

}