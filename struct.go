package main

import "fmt"

func main() {
	type student struct {
		name  string
		grade int
	}

	var s1 student
	s1.name = "john"
	s1.grade = 90

	fmt.Println("name: ", s1.name)
	fmt.Println("grade: ", s1.grade)
}
