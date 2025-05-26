package main

import (
	"flag"
	"fmt"
	"os"
	"project-app-todo-list-cli-zahra/task"
)

// function for command-line flag declaration
func main() {
	add := flag.String("add", "", "Tambah tugas baru")
	desc := flag.String("desc", "", "Deskripsi tugas")
	list := flag.Bool("list", false, "Tampilkan semua tugas")
	done := flag.String("done", "", "Tandai tugas sebagai selesai")
	deleteTask := flag.String("delete", "", "Hapus tugas")
	search := flag.String("search", "", "Cari tugas")
	flag.Parse()

	// Read task data from file
	tasks, err := task.LoadTasks()
	// Error Handling
	if err != nil {
		fmt.Println("Gagal memuat tugas:", err)
		os.Exit(1)
	}

	// Switch to handle each operation
	if *add != "" {
		err = task.AddTask(&tasks, *add, *desc)
	} else if *list {
		task.PrintTasks(tasks)
	} else if *done != "" {
		err = task.MarkDone(&tasks, *done)
	} else if *deleteTask != "" {
		err = task.DeleteTask(&tasks, *deleteTask)
	} else if *search != "" {
		task.SearchTasks(tasks, *search)
	} else {
		fmt.Println("Gunakan --help untuk melihat opsi.")
	}

	// Error Handling
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
