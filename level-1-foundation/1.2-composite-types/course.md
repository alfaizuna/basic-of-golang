# Course 1.2: Composite Types

## Tujuan
Menguasai tipe data komposit di Go: Array, Slice, Map, Struct, dan Method. Kamu akan belajar bagaimana mengelola kumpulan data secara efisien, memahami alokasi memori internal slice (len vs cap, backing array), bekerja dengan key-value store (map), serta mendefinisikan tipe data kustom dan perilakunya (struct & method). Course ini diakhiri dengan proyek capstone membuat *Mini Inventory System*.

## Prasyarat
Telah menyelesaikan Course 1.1: Go Basics.

## Daftar Lesson

| # | Chapter | Lesson | Fokus | Status |
|---|---------|--------|-------|--------|
| 01 | Slices & Arrays | Arrays: Fixed Size | array literal, panjang adalah bagian dari tipe, array disalin saat di-assign | ✅ |
| 02 | Slices & Arrays | Slice Basics | slice literal, indexing, `len`, `for range` pada slice | ✅ |
| 03 | Slices & Arrays | len vs cap | beda panjang dan kapasitas, cara membacanya | ✅ |
| 04 | Slices & Arrays | Creating Slices with make | `make([]T, len, cap)`, kapan pakai `make` vs `var` | 🔓 sedang dikerjakan |
| 05 | Slices & Arrays | Growing with append | `append`, pertumbuhan kapasitas, kenapa hasilnya harus ditampung ulang | 🔒 |
| 06 | Slices & Arrays | Slicing & Shared Backing Arrays | `s[a:b]`, dua slice yang berbagi array yang sama dan efek sampingnya | 🔒 |
| 07 | Slices & Arrays | copy | `copy`, membuat salinan independen | 🔒 |
| 08 | Slices & Arrays | Nil Slices & Slice Recipes | nil vs empty slice, hapus/sisip elemen, filter | 🔒 |
| 09 | Maps | Map Basics | map literal, `make`, baca/tulis/`delete`, `len` | 🔒 |
| 10 | Maps | Presence Checks | pola `v, ok := m[k]`, beda "tidak ada" vs zero value | 🔒 |
| 11 | Maps | Counting & Grouping | menghitung frekuensi, mengelompokkan data | 🔒 |
| 12 | Maps | Safe Iteration | `range` pada map, urutan acak, iterasi berurutan lewat key yang diurutkan | 🔒 |
| 13 | Maps | Nil Maps & Maps of Slices | menulis ke nil map = panic, `map[string][]T` | 🔒 |
| 14 | Structs & Methods | Struct Basics | deklarasi `type ... struct`, struct literal, akses field | 🔒 |
| 15 | Structs & Methods | Nested Structs | struct di dalam struct, slice/map sebagai field | 🔒 |
| 16 | Structs & Methods | Embedding | struct embedding dan field/method yang "naik" | 🔒 |
| 17 | Structs & Methods | Constructor Functions | fungsi `NewXxx`, nilai default, validasi | 🔒 |
| 18 | Structs & Methods | Methods: Value vs Pointer Receivers | method pada struct, kapan memakai receiver pointer | 🔒 |
| 19 | Structs & Methods | Capstone: Mini Inventory | gabungkan slice, map, struct, dan method | 🔒 |

## Hasil Akhir
Sebuah aplikasi CLI *Mini Inventory Management* yang mengintegrasikan Slice, Map, Struct, dan Method untuk menambah, memperbarui, mengelompokkan, dan menampilkan katalog produk secara terstruktur.
