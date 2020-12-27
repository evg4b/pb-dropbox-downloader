package infrastructure

import "io"

// FileSystem interface to comunitace with file system
type FileSystem interface {
	// GetFiles return relative file paths in folder (include subfolder to)
	GetFilesInFolder(string) []string
	// CopyDataToFile copy data from reader to file
	CopyDataToFile(string, io.Reader) error
	DeleteFile(string) error
}

type RemoteFile struct {
	Path string
	Hash string
}

type Dropbox interface {
	// GetFiles return relative file paths in folder (include subfolder to)
	GetFiles() []RemoteFile
	DownloadFile(string) (io.Reader, error)
}
