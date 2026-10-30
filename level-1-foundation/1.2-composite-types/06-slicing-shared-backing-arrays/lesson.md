# Slicing & Shared Backing Arrays

Slicing `s[a:b]` tidak membuat salinan data baru, melainkan melihat ke backing array yang sama.

## Contoh
example.go
```go
package main

import "fmt"

func main() {
	original := []int{10, 20, 30, 40, 50}
	sub := original[1:3] // [20, 30]

	sub[0] = 999
	fmt.Println("Sub:", sub)         // [999 30]
	fmt.Println("Original:", original) // [10 999 30 40 50]
}
```

## Penjelasan
- `s[low:high]` → membuat slice baru dari indeks `low` sampai `high-1`.
- **Shared Backing Array**: Slice hasil slicing memegang pointer ke array memori yang sama dengan slice asal.
- **Efek samping (side effect)**: Mengubah nilai elemen melalui slice turunan akan langsung mengubah nilai pada slice asal.
- Jika ingin memodifikasi data tanpa mempengaruhi slice asal, kita harus membuat salinan independen (akan dipelajari di lesson berikutnya: `copy`).

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/06-slicing-shared-backing-arrays/main.go):
1. Buat slice `days` dengan elemen: `"Senin"`, `"Selasa"`, `"Rabu"`, `"Kamis"`, `"Jumat"`.
2. Buat sub-slice `midWeek` dari `days` yang mengambil elemen `"Selasa"`, `"Rabu"`, `"Kamis"`.
3. Ubah elemen pertama pada `midWeek` menjadi `"SUPER TUESDAY"`.
4. Cetak `days` dan `midWeek`:
   - `fmt.Println("Days:", days)`
   - `fmt.Println("MidWeek:", midWeek)`

Expected output:
```
Days: [Senin SUPER TUESDAY Rabu Kamis Jumat]
MidWeek: [SUPER TUESDAY Rabu Kamis]
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Slicing `s[a:b]` bersifat *view*, bukan *copy*. Perubahan elemen pada slice turunan akan berdampak langsung ke slice asalnya karena keduanya berbagi backing array yang sama.
