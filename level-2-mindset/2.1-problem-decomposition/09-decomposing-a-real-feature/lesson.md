# Decomposing a Real Feature: Courier Logistics

## Konsep
Penerapan menyeluruh seluruh tahapan dekomposisi pada skenario bisnis baru:
1. Membedah narasi bisnis yang ambigu.
2. Mengekstrak entitas data (kata benda) dan operasi (kata kerja).
3. Menetapkan batasan komponen independen (Single Responsibility).
4. Menyusun diagram dependensi satu arah.
5. Menentukan urutan kerja bertahap.

## Skenario
Perusahaan startup logistik **"KirimKilat"** ingin membangun modul baru: **Sistem Penugasan Kurir Instan (On-Demand Courier Dispatcher)**.

Berikut deskripsi dari Product Manager:
> *"Saat pelanggan memesan penjemputan paket, sistem harus mencari kurir aktif terdekat dalam radius maksimal 5 km dari lokasi jemput. Sistem kemudian mengirimkan tawaran order ke aplikasi kurir tersebut. Kurir memiliki waktu 30 detik untuk menerima tawaran. Jika kurir menolak atau waktu habis, sistem otomatis mencari kurir terdekat berikutnya. Jika kurir menerima, sistem mengunci order untuk kurir tersebut, memperbarui status paket menjadi 'Menunggu Penjemputan', dan mengirimkan tautan pelacakan langsung (live tracking) ke pelanggan via SMS."*

## Tugas
Lakukan dekomposisi penuh untuk fitur KirimKilat di dalam `answer.md`:
1. **Identifikasi Komponen**: Tentukan minimal 4 komponen terpisah beserta peran tunggalnya (patuh tes satu kalimat tanpa 'dan').
2. **Kontrak Input/Output**: Tuliskan input dan output untuk modul pencari kurir (`CourierFinder`) dan modul penawaran (`OfferManager`).
3. **Diagram Dependensi**: Buat diagram Mermaid alur ketergantungan komponen.
4. **Urutan Pengerjaan**: Tuliskan urutan pembangunan komponen dari fondasi hingga tahap integrasi.

## Hints (0/3)
### Hint 1
Pikirkan pemisahan antara pencarian radius lokasi (geospasial), manajemen batas waktu 30 detik (timer/timeout), dan pengiriman notifikasi SMS.

### Hint 2
Gunakan diagram Mermaid flowchart TD.

### Hint 3
Pastikan tidak ada dependensi sirkular antara CourierFinder dan OfferManager.


## Key Takeaway
- Studi kasus nyata menguji ketelitian memisahkan logika pencarian geospasial, timer timeout, dan penguncian state transaksi.
- Pisahkan orkestrator dari pelaksana tugas spesifik.
