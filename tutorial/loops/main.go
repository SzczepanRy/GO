package main

import "fmt"

func main() {

	// x := 0
	// //basicly while
	// for x < 5 {

	// 	fmt.Println("current x",x)
	// 	x++
	// }


	//reg for

	// for i :=0 ;i<5;i++{
	// 	fmt.Println("current i",i)
	// } 

	names := []string{"a","b","c","d"}

		// for i:=0; i<len(names);i++{
		// 	fmt.Println(names[i])
		// }

		// for index,val := range names{
		// 	fmt.Println(index , val)

		// }

		//or if index is not neadesd 


		for _,val := range names{
			fmt.Println(val)
			//does not update 
			//val is copu kike
			val = "new"
		}
		fmt.Println(names)

}
