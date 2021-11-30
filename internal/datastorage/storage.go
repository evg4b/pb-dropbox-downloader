package datastorage

import (
	"errors"
	"os"
	"sync"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/util"
	"github.com/kelindar/binary"
)

var ErrKeyDoesNotExists = errors.New("key does not exists")

type FileStorage struct {
	data       map[string]string
	mu         sync.Mutex
	files      billy.Filesystem
	configPath string
}

func NewFileStorage(files billy.Filesystem, configPath string) *FileStorage {
	return &FileStorage{
		data:       nil,
		files:      files,
		configPath: configPath,
		mu:         sync.Mutex{},
	}
}

func (storage *FileStorage) Get(key string) (string, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	data, err := storage.preload()
	if err != nil {
		return "", err
	}

	if value, ok := data[key]; ok {
		return value, nil
	}

	return "", ErrKeyDoesNotExists
}

func (storage *FileStorage) ToMap() (map[string]string, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	data, err := storage.preload()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (storage *FileStorage) FromMap(data map[string]string) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	storage.data = data

	if err := storage.unload(); err != nil {
		return err
	}

	return nil
}

func (storage *FileStorage) KeyExists(key string) bool {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	data, err := storage.preload()
	if err != nil {
		return false
	}

	_, ok := data[key]

	return ok
}

func (storage *FileStorage) Commit() error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	return storage.unload()
}

func (storage *FileStorage) Add(key, value string) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	storage.data[key] = value
}

func (storage *FileStorage) Remove(key string) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	delete(storage.data, key)
}

func (storage *FileStorage) unload() error {
	data, err := binary.Marshal(storage.data)
	if err != nil {
		return err
	}

	return util.WriteFile(storage.files, storage.configPath, data, os.ModePerm)
}

func (storage *FileStorage) preload() (map[string]string, error) {
	if storage.data == nil {
		storage.data = make(map[string]string)

		data, err := util.ReadFile(storage.files, storage.configPath)
		if err != nil {
			if os.IsNotExist(err) {
				return storage.data, nil
			}

			return storage.data, err
		}

		err = binary.Unmarshal(data, &storage.data)
		if err != nil {
			return storage.data, err
		}

		return storage.data, nil
	}

	return storage.data, nil
}
