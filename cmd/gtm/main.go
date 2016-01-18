package main

import (
	"fmt"
	"os"

	"github.com/ccutch/task-manager"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "GTM"
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
						fmt.Println(gtm.GetConfigYaml())
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
						if len(gtm.Tasks) == 0 {
							fmt.Println("No tasks to list out.")
						} else {
							for _, t := range gtm.Tasks {
								fmt.Println(t)
							}
						}
					},
				},
				{
					Name:    "claim",
					Aliases: []string{"c"},
					Usage:   "Claim a Task",
					Action: func(c *cli.Context) {
						id := c.Args().First()
						for i, t := range gtm.Tasks {
							if t.Id[:8] == id {
								err := gtm.Tasks[i].ClaimTask()
								if err != nil {
									fmt.Println(err.Error())
									return
								}
								fmt.Println(gtm.Tasks[i])
								gtm.SaveTasks()
								return
							}
						}
						fmt.Println("Task not found for given id")
					},
				},
				{
					Name:    "mark",
					Aliases: []string{"m"},
					Usage:   "Mark task as complete",
					Action: func(c *cli.Context) {
						id := c.Args().First()
						for i, t := range gtm.Tasks {
							if t.Id[:8] == id {
								err := gtm.Tasks[i].MarkComplete()
								if err != nil {
									fmt.Println(err.Error())
									return
								}
								fmt.Println(gtm.Tasks[i])
								gtm.SaveTasks()
								return
							}
						}
						fmt.Println("Task not found for given id")
					},
				},
				{
					Name:    "add",
					Aliases: []string{"a"},
					Usage:   "Add a new task",
					Action: func(c *cli.Context) {
						t := c.Args().First()
						d := c.String("description")
						ts, _ := gtm.NewTask(t, d)
						fmt.Println(ts)
						if ts != nil {
							gtm.AddTask(ts)
						}
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "description",
							Usage: "Description of a task",
						},
					},
				},
				{
					Name:    "remove",
					Aliases: []string{"r"},
					Usage:   "Remove a new task",
					Action: func(c *cli.Context) {
						id := c.Args().First()
						for _, t := range gtm.Tasks {
							if t.Id[:8] == id {
								gtm.RemoveTask(t.Id)
								fmt.Println(t)
								gtm.SaveTasks()
								return
							}
						}
						fmt.Println("Task not found for given id")
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
		gtm.GlobalConfig.SetUser(v)
	default:
		fmt.Println("Error unrecognised field entered to set, options are \"user\" and \"server\".")
	}
}
