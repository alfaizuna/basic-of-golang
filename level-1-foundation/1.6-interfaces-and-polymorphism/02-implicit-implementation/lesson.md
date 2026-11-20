# Implicit Implementation

Di bahasa seperti Java atau C#, suatu class harus secara eksplisit menulis kata kunci `implements MyInterface`.
Di Go, **tidak ada kata kunci `implements`**! Sebuah tipe otomatis memenuhi suatu interface jika tipe tersebut memiliki semua method yang dideklarasikan oleh interface tersebut. Ini sering disebut *structural typing* atau "Duck Typing" ala Go (*"if it walks like a duck and quacks like a duck, it's a duck"*).

Hal ini memungkinkan **Polimorfisme**: beberapa tipe berbeda dapat diperlakukan sebagai satu tipe interface yang sama.

## Contoh
example.go
```go
package main

type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// TotalArea bisa menerima slice dari berbagai bentuk shape
func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}
```

## Penjelasan
- Baik `Square` maupun `Circle` sama-sama mengimplementasikan `Shape` secara implisit karena memiliki method `Area() float64`.
- Fungsi `TotalArea(shapes []Shape)` dapat memproses `Square` dan `Circle` secara bersamaan dalam satu koleksi tanpa peduli tipe aslinya.
- Keuntungan implementasi implisit: pembuat tipe struct tidak perlu tahu tentang interface yang akan mengonsumsinya di kemudian hari.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.6-interfaces-and-polymorphism/02-implicit-implementation/main.go).

1. Definisikan interface `Notifier`:
   - Method: `Notify(message string) string`
2. Definisikan struct `EmailSender`:
   - Field: `Email string`
   - Method: `(e EmailSender) Notify(message string) string`:
     - Kembalikan `"[EMAIL to " + e.Email + "] " + message`
3. Definisikan struct `SMSSender`:
   - Field: `Phone string`
   - Method: `(s SMSSender) Notify(message string) string`:
     - Kembalikan `"[SMS to " + s.Phone + "] " + message`
4. Buat fungsi `BroadcastAlert(notifiers []Notifier, alertMessage string) []string`:
   - Panggil `Notify(alertMessage)` untuk setiap notifier di slice.
   - Kumpulkan dan kembalikan semua string pesan notifikasi.

Expected output ketika `go run main.go`:
```
Notification 1: [EMAIL to user@example.com] Server down!
Notification 2: [SMS to +62812345678] Server down!
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Di Go, pemenuhan interface bersifat implisit tanpa kata kunci tambahan. Ini menghasilkan polimorfisme yang alami dan memisahkan modul produsen data dari modul konsumen.
