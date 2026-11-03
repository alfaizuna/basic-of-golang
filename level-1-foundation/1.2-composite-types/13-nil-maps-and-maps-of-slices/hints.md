### Hint 1
Di dalam `addStudent`, panggil `append(classroom[grade], name)` dan simpan kembali ke `classroom[grade]`.

### Hint 2
Bentuk fungsi `addStudent`:
```go
classroom[grade] = append(classroom[grade], name)
```

### Hint 3
Kerangka kode lengkap:
```go
func addStudent(classroom map[string][]string, grade string, name string) {
	classroom[grade] = append(classroom[grade], name)
}

func main() {
	classroom := make(map[string][]string)
	addStudent(classroom, "Kelas A", "Budi")
	addStudent(classroom, "Kelas A", "Siti")
	addStudent(classroom, "Kelas B", "Andi")

	fmt.Println("Kelas A:", classroom["Kelas A"])
	fmt.Println("Kelas B:", classroom["Kelas B"])
}
```
