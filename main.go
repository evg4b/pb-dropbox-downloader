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

// https://www.dropbox.com/oauth2/authorize?client_id=srb8i7o71xkx08c&response_type=code

func main() {
	config := struct {
		AccessToken string `json:"access_token"`
		Folder      string `json:"folder"`
		DeleteFiles bool   `json:"delete_files"`
	}{}

	data, err := ioutil.ReadFile("./demo.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &config)

	db := dropboxLib.New(dropboxLib.NewConfig(config.AccessToken))
	account, err := db.Users.GetCurrentAccount()
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Name.DisplayName)
	fmt.Println(account.Email)

	dbClient := dropbox.NewClient(db.Files)
	fileSystem := filesystem.Local{}
	storage := datastorage.NewFileStorage(&fileSystem, "./ddd.bin")

	syncer := synchroniser.NewSynchroniser(storage, &fileSystem, dbClient, os.Stdout)
	syncer.Sync("./data", true)
}
