### Kriteria Penilaian (Rubric)

1. **Kepatuhan Aturan 'Tanpa Dan'**:
   - Seluruh deskripsi tugas komponen benar-benar bersih dari kata sambung koordinatif ("dan", "serta", "juga").
2. **Pemisahan Perhatian (Separation of Concerns)**:
   - Membedakan entitas domain/jadwal, formatting pesan, dan klien gateway pihak ketiga.
3. **Ketepatan Reason to Change**:
   - Alasan perubahan spesifik sesuai domain masing-masing (misal: gateway berubah jika vendor WA ganti tarif/API; formatter berubah jika teks promosi diubah bagian marketing).

**Status**:
- **Lulus**: Komponen terbagi bersih, lulus tes satu kalimat, dan alasan perubahan valid.
- **Perlu revisi**: Masih ada komponen raksasa atau deskripsi yang menggabungkan beberapa aksi berbeda.