package main

import "fmt"

// TODO 1: Definisikan interface Notifier dengan method Notify(message string) string
type Notifier interface {
	Notify(message string) string
}

// TODO 2: Definisikan struct EmailSender (Email string)
type EmailSender struct {
	Email string
}

// TODO: Implementasikan method (e EmailSender) Notify(message string) string
func (e EmailSender) Notify(message string) string {
	return ""
}

// TODO 3: Definisikan struct SMSSender (Phone string)
type SMSSender struct {
	Phone string
}

// TODO: Implementasikan method (s SMSSender) Notify(message string) string
func (s SMSSender) Notify(message string) string {
	return ""
}

// TODO 4: Buat fungsi BroadcastAlert(notifiers []Notifier, alertMessage string) []string
func BroadcastAlert(notifiers []Notifier, alertMessage string) []string {
	return nil
}

func main() {
	channels := []Notifier{
		EmailSender{Email: "user@example.com"},
		SMSSender{Phone: "+62812345678"},
	}

	messages := BroadcastAlert(channels, "Server down!")
	for i, msg := range messages {
		fmt.Printf("Notification %d: %s\n", i+1, msg)
	}
}
