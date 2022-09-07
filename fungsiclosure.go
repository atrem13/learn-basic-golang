package main

import "fmt"

func main() {
	// merupakan fungsi yang nilai akhirnya disimpan pada sebuah variabel
	// kita bisa membuat fungsi didalam fungsi dengan konsep ini
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 5, 6, 6, 6, 7, 8, 9}
	var min, max = getMinMax(numbers)
	fmt.Printf("data: %v\n", numbers)
	fmt.Printf("min: %v\n", min)
	fmt.Printf("max: %v\n", max)
}
