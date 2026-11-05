# Course 1.3: Pointers & Memory

## Tujuan
Memahami bagaimana Go mengelola memori dan data melalui pointer: konsep pass-by-value, mengambil alamat memori (`&`) dan dereferencing (`*`), pointer receiver pada struct, nil safety, perbedaan stack vs heap, serta diakhiri dengan capstone membangun Feature-Flag Store in-memory.

## Prasyarat
Telah menyelesaikan Course 1.2: Composite Types.

## Daftar Lesson

| # | Chapter | Lesson | Fokus | Status |
|---|---------|--------|-------|--------|
| 01 | Addresses | Values Are Copies | Go meneruskan nilai dengan disalin; kenapa perubahan di function tidak terlihat si pemanggil | ✅ |
| 02 | Addresses | Taking Addresses: `&` and `*` | operator address-of dan dereference | ✅ |
| 03 | Addresses | Pointer Types & nil | tipe `*T`, zero value `nil`, tidak ada pointer arithmetic | ✅ |
| 04 | Addresses | `new` and `&T{}` | membuat pointer ke nilai baru dan ke struct literal | 🔓 |
| 05 | Mutating | Pointer Parameters | mengubah variabel pemanggil lewat parameter pointer | 🔒 |
| 06 | Mutating | Pointers to Structs | akses field lewat pointer (auto-dereference), mengubah struct di function | 🔒 |
| 07 | Mutating | Pointer Receivers | method yang mengubah receiver-nya | 🔒 |
| 08 | Mutating | Value vs Pointer Receivers | aturan memilih, konsistensi receiver dalam satu tipe | 🔒 |
| 09 | Mutating | Pointers, Slices & Maps | apa yang dibagi dan tidak dibagi; jebakan variabel loop dan `range` yang menyalin | 🔒 |
| 10 | Nil Safety | Nil Pointer Panics | apa yang terjadi saat dereference `nil` dan cara membaca pesan panic-nya | 🔒 |
| 11 | Nil Safety | Guarding with nil Checks | pola pengecekan nil, return `(*T, error)` dan cek error dulu | 🔒 |
| 12 | Nil Safety | Optional Values | pointer sebagai "belum diisi" (nil = unset), default konfigurasi | 🔒 |
| 13 | Memory | Returning Pointers to Locals | aman mengembalikan alamat variabel lokal di Go | 🔒 |
| 14 | Memory | Stack vs Heap | gambaran escape analysis, `go build -gcflags=-m` (pengenalan saja) | 🔒 |
| 15 | Capstone | Flag Store 1: Model | struct `Flag` dan `Store` (map nama → `*Flag`) | 🔒 |
| 16 | Capstone | Flag Store 2: Find or Create | fungsi yang mencari flag, atau membuatnya jika belum ada, lalu mengembalikan `*Flag` | 🔒 |
| 17 | Capstone | Flag Store 3: Update in Place | ubah flag lewat pointer yang dikembalikan, test untuk perilaku berbagi data | 🔒 |

## Hasil Akhir
Sebuah in-memory *Feature-Flag Store* yang menggunakan pointer untuk pencarian atau pembuatan flag (*find-or-create*) dan manipulasi langsung di tempat (*update-in-place*) tanpa menduplikasi alokasi memori.
