package main

import "fmt"

func main() {


	//arrays

	//array of 3 ints

	// var ages [3]int = [3]int{22, 21, 20}

	var ages = [3]int{22, 21, 20}
	fmt.Println(ages, len(ages))

	ages[1]=33
	//cant change the length of an array

	fmt.Println(ages, len(ages))


	names := [4]string{"tom","mark","jeff","jon"}
	fmt.Println(names, len(names))


	//slices

	var numbers = []int{10,20,30}
	fmt.Println(numbers)

	numbers = append(numbers,40)
	fmt.Println(numbers)

	//slice ranges

	rangeOne := names[1:3]
	rangeTwo := names[1:]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	rangeOne = append(rangeOne,rangeTwo...)
	fmt.Println(rangeOne)


}