# Logical Operators

Here's how to combine conditions using boolean logic: `&&`, `||`, and `!`.

## Contoh
```go
// example.go
package main

import "fmt"

func main() {
	isWeekend := true
	isRaining := false
	hasUmbrella := true

	// Operator && (DAN): bernilai true hanya jika KEDUA sisi bernilai true
	canPlayOutside := !isRaining && isWeekend

	// Operator || (ATAU): bernilai true jika SALAH SATU sisi bernilai true
	canGoOut := !isRaining || hasUmbrella

	// Operator ! (NOT / NEGASI): membalikkan true menjadi false dan sebaliknya
	fmt.Println(canPlayOutside)
	fmt.Println(canGoOut)
	fmt.Println(!isRaining)
}
```

## Penjelasan
- `&&` (Logical AND) → Menghasilkan `true` hanya jika kedua kondisi terpenuhi.
- `||` (Logical OR) → Menghasilkan `true` jika minimal satu kondisi terpenuhi.
- `!` (Logical NOT) → Membalikkan nilai boolean (`!true` menjadi `false`, `!false` menjadi `true`).
- **Short-Circuit Evaluation** → Go mengevaluasi ekspresi logika dari kiri ke kanan. Jika pada `A && B` nilai `A` sudah `false`, Go tidak akan repot-repot memeriksa `B` karena hasilnya pasti `false`. Begitu juga pada `A || B`, jika `A` bernilai `true`, Go langsung berhenti mengevaluasi.

## >_ Exercise
Buka file `main.go`. Diberikan status server berikut:
- `cpuUsage := 88`
- `memoryUsage := 60`
- `isOnline := true`

Buat logika *server monitor* untuk memicu alert yang tepat:
1. Jika server **TIDAK online** ATAU **CPU di atas 95** (`cpuUsage > 95`), cetak `"CRITICAL"`.
2. Jika server **online** DAN (**CPU di atas 80** ATAU **Memory di atas 80**), cetak `"WARNING"`.
3. Jika kondisi di atas tidak ada yang terpenuhi, cetak `"OK"`.

Expected output:
```
WARNING
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan `&&` untuk mewajibkan semua syarat terpenuhi, `||` untuk opsi alternatif, dan `!` untuk negasi. Gunakan tanda kurung `(...)` untuk mengelompokkan kondisi logika yang kompleks agar urutan evaluasinya jelas.
