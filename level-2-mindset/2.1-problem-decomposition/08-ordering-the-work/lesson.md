# Ordering the Work: Building Bottom-Up

## Konsep
Setelah sistem dipecah menjadi bagian-bagian terdefinisi, pertanyaan berikutnya: **mulai bangun dari mana?**
Urutan implementasi terbaik mengikuti prinsip **Membangun dari Bawah ke Atas (*Bottom-Up from Leaves*)**:
1. Mulai dari komponen yang **paling sedikit memiliki dependensi** (komponen daun / model data).
2. Selesaikan komponen yang memiliki **risiko teknologi tertinggi** terlebih dahulu (*de-risking early*).
3. Terakhir, bangun orkestrator tingkat tinggi dan antarmuka luar (CLI / HTTP API).
Dengan cara ini, setiap bagian baru dibangun di atas fondasi yang sudah stabil dan teruji.

## Skenario
Tim Anda memiliki waktu 1 minggu untuk menyelesaikan versi pertama (MVP) Layanan Pengingat Klinik Sehat.
Daftar komponen yang sudah dirancang:
- A. `WhatsAppGatewayClient` (integrasi API vendor pesan).
- B. `AppointmentModel` (definisi tipe data struct & validasi).
- C. `ReminderSchedulerJob` (background loop pengingat harian).
- D. `AppointmentDatabaseStore` (akses SQL / penyimpanan).
- E. `MessageTemplateEngine` (penyusun teks string pesan).

## Tugas
Di dalam file `answer.md`:
1. **Urutan Pengerjaan**: Urutkan kelima komponen di atas (dari urutan 1 yang paling pertama dikerjakan, sampai urutan 5 yang terakhir).
2. **Justifikasi Teknis**: Berikan alasan singkat mengapa Anda memilih urutan tersebut (kenapa komponen X harus mendahului Y?).
3. **Definisi Milestone MVP Hari ke-3**: Apa yang harus sudah bisa didemonstrasikan pada pertengahan minggu untuk membuktikan bahwa integrasi inti berhasil?

## Hints (0/3)
### Hint 1
Bisakah Anda membuat ReminderSchedulerJob sebelum model data dan WhatsAppClient selesai? Tidak bisa!

### Hint 2
Komponen mana yang tidak bergantung pada siapapun? Kerjakan itu terlebih dahulu.

### Hint 3
Untuk de-risking: apakah vendor WhatsApp API bekerja dengan lancar? Sebaiknya diverifikasi di awal.


## Key Takeaway
- Bangun sistem dari komponen daun yang independen menuju ke orkestrator puncak.
- Uji dan selesaikan risiko integrasi eksternal paling awal untuk menghindari kejutan fatal di akhir proyek.
