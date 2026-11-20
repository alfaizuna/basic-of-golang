# Concurrency vs Parallelism

Konkurensi adalah tentang menangani banyak hal sekaligus (struktur program), sedangkan paralelisme adalah tentang melakukan banyak hal sekaligus pada saat yang sama (eksekusi hardware).

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	"runtime"
	"time"
)

func task(name string) {
	for i := 1; i <= 2; i++ {
		fmt.Printf("Task %s: step %d\n", name, i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	fmt.Printf("Go runtime memiliki %d logical CPU core\n", runtime.NumCPU())
	task("A")
	task("B")
}

```

## Penjelasan

1. **Konkurensi vs Paralelisme**: Seperti yang dikatakan Rob Pike: *"Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once."* Konkurensi adalah cara kita menstrukturkan kode agar bagian-bagian independen dapat berjalan secara mandiri. Paralelisme bergantung pada ketersediaan banyak core CPU fisik.
2. **Goroutine Scheduler**: Go memiliki scheduler internal (M:N scheduler) yang memetakan ribuan (atau jutaan) goroutine ringan ke sejumlah kecil thread sistem operasi (OS threads).
3. **Ukuran Stack**: Thread OS biasanya memerlukan alokasi stack 1MB-2MB, sedangkan goroutine hanya berukuran awal sekitar 2KB dan dapat tumbuh dinamis sesuai kebutuhan, menjadikannya sangat murah dibuat.

## >_ Exercise

Buat fungsi `SimulateWorkload(workers int) int` yang mengembalikan total unit kerja yang diselesaikan. Jika setiap worker menyelesaikan 10 unit kerja, fungsi mengembalikan `workers * 10`.
Fungsi `main()` menampilkan perbandingan jumlah core runtime (`runtime.NumCPU()`) dan hasil dari `SimulateWorkload(4)`.

### Expected output

```
Total work: 40
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Konkurensi adalah struktur pemecahan masalah independen; paralelisme adalah eksekusi simultan di multi-core CPU.
- Go mengelola goroutine secara mandiri via runtime scheduler dengan biaya memori yang sangat minim (~2KB).
