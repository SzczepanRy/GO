package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new build obj
func newBill(name string) bill {
	b := bill{name: name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
// /recever function limits function to only bill objects
func (b *bill) format() string {
	//b is the copty of the bill (now the bill pointer *bill)

	fs := "bill brakedown: \n"
	var total float64 = 0

	//list items

	for key, val := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", key+":", val)
		total += val
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", b.tip)

	total += b.tip

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}

//update tip

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

//add

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

//save

func (b *bill) saveBill() {
	//byte slice
	data := []byte(b.format())

	err := os.WriteFile(b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("bill saved")
}
