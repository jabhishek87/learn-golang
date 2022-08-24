package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup // should be pointer in real world programs

func main() {

	defer duration(track("Job Took: "))

	// simple function it will take 15 seconds to run
	// someLongRunningJob(5)

	// Same function wil take now 5 seconds
	someLongRunningJobRoutine(5)
}

func someLongRunningJob(t int) {
	for t != 0 {
		log.Println("Sleeping for ", t)
		gotoSleep(t)
		t--
	}
}

func someLongRunningJobRoutine(t int) {
	for t != 0 {
		log.Println("Sleeping for ", t)
		go gotoSleep2(t)
		t--
		wg.Add(1)
	}
	wg.Wait()
}

func gotoSleep2(t int) {
	defer wg.Done()
	time.Sleep(time.Duration(t) * time.Second)
}

func gotoSleep(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
