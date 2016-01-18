package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ccutch/task-manager"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Task Manager"
	app.Usage = "Easy and quick task management"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:  "config",
			Usage: "Init and set config for task manager.",
			Subcommands: []cli.Command{
				{
					Name:   "set",
					Usage:  "Set config option",
					Action: setConfig,
				},
				{
					Name:  "backup",
					Usage: "Get backup copy of config",
					Action: func(c *cli.Context) {
						fmt.Println(taskmanager.GetConfigYaml())
					},
				},
			},
		},
		{
			Name:  "tasks",
			Usage: "All task commands",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "List data type",
					Action: func(c *cli.Context) {
						if len(taskmanager.Tasks) == 0 {
							fmt.Println("No tasks to list out.")
						} else {
							for i, t := range taskmanager.Tasks {
								complete := "√"
								if !t.Complete {
									complete = "X"
								}
								fmt.Printf("Task %d - %s [owner: %s] %s\n", i+1, t.Title, t.Owner, complete)
								d := t.Description
								l := len(d)
								w := 60

								if l < w {
									fmt.Println("\t" + d)
								}
								for i := w; i < l; i += w {
									fmt.Println("\t" + d[i-w:i])
								}
								fmt.Println()
							}
						}
					},
				},
				{
					Name:    "claim",
					Aliases: []string{"c"},
					Usage:   "Claim a Task",
					Action: func(c *cli.Context) {
						i, _ := strconv.Atoi(c.Args().First())
						err := taskmanager.Tasks[i-1].ClaimTask()
						if err != nil {
							fmt.Println(err.Error())
							return
						}
						taskmanager.SaveTasks()
					},
				},
				{
					Name:    "mark",
					Aliases: []string{"m"},
					Usage:   "Mark task as complete",
					Action: func(c *cli.Context) {
						i, _ := strconv.Atoi(c.Args().First())
						err := taskmanager.Tasks[i-1].MarkComplete()
						if err != nil {
							fmt.Println(err.Error())
							return
						}
						taskmanager.SaveTasks()
					},
				},
				{
					Name:    "add",
					Aliases: []string{"a"},
					Usage:   "Add a new task",
					Action: func(c *cli.Context) {
						t := c.Args().First()
						d := c.String("description")
						ts, _ := taskmanager.NewTask(t, d)
						if ts != nil {
							taskmanager.AddTask(ts)
						}
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "description",
							Usage: "Description of a task",
						},
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func setConfig(c *cli.Context) {
	f := c.Args().Get(0)
	v := c.Args().Get(1)
	switch f {
	case "user":
		taskmanager.SetUser(v)
	case "server":
		taskmanager.SetServer(v)
	default:
		fmt.Println("Error unrecognised field entered to set, options are \"user\" and \"server\".")
	}
}
