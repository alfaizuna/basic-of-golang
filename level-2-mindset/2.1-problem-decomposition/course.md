# Course 2.1: Problem Decomposition

## Tujuan
Mengembangkan pola pikir insinyur perangkat lunak (*software engineer mindset*): memecah masalah kabur menjadi potongan-potongan komponen mandiri yang masing-masing hanya mengerjakan satu hal dengan batasan yang jelas, menentukan kontrak input/output, memetakan dependensi sistem, mengeliminasi dependensi sirkular, menentukan granularitas yang seimbang, menyusun urutan eksekusi kerja (MVP), serta merancang arsitektur sebelum menulis baris kode implementasi.

## Karakteristik Level II
Level ini berfokus pada penalaran konseptual, analisis skenario nyata, dan perancangan arsitektur.
Setiap lesson memiliki struktur file:
- `lesson.md`: Konsep singkat, skenario konkret, tugas analisis, hints, dan key takeaway.
- `answer.md`: Lembar kerja tempat Anda menuliskan jawaban atau kerangka rancangan.
- `rubric.md`: Kriteria evaluasi mentor untuk menilai kualitas argumen saat perintah `/submit`.
- `hints.md`: 3 petunjuk pemandu bertahap jika Anda membutuhkan arahan.
- `reference.md`: Contoh solusi teladan dan ulasan mendalam.

## Daftar Lesson

| # | Lesson | Fokus & Skenario | Status |
|---|--------|------------------|:------:|
| 01 | What Is Decomposition? | Masalah kabur vs terdefinisi; merumuskan 3–5 pertanyaan & kebutuhan konkret | 🔒 |
| 02 | One Piece, One Job | Single Responsibility; tes "satu kalimat tanpa kata *dan*" | 🔒 |
| 03 | Finding Nouns and Verbs | Ekstraksi entitas data (kata benda) dan operasi (kata kerja) dari narasi bisnis | 🔒 |
| 04 | Drawing Boundaries | Definisi kontrak komponen: Input, Output, Responsibility, Information Hiding | 🔒 |
| 05 | Mapping Dependencies | Diagram dependensi sistem (Mermaid); arah pemanggilan yang bersih | 🔒 |
| 06 | Cycles & Tangles | Mendeteksi dan memutus siklus dependensi sirkular (*acyclic graph*) | 🔒 |
| 07 | Too Big, Too Small | Menentukan granularitas seimbang: anti Monolith Tangle & Micro-Fragmentation | 🔒 |
| 08 | Ordering the Work | Menentukan urutan implementasi berdasarkan risiko & dependensi (fondasi ke daun) | 🔒 |
| 09 | Decomposing a Real Feature | Studi kasus mandiri penuh: Sistem Logistik Pengiriman Paket Instan | 🔒 |
| 10 | Review & Redesign | Meninjau & merevisi arsitektur cacat; memodelkan kerangka Go types & signatures | 🔒 |

## Hasil Akhir
Kemampuan menganalisis kebutuhan klien yang ambigu, memecahnya menjadi komponen-komponen decoupled dengan batas tanggung jawab yang tegas, dan menuangkannya ke dalam kerangka antarmuka kode Go yang kokoh sebelum implementasi dimulai.
