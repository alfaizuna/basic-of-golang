package main

import (
	"encoding/json"
	"fmt"
)

type ServerConfig struct {
	Port  int    `json:"port"`
	Host  string `json:"host"`
	Debug bool   `json:"debug"`
}

// ParseConfig mem-parsing data JSON ke struct ServerConfig.
func ParseConfig(data []byte) (*ServerConfig, error) {
	_ = json.Unmarshal
	// TODO: Gunakan json.Unmarshal dengan pointer &cfg
	return nil, nil
}

func main() {
	raw := []byte(`{"port": 8080, "host": "localhost", "debug": true}`)
	cfg, err := ParseConfig(raw)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Port: %d, Host: %s, Debug: %t\n", cfg.Port, cfg.Host, cfg.Debug)
}
