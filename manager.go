package taskmanager

import (
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

func LoadTasks() {
	fn := globalConfig.FileName
	d, _ := ioutil.ReadFile(fn)
	yaml.Unmarshal(d, &Tasks)
	if globalConfig.Server != "" && globalConfig.User != "" {
		ConfigComplete = true
	}
}

func SaveTasks() {
	d, _ := yaml.Marshal(Tasks)
	fn := globalConfig.FileName
	ioutil.WriteFile(fn, d, 0777)
}
