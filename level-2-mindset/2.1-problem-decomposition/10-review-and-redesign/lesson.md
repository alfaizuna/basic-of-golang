# Review & Redesign: The Go Blueprint

## Konsep
Tingkat kematangan tertinggi seorang insinyur perangkat lunak adalah kemampuan **mengkritik rancangan yang cacat** dan merevisinya menjadi **kerangka arsitektur Go yang bersih (*Go Type Blueprint*)**.
Kerangka rancangan Go hanya berisi definisi struct, interface, dan function signature tanpa kode implementasi tubuh (*empty body / stubs*).
Kerangka ini harus valid secara sintaksis dan dapat dikompilasi oleh `go build`.

## Skenario
Seorang developer freelance menyerahkan rancangan arsitektur Go untuk modul notifikasi Klinik Sehat.
Berikut kode yang ia tulis:

```go
package notification

type MegaNotificationManager struct {
    DBConnectionURL string
    WhatsAppAPIKey  string
}

// Satu method raksasa melakukan semuanya
func (m *MegaNotificationManager) ProcessEverything(patientID int, doctorID int, msg string) (bool, string, error)
```

Rancangan ini memiliki cacat arsitektur berat:
- Menggabungkan konfigurasi database dan API key eksternal dalam satu struct.
- Method `ProcessEverything` melanggar Single Responsibility dan tidak fleksibel diuji.

## Tugas
Di dalam file `answer.md`:
1. **Code Review**: Tuliskan 3 kritik arsitektural yang konstruktif terhadap kode di atas.
2. **Revisi Arsitektur**: Jelaskan konsep pemisahan baru yang Anda usulkan.
3. **Kerangka Kode Go (`skeleton.go`)**:
   Lengkapi file `skeleton.go` di folder ini dengan kerangka Go types yang valid:
   - Interface `Notifier` dengan method `Notify(to string, text string) error`.
   - Struct data murni `AppointmentReminder`.
   - Struct `Service` yang menerima interface melalui constructor `NewService(n Notifier) *Service`.
   - Method `(s *Service) SendReminder(reminder AppointmentReminder) error` (cukup return `nil` sebagai stub).
   - Pastikan kode bisa dikompilasi dengan `go build`.

## Hints (0/3)
### Hint 1
Gunakan prinsip dependency injection pada constructor NewService.

### Hint 2
Pisahkan model data dari service behavior.

### Hint 3
Pastikan file skeleton.go memiliki package main atau package definition yang valid.


## Key Takeaway
- Arsitektur Go yang bersih selalu memisahkan data (struct), kontrak perilaku (interface), dan logika orkestrasi (service).
- Dependency Injection via interface membuat kode mudah diuji tanpa mocking rumit.
