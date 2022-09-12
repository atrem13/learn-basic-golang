package main

import (
	"fmt"
	"reflect"
)

func main() {
	// reflect merupakan suatu metode untuk mengetahui tipe data ataupun nilai dari sebuah variable
	var number = 23
	// mengecek nilai dari variable
	var reflectValue = reflect.ValueOf(number)

	// menampilkan tipe data dari variable
	fmt.Println("tipe  variabel :", reflectValue.Type())

	// mengecek apakah tipe data dari variable itu adalah X
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

}
