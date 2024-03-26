package main

import "fmt"

func main() {

	num:=15
	fmt.Println(num<=50)
	fmt.Println(num>=50)
	fmt.Println(num==45)
	fmt.Println(num!=50)

	if num <30{
		fmt.Println("num is lower than 30")
	}else if num <40 {
		fmt.Println("num is lower than 40")
	}else{
		fmt.Println("num higer than 40")
	}

	names:=[]string{"jon","jim","jack","jill"}

	for index , value := range names{
		if index == 1 {
			fmt.Println("current index == 1 and val == ",value)
			//brakeout of this iteration and begin the next one
			continue

		}
		if(index ==2 ){
			fmt.Println("braking")
			break
		}
		fmt.Println("index",index)
	}
}
