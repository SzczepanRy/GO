package main

import "fmt"

func main(){
	fmt.Print("A")
	fmt.Print("A\n")

	ver1 := "yow"
	fmt.Print(ver1," asd\n")

	ver2 := "whats"
	ver3 := "good"

	fmt.Printf("%v , %v %v\n",ver1,ver2,ver3)
	fmt.Printf("%v , %v %q\n",ver1,ver2,ver3)
	fmt.Printf("%v , %v %T\n",ver1,ver2,ver3)
	fmt.Printf("float %f\n",0.5)
	fmt.Printf("float %0.1f\n",0.5)

	title := fmt.Sprintf("%v , %v %v\n",ver1,ver2,ver3)
	println(title)
}