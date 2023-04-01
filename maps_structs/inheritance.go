package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "aus"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b)
}
