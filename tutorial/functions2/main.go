package main

import (
	"fmt"
	"strings"
)

func main() {

	a,b:=getInitials("tom scot")
	fmt.Println(a,b)

}

func getInitials(name string) (string, string) {

	

	name = strings.ToUpper(name)

	nameArr := strings.Split(name," ")

	var leters [] string
	for _,val := range nameArr{

		leters = append(leters, val[:1])

	}
	// return str

	if(len(leters)>1){
		return leters[0] , leters[1]
	}
	return leters[0],"_"
}