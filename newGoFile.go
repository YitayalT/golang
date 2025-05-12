package main

import "fmt"

var taskOne = "Washing clothes"
var taskTwo = "Sweeping the floor"
var taskThree = "Reading books"
var taskFour = "Cooking dinner"

var taskLists = []string{taskOne, taskTwo, taskThree, taskFour}

func main() {
	fmt.Println("The following are lists of my tasks, I'll do it at night")
	listOfTasks(taskLists)
	fmt.Println("The following are updated task lists")
	taskLists = addTask(taskLists, "Watching movies")
	taskLists = addTask(taskLists, "Learning Go")
	fmt.Println(taskLists)
}

func listOfTasks(taskItems []string) {
	for _, taskItem := range taskItems {
		fmt.Println(taskItem)
	}
}

func addTask(taskItem []string, newTask string) []string {
	var updatedTask = append(taskItem, newTask)
	return updatedTask
}
