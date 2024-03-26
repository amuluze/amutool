// Package storage
// Date: 2024/3/26 11:20
// Author: Amu
// Description:
package storage

import (
	"github.com/amuluze/amutool/errors"
	bolt "go.etcd.io/bbolt"
	"sync"
	"time"
)

var (
	defaultBucket = []byte("default")
	defaultTimout = 10 * time.Second

	Separator = "/"

	ErrKeyNotFound = errors.New("Key is not found")
)

type Storage struct {
	timeout time.Duration

	path string
	db   *bolt.DB
	ref  uint
	mu   sync.Mutex
}

func NewStorage(path string) *Storage {
	return &Storage{timeout: defaultTimout, path: path}
}

func (s *Storage) Path() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.path
}

// Update 写入操作
func (s *Storage) Update(fn func(tx *Tx) error) error {
	err := s.acquire()
	if err != nil {
		return err
	}

	defer s.release()

	return s.db.Update(func(btx *bolt.Tx) error {
		tx := &Tx{tx: btx}
		return fn(tx)
	})
}

// View 读取操作
func (s *Storage) View(fn func(tx *Tx) error) error {
	err := s.acquire()
	if err != nil {
		return err
	}
	defer s.release()

	return s.db.View(func(btx *bolt.Tx) error {
		tx := &Tx{tx: btx}
		return fn(tx)
	})
}

// Delete 删除 key
func (s *Storage) Delete(key string) error {
	return s.Update(func(tx *Tx) error {
		return tx.Delete(key)
	})
}

func (s *Storage) Keys(prefix string, recursive bool) ([]string, error) {
	var keys []string
	err := s.View(func(tx *Tx) (nErr error) {
		keys, nErr = tx.Keys(prefix, recursive)
		return
	})
	return keys, err
}

func (s *Storage) Get(key string) ([]byte, error) {
	var result []byte
	err := s.View(func(tx *Tx) (nErr error) {
		result, nErr = tx.Get(key)
		return
	})
	return result, err
}

func (s *Storage) GetString(key string) (string, error) {
	var result string
	err := s.View(func(tx *Tx) (nErr error) {
		result, nErr = tx.GetString(key)
		return
	})
	return result, err
}

func (s *Storage) GetJson(key string, out interface{}) error {
	return s.View(func(tx *Tx) error {
		return tx.GetJson(key, out)
	})
}

func (s *Storage) Put(key string, val []byte) error {
	return s.Update(func(tx *Tx) error {
		return tx.Put(key, val)
	})
}

func (s *Storage) PutString(key string, val string) error {
	return s.Update(func(tx *Tx) error {
		return tx.PutString(key, val)
	})
}

func (s *Storage) PutJson(key string, val interface{}) error {
	return s.Update(func(tx *Tx) error {
		return tx.PutJson(key, val)
	})
}

func (s *Storage) acquire() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.ref == 0 {
		var err error
		s.db, err = bolt.Open(s.path, 0600, &bolt.Options{Timeout: defaultTimout})
		if err != nil {
			return err
		}

		// create default bucket if not exists
		err = s.db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists(defaultBucket)
			return err
		})
	}
	s.ref++
	return nil
}

func (s *Storage) release() {
	for i := 0; i < 3; i++ {
		err := s.doRelease()
		if err == nil {
			return
		}
		time.Sleep(s.timeout)
	}
}

func (s *Storage) doRelease() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.ref == 0 {
		return nil
	}
	if s.ref == 1 {
		err := s.db.Close()
		if err != nil {
			return err
		}
		s.db = nil
	}
	s.ref--
	return nil
}
