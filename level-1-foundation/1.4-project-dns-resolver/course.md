# Course 1.4: Project: DNS Resolver

## Tujuan
Menerapkan pemahaman tentang pointer, struct, map, slice, nil safety, dan manipulasi data in-memory ke dalam proyek nyata yang terintegrasi: membangun sistem **In-Memory DNS Resolver**.
Topik meliputi pemodelan DNS record (A dan CNAME) menggunakan pointer, penyimpanan zone berstruktur map-of-slice-of-pointers, pencarian dan pembaruan in-place, caching dengan kedaluwarsa TTL (Time-To-Live), penelusuran rantai alias CNAME (*CNAME chaining*), deteksi siklus alias (*cycle detection*), serta replikasi zone (*zone transfer*) menggunakan konsep deep copy.

## Prasyarat
Telah menyelesaikan Course 1.1 (Go Basics), Course 1.2 (Composite Types), dan Course 1.3 (Pointers & Memory).

## Daftar Lesson

| # | Lesson | Fokus | Status |
|---|--------|-------|--------|
| 01 | Record Model | struct `Record` (Name, Type, Value, TTL) dan tipe record dasar (A, CNAME) | 🔒 |
| 02 | Zone Storage | `Zone` menyimpan record di `map[string][]*Record` | 🔒 |
| 03 | Add & Lookup | tambah record, cari record, kembalikan `*Record` atau `nil` jika tidak ada | 🔒 |
| 04 | Update in Place | ubah value/TTL record lewat pointer tanpa menyalin | 🔒 |
| 05 | Cache with TTL Expiry | cache hasil lookup, kedaluwarsa berdasarkan TTL dengan `now func() time.Time` | 🔒 |
| 06 | Following CNAME Chains | mengikuti rantai alias sampai ketemu record A | 🔒 |
| 07 | Detecting CNAME Cycles | deteksi siklus (A → B → A) dengan set `map[string]bool` dan batas kedalaman | 🔒 |
| 08 | Zone Transfers | menyalin zone ke zone lain: deep copy agar tidak berbagi pointer | 🔒 |

## Hasil Akhir
Sebuah engine *DNS Resolver* in-memory lengkap dengan manajemen zone, caching ber-TTL dinamis yang dapat diuji dengan clock mock, resolusi rekursif rantai CNAME dengan proteksi anti-loop/siklus tak berujung, serta mekanisme transfer zone yang aman dengan deep copy memori.
