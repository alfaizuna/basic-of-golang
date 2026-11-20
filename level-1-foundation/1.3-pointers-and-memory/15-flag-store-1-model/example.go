package main

// example.go — hanya untuk referensi, jangan diubah

type ConfigItem struct {
	Key   string
	Value string
}

type ConfigStore struct {
	items map[string]*ConfigItem
}

func NewConfigStore() *ConfigStore {
	return &ConfigStore{
		items: make(map[string]*ConfigItem),
	}
}
