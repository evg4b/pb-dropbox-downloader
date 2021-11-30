package dropbox

import "github.com/tj/go-dropbox"

type dropboxFiles interface {
	ListFolder(in *dropbox.ListFolderInput) (out *dropbox.ListFolderOutput, err error)
	Download(in *dropbox.DownloadInput) (out *dropbox.DownloadOutput, err error)
}

// RemoteFile is structure to describe file in dropbox.
type RemoteFile struct {
	Path string
	Hash string
	Size uint64
}
