### Hint 1
Fungsi yang mengembalikan lebih dari satu nilai ditulis dengan membungkus tipe kembaliannya di dalam tanda kurung: `func calculate(a, b int) (int, int)`.

### Hint 2
Struktur fungsi `calculate`:
```go
func calculate(a, b int) (int, int) {
	return a + b, a - b
}
```
Untuk menampung 2 nilai kembalian sekaligus di `main()`:
```go
sum, diff := calculate(20, 5)
```

### Hint 3
Untuk langkah ke-3, abaikan nilai selisih dengan menaruh karakter `_` di posisi kedua:
```go
total, _ := calculate(50, 10)
fmt.Println(total)
```
