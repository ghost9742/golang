package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	statePopulation := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
	}
	fmt.Println(statePopulation)
	fmt.Println(statePopulation["Texas"])
	statePopulation["Texas"] = 123123123
	fmt.Println(statePopulation["Texas"])
	statePopulation["Georgia"] = 10
	fmt.Println(statePopulation)
	delete(statePopulation, "Texas")
	fmt.Println(statePopulation)

	_, ok := statePopulation["Ohio"]
	fmt.Println(ok)
	fmt.Println(len(statePopulation))

	sp := statePopulation
	delete(sp, "California")
	fmt.Println(statePopulation)

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon",
		companions: []string{
			"Liz",
			"Jo",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)

	aDoctor2 := struct{ name string }{name: "John"}
	fmt.Println(aDoctor2)
}
