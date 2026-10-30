package main

import "fmt"

func main() {
	// TODO:
	// 1. Buat slice days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}

	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}
	// 2. Buat sub-slice midWeek dari days yang mengambil "Selasa", "Rabu", "Kamis"
	midWeek := days[1:4]
	// 3. Ubah elemen pertama midWeek menjadi "SUPER TUESDAY"
	midWeek[0] = "SUPER TUESDAY"
	// 4. Cetak: fmt.Println("Days:", days)
	fmt.Println("Days:", days)
	// 5. Cetak: fmt.Println("MidWeek:", midWeek)
	fmt.Println("MidWeek:", midWeek)
}
