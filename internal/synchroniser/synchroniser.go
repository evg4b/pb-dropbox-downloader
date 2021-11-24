package synchroniser

import (
	"fmt"
	"io"
	"pb-dropbox-downloader/infrastructure"
	"pb-dropbox-downloader/internal"
)

// DropboxSynchroniser Dropbox data synchroniser app structure.
type DropboxSynchroniser struct {
	storage        internal.DataStorage
	files          infrastructure.FileSystem
	dropbox        infrastructure.Dropbox
	maxParallelism int
	output         io.Writer
}

// NewSynchroniser creates and initialize new instance of DropboxSynchroniser create.
func NewSynchroniser(
	storage internal.DataStorage,
	files infrastructure.FileSystem,
	dropbox infrastructure.Dropbox,
	output io.Writer,
	maxParallelism int,
) *DropboxSynchroniser {
	return &DropboxSynchroniser{
		storage:        storage,
		files:          files,
		dropbox:        dropbox,
		maxParallelism: maxParallelism,
		output:         output,
	}
}

func (db *DropboxSynchroniser) printf(format string, a ...interface{}) {
	fmt.Fprintln(db.output, fmt.Sprintf(format, a...))
}
