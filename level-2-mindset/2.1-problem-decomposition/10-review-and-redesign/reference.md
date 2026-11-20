## Contoh Jawaban Referensi

### 1. 3 Kritik Arsitektural Konstruktif
1. **God Object Anti-Pattern**: Struct `MegaNotificationManager` mencampuradukkan kredensial database dan API gateway WhatsApp. Jika kredensial WA berubah, komponen database ikut berisiko.
2. **Ketiadaan Abstraksi (Untestable)**: Mengunci pengiriman ke implementasi konkret tanpa interface membuat unit testing mustahil dilakukan tanpa mengirim WhatsApp sungguhan.
3. **Parameter Primitive Obsession**: Meneruskan `patientID int`, `doctorID int`, `msg string` secara terpisah membuat signature rapuh. Seharusnya dibungkus dalam struct model domain.

### 2. Revisi Arsitektur
Pisahkan menjadi:
- Model domain `AppointmentReminder`.
- Interface `Notifier` yang dapat diimplementasikan oleh mock atau WhatsApp gateway riil.
- Service orkestrator yang menerima `Notifier` lewat Dependency Injection.