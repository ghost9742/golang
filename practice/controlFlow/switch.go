package main

import "fmt"

func main() {
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}

	j := 5
	switch {
	case j <= 10:
		fmt.Println("less than 10")
		fallthrough
	case j <= 20:
		fmt.Println("less and equal than 20")
	default:
		fmt.Println("not less than 10")
	}

	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is an float65")
	default:
		fmt.Println("i is another type")
	}
}
