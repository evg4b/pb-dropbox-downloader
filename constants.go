package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
