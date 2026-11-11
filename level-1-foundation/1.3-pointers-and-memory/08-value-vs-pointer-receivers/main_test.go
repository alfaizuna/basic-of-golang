package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestThermostatMethods(t *testing.T) {
	t.Run("SetTemp sets celsius", func(t *testing.T) {
		th := Thermostat{}
		th.SetTemp(20)
		if th.Celsius != 20 {
			t.Errorf("SetTemp(20): ingin Celsius=20, dapat %.1f", th.Celsius)
		}
	})

	t.Run("Warmup increases celsius", func(t *testing.T) {
		th := Thermostat{Celsius: 20}
		th.Warmup(5)
		if th.Celsius != 25 {
			t.Errorf("Warmup(5) dari 20: ingin Celsius=25, dapat %.1f", th.Celsius)
		}
	})

	t.Run("ToFahrenheit converts correctly", func(t *testing.T) {
		th := Thermostat{Celsius: 25}
		got := th.ToFahrenheit()
		if got != 77.0 {
			t.Errorf("ToFahrenheit() dari 25°C: ingin 77.0, dapat %.1f", got)
		}
	})

	t.Run("ToFahrenheit does not modify struct", func(t *testing.T) {
		th := Thermostat{Celsius: 25}
		th.ToFahrenheit()
		if th.Celsius != 25 {
			t.Errorf("ToFahrenheit() seharusnya tidak mengubah Celsius, dapat %.1f", th.Celsius)
		}
	})

	t.Run("full scenario output", func(t *testing.T) {
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		th := Thermostat{}
		th.SetTemp(20)
		th.Warmup(5)
		fmt.Printf("Celsius   : %.1f\n", th.Celsius)
		fmt.Printf("Fahrenheit: %.1f\n", th.ToFahrenheit())

		w.Close()
		os.Stdout = old
		var buf bytes.Buffer
		io.Copy(&buf, r)

		got := strings.TrimSpace(buf.String())
		want := "Celsius   : 25.0\nFahrenheit: 77.0"
		if got != want {
			t.Errorf("output salah:\ningin:\n%s\ndapat:\n%s", want, got)
		}
	})
}
