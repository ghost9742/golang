package main

import "fmt"

func main() {
	var a int = 42
	var b *int
	b = &a
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)

	j := [3]int{1, 2, 3}
	k := &j[0]
	c := &j[1]
	fmt.Printf("%v %p %p\n", j, k, c)
}
