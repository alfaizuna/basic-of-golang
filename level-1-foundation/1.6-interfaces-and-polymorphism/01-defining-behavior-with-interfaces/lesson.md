# Defining Behavior with Interfaces

Di Go, **interface** adalah sekumpulan tanda tangan method (*method signatures*).
Tidak seperti struct yang mendefinisikan *data/state*, interface mendefinisikan **perilaku (*behavior*)**: apa yang bisa dilakukan oleh suatu tipe, bukan bagaimana data tersebut disimpan.

## Contoh
example.go
```go
package main

import "fmt"

// Greeter mendefinisikan perilaku menyapa
type Greeter interface {
	Greet() string
}

// User mengimplementasikan Greeter karena memiliki method Greet() string
type User struct {
	Name string
}

func (u User) Greet() string {
	return "Halo, saya " + u.Name
}

// Welcome menerima interface Greeter sebagai parameter
func Welcome(g Greeter) {
	fmt.Println(g.Greet())
}
```

## Penjelasan
- `type Greeter interface { Greet() string }`: deklarasi interface bernama `Greeter`. Tipe apa pun yang memiliki method `Greet() string` secara otomatis dianggap sebagai seorang `Greeter`.
- Fungsi `Welcome(g Greeter)` tidak peduli tipe data konkret apa yang dikirimkan, asalkan memiliki method `Greet() string`.
- Interface membuat kode kita terpisah (*decoupled*) dari detail implementasi struct konkret.

## >_ Exercise
Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.6-interfaces-and-polymorphism/01-defining-behavior-with-interfaces/main.go).

1. Definisikan interface `Speaker`:
   - Memiliki satu method: `Speak() string`
2. Definisikan struct `Robot` dengan field:
   - `Model` (`string`)
3. Buat method pada `Robot`:
   - `(r Robot) Speak() string`: mengembalikan string `"Beep boop! I am robot " + r.Model`
4. Buat fungsi umum `Broadcast(s Speaker) string`:
   - Memanggil `s.Speak()` dan mengembalikan hasilnya.

Expected output ketika `go run main.go`:
```
Robot speaks: Beep boop! I am robot RX-78
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Interface di Go fokus pada apa yang bisa dilakukan suatu objek (*behavior*), bukan apa isi datanya. Menggunakan interface sebagai parameter fungsi membuat kode fleksibel dan mudah diuji.
