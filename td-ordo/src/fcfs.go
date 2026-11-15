package main

import "fmt"

var tasks []task

// Ordonnanceur premier arrivé premier servi (first come first served)
func fcfs(newTasks []task) (currentTask *task) {
		tasks = append(tasks,newTasks...)
		fmt.Println(tasks)

		if(len(tasks)==0) {return nil}
		if (tasks[0].Duration<=0){
			tasks = tasks[1:]
		}
		
		if(len(tasks)==0) {return nil}
		currentTask = &tasks[0]
		// fmt.Println(currentTask.ID)
	return
}
