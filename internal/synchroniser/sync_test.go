package synchroniser_test

import (
	"io/ioutil"
	"path/filepath"
	infs "pb-dropbox-downloader/infrastructure"
	"pb-dropbox-downloader/infrastructure/filesystem"
	"pb-dropbox-downloader/internal/datastorage"
	sync "pb-dropbox-downloader/internal/synchroniser"
	"pb-dropbox-downloader/utils"
	"strings"
	"testing"

	"pb-dropbox-downloader/mocks"

	"github.com/stretchr/testify/assert"
)

var book1 = infs.RemoteFile{Path: "book1.epub", Hash: "00001"}
var book2 = infs.RemoteFile{Path: "book2.epub", Hash: "00002"}
var book3 = infs.RemoteFile{Path: "book3.epub", Hash: "00003"}
var book4 = infs.RemoteFile{Path: "book4.epub", Hash: "00004"}
var book5 = infs.RemoteFile{Path: "book5.epub", Hash: "00005"}

func TestDropboxSynchroniser_Sync(t *testing.T) {
	folder := "/mnt/ext1/dropbox"

	storage := datastorage.NewFileStorage(filesystem.NewFileSystem(), filepath.Join(t.TempDir(), "test.bin"))

	dataReader1 := ioutil.NopCloser(strings.NewReader("This is book #1"))
	dataReader3 := ioutil.NopCloser(strings.NewReader("This is book #3"))
	dataReader5 := ioutil.NopCloser(strings.NewReader("This is book #5"))
	dropboxMocks := mocks.NewDropboxMock(t).
		GetFilesMock.Return([]infs.RemoteFile{book1, book3, book5}, nil).
		DownloadFileMock.When(book3.Path).Then(dataReader3, nil).
		DownloadFileMock.When(book1.Path).Then(dataReader1, nil).
		DownloadFileMock.When(book5.Path).Then(dataReader5, nil)

	filesMock := mocks.NewFileSystemMock(t).
		GetFilesInFolderMock.Return([]string{book1.Path, book2.Path, book3.Path, book4.Path}).
		CopyDataToFileMock.When(utils.JoinPath(folder, book1.Path), dataReader1).Then(nil).
		CopyDataToFileMock.When(utils.JoinPath(folder, book3.Path), dataReader3).Then(nil).
		CopyDataToFileMock.When(utils.JoinPath(folder, book5.Path), dataReader5).Then(nil).
		DeleteFileMock.When(utils.JoinPath(folder, book1.Path)).Then(nil).
		DeleteFileMock.When(utils.JoinPath(folder, book2.Path)).Then(nil).
		DeleteFileMock.Expect(utils.JoinPath(folder, book4.Path)).Return(nil)

	synchroniser := sync.NewSynchroniser(storage, filesMock, dropboxMocks, ioutil.Discard, 3)

	err := synchroniser.Sync(folder, true)

	assert.NoError(t, err)
}

func TestDropboxSynchroniser_Sync_WithoutDelete(t *testing.T) {
	folder := "/mnt/ext1/dropbox"

	storage := datastorage.NewFileStorage(filesystem.NewFileSystem(), filepath.Join(t.TempDir(), "test.bin"))

	dataReader1 := ioutil.NopCloser(strings.NewReader("This is book #1"))
	dataReader3 := ioutil.NopCloser(strings.NewReader("This is book #3"))
	dataReader5 := ioutil.NopCloser(strings.NewReader("This is book #5"))
	dropboxMocks := mocks.NewDropboxMock(t).
		GetFilesMock.Return([]infs.RemoteFile{book1, book3, book5}, nil).
		DownloadFileMock.When(book1.Path).Then(dataReader1, nil).
		DownloadFileMock.When(book3.Path).Then(dataReader3, nil).
		DownloadFileMock.When(book5.Path).Then(dataReader5, nil)

	filesMock := mocks.NewFileSystemMock(t).
		GetFilesInFolderMock.Return([]string{book1.Path, book2.Path, book3.Path, book4.Path}).
		CopyDataToFileMock.When(utils.JoinPath(folder, book1.Path), dataReader1).Then(nil).
		CopyDataToFileMock.When(utils.JoinPath(folder, book3.Path), dataReader3).Then(nil).
		CopyDataToFileMock.When(utils.JoinPath(folder, book5.Path), dataReader5).Then(nil)

	synchroniser := sync.NewSynchroniser(storage, filesMock, dropboxMocks, ioutil.Discard, 3)

	err := synchroniser.Sync(folder, false)

	assert.NoError(t, err)
}
