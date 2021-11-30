package dropbox_test

import (
	"errors"
	"io/ioutil"
	"pb-dropbox-downloader/internal/dropbox"
	"pb-dropbox-downloader/testing/mocks"
	"pb-dropbox-downloader/utils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	dropboxLib "github.com/tj/go-dropbox"
)

func TestClient_GetFiles(t *testing.T) {
	input := dropboxLib.ListFolderInput{Path: "/", Recursive: true}
	output := dropboxLib.ListFolderOutput{
		HasMore: true,
		Cursor:  "",
		Entries: []*dropboxLib.Metadata{
			{PathLower: "/book1.epub", ContentHash: "00001", Tag: "file"},
			{PathLower: "/book2.epub", ContentHash: "00002", Tag: "file"},
			{PathLower: "/data", ContentHash: "", Tag: "folder"},
		},
	}
	filesMock := mocks.NewDropboxFilesMock(t).
		ListFolderMock.Expect(&input).Return(&output, nil)
	client := dropbox.NewClient(dropbox.WithFiles(filesMock))

	files, err := client.GetFiles()

	assert.NoError(t, err)
	assert.Equal(t, []dropbox.RemoteFile{
		{Path: "book1.epub", Hash: "00001"},
		{Path: "book2.epub", Hash: "00002"},
	}, files)
}

func TestClient_GetFiles_Error(t *testing.T) {
	expectedError := errors.New("test error")
	input := dropboxLib.ListFolderInput{Path: "/", Recursive: true}
	filesMock := mocks.NewDropboxFilesMock(t).
		ListFolderMock.Expect(&input).Return(nil, expectedError)
	client := dropbox.NewClient(dropbox.WithFiles(filesMock))

	files, err := client.GetFiles()

	assert.Equal(t, expectedError, err)
	assert.Nil(t, files)
}

func TestClient_DownloadFile(t *testing.T) {
	file := "book1.epub"
	input := dropboxLib.DownloadInput{Path: utils.JoinPath("/", file)}
	expectedReader := ioutil.NopCloser(strings.NewReader(""))
	output := dropboxLib.DownloadOutput{
		Body:   expectedReader,
		Length: 10,
	}
	filesMock := mocks.NewDropboxFilesMock(t).
		DownloadMock.Expect(&input).Return(&output, nil)

	client := dropbox.NewClient(dropbox.WithFiles(filesMock))

	reader, err := client.DownloadFile(file)

	assert.NoError(t, err)
	assert.Equal(t, expectedReader, reader)
}

func TestClient_DownloadFile_Error(t *testing.T) {
	file := "book1.epub"
	input := dropboxLib.DownloadInput{Path: utils.JoinPath("/", file)}
	expectedError := errors.New("test error")
	filesMock := mocks.NewDropboxFilesMock(t).
		DownloadMock.Expect(&input).Return(nil, expectedError)
	client := dropbox.NewClient(dropbox.WithFiles(filesMock))

	reader, err := client.DownloadFile(file)

	assert.Equal(t, expectedError, err)
	assert.Nil(t, reader)
}

func TestClient_AccountDisplayName(t *testing.T) {
	client := dropbox.NewClient(dropbox.WithAccount(&mocks.Account))

	assert.Equal(t, mocks.Account.Name.DisplayName, client.AccountDisplayName())
}

func TestClient_AccountEmail(t *testing.T) {
	client := dropbox.NewClient(dropbox.WithAccount(&mocks.Account))

	assert.Equal(t, mocks.Account.Email, client.AccountEmail())
}
