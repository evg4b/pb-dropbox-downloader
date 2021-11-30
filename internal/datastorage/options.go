package datastorage

import "github.com/go-git/go-billy/v5"

type storageOption = func(s *FileStorage)

func WithMarshalFunc(marshal MarshalFunc) storageOption {
	return func(s *FileStorage) {
		s.marshal = marshal
	}
}

func WithUnmarshalFunc(unmarshal UnmarshalFunc) storageOption {
	return func(s *FileStorage) {
		s.unmarshal = unmarshal
	}
}

func WithConfigPath(configPath string) storageOption {
	return func(s *FileStorage) {
		s.configPath = configPath
	}
}

func WithFilesystem(files billy.Filesystem) storageOption {
	return func(s *FileStorage) {
		s.files = files
	}
}
