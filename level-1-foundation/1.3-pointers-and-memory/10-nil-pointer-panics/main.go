package main

import "fmt"

type Profile struct {
	Bio string
}

type ServerConfig struct {
	Host string
	Port int
}

type AppConfig struct {
	AppName string
	Server  *ServerConfig
}

// createProfile seharusnya membuat dan mengembalikan *Profile yang valid.
// BUG: p belum dialokasikan memori (masih nil), mencoba mengisi p.Bio memicu panic!
func createProfile(bio string) *Profile {
	return &Profile{Bio: bio}
}

// loadConfig seharusnya membuat dan mengembalikan *AppConfig lengkap dengan ServerConfig.
// BUG: cfg.Server masih nil saat field Host dan Port diakses!
func loadConfig(appName string, host string, port int) *AppConfig {
	return &AppConfig{
		AppName: appName,
		Server: &ServerConfig{
			Host: host,
			Port: port,
		},
	}
}

func main() {
	// 1. Uji createProfile
	prof := createProfile("Software Engineer")
	fmt.Println("Profile:", prof.Bio)

	// 2. Uji loadConfig
	cfg := loadConfig("WebApp", "localhost", 8080)
	fmt.Printf("Config: %s on %s:%d\n", cfg.AppName, cfg.Server.Host, cfg.Server.Port)
}
