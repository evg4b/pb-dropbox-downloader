package dropbox

import (
	"io"
	"path/filepath"
	"pb-dropbox-downloader/infrastructure"
	"pb-dropbox-downloader/utils"

	"github.com/tj/go-dropbox"
)

const rootDir = "/"

// Client is main structure to dropbox client wrapper.
type Client struct {
	files dropboxFiles
}

// NewClient creates new instance of dropbox client wrapper.
func NewClient(files dropboxFiles) *Client {
	return &Client{files}
}

// GetFiles returns files in application folder (include subfolder to).
func (client *Client) GetFiles() ([]infrastructure.RemoteFile, error) {
	out, err := client.files.ListFolder(&dropbox.ListFolderInput{
		Path:      rootDir,
		Recursive: true,
	})
	if err != nil {
		return nil, err
	}

	mappedFiles := []infrastructure.RemoteFile{}
	for _, entry := range out.Entries {
		if isFile(entry) {
			mappedFiles = append(mappedFiles, infrastructure.RemoteFile{
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
	out, err := client.files.Download(&dropbox.DownloadInput{
		Path: utils.JoinPath(rootDir, path),
	})

	if err != nil {
		return nil, err
	}

	return out.Body, nil
}

func isFile(metadata *dropbox.Metadata) bool {
	return metadata.Tag == "file"
}
