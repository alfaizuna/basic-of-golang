# Course 1.11: Real-World Patterns

## Tujuan
Menguasai pola-pola praktis yang dipakai setiap hari oleh software engineer Go di dunia industri: operasi File I/O (`os.ReadFile`, `os.Create`, streaming dengan `bufio.Scanner`), serialisasi dan deserialisasi JSON (`encoding/json` dan struct tags), komunikasi jaringan HTTP Client (GET, POST JSON, timeouts), perancangan HTTP Server dan REST API standar (`net/http`, `ServeMux`, query params), serta teknik pengujian profesional menggunakan table-driven tests dan `net/http/httptest`.

## Prasyarat
Telah menyelesaikan Course 1.1 s.d. Course 1.10.

## Daftar Lesson

| # | Lesson | Fokus | Status |
|---|--------|-------|--------|
| 01 | Reading Files | `os.ReadFile`, penanganan error file tidak ada | 🔒 |
| 02 | Writing Files | `os.WriteFile`, `os.Create`, `defer f.Close()` | 🔒 |
| 03 | Line-by-Line with bufio | `bufio.Scanner`, membaca stream baris efisien | 🔒 |
| 04 | JSON Encoding | `json.Marshal`, struct tags, omitempty | 🔒 |
| 05 | JSON Decoding | `json.Unmarshal`, parsing JSON ke struct | 🔒 |
| 06 | HTTP Client: GET | `http.Get`, status code, `defer resp.Body.Close()` | 🔒 |
| 07 | HTTP Client: POST & Timeouts | POST JSON, kustom `http.Client` dengan timeout | 🔒 |
| 08 | HTTP Server: Handlers | `http.HandleFunc`, `ResponseWriter`, `*http.Request` | 🔒 |
| 09 | Routing & Request Data | `http.NewServeMux`, query params, verifikasi method | 🔒 |
| 10 | JSON APIs | server REST yang menerima dan membalas JSON | 🔒 |
| 11 | Your First Tests | framework `testing`, `t.Errorf`, `t.Fatalf` | 🔒 |
| 12 | Table-Driven Tests & httptest | pola table-driven, `httptest.NewRecorder` | 🔒 |

## Hasil Akhir
Kemampuan membangun aplikasi web backend dan RESTful API di Go menggunakan standard library tanpa memerlukan external web framework.
