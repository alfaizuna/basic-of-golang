# Hints

### Hint 1
Deklarasikan `var cfg ServerConfig`.

### Hint 2
Panggil `err := json.Unmarshal(data, &cfg)`.

### Hint 3
Jika ada error, `return nil, err`. Jika berhasil, `return &cfg, nil`.
