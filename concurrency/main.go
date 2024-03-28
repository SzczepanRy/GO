package main

import "fmt"

//example of sincronisation

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//morker is a new thread
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}
}

//only receve from jobs       only send to results
func worker(jobs <-chan int, results chan int) {

	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
