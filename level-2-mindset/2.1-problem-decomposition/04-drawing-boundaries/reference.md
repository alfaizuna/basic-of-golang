## Contoh Jawaban Referensi

### Komponen 1: `ScheduleInspector`
- **Tanggung Jawab Inti**: Mencari dan mengumpulkan semua janji temu pasien yang jadwalnya jatuh tempo untuk dikirimi pengingat pada tanggal target.
- **Input**:
  - `targetDate` (tanggal evaluasi)
  - `statusFilter` (misal: hanya yang berstatus 'Pending')
- **Output**:
  - Daftar objek `[]Appointment` yang memenuhi syarat.
  - Nilai `error` jika akses penyimpanan gagal.
- **Detail yang Disembunyikan (Secrets)**:
  - Jenis database (PostgreSQL, MySQL, atau in-memory map).
  - Kueri SQL, indexing tabel, dan struktur skema kolom database.

---

### Komponen 2: `MessageTransmitter`
- **Tanggung Jawab Inti**: Mengantarkan teks pesan pengingat ke nomor telepon penerima melalui jaringan komunikasi eksternal.
- **Input**:
  - `recipientPhone` (string nomor telepon berformat E.164, misal `+62812...`)
  - `messageText` (string pesan lengkap siap baca)
- **Output**:
  - `receiptID` (identitas pengiriman unik dari vendor)
  - Nilai `error` jika pengiriman gagal lewat batas retry.
- **Detail yang Disembunyikan (Secrets)**:
  - Alamat URL API vendor gateway dan HTTP headers.
  - API Key / Secret Token otentikasi pihak ketiga.
  - Format serialisasi JSON spesifik dari vendor.
  - Mekanisme timeout dan pengulangan kirim ulang (*exponential backoff retry*).