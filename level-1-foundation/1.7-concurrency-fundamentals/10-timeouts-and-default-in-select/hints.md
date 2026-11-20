# Hints

### Hint 1
Gunakan statement `select` dengan dua case.

### Hint 2
Case pertama: `case val := <-ch: return val, nil`.

### Hint 3
Case kedua: `case <-time.After(timeout): return "", fmt.Errorf("operation timed out")`.
