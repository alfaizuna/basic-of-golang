# Hints

### Hint 1
Panggil `data, err := os.ReadFile(path)`.

### Hint 2
Periksa `if errors.Is(err, os.ErrNotExist) { return "", fmt.Errorf("file not found") }`.

### Hint 3
Jika sukses, `return string(data), nil`.
