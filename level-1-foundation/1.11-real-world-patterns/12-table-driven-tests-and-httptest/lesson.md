# Table-Driven Tests & httptest

Pola Table-Driven Test adalah standar emas penulisan test di ekosistem Go, menguji banyak skenario input/output secara deklaratif dan terstruktur.

## Contoh

```go
//go:build ignore

package main

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"bilangan genap positif", 4, true},
		{"bilangan ganjil positif", 5, false},
		{"nol adalah genap", 0, true},
		{"bilangan ganjil negatif", -3, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// got := IsEven(tc.input)
			// if got != tc.expected { t.Errorf(...) }
		})
	}
}

```

## Penjelasan

1. **Anatomi Table-Driven Test**:
   - Definisikan slice dari anonymous struct `tests := []struct { name string; input ...; want ... }{ ... }`.
   - Iterasi slice dengan loop `for _, tc := range tests`.
   - Gunakan `t.Run(tc.name, func(t *testing.T) { ... })` agar setiap skenario menjadi sub-test terisolasi yang mudah dilacak jika gagal.
2. **Manfaat**: Menambah test case baru hanya membutuhkan penambahan 1 baris di tabel data tanpa menduplikasi kode logika pengujian.

## >_ Exercise

Buat fungsi validasi email sederhana: `ValidateEmail(email string) bool`.
Sebuah email dianggap valid jika:
1. Tidak kosong.
2. Mengandung karakter `'@'`.
3. Mengandung karakter `'.'`.
4. Panjangnya minimal 5 karakter.

### Expected output

```
Email valid: true
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Table-driven tests mengeliminasi duplikasi kode dan membuat penambahan skenario uji menjadi sangat cepat.
- `t.Run()` memberikan laporan granular untuk setiap kasus uji.
