package main

import "fmt"

func main() {
	// 1. Deklarasikan variabel 'app' bertipe string bernilai "Server" menggunakan 'var'
	var app string = "Server"
	

	// 2. Deklarasikan variabel 'port' bernilai 8080 menggunakan ':='
	port := 8080
	// 3. Ubah nilai 'port' menjadi 9090 menggunakan '='
	port = 9090
	// 4. Cetak 'app' dan 'port' di baris terpisah
	fmt.Println(app)
	fmt.Println(port)
}
