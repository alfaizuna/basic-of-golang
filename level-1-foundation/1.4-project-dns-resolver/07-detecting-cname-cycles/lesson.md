# Detecting CNAME Cycles

Bagaimana jika administrator salah mengonfigurasi DNS sehingga terjadi perputaran alias melingkar?
Contohnya:
- `a.example.com` CNAME ➡️ `b.example.com`
- `b.example.com` CNAME ➡️ `a.example.com`

Tanpa perlindungan, fungsi `Resolve` yang kita buat di pelajaran sebelumnya akan mengalami **infinite loop** (perulangan tanpa henti) dan membuat program hang!
Pada pelajaran ketujuh, kita menambahkan deteksi siklus menggunakan `map[string]bool` sebagai set penanda domain yang sudah dikunjungi, serta batas maksimal kedalaman penelusuran (`maxDepth`).

## Contoh
example.go
```go
package main

import "errors"

const maxCNAMEDepth = 10

func (z *Zone) Resolve(name string) (*Record, error) {
	visited := make(map[string]bool)
	current := name
	depth := 0

	for {
		if depth >= maxCNAMEDepth {
			return nil, errors.New("maximum CNAME depth exceeded")
		}
		if visited[current] {
			return nil, errors.New("cname loop detected")
		}
		visited[current] = true

		if rec := z.Lookup(current, TypeA); rec != nil {
			return rec, nil
		}

		cname := z.Lookup(current, TypeCNAME)
		if cname == nil {
			return nil, errors.New("record not found")
		}

		current = cname.Value
		depth++
	}
}
```

## Penjelasan
- `visited := make(map[string]bool)`: idiom standar di Go untuk membuat *Set* unik.
- Jika `visited[current]` bernilai `true`, artinya domain ini sudah pernah disinggahi sebelumnya dalam rantai pencarian yang sama. Ini tanda pasti adanya siklus/looping!
- `depth >= maxCNAMEDepth`: pengaman tambahan jika rantai alias terlalu panjang, untuk mencegah kehabisan sumber daya komputasi.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.4-project-dns-resolver/07-detecting-cname-cycles/main.go).

Perbarui method `Resolve(name string) (*Record, error)`:
1. Buat set pelacak `visited := make(map[string]bool)` dan penghitung kedalaman `depth := 0`.
2. Di awal setiap iterasi loop:
   - Jika `depth >= 10`, kembalikan `nil, errors.New("maximum CNAME depth exceeded")`.
   - Jika `visited[current]`, kembalikan `nil, errors.New("cname loop detected")`.
   - Tandai `visited[current] = true`.
3. Lanjutkan pencarian `TypeA` dan `TypeCNAME`. Jangan lupa tambahkan `depth++` saat berpindah ke CNAME berikutnya.

Expected output ketika `go run main.go`:
```
Resolving loop-a: error: cname loop detected
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Menggunakan map sebagai set penanda (`map[string]bool`) dan menetapkan batas kedalaman maksimum (*max depth*) adalah pertahanan standar algoritma graf dan penelusuran pointer untuk mencegah infinite loop.
