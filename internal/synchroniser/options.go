package synchroniser

import (
	"io"
	"pb-dropbox-downloader/internal"

	"github.com/go-git/go-billy/v5"
)

type synchroniserOption = func(sync *DropboxSynchroniser)

func WithStorage(storage internal.DataStorage) synchroniserOption {
	return func(ds *DropboxSynchroniser) {
		ds.storage = storage
	}
}

func WithFileSystem(files billy.Filesystem) synchroniserOption {
	return func(ds *DropboxSynchroniser) {
		ds.files = files
	}
}

func WithDropboxClient(dropbox internal.Dropbox) synchroniserOption {
	return func(ds *DropboxSynchroniser) {
		ds.dropbox = dropbox
	}
}

func WithOutput(output io.Writer) synchroniserOption {
	return func(ds *DropboxSynchroniser) {
		ds.output = output
	}
}

func WithMaxParallelism(maxParallelism int) synchroniserOption {
	return func(ds *DropboxSynchroniser) {
		ds.maxParallelism = maxParallelism
	}
}
