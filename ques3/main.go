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
	"log"
	"math/rand"
	"time"
)

type Task struct {
	Id          string
	IsCompleted bool   // have a random function to mark the IsCompleted after a random period
	Status      string //completed, failed, timeout
}

type MyTask struct {
	started time.Time
	task    *Task
}

const statusCompleted string = "completed"
const statusFailed string = "failed"
const statusTimeout string = "timeout"
const timeout time.Duration = time.Duration(20000 * time.Millisecond)

func getDummyTasks() []Task {
	return []Task{
		{Id: "101", IsCompleted: false, Status: statusFailed},
		{Id: "102", IsCompleted: false, Status: statusFailed},
		{Id: "103", IsCompleted: false, Status: statusFailed},
		{Id: "104", IsCompleted: false, Status: statusFailed},
		{Id: "105", IsCompleted: false, Status: statusFailed},
		{Id: "106", IsCompleted: false, Status: statusFailed},
		{Id: "107", IsCompleted: false, Status: statusFailed},
		{Id: "108", IsCompleted: false, Status: statusFailed},
		{Id: "109", IsCompleted: false, Status: statusFailed},
	}
}
func main() {
	tasks := getDummyTasks()

	go func(tasks []Task) {
		var indices []int
		for {
			idx := rand.Intn(len(tasks))
			indices = append(indices, idx)
			log.Printf("rand index: %d", idx)
			if !tasks[idx].IsCompleted {
				log.Printf("completing the task: %s", tasks[idx].Id)
				tasks[idx].IsCompleted = true
			}
			time.Sleep(2 * time.Second)
		}
	}(tasks)

	q := make(chan *MyTask, 10)

	// submitting all tasks to queue
	log.Println("main: Adding tasks to the queue")
	for index, _ := range tasks {

		q <- &MyTask{started: time.Now(), task: &tasks[index]}
	}
	log.Println("main: Tasks added to the queue, now processing te queue")

	var completedTasks []string
	for len(q) > 0 {
		mt := <-q
		t := mt.task
		if t.IsCompleted {
			t.Status = statusCompleted
			completedTasks = append(completedTasks, t.Id)
			log.Printf("task: %s is [completed](Tasks completed: %v), len(queue): %d", t.Id, completedTasks, len(q))
		} else {
			if time.Now().Sub(mt.started) > timeout {
				log.Printf("TIMEOUT: discarding id: %s", t.Id)
				t.Status = statusTimeout
				continue
			}
			q <- mt // push back if not completed
			log.Printf("task: %s is NOT completed, pushing back to queue...(Tasks completed: %v), len(queue): %d", t.Id, completedTasks, len(q))
		}
		log.Println()
		time.Sleep(500 * time.Millisecond)
	}

	log.Printf("%+v", tasks)
	log.Println("Thank you!!")
}
