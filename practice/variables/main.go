package main

import (
	"fmt"
	"strconv"
)

// uppercase exported globally
var k int = 42

var (
	actorName string = "asd"
	companion string = "dsa"
)

func main() {
	//var i int = 42
	i := 42
	fmt.Println(i)
	var j int = 27
	fmt.Printf("%v, %T", j, j)
	fmt.Println(k)
	fmt.Println(actorName, companion)
	s := float32(j)
	fmt.Println(s)
	z := strconv.Itoa(i)
	fmt.Println(z)
}
