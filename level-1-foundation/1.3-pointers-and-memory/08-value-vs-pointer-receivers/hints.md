## Hint 1 (arah umum)
Ada tiga method yang perlu dibuat. Dua di antaranya mengubah `Celsius` (pakai pointer receiver), satu hanya menghitung dan mengembalikan nilai baru (cukup value receiver).

---

## Hint 2 (sintaks yang relevan)
Pola pointer receiver untuk yang mengubah:
```go
func (t *Thermostat) SetTemp(c float64) {
    t.Celsius = ___
}
```
Pola value receiver untuk yang hanya membaca:
```go
func (t Thermostat) ToFahrenheit() float64 {
    return ___ // gunakan t.Celsius dan rumus C×9/5 + 32
}
```

---

## Hint 3 (hampir jawaban)
```go
func (t *Thermostat) SetTemp(c float64) {
    t.Celsius = c
}

func (t *Thermostat) Warmup(delta float64) {
    t.Celsius += ___
}

func (t Thermostat) ToFahrenheit() float64 {
    return t.Celsius*___/___ + 32
}
```
Isi bagian `___` dengan angka atau variabel yang tepat.
