## Contoh Jawaban Referensi

### 1. Identifikasi Masalah Utama
Tingginya angka ketidakhadiran pasien (*no-show rate*) tanpa pemberitahuan sebelumnya, yang mengakibatkan terbuangnya slot waktu produktif dokter dan hilangnya potensi pendapatan klinik.

### 2. 5 Pertanyaan Klarifikasi Kritis
1. **Saluran Komunikasi (Channel)**: Saluran apa yang paling efektif dan biasa digunakan oleh pasien klinik saat ini (WhatsApp, SMS, atau Email)?
2. **Jendela Waktu (Timing)**: Kapan waktu terbaik pengingat dikirim (misal: H-1 pada jam 10 pagi, atau 2 jam sebelum janji temu)?
3. **Mekanisme Konfirmasi**: Apakah pesan bersifat satu arah (informasi saja), atau pasien wajib membalas konfirmasi hadir/batal?
4. **Sumber Data Jadwal**: Dari mana sistem pengingat mengambil data pasien dan janji temu (apakah sudah ada database/Google Calendar/spreadsheet klinik, atau dibuatkan modul input janji baru)?
5. **Kebijakan Pasien Berhalangan**: Apa yang harus terjadi jika pasien menyatakan batal (apakah slot janji langsung dilepas untuk pasien antrean cadangan)?

### 3. 3 Batasan Ruang Lingkup V1 (Out of Scope)
1. **Tidak Ada Fitur Reschedule Mandiri**: Pasien yang ingin mengubah jam harus tetap menghubungi resepsionis secara manual.
2. **Hanya Satu Saluran Notifikasi**: V1 hanya mendukung WhatsApp gateway (tidak membangun integrasi SMS atau Email terlebih dahulu).
3. **Tidak Menangani Rekam Medis**: Sistem murni mengelola pengingat jadwal, tidak menyimpan riwayat diagnosa atau resep obat pasien.