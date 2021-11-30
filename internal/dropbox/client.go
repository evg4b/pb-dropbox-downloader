package dropbox

import (
	"io"
	"path/filepath"
	"pb-dropbox-downloader/internal"
	"pb-dropbox-downloader/utils"

	dropboxLib "github.com/tj/go-dropbox"
)

const rootDir = "/"

// Client is main structure to dropbox client wrapper.
type Client struct {
	files   dropboxFiles
	account *dropboxLib.GetCurrentAccountOutput
}

// NewClient creates new instance of dropbox client wrapper.
func NewClient(files dropboxFiles, account *dropboxLib.GetCurrentAccountOutput) *Client {
	return &Client{files, account}
}

// GetFiles returns files in application folder (include subfolder to).
func (client *Client) GetFiles() ([]internal.RemoteFile, error) {
	out, err := client.files.ListFolder(&dropboxLib.ListFolderInput{
		Path:      rootDir,
		Recursive: true,
	})
	if err != nil {
		return nil, err
	}

	mappedFiles := []internal.RemoteFile{}
	for _, entry := range out.Entries {
		if isFile(entry) {
			mappedFiles = append(mappedFiles, internal.RemoteFile{
				Path: filepath.ToSlash(entry.PathLower[1:]),
				Hash: entry.ContentHash,
				Size: entry.Size,
			})
		}
	}

	return mappedFiles, nil
}

// DownloadFile downloaded file by path.
func (client *Client) DownloadFile(path string) (io.ReadCloser, error) {
	out, err := client.files.Download(&dropboxLib.DownloadInput{
		Path: utils.JoinPath(rootDir, path),
	})

	if err != nil {
		return nil, err
	}

	return out.Body, nil
}

// AccountDisplayName returns account display name.
func (client *Client) AccountDisplayName() string {
	return client.account.Name.DisplayName
}

// AccountEmail returns account display email.
func (client *Client) AccountEmail() string {
	return client.account.Email
}

func isFile(metadata *dropboxLib.Metadata) bool {
	return metadata.Tag == "file"
}
