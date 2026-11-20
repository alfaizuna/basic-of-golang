# Stack vs Heap

Di Go, kamu tidak perlu mengalokasikan dan membebaskan memori secara manual seperti `malloc` dan `free` di C. Compiler Go secara otomatis memutuskan apakah sebuah variabel ditempatkan di **Stack** atau di **Heap** melalui proses yang disebut **Escape Analysis**.

## Contoh
example.go
```go
package main

type Point struct {
	X, Y int
}

// staysOnStack: p dialokasikan di stack frame fungsi.
// Saat fungsi selesai, memori p otomatis dibersihkan dengan sangat cepat tanpa campur tangan Garbage Collector (GC).
func staysOnStack() int {
	p := Point{X: 10, Y: 20}
	return p.X + p.Y
}

// escapesToHeap: pointer &p dikembalikan ke pemanggil,
// sehingga p harus tetap hidup melebihi siklus hidup (stack frame) fungsi ini.
// Compiler Go mendeteksi ini dan memindahkan p ke heap ("escapes to heap").
func escapesToHeap() *Point {
	p := Point{X: 10, Y: 20}
	return &p
}
```

## Penjelasan
- **Stack Memory**:
  - Sangat cepat (alokasi dan dealokasi hanya memajukan/memundurkan pointer stack CPU).
  - Terikat dengan eksekusi suatu fungsi (*stack frame*).
  - Variabel yang ukurannya diketahui dan tidak dirujuk lagi setelah fungsi berakhir akan ditaruh di stack.
- **Heap Memory**:
  - Digunakan untuk data dinamis atau data yang masa hidupnya melebihi fungsi yang membuatnya (*escapes*).
  - Dikelola oleh **Garbage Collector (GC)** Go.
  - Memiliki *overhead* lebih tinggi dibandingkan stack.
- **Melihat Escape Analysis**:
  Kamu bisa melihat keputusan compiler Go dengan flag kompilasi:
  ```bash
  go build -gcflags="-m" main.go
  ```
  Output compiler akan menampilkan catatan seperti:
  - `... moved to heap: p`
  - `... &p escapes to heap`
  - `... does not escape`

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/14-stack-vs-heap/main.go).

Terdapat dua fungsi yang perlu kamu lengkapi:
1. [`SumValues(a, b int) int`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/14-stack-vs-heap/main.go#L11-L15):
   - Buat variabel lokal `total := a + b`.
   - Kembalikan nilainya secara **value** (`total`), bukan pointer.
   - Variabel ini tetap berada di **stack** karena nilainya disalin (*copied*) dan tidak bocor ke luar fungsi sebagai pointer.
2. [`CreateBuffer(size int) *[]byte`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/14-stack-vs-heap/main.go#L18-L22):
   - Buat slice byte lokal menggunakan `buf := make([]byte, size)`.
   - Kembalikan pointer ke slice tersebut (`&buf`).
   - Karena pointernya dibagikan ke luar, buffer ini akan **escape ke heap**.

Setelah mengimplementasikannya, jalankan program:
```bash
go run main.go
```
Dan jalankan escape analysis untuk melihat laporan compiler:
```bash
go build -gcflags="-m" main.go
```

Expected output:
```
Sum: 30
Buffer len: 1024
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Stack cepat dan otomatis dibersihkan saat fungsi selesai, sedangkan Heap menampung data yang lolos (*escaped*) dari stack frame dan dikelola oleh Garbage Collector. Kamu bisa mengintip analisis compiler kapan saja dengan `go build -gcflags="-m"`.
