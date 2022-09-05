package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// fungsi print
	var names = []string{"foo", "bar"}
	printMessage("halo", names)

	// fungsi calculate
	var diameter float64 = 15
	var luas, keliling = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", luas)
	fmt.Printf("keliling lingkaran\t\t: %.2f \n", keliling)
}

// fungsi untuk print text
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// fungsi untuk menghitung diameter
func calculate(d float64) (float64, float64) {
	var luas = math.Pi * math.Pow(d/2, 2)
	var keliling = math.Pi * d

	return luas, keliling

}
