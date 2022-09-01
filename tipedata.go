// data number
// integer
// int8
// int16
// int32
// int64

// unt berarti tidak ada bilangan negatif
// uint8
// uint16
// uint32
// uint64

// floag
// float32
// float64
// complex64
// complex128

// boolean / bool
// true or false

// string
// kumpulan huruf/karakter

package main

import "fmt"

func main() {
	// angka
	// fmt.Println("satu =", 1)
	// fmt.Println("dua =", 2)
	// fmt.Println("tiga koma lima =", 3.5)

	// bool
	// fmt.Println("benar =", true)
	// fmt.Println("salah =", false)

	// string
	fmt.Println(len("ini contoh"))
	fmt.Println("ini contoh dua"[2]) //yang muncul ukuran byte dari huruf 'i'
	fmt.Println("ini contoh tiga")
}
