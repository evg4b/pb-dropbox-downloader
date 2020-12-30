package synchroniser_test

import (
	"io/ioutil"
	"os"
	"path"
	infs "pb-dropbox-downloader/infrastructure"
	sync "pb-dropbox-downloader/internal/synchroniser"
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
	fakeStorage := map[string]string{
		book1.Path: book1.Hash,
		book2.Path: book2.Hash,
		book5.Path: "other key",
	}
	fakeFromMock := func(data map[string]string) error {
		fakeStorage = data
		return nil
	}
	fakeExistMock := func(key string) bool {
		_, ok := fakeStorage[key]
		return ok
	}
	fakeGet := func(key string) (string, bool) {
		value, ok := fakeStorage[key]
		return value, ok
	}
	storageMock := mocks.NewDataStorageMock(t).
		GetMock.Set(fakeGet).
		ToMapMock.Return(fakeStorage, nil).
		FromMapMock.Set(fakeFromMock).
		KeyExistsMock.Set(fakeExistMock).
		CommitMock.Return(nil).
		AddMock.Return()

	dataReader1 := ioutil.NopCloser(strings.NewReader("This is book #3"))
	dataReader2 := ioutil.NopCloser(strings.NewReader("This is book #5"))
	dropboxMocks := mocks.NewDropboxMock(t).
		GetFilesMock.Return([]infs.RemoteFile{book1, book3, book5}, nil).
		DownloadFileMock.When(book3.Path).Then(dataReader1, nil).
		DownloadFileMock.When(book5.Path).Then(dataReader2, nil)

	filesMock := mocks.NewFileSystemMock(t).
		GetFilesInFolderMock.Return([]string{book1.Path, book2.Path, book3.Path, book4.Path}).
		CopyDataToFileMock.When(path.Join(folder, book3.Path), dataReader1).Then(nil).
		CopyDataToFileMock.When(path.Join(folder, book5.Path), dataReader2).Then(nil).
		DeleteFileMock.When(path.Join(folder, book1.Path)).Then(nil).
		DeleteFileMock.When(path.Join(folder, book2.Path)).Then(nil)

	synchroniser := sync.NewSynchroniser(storageMock, filesMock, dropboxMocks, os.Stdout)

	err := synchroniser.Sync(folder, true)

	assert.NoError(t, err)
}

func TestDropboxSynchroniser_Sync_WithoutDelete(t *testing.T) {
	folder := "/mnt/ext1/dropbox"
	fakeStorage := map[string]string{
		book1.Path: book1.Hash,
		book2.Path: book2.Hash,
		book5.Path: "other key",
	}
	fakeFromMock := func(data map[string]string) error {
		fakeStorage = data
		return nil
	}
	fakeExistMock := func(key string) bool {
		_, ok := fakeStorage[key]
		return ok
	}
	fakeGet := func(key string) (string, bool) {
		value, ok := fakeStorage[key]
		return value, ok
	}
	storageMock := mocks.NewDataStorageMock(t).
		GetMock.Set(fakeGet).
		ToMapMock.Return(fakeStorage, nil).
		FromMapMock.Set(fakeFromMock).
		KeyExistsMock.Set(fakeExistMock).
		CommitMock.Return(nil).
		AddMock.Return()

	dataReader1 := ioutil.NopCloser(strings.NewReader("This is book #3"))
	dataReader2 := ioutil.NopCloser(strings.NewReader("This is book #5"))
	dropboxMocks := mocks.NewDropboxMock(t).
		GetFilesMock.Return([]infs.RemoteFile{book1, book3, book5}, nil).
		DownloadFileMock.When(book3.Path).Then(dataReader1, nil).
		DownloadFileMock.When(book5.Path).Then(dataReader2, nil)

	filesMock := mocks.NewFileSystemMock(t).
		GetFilesInFolderMock.Return([]string{book1.Path, book2.Path, book3.Path, book4.Path}).
		CopyDataToFileMock.When(path.Join(folder, book3.Path), dataReader1).Then(nil).
		CopyDataToFileMock.When(path.Join(folder, book5.Path), dataReader2).Then(nil)

	synchroniser := sync.NewSynchroniser(storageMock, filesMock, dropboxMocks, os.Stdout)

	err := synchroniser.Sync(folder, false)

	assert.NoError(t, err)
}
