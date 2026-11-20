## Hint 1 (arah umum)
Definisikan struct `CacheEntry` dan `DNSCache`.
Di konstruktor `NewDNSCache`, inisialisasi map dengan `make(map[string]*CacheEntry)` dan fungsi waktu `now: time.Now`.

---

## Hint 2 (sintaks yang relevan)
Untuk method `Set`:
```go
key := string(rec.Type) + ":" + rec.Name
expiresAt := c.now().Add(time.Duration(rec.TTL) * time.Second)
c.entries[key] = &CacheEntry{
    Record:    rec,
    ExpiresAt: expiresAt,
}
```

---

## Hint 3 (hampir jawaban)
```go
func (c *DNSCache) Get(name string, recordType RecordType) *Record {
	if c == nil || c.entries == nil {
		return nil
	}
	key := string(recordType) + ":" + name
	entry, ok := c.entries[key]
	if !ok {
		return nil
	}
	if c.now().Before(entry.ExpiresAt) {
		return entry.Record
	}
	delete(c.entries, key)
	return nil
}
```
