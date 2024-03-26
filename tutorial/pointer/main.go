package main

import "fmt"

func updateName(n *string) {

	*n="jerry"
}

func main() {

	name := "tom"

	pointerName := &name

	// updateName(name)

	fmt.Println("mem addres of name ", &name)
	fmt.Println("val of name adress ", *pointerName)
	fmt.Println(name)


	updateName(pointerName)

	fmt.Println("mem addres of name ", &name)
	fmt.Println("val of name adress ", *pointerName)
	fmt.Println(name)


}
