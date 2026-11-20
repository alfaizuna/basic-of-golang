# Cycles & Tangles

## Konsep
Dependensi sirkular (*circular dependency*) terjadi ketika Komponen A bergantung pada B, dan Komponen B secara langsung atau tidak langsung bergantung kembali pada A (A -> B -> A).
Di Go, circular import akan **langsung ditolak kompilator** saat build time.
Di level arsitektur, siklus dependensi membuat modul tidak bisa dipisahkan, mustahil diuji secara terisolasi, dan menimbulkan efek samping tersembunyi.
Solusi standar memutus siklus:
1. **Ekstraksi Komponen Ketiga**: Pindahkan data atau fungsi bersama ke modul baru (C).
2. **Inversi Dependensi (Interface)**: Definisikan interface pada modul pemanggil.

## Skenario
Seorang developer menambahkan fitur baru di sistem Klinik Sehat:
- `AppointmentManager` memanggil `NotificationSender.Send()` untuk mengirim pengingat.
- Namun di dalam `NotificationSender`, developer tersebut menambahkan kode: jika pengiriman pesan berhasil, `NotificationSender` memanggil kembali `AppointmentManager.MarkNotificationSent(appointmentID)`.

Akibatnya terjadi lingkaran setan:
`AppointmentManager` ➔ butuh ➔ `NotificationSender`
`NotificationSender` ➔ butuh ➔ `AppointmentManager`

## Tugas
Di dalam file `answer.md`:
1. **Identifikasi Masalah**: Jelaskan mengapa ketergantungan bolak-balik di atas berbahaya dan bagaimana kompilator Go meresponsnya jika keduanya berada di package berbeda.
2. **Solusi 1 (Pola Return Value / Orkestrator)**: Jelaskan bagaimana cara memutus siklus tersebut cukup dengan mengubah nilai kembalian (*return value*) dari `NotificationSender.Send()`.
3. **Solusi 2 (Diagram Mermaid Bersih)**: Gambarkan diagram dependensi baru yang sudah bebas siklus (*acyclic*).

## Hints (0/3)
### Hint 1
Kenapa NotificationSender harus repot mengupdate status ke database? Bukankah ia hanya bertugas mengirim pesan?

### Hint 2
Cukup kembalikan status berhasil/gagal dari `Send()` ke pemanggil, lalu biarkan pemanggil yang mengupdate status!

### Hint 3
Pastikan arah panah Mermaid mengalir tanpa membentuk lingkaran tertutup.


## Key Takeaway
- Dependensi sirkular adalah tanda bahwa batas tanggung jawab komponen bocor.
- Putus siklus dengan mengembalikan hasil via return value atau mengekstrak dependensi bersama ke komponen terpisah.
