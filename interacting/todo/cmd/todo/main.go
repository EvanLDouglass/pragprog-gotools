package main

import (
	"fmt"
	"os"
	"strings"

	"pragprog.com/gocmd/interacting/todo"
)

// TODO: make this flexible
const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	// Use Get to check if there are current items
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Check number of args for todo behavior
	switch {
	// No extra args: list current items
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	// Given args: concat and use as name for a new task
	default:
		// Concat with space
		item := strings.Join(os.Args[1:], " ")
		// Add and save
		l.Add(item)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
