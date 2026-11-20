//go:build ignore

package main

import (
	"fmt"
	"runtime"
	"time"
)

func task(name string) {
	for i := 1; i <= 2; i++ {
		fmt.Printf("Task %s: step %d\n", name, i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	fmt.Printf("Go runtime memiliki %d logical CPU core\n", runtime.NumCPU())
	task("A")
	task("B")
}
