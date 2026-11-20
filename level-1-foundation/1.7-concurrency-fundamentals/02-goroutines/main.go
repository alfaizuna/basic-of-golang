package main

import (
	"fmt"
)

// LaunchWorker meluncurkan goroutine yang mengirim status selesai ke channel out.
func LaunchWorker(id int, out chan<- string) {
	// TODO: Luncurkan goroutine dengan kata kunci 'go' yang mengirim "Worker <id> done" ke out
}

func main() {
	_ = fmt.Println
	out := make(chan string)
	_ = out
	// TODO: Panggil LaunchWorker(1, out) dan terima pesannya lalu cetak
}
