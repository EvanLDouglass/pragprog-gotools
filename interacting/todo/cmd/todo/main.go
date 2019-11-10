package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"pragprog.com/gocmd/interacting/todo"
)

// Default filename
var todoFileName = ".todo.json"

func main() {
	// Customize help message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for The Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(),
			"Copyright 2020\n")
		fmt.Fprintf(flag.CommandLine.Output(),
			"Usage information (only supports one option per call):\n")
		flag.PrintDefaults()
	}

	// Get cmd line flags
	add := flag.Bool("add", false, "Add task to the todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to mark as complete")
	flag.Parse()

	// Look for env variable for filename
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

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
	case *add:
		// Add a new task to the list
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(task)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *list:
		// List items
		fmt.Print(l)
	default:
		// Invalid flag
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

// getTask function decides where to get the description for a new task: from
// arguments or STDIN
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return s.Text(), nil
}
