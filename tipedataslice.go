package main

import "fmt"

func main() {
	// slice merupakan potongan / bagian dari array
	// slice dan array saling terhubunga, jadi jika ada perubahan di salah satunya, maka yang lainnya juga akan ikut terpengaruh

	// inisiasi array
	var dataArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// inisiasi slice
	var dataSlice = dataArray[5:8]
	// dataSlice[0] = 66

	fmt.Println(dataSlice)

	var dataAppend = append(dataSlice, 11)
	fmt.Println(dataAppend)

	fmt.Println(dataArray)
}
