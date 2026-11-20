## Hint 1 (arah umum)
Untuk `Enable()` dan `Disable()`:
Tambahkan *nil guard* `if f == nil { return }` di awal method, lalu ubah `f.Enabled = true` atau `f.Enabled = false`.

---

## Hint 2 (sintaks yang relevan)
Untuk `SetRollout(percentage int) error`:
```go
if f == nil {
    return errors.New("flag is nil")
}
if percentage < 0 || percentage > 100 {
    return errors.New("rollout must be between 0 and 100")
}
f.Rollout = percentage
return nil
```

---

## Hint 3 (hampir jawaban)
```go
func (f *Flag) Enable() {
	if f == nil {
		return
	}
	f.Enabled = true
}

func (f *Flag) Disable() {
	if f == nil {
		return
	}
	f.Enabled = false
}

func (f *Flag) SetRollout(percentage int) error {
	if f == nil {
		return errors.New("flag is nil")
	}
	if percentage < 0 || percentage > 100 {
		return errors.New("rollout must be between 0 and 100")
	}
	f.Rollout = percentage
	return nil
}
```
