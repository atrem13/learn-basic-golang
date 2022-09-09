package main

import (
	"fmt"
	"math"
)

// inisiasi interface
type hitung interface { //ini variable interface
	luas() float64
	keliling() float64
}

// inisiasi struct
type lingkaran struct { //ini variable struct
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}
func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}
func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}
func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func main() {
	// sederhananya ini merupakan sebuah cara implementasi design pattern
	// alurnya:
	// buat beberapa kelas dengan fungsi yang sama
	// dari pada import banyak kelas kita bisa gunakan interface sebagai perantara

	var bangunDatar hitung // implementasi penggunaan interface
	bangunDatar = persegi{5}
	fmt.Println("Persegi")
	//fungsi yang bisa diambil harus di decalre pada interface
	// semua anak yang mewarisi interface harus memiliki fungsi yang di declare pada interface
	fmt.Println("Luas: ", bangunDatar.luas())
	fmt.Println("Keliling: ", bangunDatar.keliling())

}
