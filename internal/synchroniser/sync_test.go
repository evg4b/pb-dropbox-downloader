// nolint: dupl

package synchroniser_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/fs"
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
	type testCase struct {
		name        string
		remove      bool
		storage     synchroniser.DataStorage
		dropbox     synchroniser.Dropbox
		fs          billy.Filesystem
		expectedErr string
		expected    map[string]string
	}

	callCount := 0
	tests := []testCase{
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
				ReadDirMock.Return(nil, errors.New("fs error")).
				MkdirAllMock.Return(nil),
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
			name: "download new files",
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
		{
			name: "download changed files",
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
				DownloadFileMock.When("demo1").Then(readCloser("new content"), nil).
				DownloadFileMock.When("demo2").Then(readCloser("test file"), nil).
				GetFilesMock.Return(
				[]dropbox.RemoteFile{
					{Path: "demo1", Hash: "0003", Size: 2132},
					{Path: "demo2", Hash: "0002", Size: 2132},
				},
				nil,
			),
			fs: testutils.FsFromMap(t, map[string]string{
				"./dropbox/demo1": "content",
			}),
			expected: map[string]string{
				"./dropbox/demo1": "new content",
				"./dropbox/demo2": "test file",
			},
		},
		{
			name: "no files to sync",
			storage: mocks.NewDataStorageMock(t).
				ToMapMock.Return(map[string]string{"demo1": "0001", "demo2": "0002"}, nil).
				GetMock.When("demo1").Then("0001", nil).
				GetMock.When("demo2").Then("0002", nil).
				AddMock.Return(nil).
				FromMapMock.Return().
				CommitMock.Return(nil),
			dropbox: mocks.NewDropboxMock(t).
				AccountEmailMock.Return("test").
				AccountDisplayNameMock.Return("test").
				GetFilesMock.Return(
				[]dropbox.RemoteFile{
					{Path: "demo1", Hash: "0001", Size: 2132},
					{Path: "demo2", Hash: "0002", Size: 2132},
				},
				nil,
			),
			fs: testutils.FsFromMap(t, map[string]string{
				"./dropbox/demo1": "content1",
				"./dropbox/demo2": "content2",
			}),
			expected: map[string]string{
				"./dropbox/demo1": "content1",
				"./dropbox/demo2": "content2",
			},
		},
		{
			name: "drop box get files error",
			storage: mocks.NewDataStorageMock(t).
				ToMapMock.Return(map[string]string{"demo1": "0001", "demo2": "0002"}, nil).
				FromMapMock.Return().
				CommitMock.Return(nil),
			dropbox: mocks.NewDropboxMock(t).
				AccountEmailMock.Return("test").
				AccountDisplayNameMock.Return("test").
				GetFilesMock.Return(nil, errors.New("drop box error")),
			fs:          memfs.New(),
			expectedErr: "drop box error",
		},
		{
			name: "download file error",
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
				DownloadFileMock.When("demo2").Then(nil, errors.New("filed to download demo2")).
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
			expectedErr: "1 error occurred:\n\t* filed to download demo2\n\n",
		},
		{
			name: "download file error",
			storage: mocks.NewDataStorageMock(t).
				ToMapMock.Return(map[string]string{"demo1": "0001"}, nil).
				GetMock.When("demo1").Then("0001", nil).
				GetMock.When("demo2").Then("", datastorage.ErrKeyDoesNotExists).
				AddMock.Return(errors.New("failed to add data")).
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
			expectedErr: "1 error occurred:\n\t* failed to add data\n\n",
		},
		{
			name: "download file reading error",
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
				DownloadFileMock.When("demo2").Then(testutils.ErrorCloser(errors.New("io.Reader error")), nil).
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
			expectedErr: "1 error occurred:\n\t* io.Reader error\n\n",
		},
		{
			name: "download file writing error",
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
			fs: mocks.NewFilesystemMock(t).
				ReadDirMock.Return([]fs.FileInfo{}, nil).
				CreateMock.Return(nil, errors.New("file system write error")).
				MkdirAllMock.Return(nil),
			expectedErr: "1 error occurred:\n\t* file system write error\n\n",
		},
		{
			name: "committing to datastorage error",
			storage: mocks.NewDataStorageMock(t).
				ToMapMock.Return(map[string]string{"demo1": "0001"}, nil).
				GetMock.When("demo1").Then("0001", nil).
				GetMock.When("demo2").Then("", datastorage.ErrKeyDoesNotExists).
				AddMock.Return(nil).
				FromMapMock.Return().
				CommitMock.Set(func() (err error) {
				if callCount > 0 {
					return errors.New("committing error")
				}

				callCount++

				return nil
			}),
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
			expectedErr: "committing error",
		},
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
			fs:     memfs.New(),
			remove: true,
		},
		(func() testCase {
			fs := testutils.FsFromMap(t, map[string]string{
				"./dropbox/file1": "content",
				"./dropbox/file2": "content",
				"./dropbox/file3": "content",
			})

			return testCase{
				name: "files deleted successful",
				storage: datastorage.NewFileStorage(
					datastorage.WithFilesystem(fs),
				),
				dropbox: mocks.NewDropboxMock(t).
					AccountEmailMock.Return("test").
					AccountDisplayNameMock.Return("test").
					GetFilesMock.Return([]dropbox.RemoteFile{}, nil),
				fs:       fs,
				remove:   true,
				expected: map[string]string{},
			}
		}()),
		{
			name: "delete failed by storage error",
			storage: mocks.NewDataStorageMock(t).
				ToMapMock.Return(map[string]string{}, nil).
				FromMapMock.Return().
				CommitMock.Return(nil).
				KeyExistsMock.Return(false, errors.New("key reading error")),
			dropbox: mocks.NewDropboxMock(t).
				AccountEmailMock.Return("test").
				AccountDisplayNameMock.Return("test").
				GetFilesMock.Return([]dropbox.RemoteFile{}, nil),
			fs: testutils.FsFromMap(t, map[string]string{
				"./dropbox/file1": "content",
				"./dropbox/file2": "content",
				"./dropbox/file3": "content",
			}),
			remove:      true,
			expectedErr: "key reading error",
		},
		{
			name: "delete failed by fs error",
			storage: mocks.NewDataStorageMock(t).
				ToMapMock.Return(map[string]string{}, nil).
				FromMapMock.Return().
				CommitMock.Return(nil).
				KeyExistsMock.Return(false, nil),
			dropbox: mocks.NewDropboxMock(t).
				AccountEmailMock.Return("test").
				AccountDisplayNameMock.Return("test").
				GetFilesMock.Return([]dropbox.RemoteFile{}, nil),
			fs: mocks.NewFilesystemMock(t).
				ReadDirMock.Return([]fs.FileInfo{
				mocks.NewFileInfo(t).NameMock.Return("file1"),
				mocks.NewFileInfo(t).NameMock.Return("file2"),
				mocks.NewFileInfo(t).NameMock.Return("file3"),
			}, nil).
				RemoveMock.Return(errors.New("remove error")).
				MkdirAllMock.Return(nil),
			remove:      true,
			expectedErr: "remove error",
		},
		{
			name: "delete failed by fs error",
			storage: mocks.NewDataStorageMock(t).
				ToMapMock.Return(map[string]string{}, nil).
				FromMapMock.Return().
				CommitMock.Return(nil).
				KeyExistsMock.Return(false, nil).
				RemoveMock.Return(errors.New("remove key error")),
			dropbox: mocks.NewDropboxMock(t).
				AccountEmailMock.Return("test").
				AccountDisplayNameMock.Return("test").
				GetFilesMock.Return([]dropbox.RemoteFile{}, nil),
			fs: mocks.NewFilesystemMock(t).
				ReadDirMock.Return([]fs.FileInfo{
				mocks.NewFileInfo(t).NameMock.Return("file1"),
				mocks.NewFileInfo(t).NameMock.Return("file2"),
				mocks.NewFileInfo(t).NameMock.Return("file3"),
			}, nil).
				RemoveMock.Return(nil).
				MkdirAllMock.Return(nil),
			remove:      true,
			expectedErr: "remove key error",
		},
		(func() testCase {
			fs := testutils.FsFromMap(t, map[string]string{
				"./dropbox/lorem-file1": "lorem-file1",
				"./dropbox/lorem-file2": "lorem-file2",
				"./dropbox/lorem-file3": "lorem-file3",
				"./dropbox/lorem-file4": "lorem-file4",
				"./dropbox/lorem-file5": "lorem-file5",
				"./dropbox/lorem-file6": "lorem-file6",
			})

			return testCase{
				name: "download and delete a lot of files",
				storage: datastorage.NewFileStorage(
					datastorage.WithFilesystem(fs),
				),
				dropbox: mocks.NewDropboxMock(t).
					AccountEmailMock.Return("test").
					AccountDisplayNameMock.Return("test").
					GetFilesMock.Return([]dropbox.RemoteFile{
					{Path: "file1", Hash: "0001", Size: 1001},
					{Path: "file2", Hash: "0002", Size: 1002},
					{Path: "file3", Hash: "0003", Size: 1003},
					{Path: "file4", Hash: "0004", Size: 1004},
					{Path: "file5", Hash: "0005", Size: 1005},
					{Path: "file6", Hash: "0006", Size: 1006},
					{Path: "file7", Hash: "0007", Size: 1007},
				}, nil).
					DownloadFileMock.When("file1").Then(readCloser("content1"), nil).
					DownloadFileMock.When("file2").Then(readCloser("content2"), nil).
					DownloadFileMock.When("file3").Then(readCloser("content3"), nil).
					DownloadFileMock.When("file4").Then(readCloser("content4"), nil).
					DownloadFileMock.When("file5").Then(readCloser("content5"), nil).
					DownloadFileMock.When("file6").Then(readCloser("content6"), nil).
					DownloadFileMock.When("file7").Then(readCloser("content7"), nil),
				fs:     fs,
				remove: true,
				expected: map[string]string{
					"./dropbox/file1": "content1",
					"./dropbox/file2": "content2",
					"./dropbox/file3": "content3",
					"./dropbox/file4": "content4",
					"./dropbox/file5": "content5",
					"./dropbox/file6": "content6",
					"./dropbox/file7": "content7",
				},
			}
		}()),
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

			err := synchroniser.Sync(context.TODO(), "./dropbox", tt.remove)

			testutils.AssertError(t, tt.expectedErr, err)
			testutils.AssertFiles(t, tt.fs, "./dropbox", tt.expected)
		})
	}
}

func readCloser(content string) io.ReadCloser {
	buff := bytes.NewBufferString(content)

	return testutils.NopCloser(buff)
}
