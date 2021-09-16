package main

import (
	"fmt"
	"github.com/cookienyancloud/workerPool/pool"
	"github.com/cookienyancloud/workerPool/work"
	"log"
	"os"
)
	const WORKER_COUNT = 5
	const JOB_COUNT = 100

func main() {
	err := os.Setenv("DEBUG", "true")
	if err != nil {
		fmt.Printf("env:%s", err.Error())
	}

	log.Println("starting application...")
	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool

	for i, job := range work.CreateJobs(JOB_COUNT) {
		collector.Work <- pool.Work{Job: job, ID: i}
	}
}
