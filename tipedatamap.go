package main

import "fmt"

func main() {
	// inisiasi variable map
	// map mirip dengan array object tapi tidak bisa membaca data berdasarkan indexnya
	var chicken map[string]int //kalau hanya begini otomatis error
	chicken = map[string]int{} //harus ditambah ini baru bener
	chicken["januari"] = 50
	chicken["februari"] = 60

	// jika membuat key yang sama maka value yang lama akan di replace
	// cara lain inisiasi variable

	var dog = map[string]int{"berat": 10, "tinggi": 5}
	var cat = map[string]int{
		"berat":  10,
		"tinggi": 5, //inget tetep ada koma di item terakhir
	}

	// fmt.Println("januari ", chicken["januari"])
	// fmt.Println("januari ", chicken["januari"])

	// menampilkan data dengan looping
	for key, val := range chicken {
		fmt.Println(key, " \t:", val)
	}
	for key, val := range dog {
		fmt.Println(key, " \t:", val)
	}
	for key, val := range cat {
		fmt.Println(key, " \t:", val)
	}
}
