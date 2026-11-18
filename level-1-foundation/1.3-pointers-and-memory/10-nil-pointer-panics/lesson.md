# Nil Pointer Panics

Di Go, variabel pointer yang belum diinisialisasi memiliki *zero value* berupa `nil`.
Jika kamu mencoba mengakses nilai atau field melalui pointer yang bernilai `nil` (*dereference nil pointer*), Go tidak bisa melanjutkan dan program akan langsung berhenti seketika dengan **runtime panic**.

## Contoh Panic & Stack Trace

Perhatikan apa yang terjadi ketika kode berikut dijalankan:

```go
package main

import "fmt"

type Profile struct {
	Bio string
}

func printBio(p *Profile) {
	fmt.Println("Bio:", p.Bio) // 💥 BARIS 11: p adalah nil!
}

func main() {
	var userProfile *Profile // nil
	printBio(userProfile)
}
```

Ketika dijalankan, Go menghentikan program seketika dan mencetak:

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x...]

goroutine 1 [running]:
main.printBio(0x0)
	/path/to/main.go:11 +0x24
main.main()
	/path/to/main.go:16 +0x2b
exit status 2
```

## Cara Membaca Pesan Panic & Stack Trace

Stack trace mungkin terlihat rumit, tetapi polanya selalu konsisten:

1. **Jenis Error (Baris Pertama)**:
   `panic: runtime error: invalid memory address or nil pointer dereference`
   Artinya: program mencoba membaca/menulis ke alamat memori yang tidak valid (alamat `0x0` alias `nil`).

2. **Titik Kejadian (Puncak Stack)**:
   ```
   main.printBio(0x0)
   	/path/to/main.go:11 +0x24
   ```
   - `main.printBio`: nama fungsi tempat crash terjadi.
   - `(0x0)`: argumen pointer yang diterima bernilai `0x0` (yaitu `nil`).
   - `/path/to/main.go:11`: **file dan nomor baris pasti** penyebab crash.

3. **Alur Pemanggilan (Call Stack)**:
   ```
   main.main()
   	/path/to/main.go:16 +0x2b
   ```
   Baris-baris di bawahnya menunjukkan urutan pemanggilan fungsi: `main()` di baris 16 memanggil `printBio()`.

## Penyebab Umum Nil Pointer Panic

1. **Lupa mengalokasikan struct sebelum mengisi field**:
   ```go
   var u *User
   u.Name = "Budi" // 💥 Panic! u masih nil
   // Solusi: u = &User{Name: "Budi"}
   ```

2. **Nested struct pointer bernilai nil**:
   ```go
   type Server struct {
       Config *Config // pointer field
   }
   s := &Server{}
   s.Config.Port = 8080 // 💥 Panic! s.Config masih nil
   // Solusi: s.Config = &Config{Port: 8080}
   ```

## >_ Exercise

Buka file [main.go](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/10-nil-pointer-panics/main.go). Di dalamnya ada sistem registrasi profil pengguna dan konfigurasi server yang saat ini **crash karena nil pointer panic**.

Coba jalankan dengan `go run main.go` di terminal untuk melihat pesan panic dan membaca stack trace-nya!

Tugasmu:
1. **Perbaiki [`createProfile`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/10-nil-pointer-panics/main.go#L21-L28)**:
   Fungsi ini mendeklarasikan `var p *Profile` dan langsung mengisi `p.Bio = bio` sebelum mengalokasikan memori struct, sehingga memicu panic. Perbaiki agar mengalokasikan struct [`Profile`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/10-nil-pointer-panics/main.go#L5-L7) yang valid dan mengembalikan pointernya.
2. **Perbaiki [`loadConfig`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/10-nil-pointer-panics/main.go#L32-L42)**:
   Struct [`AppConfig`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/10-nil-pointer-panics/main.go#L14-L17) dibuat tanpa menginisialisasi nested pointer `Server *ServerConfig`. Saat kode mengakses `cfg.Server.Host` dan `cfg.Server.Port`, terjadi panic. Inisialisasi struct [`ServerConfig`](file:///Users/alfaizuna/Documents/Projects/golang/basic-of-golang/level-1-foundation/1.3-pointers-and-memory/10-nil-pointer-panics/main.go#L9-L12) tersebut dengan benar.

Expected output ketika `go run main.go`:
```
Profile: Software Engineer
Config: WebApp on localhost:8080
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Dereferencing pointer yang bernilai `nil` memicu *runtime panic*. Jangan pernah membaca atau menulis field dari pointer yang belum dialokasikan memori. Selalu inisialisasi pointer dengan `&T{...}` atau `new(T)`.
