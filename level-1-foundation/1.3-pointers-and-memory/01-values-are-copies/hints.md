### Hint 1
Di dalam `resetCount(count int)`, ubah nilai parameter dengan assignment `count = 0`, lalu cetak hasilnya dengan `fmt.Println("Di dalam resetCount:", count)`.

### Hint 2
Di dalam `main()`, panggil fungsi tersebut dengan mengoper `total` sebagai argumen: `resetCount(total)`. Setelah itu, cetak nilai `total` menggunakan `fmt.Println("Di dalam main:", total)`.

### Hint 3
Bentuk kode lengkapnya adalah:
```go
func resetCount(count int) {
	count = 0
	fmt.Println("Di dalam resetCount:", count)
}

func main() {
	total := 100
	resetCount(total)
	fmt.Println("Di dalam main:", total)
}
```
