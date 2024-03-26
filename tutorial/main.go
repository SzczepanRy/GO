package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	mybill := createBill()

	promptOptions(mybill)
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput(reader, "create a new bill name : ")

	b := newBill(name)

	fmt.Println("created bill : ", b.name)
	return b
}

func getInput(r *bufio.Reader, question string) (string, error) {
	fmt.Print(question)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput(reader, "choose option (a : add , s : save , t : tip): ")

	//	fmt.Println(opt)
	switch opt {
	case "a":

		name, _ := getInput(reader, "whats the name : ")
		price, _ := getInput(reader, "whats the price : ")

		priceFloat, err := strconv.ParseFloat(price, 64)

		if err == nil {
			b.addItem(name, priceFloat)
			fmt.Println("added item")

		} else {
			fmt.Println("price should be a float")
		}
		promptOptions(b)

		// break
	case "s":
		b.saveBill()
		fmt.Println("saved bill : ", b.name)
		// break
	case "t":
		tip, _ := getInput(reader, "whats the tip : ")

		tipFloat, err := strconv.ParseFloat(tip, 64)

		if err == nil {
			b.updateTip(tipFloat)
			fmt.Println("added tip")

		} else {
			fmt.Println("price should be a float")
		}
		promptOptions(b)

		// break
	default:
		fmt.Println("not a valid option")
		promptOptions(b)
	}

}
