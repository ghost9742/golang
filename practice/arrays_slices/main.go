package main

import "fmt"

func main() {
	grades := [...]int{97, 53, 21}
	var students [3]string
	fmt.Printf("Grades: %v", grades)
	fmt.Printf("Students: %v", students)
	students[0] = "Lisa"
	fmt.Printf("Studends: %v", students)
	fmt.Printf("Number of Students: %v\n", len(students))
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// slices
	c := []int{1, 2, 3}
	fmt.Println(c)
	fmt.Printf("Length %v\n", len(c))
	fmt.Printf("Capacity %v\n", cap(c))

	g := []int{1, 2, 3, 4, 5, 65}
	h := g[3:]
	k := g[3:6]
	fmt.Println(h)
	fmt.Println(k)

	z := make([]int, 3, 100)
	fmt.Printf("length %v", len(z))
	fmt.Printf("cap %v", cap(z))

	g = append(g, 12)
	fmt.Println(g)
	g = g[0:3]
	fmt.Println(g)

	r := append(g[:2], g[3:]...)
	fmt.Println(r)
}
