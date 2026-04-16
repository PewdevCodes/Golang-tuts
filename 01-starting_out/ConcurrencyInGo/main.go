/* 
	concurrency in Go

- Concurrency is the ability of a program to perform multiple tasks simultaneously. In Go, concurrency is achieved through goroutines and channels.

- Goroutines are lightweight threads managed by the Go runtime, allowing you to run functions concurrently without blocking the main thread.

- Channels are used for communication between goroutines, enabling them to synchronize and share data safely.	

*/

package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int , results chan<- int) {
/*
jobs <-chan int  --> This means that the worker function will receive jobs from a channel of integers. The worker will read from this channel to get the jobs it needs to process.


results chan<- int --> This means that the worker function will send results to a channel of integers. The worker will write to this channel to send the results of its processing back to the main function or another goroutine that is waiting for the results.
*/
   for j := range jobs {
       fmt.Println("Worker", id, "processing job", j)
       time.Sleep(time.Second) // Simulate time-consuming work by sleeping for a second.
       results <- j * 2 // Send the result back to the results channel. In this case, the worker is simply doubling the job number and sending it back as the result.
   }
}

func main() { 
	const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
 
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
 
    for a := 1; a <= numJobs; a++ {
        fmt.Println("Result:", <-results)
    }
}