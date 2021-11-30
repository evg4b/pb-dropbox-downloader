// nolint: dupl
package synchroniser_test

import (
	"io/ioutil"
	"path/filepath"
	"pb-dropbox-downloader/internal/datastorage"
	"pb-dropbox-downloader/internal/dropbox"
	sync "pb-dropbox-downloader/internal/synchroniser"
	"pb-dropbox-downloader/testing/testutils"
	"pb-dropbox-downloader/utils"
	"strings"
	"testing"

	"pb-dropbox-downloader/testing/mocks"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/assert"
)

var book1 = dropbox.RemoteFile{Path: "book1.epub", Hash: "00001"}
var book3 = dropbox.RemoteFile{Path: "book3.epub", Hash: "00003"}
var book5 = dropbox.RemoteFile{Path: "book5.epub", Hash: "00005"}

func TestDropboxSynchroniser_Sync(t *testing.T) {
	folder := "/mnt/ext1/dropbox"

	fs := memfs.New()

	storage := datastorage.NewFileStorage(fs, filepath.Join(t.TempDir(), "test.bin"))

	dataReader1 := ioutil.NopCloser(strings.NewReader("This is book #1"))
	dataReader3 := ioutil.NopCloser(strings.NewReader("This is book #3"))
	dataReader5 := ioutil.NopCloser(strings.NewReader("This is book #5"))

	dropboxMocks := mocks.NewDropboxMock(t).
		GetFilesMock.Return([]dropbox.RemoteFile{book1, book3, book5}, nil).
		DownloadFileMock.When(book3.Path).Then(dataReader3, nil).
		DownloadFileMock.When(book1.Path).Then(dataReader1, nil).
		DownloadFileMock.When(book5.Path).Then(dataReader5, nil).
		AccountDisplayNameMock.Return("DisplayName").
		AccountEmailMock.Return("email@mail.com")
	testutils.MakeFiles(t, fs, map[string]string{
		utils.JoinPath(folder, book1.Path): "This is book #1",
		utils.JoinPath(folder, book3.Path): "This is book #3",
		utils.JoinPath(folder, book5.Path): "This is book #5",
	})

	synchroniser := sync.NewSynchroniser(
		sync.WithStorage(storage),
		sync.WithFileSystem(fs),
		sync.WithDropboxClient(dropboxMocks),
		sync.WithMaxParallelism(3),
	)

	err := synchroniser.Sync(folder, true)

	assert.NoError(t, err)
}

func TestDropboxSynchroniser_Sync_WithoutDelete(t *testing.T) {
	folder := "/mnt/ext1/dropbox"

	fs := memfs.New()

	storage := datastorage.NewFileStorage(fs, filepath.Join(t.TempDir(), "test.bin"))

	dataReader1 := ioutil.NopCloser(strings.NewReader("This is book #1"))
	dataReader3 := ioutil.NopCloser(strings.NewReader("This is book #3"))
	dataReader5 := ioutil.NopCloser(strings.NewReader("This is book #5"))
	dropboxMocks := mocks.NewDropboxMock(t).
		GetFilesMock.Return([]dropbox.RemoteFile{book1, book3, book5}, nil).
		DownloadFileMock.When(book1.Path).Then(dataReader1, nil).
		DownloadFileMock.When(book3.Path).Then(dataReader3, nil).
		DownloadFileMock.When(book5.Path).Then(dataReader5, nil).
		AccountDisplayNameMock.Return("DisplayName").
		AccountEmailMock.Return("email@mail.com")

	testutils.MakeFiles(t, fs, map[string]string{
		utils.JoinPath(folder, book1.Path): "This is book #1",
		utils.JoinPath(folder, book3.Path): "This is book #3",
		utils.JoinPath(folder, book5.Path): "This is book #5",
	})

	synchroniser := sync.NewSynchroniser(
		sync.WithStorage(storage),
		sync.WithFileSystem(fs),
		sync.WithDropboxClient(dropboxMocks),
		sync.WithMaxParallelism(3),
	)

	err := synchroniser.Sync(folder, false)

	assert.NoError(t, err)
}
