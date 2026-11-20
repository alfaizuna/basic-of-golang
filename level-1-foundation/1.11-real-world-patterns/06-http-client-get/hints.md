# Hints

### Hint 1
Panggil `resp, err := http.Get(url)`.

### Hint 2
Lakukan `defer resp.Body.Close()`.

### Hint 3
Baca isi body dengan `bytes, err := io.ReadAll(resp.Body)`.

### Hint 4
Kembalikan `resp.StatusCode, string(bytes), nil`.
