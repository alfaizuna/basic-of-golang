# Too Big, Too Small: Finding the Sweet Spot

## Konsep
Dekomposisi adalah soal keseimbangan (*balance*):
- **Terlalu Besar (Monolith Tangle)**: Satu file atau fungsi ribuan baris yang menangani semua hal sekaligus. Sulit dibaca, rentan konflik git, dan mustahil diuji.
- **Terlalu Kecil (Over-Fragmentation / Nano-Components)**: Memecah setiap 2 baris kode menjadi modul atau fungsi terpisah. Menghasilkan puluhan file kecil yang hanya meneruskan panggilan (*pass-through*), membuat pembaca pusing melompat dari satu file ke file lain.
Komponen yang baik memiliki kohesi tinggi (*high cohesion*) dan keterikatan rendah (*loose coupling*).

## Skenario
Dua arsitek software sedang berdebat sengit merancang modul validasi nomor telepon pasien untuk SMS/WhatsApp Klinik Sehat:

- **Arsitek A (Gaya Monolith)**: Ingin menyatukan validasi nomor HP, validasi jadwal dokter, validasi usia pasien, dan koneksi database dalam satu fungsi raksasa `ValidateEverything(data any) error` sepanjang 400 baris.
- **Arsitek B (Gaya Ekstrem Mikro)**: Memecah validasi nomor HP menjadi 5 fungsi di 5 file terpisah:
  1. `CheckStringNotEmpty()`
  2. `CheckStartsWithPlus()`
  3. `StripHyphens()`
  4. `CountDigits()`
  5. `EnsureMinTenDigits()`
  Masing-masing fungsi hanya terdiri dari 1-2 baris kode.

## Tugas
Di dalam file `answer.md`:
1. **Analisis Rancangan A**: Sebutkan 2 kelemahan fatal dari pendekatan Arsitek A.
2. **Analisis Rancangan B**: Sebutkan 2 biaya tersembunyi (*hidden overhead*) dari pendekatan Arsitek B.
3. **Rancangan Jalan Tengah (*Sweet Spot*)**: Tuliskan rancangan fungsi validasi nomor telepon pasien yang seimbang menurut Anda (sebutkan nama fungsi, parameter input, return value, dan apa yang ia periksa di dalamnya).

## Hints (0/3)
### Hint 1
Pikirkan tentang keterbacaan: apakah membaca 5 file untuk memvalidasi nomor telepon mempermudah atau justru melelahkan?

### Hint 2
Pikirkan tentang maintainability: fungsi yang ideal biasanya menyelesaikan satu konsep domain utuh tanpa fragmentasi berlebihan.

### Hint 3
Validasi nomor telepon biasanya cukup satu fungsi terfokus seperti `NormalizeAndValidatePhone(raw string) (string, error)`.


## Key Takeaway
- Hindari monolith raksasa, tetapi jangan jatuh ke perangkap over-engineering komponen mikro.
- Granularitas ideal adalah komponen yang menyelesaikan satu konsep domain utuh secara tuntas.
