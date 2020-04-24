package myjobscheduler

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// from https://www.opsdash.com/blog/job-queues-in-go.html

type Job struct {
	id   int
	name string
}

func setup() {
	fmt.Println("@@@ setup")
}

func cleanup() {
	fmt.Println("@@@ cleanup")
}

func worker(jobChan <-chan Job) {
	for job := range jobChan {
		fmt.Println("runJob", job)
		time.Sleep(100 * time.Millisecond)
	}
}

// possible without blocking. Job is not enqueued in the latter case.
func TryEnqueue(job Job, jobChan chan<- Job) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}
func TestSingleJob(t *testing.T) {
	// setup()
	// fmt.Println("Name=", t.Name())
	// t.Cleanup(cleanup)

	// make a channel with a capacity of 100.
	jobChan := make(chan Job)
	//jobChan := make(chan Job, 100)
	// start the worker
	go worker(jobChan)

	// enqueue a job
	jobChan <- Job{1, "job1"}

	for i := 0; i < 20; i++ {
		//fmt.Println("i=", strconv.Itoa(i))
		//jobChan <- Job{i, "job" + strconv.Itoa(i)}
		if !TryEnqueue(Job{i, "job" + strconv.Itoa(i)}, jobChan) {
			fmt.Println("ERROR at ", i)
			return
		}
	}

	fmt.Println("@@@ TestSingleJob")

}
