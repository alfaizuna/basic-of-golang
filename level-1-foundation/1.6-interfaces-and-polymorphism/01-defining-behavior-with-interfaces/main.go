package main

import "fmt"

// TODO 1: Definisikan interface Speaker dengan method Speak() string
type Speaker interface {
	Speak() string
}

// TODO 2: Definisikan struct Robot dengan field Model string
type Robot struct {
	Model string
}

// TODO 3: Buat method (r Robot) Speak() string
// Kembalikan "Beep boop! I am robot " + r.Model
func (r Robot) Speak() string {
	return ""
}

// TODO 4: Buat fungsi Broadcast(s Speaker) string
func Broadcast(s Speaker) string {
	return ""
}

func main() {
	rx := Robot{Model: "RX-78"}
	msg := Broadcast(rx)
	fmt.Printf("Robot speaks: %s\n", msg)
}
