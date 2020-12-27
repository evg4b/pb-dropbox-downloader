package synchroniser_test

import (
	infs "pb-dropbox-downloader/infrastructure"
	sync "pb-dropbox-downloader/internal/synchroniser"
	"testing"

	"pb-dropbox-downloader/mocks"

	"github.com/stretchr/testify/assert"
)

var book1 = infs.RemoteFile{
	Path: "book1.epub",
	Hash: "00000000000000000000000000000001",
}

var book2 = infs.RemoteFile{
	Path: "book2.epub",
	Hash: "00000000000000000000000000000002",
}

var book3 = infs.RemoteFile{
	Path: "book3.epub",
	Hash: "00000000000000000000000000000003",
}

var book4 = infs.RemoteFile{
	Path: "book4.epub",
	Hash: "00000000000000000000000000000004",
}

func TestDropboxSynchroniser_Sync(t *testing.T) {
	storageMock := mocks.NewDataStorageMock(t).
		GetMock.When(book1.Path).Then(book1.Hash, true).
		GetMock.When(book2.Path).Then(book2.Hash, true).
		GetMock.When(book3.Path).Then("", true).
		GetMock.When(book4.Path).Then("", true).
		ToMapMock.Return(map[string]string{book1.Path: book1.Hash, book2.Path: book2.Hash}, nil).
		FromMapMock.Return(nil)

	filesMock := mocks.NewFileSystemMock(t).
		GetFilesInFolderMock.Return([]string{book1.Path, book2.Path, book3.Path, book4.Path})

	dropboxMocks := mocks.NewDropboxMock(t).
		GetFilesMock.Return([]infs.RemoteFile{book1, book3})

	synchroniser := sync.NewSynchroniser(storageMock, filesMock, dropboxMocks)

	err := synchroniser.Sync("/mnt/ext1/dropbox", false)

	assert.NoError(t, err)
}
