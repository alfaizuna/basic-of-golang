### Hint 1
Gunakan fungsi bawaan `new(int)` untuk mengalokasikan integer dengan nilai awal 0 (zero value), dan gunakan sintaks literal `&ServerConfig{...}` untuk membuat pointer ke struct sekaligus mengisi nilai awalnya.

### Hint 2
- Pointer int: `version := new(int)` lalu ubah nilainya dengan `*version = 1`.
- Pointer struct:
  ```go
  cfg := &ServerConfig{
      Host: "localhost",
      Port: 8080,
  }
  ```

### Hint 3
Kerangka kode lengkap di dalam `main()`:
```go
version := new(int)
*version = 1
fmt.Println("Version:", *version)

cfg := &ServerConfig{
	Host: "localhost",
	Port: 8080,
}
fmt.Printf("Server running at %s:%d\n", cfg.Host, cfg.Port)
```
