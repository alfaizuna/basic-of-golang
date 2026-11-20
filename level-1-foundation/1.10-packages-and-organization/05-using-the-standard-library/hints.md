# Hints

### Hint 1
Panggil `clean := strings.TrimSpace(raw)`.

### Hint 2
Untuk parsing integer yang berada di depan, bisa gunakan `strconv.Atoi(clean)`.

### Hint 3
Jika parsing error, `return 0, "", err`.
