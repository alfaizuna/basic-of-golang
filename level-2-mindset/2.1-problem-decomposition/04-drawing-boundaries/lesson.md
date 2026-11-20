# Drawing Boundaries

## Konsep
Sebuah komponen yang baik seperti kotak hitam (*black box*): pemanggil hanya perlu mengetahui permukaannya (input dan output), tanpa tahu detail implementasi rumit di dalamnya (*information hiding*).
Mendefinisikan batas komponen berarti menyepakati:
1. **Input**: Data apa yang mutlak dibutuhkan.
2. **Output**: Data atau hasil apa yang dijanjikan keluar.
3. **Tanggung Jawab**: Apa yang ia selesaikan.
4. **Hal yang Disembunyikan**: Detail teknis internal yang tidak boleh dibocorkan ke luar.

## Skenario
Kita akan memformalkan batas untuk komponen **`ReminderDispatcher`** pada sistem Klinik Sehat.
Komponen ini bertugas mengambil janji temu yang harus diingatkan hari ini, menyiapkan pesan yang ramah, dan memastikannya terkirim ke WhatsApp pasien tanpa membuat komponen lain repot memikirkan format payload JSON vendor WhatsApp.

## Tugas
Di dalam file `answer.md`, tuliskan kontrak batas (*component contract*) untuk 2 komponen berikut:
1. **`ScheduleInspector`** (Pemeriksa Jadwal yang perlu diingatkan)
2. **`MessageTransmitter`** (Pengirim Pesan ke Vendor Eksternal)

Untuk masing-masing komponen, definisikan secara eksplisit:
- **Tanggung Jawab Inti**: Apa tugas tunggalnya?
- **Input yang Diterima**: Data spesifik apa yang masuk?
- **Output yang Dihasilkan**: Apa nilai kembalian dan status yang keluar?
- **Detail yang Disembunyikan (Hidden Secrets)**: Hal teknis apa di dalamnya yang tidak boleh diketahui oleh komponen lain?

## Hints (0/3)
### Hint 1
Pikirkan: apakah pemanggil `MessageTransmitter` perlu tahu URL API vendor WhatsApp atau API token-nya? Tentu tidak!

### Hint 2
Pikirkan: apakah pemanggil `ScheduleInspector` perlu tahu sintaks SQL query atau nama tabel database? Tentu tidak!

### Hint 3
Gunakan format spesifikasi yang terstruktur dan ringkas.


## Key Takeaway
- Batas yang tegas mencegah kebocoran abstraksi (*leaky abstractions*).
- Selalu sembunyikan detail implementasi (SQL, HTTP payload, API keys) di balik kontrak antarmuka yang bersih.
