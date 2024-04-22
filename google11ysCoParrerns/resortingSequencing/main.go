package main

import (
	"fmt"
	"math/rand"
	"time"
)

//channal inside a channal

type Message struct {
	str  string
	wait chan bool
}

func main() {

	c := fanIn(boring("joe"), boring("ann"))
	for i := 0; i < 10; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("for ended")

}
func fanIn(input1, input2 <-chan Message) <-chan Message {
	ch := make(chan Message)
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
func boring(msg string) <-chan Message {
	ch := make(chan Message)

	waitCh := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			ch <- Message{fmt.Sprintf("%s %d", msg, i), waitCh}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitCh //waits for the change
		}
	}()
	return ch
}
