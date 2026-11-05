### Hint 1
Deklarasikan pointer tanpa inisialisasi menggunakan sintaks `var ptr *string`. Karena belum diisi, nilai default-nya (*zero value*) otomatis adalah `nil`.

### Hint 2
- Untuk mengecek nil: `if ptr == nil { fmt.Println("ptr adalah nil") }`
- Untuk menghubungkan ke variabel: `ptr = &message`
- Untuk mengecek tidak nil dan dereference: `if ptr != nil { fmt.Println("ptr tidak nil, isinya:", *ptr) }`

### Hint 3
Kerangka kode lengkap di dalam `main()`:
```go
var ptr *string
if ptr == nil {
	fmt.Println("ptr adalah nil")
}

message := "Belajar Go menyenangkan!"
ptr = &message

if ptr != nil {
	fmt.Println("ptr tidak nil, isinya:", *ptr)
}
```
