package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new build obj
func newBill(name string) bill {
	b := bill{name: name,
		items: map[string]float64{
			"pie":  5.99,
			"cake": 3.99,
		},
		tip: 0,
	}
	return b
}

// format the bill
// /recever function limits function to only bill objects
func (b bill) format() string {
	//b is the copty of the bill

	fs := "bill brakedown: \n"
	var total float64 = 0

	//list items

	for key, val := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n",key+":",val)
		total+= val
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f","total:",total)

	return fs
}