# Course 1.13: Capstone: Web Crawler

## Tujuan
Membangun proyek Capstone akhir Level 1: sebuah **Concurrent Web Crawler** dari nol yang menggabungkan seluruh keterampilan yang telah dipelajari: goroutine, channel semaphore, `sync.Mutex`, `sync.WaitGroup`, struct, interface polimorfik, parsing HTML dengan pustaka `golang.org/x/net/html`, resolusi URL dengan `net/url`, pembatasan kedalaman crawl (*depth limit*), pembatasan domain (*same-domain filter*), pengumpulan hasil dan pelaporan error, serta pengujian end-to-end dengan `net/http/httptest`.

## Prasyarat
Telah menyelesaikan Course 1.1 s.d. Course 1.12.

## Struktur Project
Proyek capstone ini bertahap mulai dari pemodelan interface pengambil (*fetcher*), parsing tag HTML, pencegahan siklus dengan visited set, konkurensi terkontrol (*bounded concurrency*), hingga evaluasi hasil akhir.

## Daftar Lesson

| # | Lesson | Fokus | Status |
|---|--------|-------|--------|
| 01 | Fetching a Page | interface `Fetcher`, HTTP fetcher dengan timeout, `FakeFetcher` | 🔒 |
| 02 | Extracting Links | parse HTML (`golang.org/x/net/html`), resolusi URL absolut | 🔒 |
| 03 | Sequential Crawl | algoritma crawl rekursif, visited set (`map[string]bool`) | 🔒 |
| 04 | Depth Limit & Same-Domain | batas kedalaman tree crawl, filter isolasi hostname | 🔒 |
| 05 | Concurrent Crawling | paralelisme goroutine + `sync.WaitGroup` | 🔒 |
| 06 | Safe Shared State | thread-safe visited tracker (`sync.Mutex`), rate limiter token/semaphore | 🔒 |
| 07 | Results & Error Handling | agregasi metrik, penanganan URL rusak, end-to-end test | 🔒 |

## Hasil Akhir
Sistem crawler web konkuren yang tangguh, cepat, hemat memori, dan mampu menjelajah situs web tanpa mengalami deadlock, race condition, atau infinite loop.
