package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"pb-dropbox-downloader/infrastructure/dropbox"
	"pb-dropbox-downloader/infrastructure/filesystem"
	"pb-dropbox-downloader/infrastructure/pocketbook"
	"pb-dropbox-downloader/internal/datastorage"
	"pb-dropbox-downloader/internal/synchroniser"
	"pb-dropbox-downloader/utils"

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

func main() {
	defer utils.PanicInterceptor(os.Exit, fatalExitCode)

	logfile, err := os.OpenFile(pocketbook.Share(logFileName), os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}

	defer logfile.Close()
	log.SetOutput(logfile)

	config, err := loadConfig(pocketbook.ConfigPath(configFileName))
	if err != nil {
		panic(err)
	}

	dropboxLibClient := dropboxLib.New(dropboxLib.NewConfig(config.AccessToken))
	account, err := dropboxLibClient.Users.GetCurrentAccount()
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Name.DisplayName)
	fmt.Println(account.Email)

	dropboxClient := dropbox.NewClient(dropboxLibClient.Files)
	fileSystem := filesystem.NewFileSystem()
	storage := datastorage.NewFileStorage(fileSystem, pocketbook.Share(databaseFileName))

	synchroniser := synchroniser.NewSynchroniser(storage, fileSystem, dropboxClient, os.Stdout, parallelism)

	folder := pocketbook.Internal(config.Folder)
	if config.OnSdCard {
		folder = pocketbook.SdCard(config.Folder)
	}

	err = synchroniser.Sync(folder, config.AllowDeleteFiles)
	if err != nil {
		panic(err)
	}
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
