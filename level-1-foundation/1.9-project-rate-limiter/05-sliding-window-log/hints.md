# Hints

### Hint 1
Gunakan slice filtering: `var valid []time.Time; for _, t := range s.timestamps { if t.After(cutoff) { valid = append(valid, t) } }; s.timestamps = valid`.

### Hint 2
Lalu periksa `if len(s.timestamps) < s.limit { s.timestamps = append(s.timestamps, cur); return true }`.
