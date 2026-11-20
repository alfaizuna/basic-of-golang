# Finding Nouns and Verbs

## Konsep
Saat membaca narasi kebutuhan sistem, komponen perangkat lunak tersembunyi di dalam tata bahasa:
- **Kata Benda (Nouns)** biasanya mengindikasikan entitas data, model struct, atau penyimpan status (*state*).
- **Kata Kerja (Verbs)** biasanya mengindikasikan operasi, algoritma, method, atau interface behavior.
Mengekstrak kata benda dan kata kerja secara terstruktur adalah langkah pertama mengubah cerita bisnis menjadi rancangan teknis.

## Skenario
Berikut adalah wawancara rinci dengan Kepala Resepsionis Klinik Sehat:

> *"Setiap pagi, resepsionis mencatat **jadwal janji temu** baru yang berisi data **pasien**, pilihan **dokter**, dan **waktu kunjungan**. Sistem harus bisa memverifikasi ketersediaan jadwal dokter tersebut agar tidak bentrok. Satu hari sebelum kunjungan, sistem otomatis menyusun **daftar pengingat** dan mengirimkan **notifikasi pesan** ke nomor telepon pasien. Pasien dapat membalas pesan tersebut untuk mengonfirmasi kehadiran. Jika pasien mengonfirmasi hadir, sistem mengubah **status janji** menjadi terkonfirmasi. Jika pasien membatalkan, sistem mengosongkan slot jadwal dan mencatat pembatalan ke dalam **laporan harian** klinik."*

## Tugas
Di dalam file `answer.md`:
1. **Daftar Kata Benda (Entities / State)**: Ekstrak minimal 4 kata benda utama yang merepresentasikan data penting, dan sebutkan 2-3 atribut field yang melekat padanya.
2. **Daftar Kata Kerja (Operations / Behaviors)**: Ekstrak minimal 4 kata kerja utama yang merepresentasikan aksi bisnis penting, serta sebutkan siapa yang melakukan aksi tersebut dan apa dampaknya.
3. **Pemetaan**: Pasangkan kata kerja tersebut ke entitas kata benda yang paling bertanggung jawab menjalankannya.

## Hints (0/3)
### Hint 1
Contoh kata benda: Janji Temu (Appointment), Pasien (Patient), Dokter (Doctor).

### Hint 2
Contoh kata kerja: Memverifikasi (Verify/Check), Mengirim (Send), Mengonfirmasi (Confirm), Membatalkan (Cancel).

### Hint 3
Perhatikan kata kerja yang pasif: 'disusun' atau 'dikosongkan' tetap merupakan kata kerja sistem.


## Key Takeaway
- Kata benda dari spesifikasi menjadi struct dan database tables.
- Kata kerja dari spesifikasi menjadi methods, functions, dan interface contracts.
