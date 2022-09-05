package main

import (
	"fmt"
)

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

	// menghapus data key pada variable map
	delete(chicken, "januari")

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

	// mengecek apakah key itu ada pada map atau tidak
	var result, isExist = dog["tinggi"]
	if isExist {
		fmt.Println("tinggi exist: ", result)
	} else {
		fmt.Println("tinggi not exist")
	}

	// kombinasi slice dan map
	var fishs = []map[string]string{ //disini terlihat penggunaan [] yang menandakan kita membuat variable slice
		map[string]string{"name": "ikan 1", "gender": "male"},
		map[string]string{"name": "ikan 2", "gender": "female"},
		map[string]string{"name": "ikan 3", "gender": "male"},
	}

	// looping lagi
	for _, fish := range fishs {
		fmt.Println(fish["gender"], fish["name"])
	}
}
