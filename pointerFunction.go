package main

import "fmt"

type Person struct {
	Name, Hobby string
}

func ChangeHobby(person Person) {
	person.Hobby = "Sleep"
}

func ChangeHobby2(person *Person) { //kita mengarah ke pointer dari variable
	person.Hobby = "Sleep"
}

func main() {
	// ketika kita mengedit variable dalam fungsi maka yang berubah hanyalah variable didalam fungsi, bukan variable diluarnya
	// oleh sebab itu sebaiknya ketika kita ingin mengedit variable dengan fungsi maka kita gunakan pointernya
	// contoh
	var person1 Person = Person{Name: "Merta", Hobby: "Play"}
	ChangeHobby(person1) //berusaha mengubah data hobby dari person
	fmt.Println(person1.Name)
	fmt.Println(person1.Hobby) //ternyata tidak berubah

	ChangeHobby2(&person1) // kita mengirimkan alamat memori dari variable
	fmt.Println(person1.Name)
	fmt.Println(person1.Hobby) //berhasil berubah berubah

}
