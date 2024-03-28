package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	//unbuffered chanal
	// dataChan := make(chan int)
	//buffered channel
	dataChan := make(chan int, 1) //1 is the number of valus that can be stored in the channel

	//bouth of the send and get lines have to execute simutainiesly

	//case for unbuffered channel
	// go func() {
	// 	//send data to channel

	// 	dataChan <- 710
	// }()

	//case for buffered
	dataChan <- 710

	//get data from channel

	n := <-dataChan

	fmt.Println("data form channel ", n)

	////////////////////////////////////////
	dataCh := make(chan int)
	//bg thread
	go func() {

		//track when routene is complete
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			//new go rotene
			go func() {
				defer wg.Done()
				data := work()
				dataCh <- data
			}()

		}
		wg.Wait()
		close(dataCh)
	}()
	//mainthred
	for n := range dataCh {
		fmt.Println("data from channel ", n)
	}
}
func work() int {
	time.Sleep(time.Second)
	return rand.IntN(100)
}
