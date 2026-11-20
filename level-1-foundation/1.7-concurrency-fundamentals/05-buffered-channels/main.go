package main

import "fmt"

// QueueJobs membuat buffered channel berisi semua jobs dan mengembalikan channel beserta kapasitasnya.
func QueueJobs(jobs []string) (<-chan string, int) {
	// TODO: Buat buffered channel dengan kapasitas len(jobs)
	// TODO: Masukkan semua item dari jobs ke channel
	// TODO: Kembalikan channel dan cap(ch)
	return nil, 0
}

func main() {
	jobList := []string{"Job1", "Job2", "Job3"}
	ch, capacity := QueueJobs(jobList)
	fmt.Printf("Buffered jobs count: %d\n", capacity)
	_ = ch
}
