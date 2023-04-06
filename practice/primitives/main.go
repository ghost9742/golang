package main

import "fmt"

func main() {
	// var n bool = true
	// fmt.Printf("%v, %T", n, n)
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T", n, n)
	fmt.Printf("%v, %T", m, m)
	var z uint16 = 42
	fmt.Printf("%v, %T", z, z)
	a := 10 // 1010
	b := 3  // 0011
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a << 3)
	fmt.Println(a >> 3)
	var k complex64 = 1 + 2i
	var j complex64 = 2 + 3i
	fmt.Printf("%v, %T", k, k)
	fmt.Println(k + j)
	fmt.Println(real(k) + real(j))

	s := "this is a string"
	fmt.Printf("%v, %T", s, s)
	fmt.Print("\n")
	fmt.Println(string(s[1]))

	//r := 'a'
	var r rune = 'a'
	fmt.Printf("%v, %T", r, r)
}
