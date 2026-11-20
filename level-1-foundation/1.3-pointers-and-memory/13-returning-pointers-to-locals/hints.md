## Hint 1 (arah umum)
Untuk `NewCounter`, kamu cukup membuat variabel lokal `val := initial`, kemudian return alamat memori dari variabel tersebut menggunakan operator `&`.

Untuk `NewProfile`, kamu bisa membuat struct literal `UserProfile{...}` dan mengembalikan pointernya langsung dengan `&` di depannya.

---

## Hint 2 (sintaks yang relevan)
```go
// Mengambil alamat variabel lokal int
val := initial
return &val

// Mengembalikan pointer ke struct literal
return &UserProfile{
    Username: username,
    Email:    email,
}
```

---

## Hint 3 (hampir jawaban)
```go
func NewCounter(initial int) *int {
	val := initial
	return &val
}

func NewProfile(username, email string) *UserProfile {
	return &UserProfile{
		Username: username,
		Email:    email,
	}
}
```
