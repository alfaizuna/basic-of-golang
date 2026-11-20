# Course 1.5: Project: YAML Data Pipeline

## Tujuan
Membangun sebuah data pipeline in-memory yang lengkap di Go untuk memproses dokumen berformat YAML: mem-parse dokumen tunggal dan multi-dokumen (menggunakan pemisah `---`), melakukan pembersihan dan transformasi data baris demi baris, memfilter data, melakukan kalkulasi agregasi (*count*, *sum*, *average*, dan *group-by*), hingga merangkai beberapa fungsi tahapan pipeline (*pipeline staging*) menggunakan fungsi tingkat tinggi (*first-class functions*).

## Prasyarat
Telah menyelesaikan Course 1.1 (Go Basics), Course 1.2 (Composite Types), Course 1.3 (Pointers & Memory), dan Course 1.4 (Project: DNS Resolver).

## Library Eksternal
Proyek ini menggunakan library pihak ketiga standar industri untuk YAML: `gopkg.in/yaml.v3`. Library ini dicatat di dalam `go.mod` dan `go.sum` proyekmu.

## Daftar Lesson

| # | Lesson | Fokus | Status |
|---|--------|-------|--------|
| 01 | Parsing One Document | decode YAML ke struct (`Row`/`Dataset`), tag `yaml:"..."` | 🔒 |
| 02 | Multi-Document YAML | pemisah dokumen `---`, decode berulang dengan `yaml.NewDecoder` sampai `io.EOF` | 🔒 |
| 03 | Transforming Rows | memodifikasi/menghitung ulang field pada slice of struct (`[]Row`) | 🔒 |
| 04 | Filtering & Selecting | menyaring baris berdasarkan predikat dan memilih data tertentu | 🔒 |
| 05 | Aggregations | menghitung `count`, `sum`, `avg`, dan pengelompokan (*group-by*) menggunakan map | 🔒 |
| 06 | Chaining Stages | merancang tipe `Stage func([]Row) []Row`, merangkai tahapan pipeline, dan test end-to-end | 🔒 |

## Hasil Akhir
Sebuah engine *YAML Data Processing Pipeline* yang modular, fleksibel, dan teruji secara otomatis yang mampu mengimpor data YAML mentah, mentransformasikannya melalui rantai fungsi (*pipeline execution*), dan menghasilkan laporan metrik agregasi ringkas.
