// nolint: dupl

package synchroniser_test

import (
	"bytes"
	"errors"
	"io"
	"pb-dropbox-downloader/internal/datastorage"
	"pb-dropbox-downloader/internal/dropbox"
	"pb-dropbox-downloader/internal/synchroniser"
	"pb-dropbox-downloader/testing/mocks"
	"pb-dropbox-downloader/testing/testutils"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
)

func TestDropboxSynchroniser_Sync(t *testing.T) {
	tests := []struct {
		name        string
		remove      bool
		storage     synchroniser.DataStorage
		dropbox     synchroniser.Dropbox
		fs          billy.Filesystem
		expectedErr string
		expected    map[string]string
	}{
		{
			name: "no files to sync",
			storage: mocks.NewDataStorageMock(t).
				ToMapMock.Return(map[string]string{}, nil).
				FromMapMock.Return().
				CommitMock.Return(nil),
			dropbox: mocks.NewDropboxMock(t).
				AccountEmailMock.Return("test").
				AccountDisplayNameMock.Return("test").
				GetFilesMock.Return([]dropbox.RemoteFile{}, nil),
			fs: memfs.New(),
		},
		{
			name: "files system error",
			storage: mocks.NewDataStorageMock(t).
				ToMapMock.Return(map[string]string{}, nil).
				FromMapMock.Return().
				CommitMock.Return(nil),
			dropbox: mocks.NewDropboxMock(t).
				AccountEmailMock.Return("test").
				AccountDisplayNameMock.Return("test").
				GetFilesMock.Return([]dropbox.RemoteFile{}, nil),
			fs: mocks.NewFilesystemMock(t).
				ReadDirMock.Return(nil, errors.New("fs error")),
			expectedErr: "fs error",
		},
		{
			name: "storage error",
			storage: mocks.NewDataStorageMock(t).
				ToMapMock.Return(map[string]string{}, errors.New("storage error")).
				FromMapMock.Return().
				CommitMock.Return(nil),
			dropbox: mocks.NewDropboxMock(t).
				AccountEmailMock.Return("test").
				AccountDisplayNameMock.Return("test").
				GetFilesMock.Return([]dropbox.RemoteFile{}, nil),
			fs:          memfs.New(),
			expectedErr: "storage error",
		},
		{
			name: "storage error",
			storage: mocks.NewDataStorageMock(t).
				ToMapMock.Return(map[string]string{"demo1": "0001"}, nil).
				GetMock.When("demo1").Then("0001", nil).
				GetMock.When("demo2").Then("", datastorage.ErrKeyDoesNotExists).
				AddMock.Return(nil).
				FromMapMock.Return().
				CommitMock.Return(nil),
			dropbox: mocks.NewDropboxMock(t).
				AccountEmailMock.Return("test").
				AccountDisplayNameMock.Return("test").
				DownloadFileMock.When("demo2").Then(readCloser("test file"), nil).
				GetFilesMock.Return(
				[]dropbox.RemoteFile{
					{Path: "demo1", Hash: "0001", Size: 2132},
					{Path: "demo2", Hash: "0002", Size: 2132},
				},
				nil,
			),
			fs: testutils.FsFromMap(t, map[string]string{
				"./dropbox/demo1": "content",
			}),
			expected: map[string]string{
				"./dropbox/demo1": "content",
				"./dropbox/demo2": "test file",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synchroniser := synchroniser.NewSynchroniser(
				synchroniser.WithStorage(tt.storage),
				synchroniser.WithFileSystem(tt.fs),
				synchroniser.WithDropboxClient(tt.dropbox),
				synchroniser.WithOutput(io.Discard),
				synchroniser.WithMaxParallelism(3),
			)

			err := synchroniser.Sync("./dropbox", tt.remove)

			testutils.AssertError(t, tt.expectedErr, err)
			testutils.AssertFiles(t, tt.fs, "./dropbox", tt.expected)
		})
	}
}

func readCloser(content string) io.ReadCloser {
	buff := bytes.NewBufferString(content)

	return testutils.NopCloser(buff)
}
