package dropbox

import (
	"io"
	"pb-dropbox-downloader/infrastructure"

	"github.com/tj/go-dropbox"
)

const rootDir = "/"

type Client struct {
	files DropboxFiles
}

func NewClient(db DropboxFiles) *Client {
	return &Client{db}
}

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
		mappedFiles = append(mappedFiles, infrastructure.RemoteFile{
			Path: entry.PathLower,
			Hash: entry.ContentHash,
		})
	}

	return mappedFiles, nil
}

func (client *Client) DownloadFile(path string) (io.ReadCloser, error) {
	out, err := client.files.Download(&dropbox.DownloadInput{
		Path: path,
	})

	if err != nil {
		return nil, err
	}

	return out.Body, nil
}
