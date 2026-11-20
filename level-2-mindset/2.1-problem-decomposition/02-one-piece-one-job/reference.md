## Contoh Jawaban Referensi

### 1. `AppointmentStore`
- **Tugas**: Mengelola penyimpanan data janji temu pasien dan dokter.
- **Alasan Berubah**: Struktur tabel janji temu berubah atau sistem migrasi database (misal dari SQLite ke PostgreSQL).

### 2. `ReminderFormatter`
- **Tugas**: Mengonversi data janji temu menjadi teks pesan pengingat siap kirim.
- **Alasan Berubah**: Tim humas/marketing klinik mengubah susunan kata, bahasa, atau format sapaan pesan WhatsApp.

### 3. `WhatsAppGateway`
- **Tugas**: Mengirimkan muatan pesan ke server penyedia API WhatsApp.
- **Alasan Berubah**: Vendor WhatsApp API mengubah endpoint, format otentikasi token, atau klinik berganti vendor gateway lain.

### 4. `ScheduleChecker`
- **Tugas**: Mengidentifikasi daftar janji temu yang jadwalnya jatuh tempo untuk diingatkan.
- **Alasan Berubah**: Kebijakan waktu pengingat klinik berubah (misal dari H-1 menjadi H-2 hari sebelum jadwal).