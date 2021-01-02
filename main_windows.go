package main

import (
	"log"
	"os"
	"path/filepath"
	"pb-dropbox-downloader/utils"
)

func main() {
	defer utils.PanicInterceptor(os.Exit, fatalExitCode)

	logfile, err := openLogFile(filepath.Join("./testdata", logFileName))
	if err != nil {
		panic(err)
	}

	defer logfile.Close()
	log.SetOutput(logfile)

	config, err := loadConfig(filepath.Join("./testdata", configFileName))
	if err != nil {
		panic(err)
	}

	synchroniser, err := createSynchroniser(config.AccessToken, filepath.Join("./testdata", databaseFileName))
	if err != nil {
		panic(err)
	}

	folder := filepath.Join("./testdata/internal", config.Folder)
	if config.OnSdCard {
		folder = filepath.Join("./testdata/sdcard", config.Folder)
	}

	err = synchroniser.Sync(folder, config.AllowDeleteFiles)
	if err != nil {
		panic(err)
	}
}
