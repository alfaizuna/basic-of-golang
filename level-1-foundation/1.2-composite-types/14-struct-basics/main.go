package main

import "fmt"

// TODO: Definisikan type Book struct dengan field: Title (string), Author (string), Pages (int)
type Book struct {
	Title string
	Author string
	Pages int
}

func main() {
	// TODO:
	// 1. Buat instance Book:
	//    Title: "The Go Programming Language"
	//    Author: "Alan Donovan"
	//    Pages: 380
	b := Book{
		Title: "The Go Programming Language",
	    Author: "Alan Donovan",
	    Pages: 380,
	}
	
	// 2. Cetak: fmt.Printf("%s karya %s (%d halaman)\n", b.Title, b.Author, b.Pages)
	fmt.Printf("%s karya %s (%d halaman)\n", b.Title, b.Author, b.Pages)
	
	_ = fmt.Println
}
