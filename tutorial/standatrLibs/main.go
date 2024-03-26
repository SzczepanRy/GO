package main

import (
	"fmt"
	"sort"
)

func main() {

	// //string package
	// greeting := "hello there"

	// fmt.Println(strings.Contains(greeting,"ll"))
	// fmt.Println(strings.ReplaceAll(greeting,"e","c"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting,"l"))
	// fmt.Println(strings.Split(greeting," "))

	ages := []int{1, 4, 3, 2, 5, 4, 7, 8, 7, 9}

	sort.Ints(ages)
	fmt.Println(ages)

	//if not included than returns the length +1 of the slice

	index := sort.SearchInts(ages,7)

	fmt.Println(index)


	names := []string{"tom","jeff","jon","tim","micheal"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names,"jon"))



}
