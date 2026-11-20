# Hints

### Hint 1
Deklarasikan interface `Fetcher` dan struct `FakeFetcher`.

### Hint 2
Di method Fetch: `res, ok := f.Data[url]`. Jika `!ok` return `"", nil, fmt.Errorf("not found: %s", url)`.

### Hint 3
Jika ada, return `res.Body, res.URLs, nil`.
