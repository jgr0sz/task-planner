package main

import (
	"fmt"
	"jgr0sz/taskplanner/storage"
	"jgr0sz/taskplanner/task"
)

var wantsToExit bool

func main() {
	//Adding previously-saved tasks to the array being used
	if storage.TasksExist() {
		if storage.LoadTasks() != nil {
			fmt.Printf("Error retrieving saved tasks.")
		}
	}
	cmd := " "
	//Main loop
	for !wantsToExit {
		fmt.Print("[a] - add a new task\n[v] - view tasks\n[e] - exit\n[m] - modify: update/ an existing task\n->")
		fmt.Scan(&cmd)
		handleCmd(cmd)
	}
}

//Handler function for cmd input
func handleCmd(cmd string) {
	switch cmd {
	case "a":
		task.AddTask()
	case "e":
		storage.SaveTasks()
		wantsToExit = true
	case "v":
		task.ViewTask()
	case "m":
		task.ModifyCommandHandler()
	default:
		fmt.Printf("No command available for %q", cmd)
	}
}
