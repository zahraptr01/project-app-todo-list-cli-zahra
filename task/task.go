package task

import (
	"fmt"
	"strings"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (t Task) String() string {
	status := "✗"
	if t.Completed {
		status = "✓"
	}
	return fmt.Sprintf("%s\t%s\t%s", status, t.Title, t.Description)
}

// function to check duplication of task title data
func IsDuplicate(tasks []Task, title string) bool {
	for _, t := range tasks {
		if strings.EqualFold(t.Title, title) {
			return true
		}
	}
	return false
}
