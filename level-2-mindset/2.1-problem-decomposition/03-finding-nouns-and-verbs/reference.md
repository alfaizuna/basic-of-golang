## Contoh Jawaban Referensi

### 1. Kata Benda (Entities & Attributes)
- **Appointment (Janji Temu)**: `ID`, `PatientID`, `DoctorID`, `ScheduledTime`, `Status` (Pending/Confirmed/Cancelled).
- **Patient (Pasien)**: `ID`, `Name`, `PhoneNumber`.
- **Doctor (Dokter)**: `ID`, `Name`, `Specialization`, `WorkSchedule`.
- **ReminderNotification (Notifikasi Pengingat)**: `AppointmentID`, `Channel`, `SentAt`, `DeliveryStatus`.

### 2. Kata Kerja (Operations & Actions)
- **Memverifikasi Jadwal (`CheckAvailability`)**: Dilakukan sistem untuk memastikan dokter tidak memiliki jadwal ganda pada jam yang sama.
- **Menyusun Pengingat (`FindUpcomingAppointments`)**: Dilakukan sistem terjadwal untuk menyaring janji temu H-1.
- **Mengirim Pengingat (`SendReminder`)**: Dilakukan gateway pesan untuk menyampaikan notifikasi ke nomor pasien.
- **Mengonfirmasi Kehadiran (`ConfirmAttendance`)**: Dipicu oleh respons pasien untuk mengubah status janji menjadi Confirmed.
- **Membatalkan Janji (`CancelAppointment`)**: Dipicu oleh pasien/resepsionis untuk melepaskan slot jadwal dokter.

### 3. Pemetaan Operasi ke Entitas
| Entitas Bertanggung Jawab | Operasi (Method Candidate) |
|---|---|
| `ScheduleService` | `CheckAvailability(doctorID, time)`, `ReleaseSlot(appointmentID)` |
| `AppointmentStore` | `Save(appointment)`, `FindUpcoming(date)`, `UpdateStatus(id, status)` |
| `Notifier` | `SendReminder(patient, appointment)` |