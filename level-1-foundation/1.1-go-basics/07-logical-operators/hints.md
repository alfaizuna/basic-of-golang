### Hint 1
Gunakan operator negasi `!isOnline` untuk mengecek apakah server offline, dan gabungkan dengan `cpuUsage > 95` menggunakan `||`.

### Hint 2
Untuk kondisi kedua, gunakan tanda kurung untuk mengelompokkan pengecekan CPU dan Memory:
`isOnline && (cpuUsage > 80 || memoryUsage > 80)`

### Hint 3
Lengkapi `func main()` dengan struktur berikut:
```go
if !isOnline || cpuUsage > 95 {
	fmt.Println("CRITICAL")
} else if isOnline && (cpuUsage > 80 || memoryUsage > 80) {
	fmt.Println("WARNING")
} else {
	fmt.Println("OK")
}
```
