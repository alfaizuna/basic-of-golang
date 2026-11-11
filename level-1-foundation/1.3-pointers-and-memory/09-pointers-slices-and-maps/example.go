package main

import "fmt"

// example.go — hanya untuk referensi, jangan diubah

func doubleAll(nums []int) {
	for i := range nums {
		nums[i] *= 2 // ✅ ubah elemen asli lewat index
	}
}

var _ = fmt.Sprintf // pastikan import tidak unused
