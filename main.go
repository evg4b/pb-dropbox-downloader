package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"pb-dropbox-downloader/infrastructure/dropbox"
	"pb-dropbox-downloader/infrastructure/pocketbook"
	"pb-dropbox-downloader/internal/datastorage"
	"pb-dropbox-downloader/internal/synchroniser"
	"pb-dropbox-downloader/utils"

	"github.com/go-git/go-billy/v5/osfs"
	dropboxLib "github.com/tj/go-dropbox"
)

const (
	fatalExitCode    = 500
	parallelism      = 3
	logFileName      = "pb-dropbox-downloader.log"
	databaseFileName = "pb-dropbox-downloader.bin"
	configFileName   = "pb-dropbox-downloader-config.json"
)

type pbSyncConfig struct {
	AccessToken      string `json:"access_token"`
	Folder           string `json:"folder"`
	AllowDeleteFiles bool   `json:"allow_delete_files"`
	OnSdCard         bool   `json:"on_sd_card"`
}

func main() {
	defer utils.PanicInterceptor(os.Exit, fatalExitCode)

	const logfilePerm = 0755
	logfile, err := os.OpenFile(pocketbook.Share(logFileName), os.O_CREATE|os.O_APPEND, logfilePerm)
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

	fs := osfs.New("")

	storage := datastorage.NewFileStorage(fs, pocketbook.Share(databaseFileName))

	synchroniser := synchroniser.NewSynchroniser(
		synchroniser.WithStorage(storage),
		synchroniser.WithFileSystem(fs),
		synchroniser.WithDropboxClient(dropboxClient),
		synchroniser.WithOutput(os.Stdout),
		synchroniser.WithMaxParallelism(parallelism),
	)

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
	config := pbSyncConfig{}
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	log.Printf("loaded configuration from '%s'", configPath)

	err = json.Unmarshal(configData, &config)

	return config, err
}
