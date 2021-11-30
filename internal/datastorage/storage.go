package datastorage

import (
	"errors"
	"os"
	"sync"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/util"
	"github.com/kelindar/binary"
)

type MarshalFunc = func(interface{}) ([]byte, error)
type UnmarshalFunc = func([]byte, interface{}) error

var ErrKeyDoesNotExists = errors.New("key does not exists")

type FileStorage struct {
	data       map[string]string
	mu         sync.Mutex
	files      billy.Filesystem
	configPath string
	unmarshal  UnmarshalFunc
	marshal    MarshalFunc
}

func NewFileStorage(options ...storageOption) *FileStorage {
	storage := &FileStorage{
		mu:         sync.Mutex{},
		unmarshal:  binary.Unmarshal,
		marshal:    binary.Marshal,
		configPath: "storage.bin",
	}

	for _, option := range options {
		option(storage)
	}

	return storage
}

func (s *FileStorage) Get(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := s.load()
	if err != nil {
		return "", err
	}

	if value, ok := data[key]; ok {
		return value, nil
	}

	return "", ErrKeyDoesNotExists
}

func (s *FileStorage) ToMap() (map[string]string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := s.load()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *FileStorage) FromMap(data map[string]string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = data
}

func (s *FileStorage) KeyExists(key string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := s.load()
	if err != nil {
		return false, err
	}

	_, ok := data[key]

	return ok, nil
}

func (s *FileStorage) Commit() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := s.marshal(s.data)
	if err != nil {
		return err
	}

	return util.WriteFile(s.files, s.configPath, data, os.ModePerm)
}

func (s *FileStorage) Add(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, err := s.load(); err != nil {
		return err
	}

	s.data[key] = value

	return nil
}

func (s *FileStorage) Remove(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, err := s.load(); err != nil {
		return err
	}

	delete(s.data, key)

	return nil
}

func (s *FileStorage) load() (map[string]string, error) {
	if s.data == nil {
		s.data = make(map[string]string)

		data, err := util.ReadFile(s.files, s.configPath)
		if err != nil {
			if os.IsNotExist(err) {
				return s.data, nil
			}

			return s.data, err
		}

		err = s.unmarshal(data, &s.data)
		if err != nil {
			return s.data, err
		}

		return s.data, nil
	}

	return s.data, nil
}
