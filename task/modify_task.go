package task

import (
	"fmt"
	"strconv"
	"strings"
)

var taskFields[] string
var updateFields[] string

func ModifyCommandHandler() {
	//We prompt a multi-argumented response, which is split up in the taskFields array
	fmt.Printf("[r [number]] - Remove a task by its number\n[u [number]] - Update a task by its number\n")
	infoScanner.Scan()

	if (isScannerErr(*infoScanner)) {
		return 
	}

	taskFields = strings.Fields(infoScanner.Text())
	//Converting our task number into an int
	taskNum, err := strconv.Atoi(taskFields[1])
	//While error checking, we also limit our number between 1 and what TaskArr contains
	if err != nil || taskNum < 1 || taskNum > len(TaskArr) {
		fmt.Printf("Invalid task number: %s\n", err)
		return 
	}

	//Handling our given modify command
	switch taskFields[0] {
	case "u":
		UpdateTask(taskNum-1)
	case "r": 
		RemoveTask(taskNum-1)
	default:
		return
	}
}

func UpdateTask(taskNum int) {
	fmt.Printf("Updating a task:\n[d [description]] - Update description\n[p [priority]] - Update priority\n[c [yes/no, case insensitive]] - Update completion status\n")
	infoScanner.Scan()

	if (isScannerErr(*infoScanner)) {
		return 
	}

	//Similar to the modify command options, we make a Fields array with the command + updated info
	updateFields = strings.Fields(infoScanner.Text())
	switch updateFields[0] {
	case "d":
		//For our description, we instead rely on slicing infoScanner.Text(), since strings.Fields is delimited by whitespace.
		newTaskDescription := infoScanner.Text()
		newTaskDescription = newTaskDescription[2:]
		TaskArr[taskNum].Description = newTaskDescription
	case "p":
		//Conversion to int, error checking before updating
		newTaskPriority, err := strconv.Atoi(updateFields[1]) 
		if (!isValidPriority(err, newTaskPriority)) {
			return
		}
		TaskArr[taskNum].Priority = newTaskPriority
	//For our boolean field, we lowercase the string input, and if it doesn't match yes/no it's assumed to be invalid.
	case "c":
		switch strings.ToLower(updateFields[1]) {
		case "yes":
			TaskArr[taskNum].IsComplete = true
		case "no":
			TaskArr[taskNum].IsComplete = false
		default:
			fmt.Printf("Invalid completion status: %q\n", updateFields[1])
		}
	}
}

func RemoveTask(taskNum int) {
	//We reappend the task array by excluding the task to remove; it resizes automatically
	TaskArr = append(TaskArr[:taskNum], TaskArr[taskNum+1:]...)  
}
