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

type RemoteFile struct {
	Path string
	Hash string
}

type Dropbox interface {
	// GetFiles return relative file paths in folder (include subfolder to)
	GetFiles() ([]RemoteFile, error)
	DownloadFile(string) (io.ReadCloser, error)
}
