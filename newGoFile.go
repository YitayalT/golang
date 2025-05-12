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
	addTask(taskLists, "Watching movies")
}

func listOfTasks(taskItems []string) {
	for _, taskItem := range taskItems {
		fmt.Println(taskItem)
	}
}

func addTask(taskItem []string, newTask string) {
	var updatedTask = append(taskItem, newTask)
	listOfTasks(updatedTask)
}
