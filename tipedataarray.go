package main

import "fmt"

func main() {
	// ukuran array harus di declare dari awal
	// tipe data dari array juga harus di declare
	// var tampungData [10]string
	// tampungData[0] = "merta"
	// tampungData[1] = "yoga"
	// tampungData[2] = "kocak"

	// var listAngka = [3]int{
	// 	90,
	// 	85,
	// 	80,
	// }

	listKata := [...]string{"ini", "merupakan", "daftar", "kata"}

	// fmt.Println(tampungData[1])
	// fmt.Println(listAngka[1])

	// fungsi array
	// len untuk mengetahui kapasitas dari array
	// fmt.Println(len(listAngka))

	// ga bisa nambah item array melebihi kapasitas yang sudah di declare di awal
	// listKata[5] = "hehe"
	fmt.Println(len(listKata))
	fmt.Println(listKata[3])

}
