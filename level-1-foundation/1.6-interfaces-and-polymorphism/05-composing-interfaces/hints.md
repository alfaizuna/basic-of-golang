# Hints

### Hint 1
Komposisikan interface cukup dengan menuliskan nama interface di dalamnya:
```go
type Transceiver interface {
    Sender
    Receiver
}
```

### Hint 2
Method receiver ditulis dengan:
```go
func (r *RadioDevice) Send(msg string) error {
    r.lastMessage = msg
    return nil
}

func (r *RadioDevice) Receive() (string, error) {
    return r.lastMessage, nil
}
```

### Hint 3
Di `BroadcastPing`:
```go
func BroadcastPing(t Transceiver, msg string) (string, error) {
    if err := t.Send(msg); err != nil {
        return "", err
    }
    return t.Receive()
}
```
