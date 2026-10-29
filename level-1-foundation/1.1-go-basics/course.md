# Course 1.1: Go Basics

## Tujuan
Memahami fondasi bahasa Go: struktur program, variabel, tipe data dasar, kontrol alur logika (if/else, switch), perulangan (for loop), pembuatan fungsi, error handling dasar, dan representasi teks (bytes & runes). Di akhir course, kamu akan mengasah kemampuan ini lewat proyek berkelanjutan: membuat *slug generator*.

## Prasyarat
Tidak ada. Course ini dirancang untuk pemula dari nol.

## Daftar Lesson

| # | Lesson | Fokus | Status |
|---|--------|-------|--------|
| 01 | Your First Line of Go | `package main`, `import`, `func main`, `fmt.Println` | ✅ |
| 02 | The Building Blocks | tipe dasar & literal: int, float64, string, bool | ✅ |
| 03 | Remembering Things | variabel: `var`, `:=`, assignment | ✅ |
| 04 | Constants and Zero Values | `const`, `iota`, zero value tiap tipe | ✅ |
| 05 | Math That Makes Sense | operator aritmatika, presedensi, pembagian integer, konversi tipe angka | ✅ |
| 06 | Making Decisions | `if` / `else if` / `else`, aturan kurung kurawal, `if` dengan initializer | ✅ |
| 07 | Logical Operators | `&&`, `\|\|`, `!`, short-circuit | ✅ |
| 08 | Build Your Own Functions | deklarasi function, parameter, return value | ✅ |
| 09 | Multiple Return Values | return lebih dari satu nilai, blank identifier `_` | ✅ |
| 10 | The Clean Switch | `switch` (tanpa fallthrough otomatis), switch tanpa ekspresi | ✅ |
| 11 | Repeating Things | `for` (tiga bentuk), `break` / `continue` | ✅ |
| 12 | Loop Variations | `for range`, loop bersarang, label | ✅ |
| 13 | Errors Are Values | tipe `error`, `errors.New`, pola `if err != nil` | ✅ |
| 14 | Bytes, Runes & Sized Types | `byte`, `rune`, int8/16/32/64, overflow | ✅ |
| 15 | Strings Under the Hood | string = byte immutable, UTF-8, iterasi string, `strings` package | ✅ |
| 16 | Format Like a Pro | `fmt.Printf` / `Sprintf`, verb (`%v`, `%d`, `%s`, `%q`, `%T`), lebar & presisi | ✅ |
| 17 | Slug Generator 1: Detect | project: deteksi karakter valid/tidak valid pada teks | ✅ |
| 18 | Slug Generator 2: Clean | project: bersihkan teks (huruf kecil, buang simbol, rapikan spasi) | ✅ |
| 19 | Slug Generator 3: Transform | project: ubah teks menjadi slug (pemisah `-`, kasus Unicode/aksen) | ✅ |
| 20 | Slug Generator 4: Ship It | project: gabungkan semua tahap, error handling, test akhir | ✅ |

## Hasil Akhir
Sebuah aplikasi *slug generator* CLI yang siap pakai, dapat mengubah sembarang teks/judul artikel menjadi slug URL yang bersih dan aman (misalnya: `" Belajar Go: Dasar #1! "` → `"belajar-go-dasar-1"`).
