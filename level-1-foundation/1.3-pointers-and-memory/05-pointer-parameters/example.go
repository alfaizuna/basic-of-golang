//go:build ignore

package main

import "fmt"

// addBonus menerima pointer ke int (*int), sehingga dapat mengubah nilai variabel asli
func addBonus(salary *int, bonus int) {
	if salary != nil {
		*salary += bonus
	}
}

func main() {
	mySalary := 5000000
	fmt.Println("Gaji awal:", mySalary) // 5000000

	// Mengirim alamat memori (&mySalary) ke fungsi
	addBonus(&mySalary, 1500000)

	fmt.Println("Gaji setelah bonus:", mySalary) // 6500000 (berubah!)
}
