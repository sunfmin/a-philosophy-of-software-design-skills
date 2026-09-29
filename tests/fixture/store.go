package store

import (
	"errors"
	"os"
	"strings"
)

// ConfigReader reads the config file.
type ConfigReader struct{ Path string }

// Read reads the file and splits "key=value" lines into a map.
func (r ConfigReader) Read() (map[string]string, error) {
	b, err := os.ReadFile(r.Path)
	if err != nil {
		return nil, err
	}
	m := map[string]string{}
	for _, line := range strings.Split(string(b), "\n") {
		if kv := strings.SplitN(line, "=", 2); len(kv) == 2 {
			m[kv[0]] = kv[1]
		}
	}
	return m, nil
}

// ConfigWriter writes the config file as "key=value" lines.
type ConfigWriter struct{ Path string }

func (w ConfigWriter) Write(m map[string]string) error {
	var sb strings.Builder
	for k, v := range m {
		sb.WriteString(k + "=" + v + "\n")
	}
	return os.WriteFile(w.Path, []byte(sb.String()), 0o644)
}

// Store wraps a Cache.
type Store struct {
	cache *Cache
	count int // count
}

// Get gets.
func (s *Store) Get(key string) (string, error) { return s.cache.Get(key) }

// Put puts.
func (s *Store) Put(key, val string) error { return s.cache.Put(key, val) }

var ErrNotFound = errors.New("not found")

// Delete removes key. Returns ErrNotFound if the key does not exist.
func (s *Store) Delete(key string) error {
	if _, err := s.cache.Get(key); err != nil {
		return ErrNotFound
	}
	s.count--
	return s.cache.Delete(key)
}

// Params returns the internal map.
func (s *Store) Params() map[string]string { return s.cache.data }

type Cache struct{ data map[string]string }

func (c *Cache) Get(k string) (string, error) {
	v, ok := c.data[k]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}
func (c *Cache) Put(k, v string) error { c.data[k] = v; return nil }
func (c *Cache) Delete(k string) error { delete(c.data, k); return nil }
