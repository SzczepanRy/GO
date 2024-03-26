package main

import (
	"fmt"
	"math"
)

func main() {

	names := []string{"a","b","c","d"}
	cycleNames(names,greet)
	cycleNames(names,goodbye)

	p :=circleArea(2.5)
	fmt.Println(p)
}

func greet(name string) {
	fmt.Printf("hello %v\n",name)
}
func goodbye(name string) {
	fmt.Printf("goodbye %v\n",name)
}

func cycleNames(names []string,f func(string) ){
	for _,name := range names{
		f(name)
	}
}

func circleArea(r float64) float64{
	return math.Pi *r*r
}