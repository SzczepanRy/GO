package main

import (
	"fmt"
	"math/rand"
	"time"
)

// could call int a generation function from a grneration funvc
// basicly creating a chanal that holds the values of the two chanels, witch is non blocking
func main() {

	c := fanIn(boring("joe"), boring("ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("for ended")

}
func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- <-input1
		}
	}()
	go func() {
		for {
			ch <- <-input2
		}
	}()
	return ch
}

func boring(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}
