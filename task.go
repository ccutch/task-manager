package gtm

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

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

func NewTask(title, description string) (*Task, error) {
	if ConfigComplete {
		return nil, errors.New("Config not complete, set user name and server to connect to.")
	}
	t := &Task{
		Id: uuid.NewV4().String(),

		Title:       title,
		Description: description,
		Complete:    false,
	}
	return t, nil
}

func (t *Task) MarkComplete() error {
	if t.Owner != globalConfig.User {
		return errors.New("This is not you task.")
	}
	t.Complete = true
	return t.SaveToServer()
}

func (t *Task) ClaimTask() error {
	t.Owner = globalConfig.User
	return t.SaveToServer()
}

func (t *Task) SaveToServer() error {
	if !ConfigComplete {
		return nil
	}
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	resp, err := http.Post(globalConfig.Server, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		s := buf.String()
		return errors.New("Saving failed, [status " + resp.Status + "] " + s)
	}
	t.Saved = true
	return nil
}
