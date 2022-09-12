package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) { //disini dilihat fungsi pembagi mempunyai 2 return, satunya int dan satunya adalah error
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	// merupakan salah satu cara untuk mengecek apakah terjadi error pada fungsi program
	// konsepnya mirip sebagai catch exception pada php
	// contoh
	// hasil, err := pembagi(100, 10) //contoh tidak error
	hasil, err := pembagi(100, 0) //contoh error
	if err == nil {
		fmt.Println("Hasil: ", hasil)
	} else {
		fmt.Println(err.Error())
	}
}
