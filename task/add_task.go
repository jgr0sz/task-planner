package task

import (
	"fmt"
	"strconv"
)

func AddTask() {
	//Scans text until enter is hit
	fmt.Print("What task would you like to add? ")
	infoScanner.Scan()

	//Checking if input isn't errored
	if (isScannerErr(*infoScanner)) {
		return
	}

	taskInfo = infoScanner.Text()	
	fmt.Print("What priority would you like this task to have (1-10)? ")
	infoScanner.Scan()

	//Checking again if input isn't errored
	if (isScannerErr(*infoScanner)) {
		return
	}
	priorityStr = infoScanner.Text()

	//Converting our priority into an int, checking its validity
	priority, err := strconv.Atoi(priorityStr)
	if (!isValidPriority(err, priority)) {
		return
	}

	//Creates and stores task struct
	newTask := Task {
		Description: taskInfo,
		IsComplete: false,
		Priority: priority,
	}
	TaskArr = append(TaskArr, newTask)
}
