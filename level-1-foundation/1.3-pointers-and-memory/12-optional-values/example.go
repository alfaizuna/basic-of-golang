package main

// example.go — hanya untuk referensi materi (read-only)

// NotificationConfig mendemonstrasikan pemakaian pointer untuk nilai opsional.
// Jika tipe biasa (misal bool), zero value-nya adalah false.
// Kita tidak bisa membedakan: "pengguna sengaja memilih false" atau "pengguna belum mengisinya".
// Dengan pointer (*bool), nil berarti "belum diatur" (gunakan default).
type NotificationConfig struct {
	EmailAlerts *bool // nil = default (true), true = ON, false = OFF
	MaxRetries  *int  // nil = default (3)
}

func newBool(b bool) *bool {
	return &b
}

func newInt(n int) *int {
	return &n
}

type ResolvedNotification struct {
	EmailAlerts bool
	MaxRetries  int
}

func ResolveNotification(cfg NotificationConfig) ResolvedNotification {
	res := ResolvedNotification{
		EmailAlerts: true, // nilai default
		MaxRetries:  3,    // nilai default
	}

	if cfg.EmailAlerts != nil {
		res.EmailAlerts = *cfg.EmailAlerts
	}
	if cfg.MaxRetries != nil {
		res.MaxRetries = *cfg.MaxRetries
	}

	return res
}
