package main

import (
	"fmt"
	"net/http"
)

// UserSearchHandler menangani pencarian user via query param q dengan validasi method GET.
func UserSearchHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Validasi method GET dan baca query param q (default "all")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", UserSearchHandler)
	fmt.Println("Mux configured")
}
