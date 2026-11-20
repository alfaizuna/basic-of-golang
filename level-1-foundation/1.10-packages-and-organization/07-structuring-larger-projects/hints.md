# Hints

### Hint 1
Gunakan fungsi rekursif dengan parameter `visited map[string]bool`.

### Hint 2
Tandai `visited[node] = true`. Telusuri anak-anak dependensi `dependencies[node]`.

### Hint 3
Jika anak sudah ada di `visited`, maka cycle terdeteksi (`return true`).
