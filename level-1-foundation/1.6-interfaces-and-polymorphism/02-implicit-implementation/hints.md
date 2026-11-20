## Hint 1 (arah umum)
Definisikan struct `EmailSender` dan struct `SMSSender`.
Berikan method `Notify(message string) string` pada kedua struct tersebut.

---

## Hint 2 (sintaks yang relevan)
```go
func (e EmailSender) Notify(message string) string {
	return "[EMAIL to " + e.Email + "] " + message
}

func (s SMSSender) Notify(message string) string {
	return "[SMS to " + s.Phone + "] " + message
}
```

---

## Hint 3 (hampir jawaban)
```go
func BroadcastAlert(notifiers []Notifier, alertMessage string) []string {
	var results []string
	for _, n := range notifiers {
		results = append(results, n.Notify(alertMessage))
	}
	return results
}
```
