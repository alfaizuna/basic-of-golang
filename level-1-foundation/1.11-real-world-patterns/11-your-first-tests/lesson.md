# Your First Tests

Menulis unit test profesional di Go menggunakan built-in toolchain `go test` dan paket `testing`.

## Contoh

```go
//go:build ignore

package main

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

```

## Penjelasan

1. **Konvensi Penamaan**:
   - File test selalu berakhiran `_test.go` (misalnya `math_test.go`).
   - Fungsi test selalu diawali dengan `Test` dan menerima pointer `t *testing.T` (misalnya `TestCalculate`).
2. **Tanpa Assertion Library**: Filosofi Go mengutamakan kode Go biasa (`if got != want`) daripada pustaka DSL assertion yang rumit.
3. **`t.Errorf` vs `t.Fatalf`**:
   - `t.Errorf`: mencatat kegagalan test namun **tetap melanjutkan** pengujian berikutnya.
   - `t.Fatalf`: mencatat kegagalan test dan **langsung menghentikan** fungsi test tersebut saat itu juga.

## >_ Exercise

Buat fungsi matematika sederhana:
1. `Multiply(a, b int) int`: mengembalikan `a * b`.
2. `Divide(a, b int) (int, error)`: jika `b == 0`, kembalikan `0, fmt.Errorf("division by zero")`, selain itu kembalikan `a / b, nil`.

### Expected output

```
Multiply 4*5 = 20
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- `go test` terintegrasi langsung di Go tanpa memerlukan framework pengujian eksternal.
- Gunakan `t.Errorf` untuk melanjutkan pengujian dan `t.Fatalf` untuk kegagalan kritis.
