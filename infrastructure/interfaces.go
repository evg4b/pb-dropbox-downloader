package infrastructure

import (
	"io"
)

// FileSystem interface to comunitace with file system
type FileSystem interface {
	// GetFilesInFolder return relative file paths in folder (include subfolder to)
	GetFilesInFolder(string) []string
	// CopyDataToFile copy data from reader to file
	CopyDataToFile(string, io.Reader) error
	// DeleteFile delete file from target filesystem
	DeleteFile(string) error
	// ReadFile read file content from target filesystem
	ReadFile(filename string) ([]byte, error)
	// WriteFile writes content to file in target filesystem
	WriteFile(filename string, data []byte) error
}

// RemoteFile is structure to describe file in dropbox
type RemoteFile struct {
	Path string
	Hash string
}

// Dropbox is intereface to dropbox client wrapper
type Dropbox interface {
	// GetFiles returns files in application folder (include subfolder to)
	GetFiles() ([]RemoteFile, error)
	// DownloadFile downloaded file by path
	DownloadFile(string) (io.ReadCloser, error)
}
