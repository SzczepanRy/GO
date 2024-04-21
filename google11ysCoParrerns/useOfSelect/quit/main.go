package main

import (
	"fmt"
	"math/rand"
)

//channal inside a channal

type Message struct {
	str  string
	wait chan bool
}

func main() {

	quit := make(chan string)
	c := boring("joo", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}

	quit <- "bye1"
	fmt.Println("for ended", <-quit)

}

// func fanIn(input1, input2 <-chan Message) <-chan Message {
// 	ch := make(chan Message)
// 	//pass arguments on one go routene
// 	go func() {
// 		for {

// 			select {
// 			case s := <-input1:
// 				ch <- s
// 			case s := <-input2:
// 				ch <- s
// 			}

// 		}
// 	}()

//		return ch
//	}
func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
			case <-quit:
				//clenup()
				quit <- "bye2"
				return
			}
		}
	}()
	return c
}
