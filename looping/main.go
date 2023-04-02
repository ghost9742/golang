package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}

Loop:
	for k := 1; k <= 3; k++ {
		for l := 1; l <= 3; l++ {
			fmt.Println(k * l)
			if k*l >= 3 {
				break Loop
			}
		}
	}

	s := []int{1, 2, 3}
	fmt.Println(s)
	for k, v := range s {
		fmt.Println(k, v)
	}

	s2 := "Hello Go!"
	for k2, v2 := range s2 {
		fmt.Println(k2, string(v2))
	}

	for _, v3 := range s2 {
		fmt.Println(v3)
	}
}
