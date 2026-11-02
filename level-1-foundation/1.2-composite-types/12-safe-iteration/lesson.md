# Safe Iteration

Secara desain, Go mengacak urutan iterasi `range` pada map untuk mencegah ketergantungan pada urutan memori internal. Untuk mendapatkan urutan yang konsisten, kumpulkan semua key ke slice lalu urutkan.

## Contoh
example.go
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"Charlie": 1, "Alice": 2, "Bob": 3}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, m[k])
	}
}
```

## Penjelasan
- `for k, v := range m` → urutan pasangan key-value yang dihasilkan bersifat **acak** di setiap eksekusi program.
- Pola iterasi terurut (deterministic order):
  1. Kumpulkan semua key ke sebuah slice `[]string`.
  2. Urutkan slice tersebut (misal dengan `sort.Strings(keys)`).
  3. Lakukan loop pada slice key yang sudah rapi, dan ambil nilainya dari `m[k]`.

## >_ Exercise
Buka [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.2-composite-types/12-safe-iteration/main.go):
1. Diberikan map nilai:
   `scores := map[string]int{"Charlie": 75, "Alice": 95, "Bob": 85}`
2. Kumpulkan semua key dari `scores` ke dalam slice `keys`.
3. Urutkan `keys` menggunakan `sort.Strings(keys)`.
4. Lakukan loop pada `keys` dan cetak setiap pasangan dengan format:
   `fmt.Printf("%s: %d\n", k, scores[k])`

Expected output:
```
Alice: 95
Bob: 85
Charlie: 75
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Urutan iterasi map di Go sengaja diacak demi keamanan desain. Jika aplikasi membutuhkan urutan pasti (misalnya saat mencetak laporan atau menghasilkan JSON yang konsisten), selalu urutkan kumpulan key-nya terlebih dahulu.
