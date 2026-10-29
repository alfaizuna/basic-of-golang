### Hint 1
Untuk mengubah nilai bertipe `rune` menjadi `byte` saat nilainya `<= 255`, gunakan konversi tipe eksplisit: `byte(r)`.

### Hint 2
Bentuk fungsi `toByte` dan `addInt8`:
```go
func toByte(r rune) (byte, bool) {
	if r <= 255 {
		return byte(r), true
	}
	return 0, false
}

func addInt8(a int8, b int8) int8 {
	return a + b
}
```

### Hint 3
Di dalam fungsi `main()`:
```go
b2, ok2 := toByte('世')
fmt.Println(b2, ok2)

fmt.Println(addInt8(127, 1))
```
