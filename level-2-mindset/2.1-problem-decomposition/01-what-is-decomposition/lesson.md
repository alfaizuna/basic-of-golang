# What Is Decomposition?

## Konsep
Dekomposisi adalah seni memecah masalah besar dan kabur menjadi bagian-bagian kecil yang dapat dipahami, diuji, dan diselesaikan secara independen.
Kode yang buruk jarang berakar dari salah ketik sintaks; kode buruk hampir selalu berakar dari melompat langsung ke editor kode sebelum benar-benar memahami apa masalah yang ingin dipecahkan.
Seorang software engineer sejati tidak langsung membuat struct atau function saat menerima permintaan fitur; mereka mengklarifikasi ambiguitas hingga batas-batas masalah menjadi tajam dan terdefinisi.

## Skenario
Klien Anda (Direktur Operasional Klinik Sehat) mengirimkan pesan singkat via WhatsApp:
*"Halo mas, tolong buatin sistem pengingat otomatis ya buat klinik kami. Soalnya banyak pasien yang sering lupa jadwal janji dokter, akibatnya dokter kami sering bengong nungguin pasien yang bolos dan klinik rugi waktu!"*

Jika Anda langsung mulai mengetik kode Go sekarang, Anda pasti akan membuat asumsi liar: Pengingat lewat apa? Berapa jam sebelum jadwal? Bagaimana jika nomor tidak aktif? Bagaimana jika pasien ingin membatalkan atau ubah jadwal?

## Tugas
Analisis permintaan kabur di atas. Di dalam file `answer.md`, tuliskan:
1. **Identifikasi Masalah Utama**: 1-2 kalimat yang merangkum inti masalah bisnis yang sebenarnya.
2. **5 Pertanyaan Klarifikasi Kritis**: 5 pertanyaan terarah yang wajib Anda tanyakan kepada pihak klinik sebelum mendesain sistem.
3. **3 Batasan Awal (Scope Constraints)**: Tentukan 3 hal yang secara eksplisit **TIDAK** akan dikerjakan pada versi pertama (V1) agar cakupan proyek tidak melebar liar (*scope creep*).

## Hints (0/3)
### Hint 1
Pikirkan tentang media komunikasi (SMS, WhatsApp, Email), waktu pengiriman, dan biaya operasional.

### Hint 2
Pikirkan tentang alur interaksi dua arah: apakah pengingat hanya satu arah (broadcast) atau pasien bisa membalas 'YA' / 'BATAL'?

### Hint 3
Batasan awal yang baik misalnya: tidak membuat aplikasi mobile baru untuk pasien, atau tidak mengintegrasikan sistem pembayaran di v1.


## Key Takeaway
- Jangan pernah menulis kode dari instruksi yang ambigu; definisikan masalah sebelum merancang solusi.
- Menetapkan apa yang **tidak dibuat** (out-of-scope) sama pentingnya dengan apa yang akan dibuat.
