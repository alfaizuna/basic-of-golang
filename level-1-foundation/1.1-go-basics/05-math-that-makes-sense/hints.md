### Hint 1
Gunakan operator `/` untuk pembagian integer dan `%` untuk sisa pembagian.

### Hint 2
Ubah tipe angka ke `float64` menggunakan `float64(namaVariabel)`. Contoh: `float64(totalItems) / float64(bags)` dan `price * float64(qty)`.

### Hint 3
Lengkapi `func main()` dengan baris berikut:
```go
itemsPerBag := totalItems / bags
remainder := totalItems % bags
avgItems := float64(totalItems) / float64(bags)
totalPrice := price * float64(qty)

fmt.Println(itemsPerBag)
fmt.Println(remainder)
fmt.Println(avgItems)
fmt.Println(totalPrice)
```
