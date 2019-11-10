package todo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"pragprog.com/gocmd/interacting/todo"
)

// TestAdd test the Add method of the List type.
func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Add(%q) = %q; Expected %q", taskName, l[0].Task, taskName)
	}
}

// TestSaveGet tests the Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Add(%q) = %q; Expected %q", taskName, l1[0].Task, taskName)
	}

	// temp file for testing
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	// try Save
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	// try Get and put into different List
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	// test values
	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task.", l1[0].Task, l2[0].Task)
	}
}
