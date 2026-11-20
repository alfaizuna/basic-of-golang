package main

import (
	"errors"
	"fmt"
	"os"
)

// LoadConfigText membaca berkas teks dan mengembalikan string atau error yang sesuai.
func LoadConfigText(path string) (string, error) {
	_ = errors.Is
	_ = os.ReadFile
	// TODO: Baca file dengan os.ReadFile dan cek error os.ErrNotExist
	return "", nil
}

func main() {
	_, err := LoadConfigText("non_existent_file.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
