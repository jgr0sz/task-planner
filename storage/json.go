package storage

import (
	"encoding/json"
	"fmt"
	"jgr0sz/taskplanner/task"
	"log"
	"os"
)

var savePath = "data/tasks.json" 

//Saves TaskArr into one large JSON file for persistent storage
func SaveTasks() {
	//Converting/saving TaskArr in JSON format
	marshalledTask, err := json.MarshalIndent(task.TaskArr, "", "  ")
	if err != nil {
		fmt.Printf("Unexpected error while saving tasks.")
		panic(err)
	} 
	err = os.WriteFile(savePath, marshalledTask, 0644)
	if err != nil {
		fmt.Printf("Unexpected error while writing to save file.")
		log.Fatal(err)
	}
}

//Retrieves the saved taskArr JSON file, unmarshalling and adding its data to our main TaskArr array.
func LoadTasks() error {	
	taskData, err := os.ReadFile(savePath)
	if err != nil {
		fmt.Printf("Unexpected error while reading from save file.")
		log.Fatal(err)
	}
	
	var loadedTaskArr []task.Task
	err = json.Unmarshal(taskData, &loadedTaskArr)
	if err != nil {
		fmt.Printf("Unexpected error while retrieving task data.")
		log.Fatal(err)
	}
	task.TaskArr = loadedTaskArr
	return nil
}

//Helper function during save file retrieval to check if there are stored tasks before unmarshalling
func TasksExist() bool {
	_, err := os.Stat(savePath)
	return !os.IsNotExist(err)
}
