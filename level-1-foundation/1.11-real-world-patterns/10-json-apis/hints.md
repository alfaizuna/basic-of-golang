# Hints

### Hint 1
Set header: `w.Header().Set("Content-Type", "application/json")`.

### Hint 2
Decode: `var req SumRequest; if err := json.NewDecoder(r.Body).Decode(&req); err != nil { w.WriteHeader(http.StatusBadRequest); return }`.

### Hint 3
Encode: `json.NewEncoder(w).Encode(SumResponse{Sum: req.A + req.B})`.
