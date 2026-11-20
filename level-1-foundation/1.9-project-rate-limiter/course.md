# Course 1.9: Project: Rate Limiter

## Tujuan
Membangun sistem pembatas laju trafik (*Rate Limiter*) produksi dari nol menggunakan struct, pointer receiver, map, slice, dan mutex. Pelajari tiga algoritma inti: Fixed Window Counter, Sliding Window Log, dan Token Bucket. Pelajari teknik injeksi waktu (*clock injection*) untuk pengujian deterministik, pembersihan memori klien idle (*garbage collection in-memory*), serta pengamanan konkurensi dengan `sync.Mutex` dan verifikasi bebas race condition (`-race`).

## Prasyarat
Telah menyelesaikan Course 1.1 s.d. Course 1.8.

## Struktur Project
Setiap lesson membangun tahap pemahaman dan komponen rate limiter yang terstruktur dan dapat diuji secara mandiri maupun terintegrasi.

## Daftar Lesson

| # | Lesson | Fokus | Status |
|---|--------|-------|--------|
| 01 | Limiter Model | struct `Limiter`, clock injection `now()`, method `Allow` | 🔒 |
| 02 | Fixed Window Counter | pembatasan laju dengan jendela waktu tetap | 🔒 |
| 03 | Resetting Windows | pergantian jendela dan penanganan batas waktu | 🔒 |
| 04 | Per-Client Limits | pelacakan terisolasi per identifier klien dengan `map[string]*ClientState` | 🔒 |
| 05 | Sliding Window Log | akurasi tinggi dengan log timestamp berbasis slice | 🔒 |
| 06 | Token Bucket | refill token kontinu dan toleransi burst trafik | 🔒 |
| 07 | Cleaning Up Idle Clients | pencegahan memory leak dengan membersihkan klien pasif | 🔒 |
| 08 | Concurrency Safety | sinkronisasi multi-goroutine dengan `sync.Mutex` (`-race`) | 🔒 |
| 09 | Wrap-Up | interface `RateLimiter` terpadu, simulasi trafik, perbandingan performa | 🔒 |

## Hasil Akhir
Rate limiter modular in-memory siap pakai yang aman terhadap konkurensi tinggi dan efisien dalam penggunaan memori.
