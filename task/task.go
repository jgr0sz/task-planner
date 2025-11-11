package task

import (
	"bufio"
	"fmt"
	"os"
)

//Simple definition of our tasks
type Task struct {
	Description string
	IsComplete bool
	Priority int
}

//Slice (dynamic array) for our tasks
var TaskArr []Task

//Our main scanner
var infoScanner = bufio.NewScanner(os.Stdin)
var taskInfo string
var priorityStr string

//Helper function to turn our IsComplete bool to string output
func statusToString(taskStatus bool) string {
	if (!taskStatus) {
		return "Incomplete"
	}
	return "Complete"
}

//Helper function for a standard bufio.Scanner() error check
func isScannerErr(infoScanner bufio.Scanner) bool {
	if infoScanner.Err() != nil {
		fmt.Printf("Input error: %s\n", infoScanner.Err())
		return true 
	}
	return false
}

//Helper function to make sure we didn't encounter an Atoi error from input conversion and that priority is within range
func isValidPriority(err error, priority int) bool {
	if err != nil || priority > 10 || priority < 1 {
		fmt.Printf("Invalid priority value: %s\n", priorityStr)
		return false
	}
	return true
}
