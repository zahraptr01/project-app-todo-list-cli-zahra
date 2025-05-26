package task

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// function to add tasks
func AddTask(tasks *[]Task, title, desc string) error {
	// data validation
	if strings.TrimSpace(title) == "" {
		return errors.New("judul tugas tidak boleh kosong")
	}
	if IsDuplicate(*tasks, title) {
		return errors.New("tugas dengan judul ini sudah ada")
	}
	*tasks = append(*tasks, Task{Title: title, Description: desc})
	return SaveTasks(*tasks)
}

// function to mark the task as completed
func MarkDone(tasks *[]Task, title string) error {
	// data validation
	for i, t := range *tasks {
		if strings.EqualFold(t.Title, title) {
			(*tasks)[i].Completed = true
			return SaveTasks(*tasks)
		}
	}
	return errors.New("tugas tidak ditemukan")
}

// function to delete tasks in the task list
func DeleteTask(tasks *[]Task, title string) error {
	newTasks := []Task{}
	found := false
	// data validation
	for _, t := range *tasks {
		if !strings.EqualFold(t.Title, title) {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}
	if !found {
		return errors.New("tugas tidak ditemukan")
	}
	*tasks = newTasks
	return SaveTasks(*tasks)
}

// function to search for tasks based on keywords
func SearchTasks(tasks []Task, keyword string) {
	fmt.Println("Hasil pencarian:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Status\tJudul\tDeskripsi")
	// data validation
	for _, t := range tasks {
		if strings.Contains(strings.ToLower(t.Title), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(t.Description), strings.ToLower(keyword)) {
			fmt.Fprintln(w, t.String())
		}
	}
	w.Flush()
}

// function to display task list
func PrintTasks(tasks []Task) {
	// data validation
	if len(tasks) == 0 {
		fmt.Println("Tidak ada tugas.")
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Status\tJudul\tDeskripsi")
	for _, t := range tasks {
		fmt.Fprintln(w, t.String())
	}
	w.Flush()
}
