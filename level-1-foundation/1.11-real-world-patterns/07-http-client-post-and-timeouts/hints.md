# Hints

### Hint 1
Panggil `resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))`.

### Hint 2
Periksa `if err != nil { return 0, err }`.

### Hint 3
`defer resp.Body.Close()` dan kembalikan `resp.StatusCode, nil`.
