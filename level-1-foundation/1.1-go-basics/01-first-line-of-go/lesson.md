# Your First Line of Go

Here's your first step into writing Go code.

## Contoh
```go
// example.go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## Penjelasan
- `package main` → Memberitahu Go bahwa file ini adalah program utama yang bisa langsung dijalankan (bukan pustaka/library). Analogi: *package* seperti kotak perkakas yang mengelompokkan kode saling terkait.
- `import "fmt"` → Meminjam paket `fmt` (*format*), yaitu kotak perkakas bawaan Go untuk mencetak dan memformat teks.
- `func main()` → Pintu masuk utama program; kode di dalam `{ ... }` adalah hal pertama yang dieksekusi Go saat berjalan.
- `fmt.Println(...)` → Memanggil fungsi `Println` dari paket `fmt` untuk mencetak teks ke layar dan otomatis menambahkan baris baru di akhirnya.

## >_ Exercise
Buka file `main.go`, lalu tambahkan satu baris di dalam fungsi `main()` agar program mencetak teks berikut:

Expected output:
```
Hello, Go!
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Setiap program Go mandiri selalu dimulai dari `package main` dan fungsi `main()`. Gunakan paket `fmt` dan fungsi `fmt.Println` untuk menampilkan teks ke layar.
