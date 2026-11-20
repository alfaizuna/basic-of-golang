# Course 1.6: Interfaces & Polymorphism

## Tujuan
Memahami konsep antarmuka (*interface*) di Go untuk mendefinisikan perilaku (*behavior*), polimorfisme melalui implementasi implisit (*duck typing* ala Go), pemanfaatan empty interface (`any` / `interface{}`), pemeriksaan tipe dinamis (*type assertion* dan *type switch*), serta menyusun interface yang fleksibel melalui penggabungan (*interface composition*).

## Prasyarat
Telah menyelesaikan Course 1.1 s.d. Course 1.5.

## Daftar Lesson

| # | Lesson | Fokus | Status |
|---|--------|-------|--------|
| 01 | Defining Behavior with Interfaces | deklarasi `interface`, memakai interface sebagai parameter function | 🔒 |
| 02 | Implicit Implementation | tipe memenuhi interface tanpa kata `implements`, beberapa tipe satu interface (polimorfisme) | 🔒 |
| 03 | The Empty Interface | `any` / `interface{}`, kapan boleh dan kapan dihindari | 🔒 |
| 04 | Type Assertions & Type Switches | `v, ok := x.(T)`, `switch x.(type)` | 🔒 |
| 05 | Composing Interfaces | menggabungkan interface kecil jadi interface lebih besar (gaya `io.ReadWriter`) | 🔒 |

## Hasil Akhir
Pemahaman mendalam tentang abstraksi berbasis perilaku di Go, kemampuan menulis kode yang decoupled dan mudah diuji (*testable*), serta kemampuan memanipulasi data dinamis dengan aman.
