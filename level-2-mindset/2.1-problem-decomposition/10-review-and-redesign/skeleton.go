package main

import "fmt"

// Notifier mendefinisikan kontrak pengiriman notifikasi pesan.
type Notifier interface {
	Notify(to string, text string) error
}

// AppointmentReminder menyimpan data pengingat janji temu.
type AppointmentReminder struct {
	PatientPhone string
	MessageText  string
}

// ReminderService mengorkestrasi pengiriman pengingat menggunakan interface Notifier.
type ReminderService struct {
	notifier Notifier
}

// NewReminderService adalah constructor dengan Dependency Injection.
func NewReminderService(n Notifier) *ReminderService {
	return &ReminderService{notifier: n}
}

// SendReminder mengeksekusi pengiriman notifikasi.
func (s *ReminderService) SendReminder(r AppointmentReminder) error {
	if s.notifier == nil {
		return fmt.Errorf("notifier is required")
	}
	return s.notifier.Notify(r.PatientPhone, r.MessageText)
}

func main() {
	fmt.Println("Go Blueprint skeleton compiled successfully")
}
