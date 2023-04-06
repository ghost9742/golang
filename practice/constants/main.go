package main

import "fmt"

const (
	i = iota
	j = iota
	k
)

func main() {
	const myConst int = 42
	fmt.Printf("%v, %T", myConst, myConst)
	const b string = "ad"
	const c float32 = 42.2
	const d bool = true
	fmt.Println(b, c, d)
	fmt.Println(i, j, k)
}
