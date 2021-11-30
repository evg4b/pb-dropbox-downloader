package internal

import (
	"io"
	"pb-dropbox-downloader/internal/dropbox"
)

// DataStorage interface to storage key-value data.
type DataStorage interface {
	Get(string) (string, error)
	ToMap() (map[string]string, error)
	FromMap(map[string]string) error
	KeyExists(string) bool
	Commit() error
	Add(string, string)
	Remove(string)
}

// Dropbox is intereface to dropbox client wrapper.
type Dropbox interface {
	// GetFiles returns files in application folder (include subfolder to).
	GetFiles() ([]dropbox.RemoteFile, error)
	// DownloadFile downloaded file by path.
	DownloadFile(string) (io.ReadCloser, error)
	// AccountDisplayName returns account display name.
	AccountDisplayName() string
	// AccountEmail returns account display email.
	AccountEmail() string
}
