### Hint 1
Gunakan operator `&` di depan nama variabel untuk mengambil alamat memorinya, dan gunakan operator `*` di depan pointer untuk mengakses (*dereference*) atau mengubah nilainya.

### Hint 2
- Ambil alamat: `ptr := &score`
- Cetak nilai lewat pointer: `fmt.Println("Nilai awal lewat pointer:", *ptr)`
- Ubah nilai lewat pointer: `*ptr = 99`

### Hint 3
Kerangka kode di dalam `main()`:
```go
score := 42

// 1. Ambil alamat
ptr := &...

// 2. Cetak lewat pointer
fmt.Println("Nilai awal lewat pointer:", *...)

// 3. Ubah nilai lewat pointer
*ptr = ...

// 4. Cetak nilai score asli
fmt.Println("Nilai score setelah diubah:", score)
```
