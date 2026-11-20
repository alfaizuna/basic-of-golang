package main

import "fmt"

type ServerConfig struct {
	Host  *string // nil = default "localhost"
	Port  *int    // nil = default 8080
	Debug *bool   // nil = default false
}

type ResolvedConfig struct {
	Host  string
	Port  int
	Debug bool
}

// Helper untuk membuat pointer dari nilai primitif
func IntPtr(v int) *int {
	return &v
}

func BoolPtr(v bool) *bool {
	return &v
}

// StringPtr mengembalikan pointer ke string s.
// Di Go, kita tidak bisa langsung menulis &"localhost". Helper ini mengatasinya.
func StringPtr(s string) *string {
	// TODO: Kembalikan alamat (&) dari parameter s
	return &s
}

// ResolveServerConfig menerapkan nilai default jika pointer bernilai nil.
// Default values:
// - Host:  "localhost"
// - Port:  8080
// - Debug: false
func ResolveServerConfig(cfg ServerConfig) ResolvedConfig {
	res := ResolvedConfig{
		Host:  "localhost",
		Port:  8080,
		Debug: false,
	}

	// TODO: Periksa tiap field cfg (*Host, *Port, *Debug).
	// Jika field TIDAK nil, timpa nilai res dengan nilai dari pointer (*).
	if cfg.Host != nil {
		res.Host = *cfg.Host
	}

	if cfg.Port != nil {
		res.Port = *cfg.Port
	}

	if cfg.Debug != nil {
		res.Debug = *cfg.Debug
	}

	return res
}

func main() {
	// 1. Config kosong: semua nil -> harus menggunakan semua default
	cfg1 := ServerConfig{}
	res1 := ResolveServerConfig(cfg1)
	fmt.Printf("Default Config: %s:%d (debug: %t)\n", res1.Host, res1.Port, res1.Debug)

	// 2. Custom config: semua field diisi lewat helper pointer
	cfg2 := ServerConfig{
		Host:  StringPtr("0.0.0.0"),
		Port:  IntPtr(3000),
		Debug: BoolPtr(true),
	}
	res2 := ResolveServerConfig(cfg2)
	fmt.Printf("Custom Config: %s:%d (debug: %t)\n", res2.Host, res2.Port, res2.Debug)

	// 3. Explicit zero config: Port diisi 0 (bukan nil!), Debug diisi false
	cfg3 := ServerConfig{
		Port: IntPtr(0),
	}
	res3 := ResolveServerConfig(cfg3)
	fmt.Printf("Explicit Zero Config: %s:%d (debug: %t)\n", res3.Host, res3.Port, res3.Debug)
}
