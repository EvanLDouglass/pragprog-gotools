package main

import (
	"flag"
	"fmt"
	"os"

	"pragprog.com/gocmd/interacting/todo"
)

// TODO: make this flexible
const todoFileName = ".todo.json"

func main() {
	// Customize help message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for The Pragmatic Bookshelp\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2020\n")
		fmt.Fprintf(flag.CommandLine.Output(),
			"Usage information (only supports one option per call):\n")
		flag.PrintDefaults()
	}

	// Get cmd line flags
	task := flag.String("task", "", "Task to be included in the todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to mark as complete")
	flag.Parse()

	// Generate a list to use
	l := &todo.List{}

	// Use Get to check if there are current items
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Check flags for todo behavior
	switch {
	case *complete > 0:
		// Mark the item as complete, if it exists
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		// Add a new task to the list
		l.Add(*task)
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *list:
		// List the items not yet completed
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	default:
		// Invalid flag
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
