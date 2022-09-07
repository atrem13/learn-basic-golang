package main

import "fmt"

func main() {
	// merupakan metode dimana kita bisa mengakses memori tempat variable disimpan
	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after  :", number) // 10
}

func change(original *int, value int) {
	*original = value
}
