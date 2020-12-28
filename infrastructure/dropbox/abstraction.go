package dropbox

import "github.com/tj/go-dropbox"

type DropboxFiles interface {
	ListFolder(in *dropbox.ListFolderInput) (out *dropbox.ListFolderOutput, err error)
	Download(in *dropbox.DownloadInput) (out *dropbox.DownloadOutput, err error)
}
