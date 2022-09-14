package main

import "fmt"

type Person struct {
	Name, Hobby string
}

func ChangeHobby(person Person) {
	person.Hobby = "Sleep"
}

func main() {
	// ketika kita mengedit variable dalam fungsi maka yang berubah hanyalah variable didalam fungsi, bukan variable diluarnya
	// oleh sebab itu sebaiknya ketika kita ingin mengedit variable dengan fungsi maka kita gunakan pointernya
	// contoh
	var person1 Person = Person{Name: "Merta", Hobby: "Play"}

	fmt.Println(person1)

}
