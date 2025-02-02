// Package storage
// Date: 2024/3/26 11:42
// Author: Amu
// Description:
package storage

import (
	"encoding/json"
	bolt "go.etcd.io/bbolt"
	"strings"
)

type Tx struct {
	tx *bolt.Tx
}

func (tx *Tx) Keys(prefix string, recursive bool) ([]string, error) {
	bucket := tx.tx.Bucket(defaultBucket)
	if bucket == nil {
		return nil, ErrKeyNotFound
	}

	var rs []string
	err := bucket.ForEach(func(k, v []byte) error {
		key := string(k)
		if strings.HasPrefix(key, prefix) {
			tr := key[len(prefix):]
			if !strings.Contains(tr, Separator) || recursive {
				rs = append(rs, key)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (tx *Tx) Get(key string) ([]byte, error) {
	bucket := tx.tx.Bucket(defaultBucket)
	if bucket == nil {
		return nil, ErrKeyNotFound
	}

	val := bucket.Get([]byte(key))
	if val == nil {
		return nil, ErrKeyNotFound
	}
	return val, nil
}

func (tx *Tx) GetString(key string) (string, error) {
	val, err := tx.Get(key)
	if err != nil {
		return "", err
	}
	return string(val), err
}

func (tx *Tx) GetJson(key string, out interface{}) error {
	val, err := tx.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(val, out)
}

func (tx *Tx) Put(key string, val []byte) error {
	bucket := tx.tx.Bucket(defaultBucket)
	if bucket == nil {
		return ErrKeyNotFound
	}
	return bucket.Put([]byte(key), val)
}

func (tx *Tx) PutString(key string, val string) error {
	return tx.Put(key, []byte(val))
}

func (tx *Tx) PutJson(key string, val interface{}) error {
	raw, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return tx.Put(key, raw)
}

func (tx *Tx) Delete(key string) error {
	bucket := tx.tx.Bucket(defaultBucket)
	if bucket == nil {
		return ErrKeyNotFound
	}
	return bucket.Delete([]byte(key))
}
