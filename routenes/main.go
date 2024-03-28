package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {

	// //executes sync
	// compute(5)
	// compute(5)

	// executes async
	//turns the functions into go routenes
	go compute(5)
	go compute(5)

	fmt.Scanln()

}
