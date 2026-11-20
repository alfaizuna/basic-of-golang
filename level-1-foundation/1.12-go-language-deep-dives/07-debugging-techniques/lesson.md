# Debugging Techniques

Membaca stack trace saat panic, menggunakan `go vet` untuk analisis kode statis, dan mengisolasi bug dengan unit test.

## Contoh

```go
//go:build ignore

package main

import "fmt"

func buggyFunction(items []int) int {
	// Bug: index out of range jika slice kosong
	return items[0]
}

func main() {
	fmt.Println("Pemeriksaan alat bantu: go vet ./...")
}

```

## Penjelasan

1. **Membaca Stack Trace**:
   - Baris teratas menunjukkan titik persis terjadinya kegagalan (file dan nomor baris).
   - Baris berikutnya menelusuri rantai pemanggilan fungsi (*call chain*) dari goroutine yang bersangkutan.
2. **`go vet`**: Alat analisis statis bawaan Go yang memeriksa kesalahan umum sebelum program dijalankan (seperti argumen `Printf` salah format, unreachable code, atau struct tag salah ketik).
3. **Mengisolasi Bug dengan Test**: Alur terbaik mereparasi bug adalah membuat unit test yang menduplikasi kegagalan (*reproducing test case*), lalu memperbaiki kode sumber hingga test berubah menjadi hijau (PASS).

## >_ Exercise

Perbaiki fungsi `SafeFirst(items []int) (int, bool)` yang sebelumnya rentan panic jika slice kosong!
- Jika `len(items) == 0`: kembalikan `0, false`.
- Jika ada elemen: kembalikan elemen pertama `items[0]` dan `true`.

### Expected output

```
First: 10, OK: true
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `go vet` mendeteksi anomali kode secara statis tanpa perlu menjalankan aplikasi.
- Mengisolasi skenario edge case dengan unit test adalah metode terbaik untuk mencegah regresi bug.
