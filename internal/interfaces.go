package internal

import (
	"io"

	"github.com/nickalie/fskv"
)

// DataStorage interface to storage key-value data.
type DataStorage interface {
	Get(string) (string, bool)
	ToMap() (map[string]string, error)
	FromMap(map[string]string) error
	KeyExists(string) bool
	Commit() error
	Add(string, string)
	Remove(string)
}

// FileSystem interface to comunitace with file system.
type FileSystem interface {
	// GetFilesInFolder return relative file paths in folder (include subfolder to).
	GetFilesInFolder(string) []string
	// CopyDataToFile copy data from reader to file.
	CopyDataToFile(string, io.Reader) error
	// DeleteFile delete file from target filesystem.
	DeleteFile(string) error
	// ReadFile read file content from target filesystem.
	ReadFile(filename string) ([]byte, error)
	// WriteFile writes content to file in target filesystem.
	WriteFile(filename string, data []byte) error
}

// RemoteFile is structure to describe file in dropbox.
type RemoteFile struct {
	Path string
	Hash string
	Size uint64
}

// Dropbox is intereface to dropbox client wrapper.
type Dropbox interface {
	// GetFiles returns files in application folder (include subfolder to).
	GetFiles() ([]RemoteFile, error)
	// DownloadFile downloaded file by path.
	DownloadFile(string) (io.ReadCloser, error)
	// AccountDisplayName returns account display name.
	AccountDisplayName() string
	// AccountEmail returns account display email.
	AccountEmail() string
}

type Database interface {
	GetBucket(name string) (*fskv.Bucket, error)
	Set(key string, value []byte)
	Get(key string)
	Scan(prefix string, f func(key string, value []byte) bool)
	Remove(keys ...string) error
}
