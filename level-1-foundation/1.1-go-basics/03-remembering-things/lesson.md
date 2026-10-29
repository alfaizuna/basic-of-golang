# Remembering Things

Here's your first step into storing and managing data with variables.

## Contoh
```go
// example.go
package main

import "fmt"

func main() {
	// Deklarasi formal dengan kata kunci 'var'
	var language string = "Go"

	// Deklarasi singkat dengan ':=' (tipe otomatis disimpulkan Go)
	version := 1.26

	// Mengubah nilai variabel yang sudah ada (assignment dengan '=')
	version = 1.27

	fmt.Println(language)
	fmt.Println(version)
}
```

## Penjelasan
- `var language string = "Go"` → Mendeklarasikan variabel `language` dengan tipe `string` dan nilai awal secara eksplisit.
- `version := 1.26` → Short variable declaration (hanya bisa di dalam fungsi). Go menebak tipenya otomatis sebagai `float64`.
- `version = 1.27` → Mengubah nilai variabel yang sudah ada. Jangan gunakan `:=` lagi karena variabel sudah pernah dideklarasikan.
- **Aturan penting:** Di Go, setiap variabel lokal yang dideklarasikan **wajib dipakai**, jika tidak compiler akan menolaknya dengan error.

## >_ Exercise
Buka file `main.go`, lalu lengkapi fungsi `main()` dengan langkah berikut:
1. Buat variabel bernama `app` dengan tipe `string` dan nilai `"Server"` menggunakan kata kunci `var`.
2. Buat variabel bernama `port` bernilai `8080` menggunakan operator short declaration `:=`.
3. Ubah nilai `port` menjadi `9090` menggunakan assignment `=`.
4. Cetak variabel `app` lalu variabel `port` masing-masing dengan `fmt.Println`.

Expected output:
```
Server
9090
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan `:=` untuk membuat variabel baru secara ringkas, dan `=` untuk memperbarui nilainya. Di Go, semua variabel yang dideklarasikan wajib digunakan!
