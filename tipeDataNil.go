package main

import "fmt"

func main() {
	// nil merupakan cara inisiasi variable dengan nilai kosong
	// hanya dapat digunakan untuk beberapa jenis tipe data seperti fungsi, interface, map, slice, pointer, channel
	// contoh
	var example = map[string]string{
		"nama": "saya",
		"umur": "19",
	}
	// var example2 map[string]string
	if example == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(example)
	}
}
