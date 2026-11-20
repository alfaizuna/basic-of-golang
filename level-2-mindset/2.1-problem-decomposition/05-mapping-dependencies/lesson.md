# Mapping Dependencies

## Konsep
Arsitektur perangkat lunak ditentukan oleh **arah ketergantungan (dependencies)** antar bagian.
Jika Komponen A memanggil fungsi di Komponen B, maka **A bergantung pada B**.
Aturan emas:
1. Komponen kebijakan bisnis tingkat tinggi tidak boleh bergantung pada detail infrastruktur tingkat rendah.
2. Alur dependensi harus mengalir satu arah tanpa loop.

## Skenario
Klinik Sehat memiliki 4 komponen yang telah kita rancang sebelumnya:
1. `ReminderJob` (komponen orkestrator yang dipicu setiap jam 8 pagi).
2. `AppointmentStore` (mengambil data janji temu dari database).
3. `TemplateEngine` (mengisi template teks dengan nama pasien & dokter).
4. `WhatsAppClient` (mengirim HTTP request ke vendor WA).

Kita perlu memetakan bagaimana keempat komponen ini saling memanggil.

## Tugas
Di dalam file `answer.md`:
1. Buat **Diagram Alur Dependensi** menggunakan sintaks kode **Mermaid** (`graph TD` atau `flowchart TD`).
2. Tuliskan penjelasan teks singkat:
   - Siapa yang menjadi orkestrator utama?
   - Siapa yang tidak bergantung pada komponen lain (*leaf nodes*)?
3. Jelaskan apa yang terjadi jika format teks pengingat di `TemplateEngine` diubah: komponen mana saja yang **TIDAK** terpengaruh sama sekali?

## Hints (0/3)
### Hint 1
Gunakan format Mermaid: ```mermaid ... ```

### Hint 2
Orkestrator (ReminderJob) biasanya berada di puncak dan memanggil modul pendukung.

### Hint 3
Komponen leaf node adalah komponen yang tidak memiliki panah keluar ke komponen lain (mandiri).


## Key Takeaway
- Petakan dependensi secara visual dengan diagram alur satu arah.
- Minimalkan dependensi horizontal antar komponen sejajar agar perubahan di satu bagian tidak merembet ke bagian lain.
