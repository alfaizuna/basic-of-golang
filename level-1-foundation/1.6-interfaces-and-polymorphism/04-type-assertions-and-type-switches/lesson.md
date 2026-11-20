# Type Assertions and Type Switches

Type assertion dan type switch memungkinkan kita mengekstrak nilai konkret dari dalam interface variable.

## Contoh

```go
package main

import "fmt"

func inspect(val any) {
	switch v := val.(type) {
	case string:
		fmt.Printf("String dengan panjang %d: %q\n", len(v), v)
	case int:
		fmt.Printf("Integer kuadrat: %d\n", v*v)
	default:
		fmt.Printf("Tipe lain: %T\n", v)
	}
}

func main() {
	var x any = "gopher"
	if s, ok := x.(string); ok {
		fmt.Println("Isi string:", s)
	}
	inspect(5)
}
```

## Penjelasan

1. **Type Assertion**: Sintaks `v, ok := i.(T)` mengecek apakah interface `i` menyimpan nilai dari tipe konkret `T`.
   - Jika ya: `v` adalah nilai bertipe `T`, dan `ok` bernilai `true`.
   - Jika tidak: `v` adalah zero value dari `T`, dan `ok` bernilai `false` (tanpa panic!).
   - Jika tanpa tanda terima `ok` (`v := i.(T)`) dan tipe salah, program akan mengalami runtime **panic**.
2. **Type Switch**: Konstruksi khusus `switch v := i.(type)` yang menguji tipe antarmuka secara berurutan. Di dalam setiap blok `case T:`, variabel `v` otomatis memiliki tipe statis `T`.

## >_ Exercise

Lengkapi fungsi `FormatData(val any) string` menggunakan type switch:
- Jika `val` bertipe `string`: kembalikan string `"STR: <val>"` (contoh: `"STR: golang"`)
- Jika `val` bertipe `int`: kembalikan string `"INT: <val*2>"` (contoh jika input 10, return `"INT: 20"`)
- Jika `val` bertipe `bool`: jika `true` kembalikan `"BOOL: YES"`, jika `false` kembalikan `"BOOL: NO"`
- Jika tipe lainnya: kembalikan `"UNKNOWN"`

Di `main()`, panggil `FormatData` untuk beberapa contoh nilai dan cetak hasilnya.

### Expected output

```
STR: golang
INT: 20
BOOL: YES
BOOL: NO
UNKNOWN
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Gunakan bentuk koma-ok (`v, ok := i.(T)`) saat melakukan type assertion untuk mencegah runtime panic.
- Gunakan `switch v := i.(type)` ketika menangani beberapa kemungkinan tipe konkret secara elegan.
