//go:build ignore

package main

import (
	"fmt"
	"log"
)

type Config struct {
	Host string
	Port int
}

func main() {
	cfg := Config{Host: "localhost", Port: 8080}

	// %v: hanya nilai
	fmt.Printf("%%v:  %v\n", cfg)

	// %+v: menyertakan nama field struct
	fmt.Printf("%%+v: %+v\n", cfg)

	// %#v: representasi sintaks Go lengkap
	fmt.Printf("%%#v: %#v\n", cfg)

	log.Println("Pesan log otomatis menyertakan timestamp")
}
