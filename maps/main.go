package main

import "fmt"

func main() {

	m := map[string]int{
		"id":     2,
		"age":    3,
		"number": 5,
	}
	for i, v := range m {
		fmt.Printf("index %q value %v\n", i, v)
	}
	fmt.Println("=====================================")

	fmt.Println("using single value", m["id"])

	//updata val
	fmt.Println("=====================================")

	if v, ok := m["id"]; ok {
		fmt.Println("the value is present", v)
	} else {
		fmt.Println("the value is not present")
	}
	//adding keys

	fmt.Println("=====================================")
	//should check it with ok

	m["newId"] = 12
	for i, v := range m {
		fmt.Printf("index %q value %v\n", i, v)
	}

	//delete keys
	fmt.Println("=====================================")

	//should check it with ok
	delete(m, "id")

	for i, v := range m {
		fmt.Printf("index %q value %v\n", i, v)
	}

}
