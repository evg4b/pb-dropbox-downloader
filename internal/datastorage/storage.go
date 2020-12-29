package datastorage

import (
	"pb-dropbox-downloader/infrastructure"
	"sync"

	"github.com/kelindar/binary"
)

type FileStorage struct {
	data       map[string]string
	mu         sync.Mutex
	files      infrastructure.FileSystem
	configPath string
}

func NewFileStorage(files infrastructure.FileSystem, configPath string) *FileStorage {
	return &FileStorage{
		data:       nil,
		files:      files,
		configPath: configPath,
		mu:         sync.Mutex{},
	}
}

func (storage *FileStorage) Get(key string) (string, bool) {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	err := storage.preload()
	if err != nil {
		return "", false
	}

	value, ok := storage.data[key]

	return value, ok
}

func (storage *FileStorage) ToMap() (map[string]string, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	err := storage.preload()
	if err != nil {
		return nil, err
	}

	return storage.data, nil
}

func (storage *FileStorage) FromMap(data map[string]string) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	err := storage.preload()
	if err != nil {
		return err
	}

	storage.data = data

	err = storage.unload()
	if err != nil {
		return err
	}

	return nil
}

func (storage *FileStorage) KeyExists(key string) bool {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	err := storage.preload()
	if err != nil {
		return false
	}

	_, ok := storage.data[key]

	return ok
}

func (storage *FileStorage) Commit() error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	return storage.unload()
}
func (storage *FileStorage) unload() error {
	data, err := binary.Marshal(storage.data)
	if err != nil {
		return err
	}

	return storage.files.WriteFile(storage.configPath, data)
}

func (storage *FileStorage) preload() error {
	if storage.data == nil {
		storage.data = make(map[string]string)

		data, err := storage.files.ReadFile(storage.configPath)
		if err != nil {
			return err
		}

		return binary.Unmarshal(data, &storage.data)
	}

	return nil
}

func (storage *FileStorage) Add(key, value string) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	storage.data[key] = value
}
