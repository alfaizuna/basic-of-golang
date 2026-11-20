## Hint 1 (arah umum)
Definisikan interface `Speaker` dengan satu method `Speak() string`.
Pastikan nama method dan return type pada struct `Robot` sama persis dengan yang ada di interface.

---

## Hint 2 (sintaks yang relevan)
```go
type Speaker interface {
	Speak() string
}

type Robot struct {
	Model string
}
```

---

## Hint 3 (hampir jawaban)
```go
func (r Robot) Speak() string {
	return "Beep boop! I am robot " + r.Model
}

func Broadcast(s Speaker) string {
	return s.Speak()
}
```
