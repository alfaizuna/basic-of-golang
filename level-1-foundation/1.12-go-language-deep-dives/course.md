# Course 1.12: Go Language Deep Dives

## Tujuan
Menyelami fitur-fitur bahasa Go yang lebih mendalam dan nuansa teknisnya: fungsi variadik (`...T`) dan penyebaran slice (`s...`), pro dan kontra *named return values* serta *naked return*, mekanisme internal `defer` (urutan LIFO, evaluasi argumen di awal, dan manipulasi return value), mekanisme `panic` dan penangkapannya secara elegan menggunakan `recover`, teknik logging dan formatting debugging (`%+v` dan `%#v`), serta metodologi debugging praktis (membaca stack trace, `go vet`, isolasi bug).

## Prasyarat
Telah menyelesaikan Course 1.1 s.d. Course 1.11.

## Daftar Lesson

| # | Lesson | Fokus | Status |
|---|--------|-------|--------|
| 01 | Variadic Functions | `...T`, unpacking slice dengan `s...` | 🔒 |
| 02 | Named Return Values | deklarasi nilai kembalian bernama, naked return vs eksplisit | 🔒 |
| 03 | Defer Deep Dive | LIFO, snapshot argumen, modifikasi named return di defer | 🔒 |
| 04 | Panic | kapan menggunakan panic vs error, stack unwinding | 🔒 |
| 05 | Recover | menangkap panic di dalam deferred function, konversi ke error | 🔒 |
| 06 | Debugging with Prints & Logs | formatting `%+v`, `%#v`, log timestamps, Delve | 🔒 |
| 07 | Debugging Techniques | analisis stack trace, `go vet`, isolasi masalah dengan unit test | 🔒 |

## Hasil Akhir
Pemahaman komprehensif mengenai mekanisme internal bahasa Go dan kemampuan mendiagnosis serta memperbaiki bug kompleks di aplikasi nyata.
