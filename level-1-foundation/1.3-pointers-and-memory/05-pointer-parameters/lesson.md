# Pointer Parameters

Here's your first look at modifying a caller's variable directly by passing a pointer as a function parameter.

## Contoh
example.go
```go
//go:build ignore

package main

import "fmt"

// addBonus menerima pointer ke int (*int), sehingga dapat mengubah nilai variabel asli
func addBonus(salary *int, bonus int) {
	if salary != nil {
		*salary += bonus
	}
}

func main() {
	mySalary := 5000000
	fmt.Println("Gaji awal:", mySalary) // 5000000

	// Mengirim alamat memori (&mySalary) ke fungsi
	addBonus(&mySalary, 1500000)

	fmt.Println("Gaji setelah bonus:", mySalary) // 6500000 (berubah!)
}
```

## Penjelasan
- Di Lesson 01, kita belajar bahwa Go selalu menyalin nilai (*pass by value*).
- Ketika kita mengoper pointer `&mySalary` ke parameter `salary *int`, Go tetap menyalin nilai argumen, namun nilai yang disalin adalah **alamat memorinya**!
- Karena salinan alamat tetap menunjuk ke lokasi memori yang sama, operasi dereference `*salary += bonus` langsung mengubah data asli milik si pemanggil.
- Jangan lupa untuk selalu memeriksa `if salary != nil` sebelum melakukan dereference untuk keamanan.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/05-pointer-parameters/main.go):

1. Lengkapi fungsi `double(n *int)`:
   - Periksa jika `n != nil`.
   - Kalikan nilai yang ditunjuk oleh `n` dengan `2` (`*n = *n * 2` atau `*n *= 2`).
2. Di dalam `func main()`:
   - Diberikan variabel `score := 25`.
   - Panggil fungsi `double` dengan mengoper alamat memori dari `score` (`&score`).
   - Cetak nilai `score` yang telah termutasi:
     ```go
     fmt.Println("Score setelah didouble:", score)
     ```

Expected output:
```
Score setelah didouble: 50
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Dengan mengoper alamat variabel (`&var`) ke parameter bertipe pointer (`*T`), sebuah fungsi dapat mendereference pointer tersebut untuk memodifikasi langsung nilai variabel asli milik si pemanggil (*in-place mutation*).
