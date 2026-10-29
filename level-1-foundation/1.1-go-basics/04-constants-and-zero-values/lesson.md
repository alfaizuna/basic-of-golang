# Constants and Zero Values

Here's your first encounter with immutable values and Go's safe default values.

## Contoh
```go
// example.go
package main

import "fmt"

const Pi = 3.14

const (
	A = iota // 0
	B        // 1
	C        // 2
)

func main() {
	var count int   // zero value: 0
	var isDone bool // zero value: false

	fmt.Println(Pi)
	fmt.Println(A, B, C)
	fmt.Println(count, isDone)
}
```

## Penjelasan
- `const Pi = 3.14` → Konstanta adalah nilai tetap yang tidak dapat diubah setelah didefinisikan.
- `iota` → Generator angka urut otomatis di dalam blok `const (...)`. Selalu dimulai dari angka `0` dan otomatis bertambah `1` pada baris berikutnya.
- **Zero Value** → Di Go, jika kamu mendeklarasikan variabel tanpa nilai awal (misalnya `var count int`), Go otomatis mengisinya dengan nilai bawaan yang aman (`0` untuk angka, `false` untuk boolean, dan `""` untuk string). Tidak ada nilai sampah (*garbage value*) atau *undefined* di Go!

## >_ Exercise
Buka file `main.go`, lalu:
1. Deklarasikan konstanta `MaxScore = 100`.
2. Gunakan blok `const (...)` dengan `iota` untuk membuat konstanta `StatusDraft` (0) dan `StatusPublished` (1).
3. Di dalam `func main()`:
   - Deklarasikan variabel `defaultAttempts` bertipe `int` tanpa nilai awal (biarkan zero value).
   - Deklarasikan variabel `isReady` bertipe `bool` tanpa nilai awal (biarkan zero value).
   - Cetak berturut-turut di baris terpisah: `MaxScore`, `StatusDraft`, `StatusPublished`, `defaultAttempts`, dan `isReady`.

Expected output:
```
100
0
1
0
false
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Gunakan `const` untuk nilai yang tidak boleh berubah dan `iota` untuk membuat urutan nilai konstan secara otomatis. Go menjamin keamanan memori dengan selalu memberikan *zero value* pada variabel baru.
