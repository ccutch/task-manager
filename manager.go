package gtm

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var Tasks = []*Task{}

func init() {
	fn := "tasks.yml"
	d, _ := ioutil.ReadFile(fn)
	ioutil.WriteFile(fn, d, 0777)
	LoadTasks()
}

func AddTask(t *Task) {
	Tasks = append(Tasks, t)
	SaveTasks()
}

func RemoveTask(id string) error {
	for i, t := range Tasks {
		if t.Id == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			SaveTasks()
			return nil
		}
	}
	return errors.New("Task not found")
}

func LoadTasks() error {
	fn := GlobalConfig.FileName
	d, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(d, &Tasks)
	return err
}

func SaveTasks() error {
	d, err := yaml.Marshal(Tasks)
	if err != nil {
		return err
	}
	fn := GlobalConfig.FileName
	err = ioutil.WriteFile(fn, d, 0777)
	return err
}
