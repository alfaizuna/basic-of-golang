### Hint 1
Go tidak memerlukan kata kunci `break` di setiap `case`. Untuk memeriksa beberapa nilai dalam satu baris, pisahkan nilainya dengan tanda koma (contoh: `case "senin", "selasa":`).

### Hint 2
Struktur fungsi `dayType` dengan `switch day`:
```go
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
```

### Hint 3
Struktur fungsi `categorizeAge` dengan switch tanpa ekspresi:
```go
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
```
Lalu di `main()`, panggil `fmt.Println(dayType("sabtu"))` dan `fmt.Println(categorizeAge(17))`.
