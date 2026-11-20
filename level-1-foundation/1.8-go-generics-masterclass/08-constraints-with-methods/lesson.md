# Constraints with Methods

Type constraint dapat berupa interface reguler yang mengharuskan tipe konkret mengimplementasikan sekumpulan method tertentu.

## Contoh

```go
//go:build ignore

package main

import "fmt"

// Constraint interface dengan method
type Stringer interface {
	String() string
}

func PrintAllStringers[T Stringer](items []T) {
	for _, item := range items {
		fmt.Println(item.String())
	}
}

type User struct {
	Name string
}

func (u User) String() string {
	return "User: " + u.Name
}

func main() {
	users := []User{{Name: "Alice"}, {Name: "Bob"}}
	PrintAllStringers(users)
}

```

## Penjelasan

1. **Interface Biasa sebagai Constraint**: Setiap interface Go (seperti `fmt.Stringer`, `io.Reader`, atau interface kustom) dapat langsung dipakai sebagai constraint pada parameter tipe.
2. **Kapan Menggunakan Generics vs Interface Parameter Biasa**:
   - Jika fungsi hanya perlu memanggil method: `func PrintAll(items []Stringer)` vs `func PrintAll[T Stringer](items []T)`.
   - Menggunakan generics (`[]T`) mengembalikan slice bertipe spesifik yang sama tanpa slicing conversion, dan menghindari alokasi boxing interface di memori.

## >_ Exercise

1. Definisikan interface constraint `Measurable`:
   ```go
   type Measurable interface {
       Area() float64
   }
   ```
2. Buat struct `Square` dengan field `Side float64` dan implementasikan method `Area() float64`.
3. Buat fungsi generik `SumAreas[T Measurable](items []T) float64` yang menjumlahkan luas dari semua elemen di dalam slice.

### Expected output

```
Total Area: 25.00
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Interface standar dengan definisi method dapat berfungsi sebagai type constraint.
- Memungkinkan pemanggilan method pada type parameter di dalam fungsi generik.
