//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Pastikan Done() dipanggil saat fungsi selesai
	fmt.Printf("Worker %d mulai\n", id)
	time.Sleep(20 * time.Millisecond)
	fmt.Printf("Worker %d selesai\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Tambah counter sebelum meluncurkan goroutine
		go worker(i, &wg) // Wajib passing pointer *sync.WaitGroup
	}

	wg.Wait() // Blokir sampai counter kembali ke 0
	fmt.Println("Semua worker telah selesai")
}
