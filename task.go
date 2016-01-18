package gtm

import (
	"errors"
	"fmt"

	"github.com/satori/go.uuid"
)

type Task struct {
	// Data about the user
	Id    string `yaml:"id"`
	Owner string `yaml:"owner"`
	// Data about the task
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	// Meta data about the task
	Tags     []string `yaml:"tags"`
	Complete bool     `yaml:"complete"`
	Saved    bool     `yaml:"-"`
}

// Create a new task
func NewTask(title, description string) (*Task, error) {
	t := &Task{
		Id: uuid.NewV4().String(),

		Title:       title,
		Description: description,
		Complete:    false,
	}
	return t, nil
}

// Mark task as complete, this can only be done if the tasks Owner field
// to the global config user field
func (t *Task) MarkComplete() error {
	if t.Owner != GlobalConfig.User {
		return errors.New("This is not you task.")
	}
	t.Complete = true
	return SaveTasks()
}

// Claim task as your task, TODO: check for conflicts
func (t *Task) ClaimTask() error {
	if GlobalConfig.User == "" {
		return errors.New("No user is set us gtm config set user <your name>.")
	}
	t.Owner = GlobalConfig.User
	return SaveTasks()
}

func (t *Task) String() string {
	complete := "√"
	if !t.Complete {
		complete = "X"
	}
	s := fmt.Sprintf("Task <%s> - %s [owner: %s] %s\n", t.Id[:8], t.Title, t.Owner, complete)
	d := t.Description
	l := len(d)
	w := 60

	if l < w {
		s += "\t" + d
	}
	for i := w; i < l; i += w {
		s += "\t" + d[i-w:i] + "\n"
	}
	return s + "\n"
}
