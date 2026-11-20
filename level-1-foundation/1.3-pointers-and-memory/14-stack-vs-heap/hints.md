## Hint 1 (arah umum)
Untuk `SumValues`: lakukan penjumlahan langsung dan kembalikan hasilnya:
```go
return a + b
```

Untuk `CreateBuffer`: buat slice byte menggunakan `make([]byte, size)`, simpan di variabel, lalu kembalikan alamatnya dengan `&`.

---

## Hint 2 (sintaks yang relevan)
```go
buf := make([]byte, size)
return &buf
```

---

## Hint 3 (hampir jawaban)
```go
func SumValues(a, b int) int {
	return a + b
}

func CreateBuffer(size int) *[]byte {
	buf := make([]byte, size)
	return &buf
}
```
