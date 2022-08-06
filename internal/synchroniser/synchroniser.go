package synchroniser

import (
	"fmt"
	"io"

	"github.com/go-git/go-billy/v5"
)

// DropboxSynchroniser Dropbox data synchroniser app structure.
type DropboxSynchroniser struct {
	storage        DataStorage
	files          billy.Filesystem
	dropbox        Dropbox
	maxParallelism int
	output         io.Writer
	version        string
}

// NewSynchroniser creates and initialize new instance of DropboxSynchroniser create.
func NewSynchroniser(options ...synchroniserOption) *DropboxSynchroniser {
	ds := &DropboxSynchroniser{maxParallelism: 1, output: io.Discard}

	for _, option := range options {
		option(ds)
	}

	return ds
}

func (ds *DropboxSynchroniser) printf(format string, a ...interface{}) {
	fmt.Fprintln(ds.output, fmt.Sprintf(format, a...))
}
