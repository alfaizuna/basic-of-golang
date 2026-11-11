## Hint 1 (arah umum)
Untuk `addBonus`: masalahnya ada di cara kamu mengiterasi slice. `for _, v := range` memberikan **salinan** nilai — ubah cara loop agar kamu bisa memodifikasi elemen aslinya.

Untuk `resetMap`: kamu tidak bisa mengganti map yang diterima function dengan map baru (karena header map di pemanggil tidak berubah). Pikirkan cara menghapus isi map **satu per satu** tanpa menggantinya.

---

## Hint 2 (sintaks yang relevan)
Untuk `addBonus`, gunakan index saat iterasi:
```go
for i := range scores {
    scores[i] += bonus // ✅ mengubah elemen asli
}
```

Untuk `resetMap`, Go punya built-in `delete(m, key)`. Kamu perlu iterasi semua key dan hapus satu per satu:
```go
for k := range m {
    delete(m, k)
}
```

---

## Hint 3 (hampir jawaban)
```go
func addBonus(scores []int, bonus int) {
	for ___ := range scores {
		scores[___] += bonus
	}
}

func resetMap(m map[string]int) {
	for k := range m {
		delete(___, ___)
	}
}
```
Isi bagian `___` dengan variabel yang tepat.
