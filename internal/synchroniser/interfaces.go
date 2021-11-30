package synchroniser

import (
	"io"
	"pb-dropbox-downloader/internal/dropbox"
)

// DataStorage interface to storage key-value data.
type DataStorage interface {
	Get(string) (string, error)
	ToMap() (map[string]string, error)
	FromMap(map[string]string)
	KeyExists(string) (bool, error)
	Commit() error
	Add(string, string) error
	Remove(string) error
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
