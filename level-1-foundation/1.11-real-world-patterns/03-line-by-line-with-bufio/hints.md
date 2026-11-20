# Hints

### Hint 1
Inisialisasi `scanner := bufio.NewScanner(r)`.

### Hint 2
Lakukan loop `for scanner.Scan() { line := scanner.Text(); if strings.Contains(line, keyword) { matched = append(matched, line) } }`.

### Hint 3
Cek `if err := scanner.Err(); err != nil { return nil, err }`.
