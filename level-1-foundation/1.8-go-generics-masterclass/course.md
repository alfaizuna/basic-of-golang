# Course 1.8: Go Generics Masterclass

## Tujuan
Menguasai Go Generics (diperkenalkan sejak Go 1.18): memahami alasan di balik generics dan kelemahan interface kosong (`any`), deklarasi fungsi generik dengan type parameter `[T any]` dan *type inference*, penggunaan type constraint bawaan seperti `comparable` dan pembuatan constraint kustom, union constraint (`|`) dan tilde approximation (`~T`), pembuatan tipe data generik (seperti `Stack[T]` atau `Pair[T, U]`), pendefinisian method pada tipe generik, pola *functional helpers* (`Map`, `Filter`, `Reduce`), constraint berbasis method, serta panduan praktis kapan harus dan kapan tidak perlu menggunakan generics.

## Prasyarat
Telah menyelesaikan Course 1.1 s.d. Course 1.7.

## Daftar Lesson

| # | Lesson | Fokus | Status |
|---|--------|-------|--------|
| 01 | Why Generics | masalah duplikasi kode, kenapa `any` saja tidak cukup | 🔒 |
| 02 | Generic Functions | type parameter `[T any]`, type inference | 🔒 |
| 03 | Type Constraints | `any`, `comparable`, membuat constraint sendiri | 🔒 |
| 04 | Union Constraints & `~` | `int \| float64`, tilde, tipe kustom | 🔒 |
| 05 | Generic Types | `type Pair[T, U any] struct`, inisialisasi tipe generik | 🔒 |
| 06 | Methods on Generic Types | method pada tipe generik, receiver `*Stack[T]` | 🔒 |
| 07 | Generic Helpers | `Map` dan `Filter`, memanfaatkan type safety | 🔒 |
| 08 | Constraints with Methods | constraint berupa interface dengan method (`Area()`) | 🔒 |
| 09 | Best Practices | kapan **tidak** memakai generics, keterbacaan vs fleksibilitas | 🔒 |

## Hasil Akhir
Kemampuan menulis kode Go yang reusable, efisien tanpa boxing overhead atau runtime assertions, dan tetap mempertahankan keamanan tipe statis (*static type-safety*).
