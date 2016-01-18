# GTM (Go/GIT Task Manager)
GTM is a simple task manager writen completely in go. The purpose of this package is to provide a simple method of tracking tasks in a git repositories. All of the task are stored in a local yaml file and will continue to be stored this way, though server support will be added in the future for more flexable usability.

## Commands
`gtm <command> [command option] [args...]`


**config** Set config for your instance, including user name and server (coming soon)
 - options:
   * `set` Set a config option, `user` or `server`. example: `gtm config set user "<your name>"`
   * `backup` Get a backup copy of your config. example: `gtm config backup > <your file name>.yml`
   
**tasks** Run task commands
 - options:
   * `list, l` List all tasks for this project. example: `gtm tasks list`
   * `add, a` Add a new task to the list of tasks. example: `gtm tasks add "<Task Name>" --description="<Task Description>"`
   * `claim, c` Claim a task as something you are doing, uses your config user value. example `gtm tasks claim <project number from list>`
   * `mark, m` Mark a task as complete, must be a task that you have claimed. example `gtm tasks mark <project number from list>`

**help**

## Installation

Right now there are no distribution copies, install with go! :D `go get github.com/ccutch/task-manager`


Thank you for checking out this project, feel free to contribute! 

MIT License
