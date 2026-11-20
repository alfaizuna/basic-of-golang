### Kriteria Penilaian (Rubric)

1. **Ketajaman Definisi Input & Output**:
   - Input dan output berpengetikan logis yang jelas (bukan data samar seperti 'data apa saja').
2. **Prinsip Information Hiding**:
   - Mampu mengidentifikasi detail teknis yang berhasil disembunyikan (koneksi DB, URL vendor, format JSON payload, retry logic).
3. **Kerapian Batas Tanggung Jawab**:
   - Komponen tidak saling melangkahi wewenang (pemeriksa jadwal tidak ikut mengirim pesan).

**Status**:
- **Lulus**: Batas input/output jelas dan pemahaman information hiding tepat.
- **Perlu revisi**: Masih ada detail bocor (*leaky abstraction*) atau batas tugas tumpang-tindih.