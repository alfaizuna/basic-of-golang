### Hint 1
Untuk mengambil `"Selasa"` (indeks 1) sampai `"Kamis"` (indeks 3), gunakan slicing dengan batas bawah `1` dan batas atas `4` (karena batas atas bersifat eksklusif).

### Hint 2
Sintaks slicing:
```go
midWeek := days[1:4]
```
Untuk mengubah elemen pertama `midWeek`:
```go
midWeek[0] = "SUPER TUESDAY"
```

### Hint 3
Kerangka kode lengkap:
```go
days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}
midWeek := days[1:4]
midWeek[0] = "SUPER TUESDAY"

fmt.Println("Days:", days)
fmt.Println("MidWeek:", midWeek)
```
