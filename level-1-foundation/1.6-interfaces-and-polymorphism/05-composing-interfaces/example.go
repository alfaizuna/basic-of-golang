//go:build ignore

package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

// ReadWriter mengomposisikan Reader dan Writer
type ReadWriter interface {
	Reader
	Writer
}

type MemoryBuffer struct {
	content string
}

func (m *MemoryBuffer) Read() string {
	return m.content
}

func (m *MemoryBuffer) Write(data string) {
	m.content = data
}

func process(rw ReadWriter) {
	rw.Write("Hello Composite Interface")
	fmt.Println("Read result:", rw.Read())
}

func main() {
	buf := &MemoryBuffer{}
	process(buf)
}
