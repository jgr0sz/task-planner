package task

import "fmt"

//Formats and prints out each of the tasks within our Task array
func ViewTask() {
	if len(TaskArr) == 0 {
		fmt.Printf("No tasks currently.\n")
	}

	for i := range TaskArr {
		fmt.Printf("Task %d (%s, Priority %d): \nDescription: %s\n", 
		i+1, statusToString(TaskArr[i].IsComplete), 
		TaskArr[i].Priority, TaskArr[i].Description)  
	}
}
