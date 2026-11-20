# Hints

### Hint 1
Buka fungsi dengan struktur type switch:
```go
switch v := val.(type) {
case string:
    // ...
}
```

### Hint 2
Untuk `int`, lakukan perkalian `v * 2` dan gunakan `fmt.Sprintf("INT: %d", v*2)`.
Untuk `bool`, gunakan `if v { return "BOOL: YES" } else { return "BOOL: NO" }`.

### Hint 3
Bentuk lengkapnya:
```go
switch v := val.(type) {
case string:
    return fmt.Sprintf("STR: %s", v)
case int:
    return fmt.Sprintf("INT: %d", v*2)
case bool:
    if v {
        return "BOOL: YES"
    }
    return "BOOL: NO"
default:
    return "UNKNOWN"
}
```
