package main

// example.go — hanya untuk referensi, jangan diubah

type Item struct {
	Key   string
	Count int
}

type Cache struct {
	data map[string]*Item
}

func (c *Cache) FindOrCreate(key string) *Item {
	if it, exists := c.data[key]; exists {
		return it
	}
	newItem := &Item{Key: key, Count: 0}
	c.data[key] = newItem
	return newItem
}
