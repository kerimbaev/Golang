package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker #%v started the job #%v", id, job)
        time.Sleep(time.Second / 2)
        fmt.Printf("Worker #%v finished the job #%v", id, job)
        results<-j*2
    }
}

const nimbJobs = 10
const numbWorkers = 8

func main() {
	jobs := make(chan int, numbJobs)
	results := make(chan int, numbJobs)

  for i := 1; i <= numbWorkers; i++ {
		go worker(i, jobs, results)
	}
  for i := 1; i <= numbJobs; i++ {
		jobs <- i
	}
	close(jobs)
  for i := 1; i <= numbJobs; i++ {
		<-results
	}
}
