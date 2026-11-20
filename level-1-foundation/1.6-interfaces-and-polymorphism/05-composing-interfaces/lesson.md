# Composing Interfaces

Go mendukung komposisi interface: membuat interface baru yang lebih besar dengan menggabungkan antarmuka-antarmuka kecil yang terfokus.

## Contoh

```go
package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}
```

## Penjelasan

1. **Prinsip Antarmuka Kecil**: Idiom Go mendorong pembuatan interface kecil (seringkali hanya 1 atau 2 method, seperti `io.Reader` atau `fmt.Stringer`).
2. **Embedding**: Kita dapat menyematkan (*embed*) satu atau lebih interface ke dalam interface lain tanpa perlu mendeklarasikan ulang method-method tersebut.
3. **Kepuasan Implisit**: Sebuah struct otomatis memenuhi `ReadWriter` jika struct tersebut mengimplementasikan semua method dari `Reader` dan `Writer`.

## >_ Exercise

1. Definisikan dua interface:
   - `Sender`: memiliki method `Send(msg string) error`
   - `Receiver`: memiliki method `Receive() (string, error)`
2. Komposisikan kedua interface tersebut menjadi interface `Transceiver`:
   ```go
   type Transceiver interface {
       Sender
       Receiver
   }
   ```
3. Buat struct `RadioDevice` yang memiliki field `lastMessage string`.
4. Implementasikan method `Send` dan `Receive` pada pointer receiver `*RadioDevice`:
   - `Send(msg string) error`: simpan `msg` ke `lastMessage`, kembalikan `nil`.
   - `Receive() (string, error)`: kembalikan `lastMessage` dan `nil`.
5. Buat fungsi `BroadcastPing(t Transceiver, msg string) (string, error)` yang memanggil `t.Send(msg)` lalu mengembalikan `t.Receive()`.

### Expected output

```
Terkirim & Diterima: Ping 101
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Interface di Go dapat dikomposisikan dengan menyematkan antarmuka lain di dalamnya.
- Komposisi memungkinkan desain modular: fungsi yang hanya butuh `Sender` hanya meminta `Sender`, sedangkan fungsi yang butuh keduanya dapat meminta `Transceiver`.
