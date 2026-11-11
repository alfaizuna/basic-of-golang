package main

import "fmt"

type Thermostat struct {
	Celsius float64
}

// TODO: tambahkan method SetTemp(c float64) di sini
// (pointer receiver — ubah suhu ke nilai baru)
func (t *Thermostat) SetTemp(c float64) {
	t.Celsius = c
}

// TODO: tambahkan method Warmup(delta float64) di sini
// (pointer receiver — naikkan suhu sebesar delta)
func (t *Thermostat) Warmup(delta float64) {
	t.Celsius += delta
}

// TODO: tambahkan method ToFahrenheit() float64 di sini
// (value receiver — kembalikan suhu dalam Fahrenheit, rumus: C×9/5 + 32)
func (t Thermostat) ToFahrenheit() float64 {
	return (t.Celsius*9)/5 + 32
}
func main() {
	t := Thermostat{}
	t.SetTemp(20)
	t.Warmup(5)
	fmt.Printf("Celsius   : %.1f\n", t.Celsius)
	fmt.Printf("Fahrenheit: %.1f\n", t.ToFahrenheit())
}
