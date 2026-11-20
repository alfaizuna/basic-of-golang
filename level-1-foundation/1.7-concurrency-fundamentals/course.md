# Course 1.7: Concurrency Fundamentals

## Tujuan
Memahami fondasi konkurensi di Go: perbedaan mendasar konkurensi vs paralelisme, meluncurkan dan mengelola goroutine, sinkronisasi dengan `sync.WaitGroup`, komunikasi aman antar-goroutine menggunakan unbuffered dan buffered channels, penutupan channel dan iterasi `for range`, deteksi data races (`-race`), proteksi shared state dengan `sync.Mutex` dan `defer mu.Unlock()`, serta multiplexing event dengan `select`, timeouts, dan non-blocking channel operations.

## Prasyarat
Telah menyelesaikan Course 1.1 s.d. Course 1.6.

## Daftar Lesson

| # | Lesson | Fokus | Status |
|---|--------|-------|--------|
| 01 | Concurrency vs Parallelism | konsep dasar, kenapa Go punya goroutine | 🔒 |
| 02 | Goroutines | kata kunci `go`, program berhenti saat `main` selesai | 🔒 |
| 03 | WaitGroup | `sync.WaitGroup`: `Add`, `Done`, `Wait` | 🔒 |
| 04 | Channels Basics | channel tanpa buffer, kirim/terima, sinkronisasi | 🔒 |
| 05 | Buffered Channels | kapasitas channel, kapan kirim memblokir | 🔒 |
| 06 | Closing Channels & range | `close`, `for range ch`, `v, ok := <-ch` | 🔒 |
| 07 | Data Races | contoh race condition, `go run -race` | 🔒 |
| 08 | Mutex | `sync.Mutex`, `Lock`/`Unlock`, pengenalan singkat `defer mu.Unlock()` | 🔒 |
| 09 | select | menunggu beberapa channel sekaligus | 🔒 |
| 10 | Timeouts & Default in select | `time.After`, case `default`, pola worker sederhana | 🔒 |

## Hasil Akhir
Kemampuan menulis kode asinkron yang aman, bebas dari race condition dan deadlock, serta memahami cara mengoordinasikan goroutine dan channel secara idiomatis di Go.
