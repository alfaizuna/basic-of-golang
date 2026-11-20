## Contoh Jawaban Referensi

### 1. Kelemahan Rancangan A (Monolith)
- **Kerapuhan Tinggi**: Mengubah satu aturan validasi umur pasien dapat secara tidak sengaja merusak logika validasi jadwal dokter karena semua variabel bercampur dalam satu fungsi raksasa.
- **Sulit Diuji (Testing Nightmare)**: Untuk menguji apakah nomor HP salah, test harus menyiapkan data tiruan dokter, jadwal, dan koneksi database tiruan secara bersamaan.

### 2. Kelemahan Rancangan B (Over-Fragmentation)
- **Beban Navigasi (*Cognitive Load*)**: Membaca 5 file terpisah untuk hal sepele membuat pengembang lain frustrasi melompati deklarasi kode yang sebenarnya hanya baris sederhana.
- **Overhead Abstraksi Palsu**: Fungsi seperti `CheckStartsWithPlus()` tidak memiliki nilai bisnis mandiri di luar konteks validasi nomor telepon.

### 3. Solusi Jalan Tengah (*The Sweet Spot*)
Buat satu fungsi kohesif di package kontak/pasien:

```go
// CleanAndValidatePhone membersihkan karakter non-digit dan memastikan nomor telepon berformat internasional yang sah.
func CleanAndValidatePhone(raw string) (string, error)
```
- **Tanggung Jawab di Dalamnya**:
  1. Memotong spasi dan tanda minus.
  2. Memastikan format diawali kode negara (misal `+62`).
  3. Memastikan panjang digit berada dalam rentang valid (10-14 digit).
  4. Mengembalikan nomor bersih siap pakai atau error yang deskriptif.
- **Kenapa Seimbang**: Menyelesaikan satu konsep domain utuh (nomor telepon valid) dalam satu tempat yang mudah dibaca dan diuji.