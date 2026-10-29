package main

import "fmt"

func main() {
	cpuUsage := 88
	memoryUsage := 60
	isOnline := true

	// Buat aturan server monitor:
	// 1. Jika TIDAK online ATAU cpuUsage > 95: cetak "CRITICAL"
	// 2. Jika online DAN (cpuUsage > 80 ATAU memoryUsage > 80): cetak "WARNING"
	// 3. Selain itu: cetak "OK"

	if !isOnline || cpuUsage > 95 {
		fmt.Println("CRITICAL")
	} else if isOnline && (cpuUsage >80 || memoryUsage > 80) {
		fmt.Println("WARNING")
	} else {
		fmt.Println("OK")
	}
}
