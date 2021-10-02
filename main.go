/*
Question : 3
------------
Implement a Task queue. Have a Task struct and create Task objects and push them to a queue.
Have a go-routine which periodically checks the tasks queue and inspect if the task is completed or not.
If the task is completed then remove it from the queue, if not completed push back into the queue.
If the task is not completed after a certain amount of time then it should be removed from the queue and marked as a timeout.

type Task struct {
   Id string
   IsCompleted boolean // have a random function to mark the IsCompleted after a random period
   Status string //completed, failed, timeout
}

Implement the above mentioned logic with proper error handling and write Go unit test cases to verify the scenarios with complete code coverage
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	Id          string
	IsCompleted bool   // have a random function to mark the IsCompleted after a random period
	Status      string //completed, failed, timeout
}

const statusCompleted string = "completed"
const statusFailed string = "failed"
const statusTimeout string = "timeout"

func main() {
	fmt.Println("Hello World!!")
	q := make(chan *Task, 10)

	tasks := []Task{
		{Id: "101", IsCompleted: false, Status: statusCompleted},
		{Id: "102", IsCompleted: false, Status: statusCompleted},
		{Id: "103", IsCompleted: false, Status: statusCompleted},
		{Id: "104", IsCompleted: false, Status: statusCompleted},
		{Id: "105", IsCompleted: false, Status: statusCompleted},
	}

	go func(tasks []Task) {
		var indices []int
		for {
			idx := rand.Intn(len(tasks))
			indices = append(indices, idx)
			fmt.Printf("\nrand index: %d", idx)
			if !tasks[idx].IsCompleted {
				fmt.Printf("\ncompleting the task: %s", tasks[idx].Id)
				tasks[idx].IsCompleted = true
			}
			time.Sleep(2 * time.Second)
		}
	}(tasks)

	// submitting all tasks to queue
	for index, _ := range tasks {
		fmt.Println("main: Adding task:", tasks[index].Id, " to queue")
		q <- &tasks[index]
	}

	count := 0
	var list []string

	for len(q) > 0 {
		t := <-q
		if t.IsCompleted {
			list = append(list, t.Id)
			fmt.Printf("\ntask: %s is [completed](Tasks completed: %v), len(queue): %d", t.Id, list, len(q))
			count++
		} else {
			q <- t // push back if not completed
			fmt.Printf("\ntask: %s is NOT completed, pushing back to queue...(Tasks completed: %v), len(queue): %d", t.Id, list, len(q))
		}
		fmt.Println()
		// time.Sleep(1 * time.Second)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\nThank you!!")
}
