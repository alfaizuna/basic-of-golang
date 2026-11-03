package main

import "fmt"

func addStudent(classroom map[string][]string, grade string, name string) {
	// TODO: tambahkan name ke slice classroom[grade] menggunakan append
	classroom[grade] = append(classroom[grade], name)
}

func main() {
	// TODO:
	// 1. Buat classroom := make(map[string][]string)
	classroom := make(map[string][]string)
	// 2. addStudent(classroom, "Kelas A", "Budi")
	addStudent(classroom, "Kelas A", "Budi")
	// 3. addStudent(classroom, "Kelas A", "Siti")
	addStudent(classroom, "Kelas A", "Siti")
	// 4. addStudent(classroom, "Kelas B", "Andi")
	addStudent(classroom, "Kelas B", "Andi")

	// 5. Cetak: fmt.Println("Kelas A:", classroom["Kelas A"])
	fmt.Println("Kelas A:", classroom["Kelas A"])

	// 6. Cetak: fmt.Println("Kelas B:", classroom["Kelas B"])
	fmt.Println("Kelas B:", classroom["Kelas B"])

	_ = fmt.Println
}
