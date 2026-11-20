package main

import (
	"fmt"
	"os"
)

// SaveReport menulis konten laporan ke dalam berkas.
func SaveReport(filename string, content string) error {
	_ = os.WriteFile
	// TODO: Tulis file dengan os.WriteFile menggunakan permission 0644
	return nil
}

func main() {
	err := SaveReport("report.txt", "Status: Normal\nThroughput: 500 req/s")
	if err != nil {
		fmt.Println("Gagal menyimpan:", err)
		return
	}
	fmt.Println("Report saved successfully")
}
