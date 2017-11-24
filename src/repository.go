package ws17_pp

import "fmt"

var currentId int

var jobs Jobs

// Give us some seed data
func init() {
	RepoCreateJob(Job{Name: "Job Mock - Fibonacci"})
}

func RepoFindJob(id int) Job {
	for _, t := range jobs {
		if t.Id == id {
			return t
		}
	}
	// return empty Job if not found
	return Job{}
}

func RepoCreateJob(t Job) Job {
	currentId += 1
	t.Id = currentId
	jobs = append(jobs, t)
	return t
}

func RepoDestroyJob(id int) error {
	for i, t := range jobs {
		if t.Id == id {
			jobs = append(jobs[:i], jobs[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Job with id of %d to delete", id)
}
