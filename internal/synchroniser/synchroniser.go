package synchroniser

import (
	"fmt"
	"io"
	"pb-dropbox-downloader/internal"

	"github.com/go-git/go-billy/v5"
)

// DropboxSynchroniser Dropbox data synchroniser app structure.
type DropboxSynchroniser struct {
	storage        internal.DataStorage
	files          billy.Filesystem
	dropbox        internal.Dropbox
	maxParallelism int
	output         io.Writer
}

// NewSynchroniser creates and initialize new instance of DropboxSynchroniser create.
func NewSynchroniser(options ...synchroniserOption) *DropboxSynchroniser {
	ds := &DropboxSynchroniser{maxParallelism: 1, output: io.Discard}

	for _, option := range options {
		option(ds)
	}

	return ds
}

func (db *DropboxSynchroniser) printf(format string, a ...interface{}) {
	fmt.Fprintln(db.output, fmt.Sprintf(format, a...))
}
