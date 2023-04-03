package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello", i)
	}
	sayGreeting("Hello", "Stacy")

	x := sum(1, 2, 3, 4, 5)
	fmt.Println(x)

	d, err := devide(5.3, 2.3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	func() {
		fmt.Println("Hel")
	}()

	g := func() {
		fmt.Println("g function")
	}
	g()

	var o func() = func() {
		fmt.Println("o func")
	}
	o()

	z := greeter{
		greeting: "hello",
		name:     "go",
	}
	z.greet()
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of index is", idx)
}

func sayGreeting(greeting string, name string) {
	fmt.Println(greeting, name)
}

func sum(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func devide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("Cannot devide by zero")
	}
	return a / b, nil
}
