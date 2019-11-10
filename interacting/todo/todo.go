package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// item struct represents a todo item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of todo items
type List []item

// Add creates a new todo item and appends it to the list.
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// Complete method marks a todo item as completed by setting
// Done = true and CompletedAt to the current time.
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Save method encodes the List as JSON and saves it using
// the provided file name.
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes the JSON
// data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		// file does not exist, nothing to unmarshal, no error
		if os.IsNotExist(err) {
			return nil
		}
		// actual error
		return err
	}
	// file exists but is empty, nothing to unmarshal, no error
	if len(file) == 0 {
		return nil
	}
	// unpack JSON to l, return the error from Unmarshal
	return json.Unmarshal(file, l)
}

// String prints out a formatted list. Implements the fmt.Stringer interface.
func (l *List) String() string {
	var formatted strings.Builder

	for k, t := range *l {
		// add padding
		prefix := "  "
		if t.Done {
			// include X in padding for finished tasks
			prefix = "X "
		}
		// write prefix, task number and task to formatted
		formatted.WriteString(fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task))
	}

	return formatted.String()
}
