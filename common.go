package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"pb-dropbox-downloader/infrastructure/dropbox"
	"pb-dropbox-downloader/infrastructure/filesystem"
	"pb-dropbox-downloader/internal/datastorage"
	"pb-dropbox-downloader/internal/synchroniser"

	dropboxLib "github.com/tj/go-dropbox"
)

const fatalExitCode = 500
const parallelism = 3
const logFileName = "pb-dropbox-downloader.log"
const databaseFileName = "pb-dropbox-downloader.bin"
const configFileName = "pb-dropbox-downloader-config.json"

type pbSyncConfig struct {
	AccessToken      string `json:"access_token"`
	Folder           string `json:"folder"`
	AllowDeleteFiles bool   `json:"allow_delete_files"`
	OnSdCard         bool   `json:"on_sd_card"`
}

func openLogFile(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0755)
}

func createSynchroniser(accessToken string, database string) (*synchroniser.DropboxSynchroniser, error) {
	dropboxLibClient := dropboxLib.New(dropboxLib.NewConfig(accessToken))
	account, err := dropboxLibClient.Users.GetCurrentAccount()
	if err != nil {
		return nil, err
	}

	fmt.Println(account.Name.DisplayName)
	fmt.Println(account.Email)

	dropboxClient := dropbox.NewClient(dropboxLibClient.Files)
	fileSystem := filesystem.NewFileSystem()
	storage := datastorage.NewFileStorage(fileSystem, database)

	return synchroniser.NewSynchroniser(storage, fileSystem, dropboxClient, os.Stdout, parallelism), nil
}

func loadConfig(configPath string) (pbSyncConfig, error) {
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return pbSyncConfig{}, err
	}

	config := pbSyncConfig{}
	json.Unmarshal(configData, &config)

	return config, nil
}
