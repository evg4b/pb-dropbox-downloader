package dropbox

import (
	"io"
	"path/filepath"
	"pb-dropbox-downloader/internal/utils"

	dropboxLib "github.com/tj/go-dropbox"
)

const rootDir = "/"

// Client is main structure to dropbox client wrapper.
type Client struct {
	files   dropboxFiles
	account *dropboxLib.GetCurrentAccountOutput
}

// NewClient creates new instance of dropbox client wrapper.
func NewClient(options ...dropboxOption) *Client {
	client := &Client{}
	for _, option := range options {
		option(client)
	}

	return client
}

// GetFiles returns files in application folder (include subfolder to).
func (c *Client) GetFiles() ([]RemoteFile, error) {
	output, err := c.files.ListFolder(&dropboxLib.ListFolderInput{
		Path:      rootDir,
		Recursive: true,
	})

	if err != nil {
		return nil, err
	}

	mappedFiles := []RemoteFile{}
	for _, entry := range output.Entries {
		if isFile(entry) {
			mappedFiles = append(mappedFiles, RemoteFile{
				Path: filepath.ToSlash(entry.PathLower[1:]),
				Hash: entry.ContentHash,
				Size: entry.Size,
			})
		}
	}

	return mappedFiles, nil
}

// DownloadFile downloaded file by path.
func (c *Client) DownloadFile(path string) (io.ReadCloser, error) {
	output, err := c.files.Download(&dropboxLib.DownloadInput{
		Path: utils.JoinPath(rootDir, path),
	})

	if err != nil {
		return nil, err
	}

	return output.Body, nil
}

// AccountDisplayName returns account display name.
func (c *Client) AccountDisplayName() string {
	return c.account.Name.DisplayName
}

// AccountEmail returns account email.
func (c *Client) AccountEmail() string {
	return c.account.Email
}

func isFile(metadata *dropboxLib.Metadata) bool {
	return metadata.Tag == "file"
}
