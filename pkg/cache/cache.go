package github

import (
	"encoding/json"
	"os"

	"github.com/crl-n/github-readme-stats-go/pkg/logger"
)

type Cache[K comparable, V any] struct {
	filename string
	data     map[K]V
}

func NewCache[K comparable, V any](filename string) *Cache[K, V] {
	c := &Cache[K, V]{
		filename: filename,
		data:     make(map[K]V),
	}

	c.loadFileData()

	return c
}

func (c *Cache[K, V]) loadFileData() {
	file, err := os.Open(c.filename)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Debugf("Cache file '%s' not found.\n", c.filename)
		} else {
			logger.Errorf("Error opening file '%s': %v\n", c.filename, err)
		}
		return
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c.data)
	if err != nil {
		logger.Errorf("Error decoding cache file '%s': %v\n", c.filename, err)
		return
	}

	logger.Debugf("Cache file '%s' found and decoded succesfully\n", c.filename)
}

func (c *Cache[K, V]) fileExists() bool {
	_, err := os.Stat(c.filename)
	if err == nil {
		return true
	}
	return false
}

// Creates new cache file or overwrites previous file with new data
func (c *Cache[K, V]) saveToFile() {
	logger.Debugf("Saving cache data to file '%s'\n", c.filename)
	if c.fileExists() {
		logger.Debugf("Cache file '%s' exists and will be overwritten\n", c.filename)
	}

	file, err := os.Create(c.filename)
	if err != nil {
		logger.Errorf("Error creating file '%s': %v\n", c.filename, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(c.data)
	if err != nil {
		logger.Errorf("Error encoding cache data: %v\n", err)
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.data[key] = value
	c.saveToFile()
}

func (c *Cache[K, V]) BulkSet(pairs []struct {
	Key   K
	Value V
}) {
	for _, pair := range pairs {
		c.data[pair.Key] = pair.Value
	}
	c.saveToFile()
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	value, found := c.data[key]
	return value, found
}
