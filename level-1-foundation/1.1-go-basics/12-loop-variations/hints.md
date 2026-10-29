### Hint 1
Untuk `countVowels`, kamu bisa mengiterasi karakter string menggunakan `for _, char := range text` dan mencocokkan karakter menggunakan tanda petik tunggal seperti `'a'` (tipe rune).

### Hint 2
Bentuk fungsi `countVowels`:
```go
func countVowels(text string) int {
	total := 0
	for _, char := range text {
		switch char {
		case 'a', 'i', 'u', 'e', 'o':
			total++
		}
	}
	return total
}
```

### Hint 3
Bentuk fungsi `findTarget` menggunakan *labeled break*:
```go
func findTarget(matrixSize int, targetSum int) (int, int) {
	foundR, foundC := 0, 0
SearchLoop:
	for r := 1; r <= matrixSize; r++ {
		for c := 1; c <= matrixSize; c++ {
			if r+c == targetSum {
				foundR = r
				foundC = c
				break SearchLoop
			}
		}
	}
	return foundR, foundC
}
```
Lalu di `main()`, panggil kedua fungsi tersebut dan cetak dengan `fmt.Println`.
