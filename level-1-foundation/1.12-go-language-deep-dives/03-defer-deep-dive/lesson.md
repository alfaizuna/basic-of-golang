# Defer Deep Dive

Memahami 3 aturan fundamental eksekusi defer di Go: evaluasi argumen di muka, urutan eksekusi LIFO, dan kemampuan memodifikasi named return.

## Contoh

```go
//go:build ignore

package main

import "fmt"

func lifoExample() {
	// Eksekusi Last-In, First-Out (tumpukan terbalik)
	defer fmt.Println("1 (pertama didaftarkan, terakhir dieksekusi)")
	defer fmt.Println("2")
	defer fmt.Println("3 (terakhir didaftarkan, pertama dieksekusi)")
}

func main() {
	lifoExample()
}

```

## Penjelasan

1. **Aturan 1 - Snapshot Argumen**: Argumen fungsi deferred dievaluasi **seketika saat statement defer dibaca**, bukan saat fungsi dieksekusi.
   ```go
   i := 0
   defer fmt.Println(i) // Mencetak 0, meski i diubah setelah baris ini!
   i++
   ```
2. **Aturan 2 - LIFO (Last In, First Out)**: Jika ada beberapa defer dalam satu fungsi, mereka dieksekusi dalam urutan terbalik dari pendaftaran (seperti stack).
3. **Aturan 3 - Modifikasi Named Return**: Fungsi anonim deferred dapat membaca dan mengubah nilai `named return value` sebelum fungsi benar-benar kembali ke pemanggil.

## >_ Exercise

Buat fungsi `AddTax(amount float64) (total float64)` dengan named return `total`:
1. Tetapkan `total = amount` di awal.
2. Daftarkan deferred anonymous function: `defer func() { total *= 1.10 }()` untuk menambahkan pajak 10%.
3. Lakukan `return total`.

### Expected output

```
Total with tax: 110.00
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `defer` berjalan secara LIFO (Last-In-First-Out).
- Deferred closure dapat mengamati dan mengubah `named return values` sebelum fungsi keluar.
