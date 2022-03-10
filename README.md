## Spectro Cloud Online Assessment Solutions

Question : 1
------------

The objective of this question is to implement a custom listenable response struct.

These are response types where the operation takes more time to complete
and there are more than one result objects are present as part of the response.

In that situation, the caller can choose to wait till it is completed
(or) caller can resume the next operation with subscribing the events using the response object in a non-blocking way.

The interface Response has methods like addResult, setError, setCompleted, subscribe(...)

Eg:
(i) Where the caller decides to wait till the whole operation is completed (a blocking wait)
response := someOperation.execute();
response.waitTillComplete();

(ii) Where the caller decides not to wait, but interested in listening to the events happening within the response (non-blocking)
response := someOperation.execute();
response.subscribe(functional hook to listen to new results, functional hook to listen for error, functional hook for completion)
//exit the main process or function as subscribing is non-blocking

The task is to implement the reusable response struct providing the above-mentioned functionalities. Also, have a sample code to demonstrate both blocking wait and non-blocking execution



Question : 2
------------

Imagine a simple application where one can plug different database engines like Postgres or MongoDb as the persistence service.
Provide a sample code implementation where user can decide the database engine at the configuration level.

The task is to implement any model objects which can be persisted to the database, however, the database engine can be decided by the user configuration.



Question : 3
------------
Implement a Task queue. Have a Task struct and create Task objects and push them to a queue.
Have a go-routine which periodically checks the tasks queue and inspect if the task is completed or not.
If the task is completed then remove it from the queue, if not completed push back into the queue.
If the task is not completed after a certain amount of time then it should be removed from the queue and marked as a timeout.
```
type Task struct {
   Id string
   IsCompleted boolean // have a random function to mark the IsCompleted after a random period
   Status string //completed, failed, timeout
}
```
Implement the above mentioned logic with proper error handling and write Go unit test cases to verify the scenarios with complete code coverage
