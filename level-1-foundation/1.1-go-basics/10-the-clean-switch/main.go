package main

import "fmt"

// 1. Buat fungsi dayType(day string) string
// Gunakan switch day:
// - "senin", "selasa", "rabu", "kamis", "jumat" -> "Hari Kerja"
// - "sabtu", "minggu" -> "Akhir Pekan"
// - default -> "Tidak Dikenal"
func dayType(day string) string {
	switch day {
	case "senin", "selasa", "rabu", "kamis", "jumat":
		return "Hari Kerja"
	case "sabtu", "minggu":
		return "Akhir Pekan"
	default:
		return "Tidak Dikenal"
	}
}

// 2. Buat fungsi categorizeAge(age int) string
// Gunakan switch tanpa ekspresi (switch { ... }):
// - age < 13 -> "Anak-anak"
// - age <= 19 -> "Remaja"
// - default -> "Dewasa"
func categorizeAge(age int) string {
	switch {
	case age < 13:
		return "Anak-anak"
	case age <= 19:
		return "Remaja"
	default:
		return "Dewasa"
	}
}

func main() {
	// 3. Panggil dayType("sabtu") dan cetak hasilnya
	fmt.Println(dayType("sabtu"))

	// 4. Panggil categorizeAge(17) dan cetak hasilnya
	fmt.Println(categorizeAge(17))

}
