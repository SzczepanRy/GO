package main

import "fmt"

func main() {

	mybill := newBill("toms bill")

	mybill.addItem("item1", 1.99)
	mybill.addItem("item2", 2.99)
	mybill.addItem("item3", 3.99)
	mybill.addItem("item4", 4.99)
	mybill.updateTip(2.33)

	fmt.Println(mybill.format())
}
