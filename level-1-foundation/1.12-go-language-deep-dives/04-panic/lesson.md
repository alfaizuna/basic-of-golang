# Panic

Panic menghentikan alur eksekusi normal ketika terjadi kondisi darurat yang fatal atau bug pemrograman yang tidak dapat dipulihkan.

## Contoh

```go
//go:build ignore

package main

import "fmt"

func assertNonEmpty(s string) {
	if s == "" {
		panic("string tidak boleh kosong!")
	}
}

func main() {
	fmt.Println("Sebelum panic")
	assertNonEmpty("ok")
	// assertNonEmpty("") // Menyebabkan crash runtime
}

```

## Penjelasan

1. **Error vs Panic**:
   - **Error**: Kondisi yang sudah diperkirakan (*expected failures*), seperti file tidak ada atau timeout jaringan. Wajib dikembalikan sebagai return value `error`.
   - **Panic**: Bug logika fatal (*programmer error*) atau kondisi inisialisasi kritis di mana program mustahil melanjutkan hidup (misalnya pointer nil dereference atau regex parser statis gagal saat startup).
2. **Stack Unwinding**: Saat panic terjadi, Go segera menghentikan eksekusi fungsi saat ini, menjalankan semua fungsi `defer` di fungsi tersebut, lalu naik ke pemanggilnya sambil terus menjalankan `defer`, hingga mencapai top-level goroutine dan mencetak crash stack trace.

## >_ Exercise

Buat fungsi `RequirePositive(val int)` yang:
- Jika `val <= 0`, panggil `panic(fmt.Sprintf("invalid value %d: must be positive", val))`.
- Jika `val > 0`, tidak melakukan apa-apa.

### Expected output

```
Valid value: 42
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `panic` digunakan untuk situasi kritis yang tidak terpulihkan atau bug programmer fatal.
- Selama proses unwinding, seluruh deferred calls tetap dijamin dieksekusi.
