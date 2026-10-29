package main

import (
	"errors"
	"fmt"
)

// divide membagi a dengan b.
// Jika b == 0, kembalikan 0 dan error "cannot divide by zero".
// Jika tidak, kembalikan hasil a / b dan nil.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	// TODO: kembalikan hasil pembagian a / b dan nil (saat ini masih 0, nil)
	return a / b, nil
}

func main() {
	// Contoh penanganan error untuk divide(10, 2):
	res1, err1 := divide(10, 2)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(res1)
	}

	// TODO: Panggil divide(10, 0)
	// Periksa dengan pola 'if err2 != nil':
	// Jika terjadi error cetak err2, jika sukses cetak res2
	res2, err2 := divide(10,0)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(res2)
	}
}
