// merupakan cara membuat nama alternatif untuk tipe data
package main

import "fmt"

func main() {
	type tipeString string

	var akuvar tipeString = "ini loh"

	fmt.Println(akuvar)
}
