# Hints

### Hint 1
Periksa `if r.Method != http.MethodGet { w.WriteHeader(http.StatusMethodNotAllowed); w.Write([]byte("Method Not Allowed")); return }`.

### Hint 2
Ambil query: `q := r.URL.Query().Get("q"); if q == "" { q = "all" }`.

### Hint 3
Tulis output: `fmt.Fprintf(w, "Query: %s", q)`.
