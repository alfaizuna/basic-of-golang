# Value vs Pointer Receivers

Kamu sudah tahu cara menulis pointer receiver. Sekarang pertanyaannya: **kapan harus pakai value, kapan harus pakai pointer?**

## Contoh

```go
// example.go
package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// Value receiver — hanya membaca, tidak mengubah
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Value receiver — hanya membaca
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Pointer receiver — mengubah struct
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	r := Rectangle{Width: 4, Height: 3}
	fmt.Println(r.Area())      // 12
	r.Scale(2)
	fmt.Println(r.Area())      // 48
}
```

## Penjelasan

- Gunakan **pointer receiver** jika method **mengubah** field struct — supaya perubahan terlihat di pemanggil.
- Gunakan **value receiver** jika method hanya **membaca** — lebih aman, lebih mudah dipahami.
- **Aturan konsistensi**: jika salah satu method di tipe kamu pakai pointer receiver, **semua method sebaiknya pakai pointer receiver** juga — ini konvensi Go agar tidak membingungkan.
- Pengecualian: tipe yang sangat kecil dan tidak punya pointer receiver sama sekali (misalnya `time.Time`) boleh pakai value receiver saja.

## >_ Exercise

Kamu punya struct `Thermostat` dengan field:
- `Celsius float64` — suhu saat ini

Tambahkan **tiga method**:

1. `SetTemp(c float64)` — pointer receiver, ubah suhu ke nilai baru.
2. `Warmup(delta float64)` — pointer receiver, naikkan suhu sebesar `delta`.
3. `ToFahrenheit() float64` — value receiver, kembalikan suhu dalam Fahrenheit **(tanpa mengubah struct)**.  
   Rumus: `F = C × 9/5 + 32`

Di `main()` sudah ada kode ini — **jangan ubah**:

```go
t := Thermostat{}
t.SetTemp(20)
t.Warmup(5)
fmt.Printf("Celsius   : %.1f\n", t.Celsius)
fmt.Printf("Fahrenheit: %.1f\n", t.ToFahrenheit())
```

Expected output:
```
Celsius   : 25.0
Fahrenheit: 77.0
```

## Hints
Hints (0/3) — gunakan `/hint` untuk membuka satu per satu.

## Key Takeaway
Pilih pointer receiver saat method mengubah state; pilih value receiver saat hanya membaca. Yang terpenting: **konsisten** — satu tipe sebaiknya tidak mencampur keduanya secara acak.
