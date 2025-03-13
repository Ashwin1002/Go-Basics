package main

import (
	"fmt"
	// "math/rand"
	"sync" // this is used for go routines
	"time"
)

// they are counters for go routines
var waitGroup = sync.WaitGroup{}

// this lock
// var mutualExecution = sync.Mutex{}
// another mutex is Read/Write mutex
var m = sync.RWMutex{}

var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6"}

var results = []string{}

func main() {

	now := time.Now()
	for i := 0; i < len(dbData); i++ {
		// adding counters for the go routine
		waitGroup.Add(1)
		// this will not stop the loop untils the function completes but rather it will keep moving to next step in the loop
		go dbCall(i)
	}
	// this method wait for counter to go back to zero i.e. all tasks are completed
	waitGroup.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(now))
	fmt.Printf("\nThe results are: %v", results)
}

func dbCall(i int) {
	// Simulate db call delay
	var delay float32 = 2000 // 2 seconds delay
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// fmt.Println("The result from the database is:", dbData[i])
	save(dbData[i])
	log()
	// it decrements the waitgroup counter
	waitGroup.Done()
}

func save(result string) {
	// checking if the lock is performed by other go routines and if true, it will wait until the lock is released and set itself
	m.Lock()
	results = append(results, result)
	// lock is released using this unlock method
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are :%v", results)
	m.RUnlock()
}
