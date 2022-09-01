package main

import "fmt"

func main() {
	// inisiasi variable
	var nama string
	// memberi nilai pada variable
	nama = "merta yoga"

	// inisiasi langsung isi nilai variable
	var nama2 = "nama lainnya"

	// contoh lainnya
	nama3 := "ini nyoba yakk"

	// inisiasi multiple variables
	var (
		firstName = "dwi"
		lastName  = "permana"
	)

	// constanta
	// wajib langsung diberi values
	// tidak bisa diubah setelah diberi values
	const iniConstan = "hehe"

	fmt.Println(nama)
	fmt.Println(nama2)
	fmt.Println(nama3)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(iniConstan)
}
