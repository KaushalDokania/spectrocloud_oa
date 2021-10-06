package main

import (
	"errors"
	"log"
	"sync"
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

	go func(resp *CalcResponse) {
		log.Println("Calculating...")
		for _, v := range op.input {
			res, err := calc(v)
			if err != nil {
				resp.errorChannel <- err
			}
			log.Println("pushing", res, "to channel")
			resp.resultChannel <- *res
			log.Println("pushed", res, "to channel")
			// resp.addResult(*res)
		}
		// log.Println("iterated all inputs..")
		// time.Sleep(5 * time.Second)

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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			res := <-resp.resultChannel
			log.Println("received ", res, " adding to response")
			resp.addResult(res)
		}
	}()
	wg.Wait()
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
	op := Operation{num: 10, input: []int{10, 25, 45}}
	resp := op.execute()
	log.Println(resp)
	resp.waitTillComplete()
	// resp.subscribe(resp.addResult, resp.setError, resp.setCompleted)
	log.Println(resp)

	// ch := make(chan Result, 10)
}
