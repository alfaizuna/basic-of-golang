## Contoh Jawaban Referensi

### 1. Urutan Pengerjaan
1. **Langkah 1: `AppointmentModel`** (Struktur data dasar tanpa dependensi).
2. **Langkah 2: `WhatsAppGatewayClient`** (Mitigasi risiko integrasi vendor eksternal).
3. **Langkah 3: `MessageTemplateEngine`** (Format teks murni dan pengujian template string).
4. **Langkah 4: `AppointmentDatabaseStore`** (Penyimpanan dan kueri janji jatuh tempo).
5. **Langkah 5: `ReminderSchedulerJob`** (Integrasi akhir dan orkestrasi seluruh alur).

### 2. Justifikasi Teknis
- **`AppointmentModel` harus pertama**: Seluruh modul lain membutuhkan definisi tipe struct `Appointment` dan `Patient` sebagai bahasa komunikasi data bersama.
- **`WhatsAppGatewayClient` di awal (De-Risking)**: Integrasi pihak ketiga adalah bagian paling rawan gagal (isu API key, format token, pembatasan nomor telepon). Mengujinya lebih awal memberi waktu mitigasi jika vendor bermasalah.
- **`ReminderSchedulerJob` terakhir**: Sebagai orkestrator puncak, job ini tidak bisa berjalan tanpa adanya store untuk membaca data dan client untuk mengirim pesan.

### 3. Milestone MVP Hari ke-3
Di hari ke-3, tim wajib mendemonstrasikan **"Walking Skeleton"**:
Sebuah skrip sederhana yang mengambil satu data janji temu dummy, memformat pesannya, dan berhasil mengirimkannya ke nomor WhatsApp dokter/penguji. Dengan lolosnya milestone ini, risiko terbesar sistem telah tereliminasi.