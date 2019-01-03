package db

import (
	"errors"
	"sync"
)

type KVDB struct {
	mx sync.Mutex
	Data map[string]string
}

var NoSuchElementError = errors.New("no such element")

func (kv *KVDB) Get(key string) (string, error) {
	kv.mx.Lock()
	defer kv.mx.Unlock()

	if value, ok := kv.Data[key]; ok {
		return value, nil
	} else {
		return "", NoSuchElementError
	}
}

func (kv *KVDB) Put(key string, value string) error {
	kv.mx.Lock()
	defer kv.mx.Unlock()
	kv.Data[key] = value
	return nil
}