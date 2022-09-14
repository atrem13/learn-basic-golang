package main

import "fmt"

// dalam membuat fungsi untuk struct disarankan untuk membuatnya dalam bentuk pointer untuk menghemat penggunaan memori pada program
type Person struct {
	Name, Hobby string
}

func (person *Person) ChangeHobby() {
	person.Hobby = "He really love to " + person.Hobby
}

func main() {
	var person1 Person = Person{Name: "Merta", Hobby: "Play"}
	person1.ChangeHobby()
	fmt.Println(person1)
}
