# One Piece, One Job

## Konsep
Prinsip Tanggung Jawab Tunggal (*Single Responsibility Principle*): Setiap komponen, package, atau modul harus memiliki satu alasan untuk berubah.
Cara paling praktis mengujinya adalah **Tes Satu Kalimat Tanpa Kata Sambung 'Dan'**:
Jika Anda tidak bisa menjelaskan tugas suatu komponen dalam satu kalimat sederhana tanpa menggunakan kata "dan", "serta", atau "juga", berarti komponen tersebut menampung terlalu banyak pekerjaan dan harus dipecah.

## Skenario
Dalam sesi perancangan awal untuk Layanan Pengingat Janji Klinik Sehat, seorang programmer junior mengusulkan sebuah modul utama bernama `AppointmentManager`:

> *"Modul `AppointmentManager` ini bertugas menyimpan data janji temu ke database, mengecek kalender dokter yang kosong, memformat template pesan WhatsApp pengingat, melakukan HTTP POST request ke vendor WhatsApp API pihak ketiga, dan mengupdate status pesan di database jika pengiriman gagal."*

Perhatikan betapa banyak kata "dan" serta variasi tanggung jawab yang tercampur di dalam modul tunggal ini.

## Tugas
Pecah modul `AppointmentManager` yang serbabisa di atas menjadi beberapa komponen mandiri.
Di dalam file `answer.md`:
1. Buat daftar komponen hasil pecahan (minimal 3 komponen terpisah).
2. Tuliskan deskripsi tugas masing-masing komponen dalam **satu kalimat tegas tanpa menggunakan kata sambung 'dan', 'serta', atau 'juga'**.
3. Jelaskan untuk setiap komponen: **satu alasan spesifik** apa yang membuat komponen tersebut harus diubah di masa depan (*reason to change*).

## Hints (0/3)
### Hint 1
Pisahkan urusan penyimpanan data (storage/database) dari urusan komunikasi eksternal (WhatsApp gateway).

### Hint 2
Pisahkan logika pembuatan teks (formatting template) dari logika pengiriman jaringan (HTTP networking).

### Hint 3
Uji kalimat deskripsi Anda: jika ada kata 'dan', pecah lagi!


## Key Takeaway
- Komponen yang melakukan satu tugas mudah dipahami, mudah diuji secara terisolasi, dan aman diubah tanpa merusak bagian lain.
- Gunakan tes kalimat tanpa kata 'dan' sebagai alarm pendeteksi komponen yang kegemukan.
