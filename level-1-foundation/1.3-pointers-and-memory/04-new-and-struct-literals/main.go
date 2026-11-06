package main

import "fmt"

type ServerConfig struct {
	Host string
	Port int
}

func main() {
	// 1. Alokasikan pointer int menggunakan new(int) dan simpan ke variabel version
	// TODO: version := new(...)
	version := new(int)

	// 2. Isi nilai yang ditunjuk oleh version dengan angka 1
	// TODO: *version = 1
	*version = 1

	// 3. Cetak "Version:", *version
	// TODO: fmt.Println("Version:", ...)
	fmt.Println("Version:",  *version)

	// 4. Buat pointer struct ServerConfig menggunakan &ServerConfig{...}
	// dengan Host: "localhost" dan Port: 8080, simpan ke variabel cfg
	// TODO: cfg := &ServerConfig{...}
	cfg := &ServerConfig{
		Host: "localhost",
		Port: 8080,
	}

	// 5. Cetak dengan fmt.Printf("Server running at %s:%d\n", cfg.Host, cfg.Port)
	// TODO: fmt.Printf(...)
	fmt.Printf("Server running at %s:%d\n", cfg.Host, cfg.Port)

	_ = fmt.Println
}
