package main

import "fmt"

// Sender interface
type Sender interface {
	Send(msg string) error
}

// Receiver interface
type Receiver interface {
	Receive() (string, error)
}

// TODO: Definisikan Transceiver dengan mengomposisikan Sender dan Receiver
type Transceiver interface {
}

// RadioDevice merepresentasikan perangkat radio
type RadioDevice struct {
	lastMessage string
}

// TODO: Implementasikan Send(msg string) error pada *RadioDevice

// TODO: Implementasikan Receive() (string, error) pada *RadioDevice

// BroadcastPing mengirim pesan lewat Transceiver dan mengambil balasannya
func BroadcastPing(t Transceiver, msg string) (string, error) {
	// TODO: Panggil t.Send(msg), jika err tidak nil return "", err. Lalu kembalikan t.Receive()
	return "", nil
}

func main() {
	radio := &RadioDevice{}
	got, err := BroadcastPing(radio, "Ping 101")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Terkirim & Diterima:", got)
}
