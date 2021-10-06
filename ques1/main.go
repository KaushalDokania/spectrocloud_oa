package main

import (
	"errors"
	"log"
)

type Response interface {
	addResult(Result)
	setError(error)
	setCompleted(bool)
	waitTillComplete()
	subscribe(func(Result), func(error), func(bool))
}

type Result struct {
	input  int
	output int
}

type CalcResponse struct {
	resultChannel    chan Result
	errorChannel     chan error
	completedChannel chan bool
	isCompleted      bool
	errorMsg         string
	results          []Result
}

type Operation struct {
	num   int
	input []int
	fn    func(int) int
}

func fib(num int) int {
	if num < 2 {
		return num
	}
	return fib(num-1) + fib(num-2)
}

func calc(num int) (*Result, error) {
	if num >= 50 {
		return nil, errors.New("TIMEOUT")
	}
	return &Result{input: num, output: fib(num)}, nil
}

func (op *Operation) execute() *CalcResponse {
	// Calculating
	resp := &CalcResponse{
		resultChannel:    make(chan Result, 10),
		errorChannel:     make(chan error),
		completedChannel: make(chan bool),
		results:          []Result{},
	}

	log.Println("calling goroutine for calculation...")
	go func(resp *CalcResponse) {
		for _, v := range op.input {
			res, err := calc(v)
			if err != nil {
				log.Println("--> TIMEOUT occured..")
				resp.errorChannel <- err
			} else {
				log.Println("pushing", res, "to channel")
				resp.resultChannel <- *res
			}
		}
		close(resp.resultChannel)
		log.Println("pushing true to completedChannel.., len: ", len(resp.resultChannel))
		resp.completedChannel <- true
		// resp.setCompleted(true)
	}(resp)

	log.Println("returning new response...")
	return resp
}

func (r *CalcResponse) addResult(res Result) {
	r.results = append(r.results, res)
}

func (r *CalcResponse) setError(err error) {
	r.errorMsg = err.Error()
}

func (r *CalcResponse) setCompleted(val bool) {
	if r.completedChannel == nil {
		r.completedChannel = make(chan bool)
	}
	log.Println("Setting completed to true")
	r.isCompleted = val
	// r.completedChannel <- true
}

func (resp *CalcResponse) waitTillComplete() {
	done := false
	for {
		select {
		case res, ok := <-resp.resultChannel:
			{
				log.Println("received ", res, " adding to response")
				resp.addResult(res)
				if ok == false { // channel is closed
					log.Println("--> channel was closed, so this was last result, so returning")
					done = true
					break
				}
			}
		case err, _ := <-resp.errorChannel:
			resp.setError(err)
		}

		if done {
			break
		}
	}
	log.Println("waiting for completion..., len: ", len(resp.completedChannel))
	resp.setCompleted(<-resp.completedChannel)
}

func (resp *CalcResponse) subscribe(fnAdd func(Result), fnErr func(error), fnDone func(bool)) {
	go func() {
		for {
			res := <-resp.resultChannel
			log.Println("received new result, adding to response")
			fnAdd(res)
		}
	}()

	go func() {
		for {
			err := <-resp.errorChannel
			fnErr(err)
		}
	}()

	fnDone(<-resp.completedChannel)
}

func main() {
	log.Println("Hello World!!")
	op := Operation{num: 10, input: []int{10, 20, 45, 50}}
	resp := op.execute()
	log.Printf("%+v", resp)
	resp.waitTillComplete()
	// resp.subscribe(resp.addResult, resp.setError, resp.setCompleted)
	log.Printf("%+v", resp)
	// ch := make(chan Result, 10)
}
