package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Concurrency
	Concurrency is creating multiple process that execute independently

	Rob pike:
	concurrency is the composition of independently executing process, while
	parallelism is the simultaneous execution of (possibly related)
	computations. Concurrency is about dealing woth lots of thing at once ,
	Whereas parallelism is about doing lots of thing at once.


	GO routine vs Threads
	a) go routine are lighter than OS threads
	b) go manages goroutne whereas threads are managed by OS
		manages means creating , updating locking unlocking killing
	c) switching between threads are costly operzation
		whereas same for the go routine, but as go manages routine, i has less
		switching
	d) Concurrenct model used in go
		Actor model
		Communicatting Seqential Process (CSP)



Parallelism

*/

func blocking_func() {
	// here serial executon will take place and wait for
	// operations like , IO network .. etc

	func() {
		time.Sleep(10 * time.Second)
		fmt.Println("Hello Blocking")
	}()

	func() {
		fmt.Println("Hello blocking from 2nd")
	}()
}

func go_routine_async() {

	// here go key word make a func go routine
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)
	go func() {

		time.Sleep(10 * time.Second)
		fmt.Println("Hello  async")
		defer waitGrp.Done()
	}()

	go func() {

		fmt.Println("Hello async from 2nd")
		defer waitGrp.Done()
	}()
	waitGrp.Wait()
}

func main() {

	blocking_func()

	go_routine_async()

}
