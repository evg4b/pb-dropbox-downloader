package main

import (
	"log"
	"os"
	"pb-dropbox-downloader/infrastructure/pocketbook"
	"pb-dropbox-downloader/utils"
)

func main() {
	defer utils.PanicInterceptor(os.Exit, fatalExitCode)

	logfile, err := openLogFile(pocketbook.Share(logFileName))
	if err != nil {
		panic(err)
	}

	defer logfile.Close()
	log.SetOutput(logfile)

	config, err := loadConfig(pocketbook.ConfigPath(configFileName))
	if err != nil {
		panic(err)
	}

	synchroniser, err := createSynchroniser(config.AccessToken, pocketbook.Share(databaseFileName))
	if err != nil {
		panic(err)
	}

	folder := pocketbook.Internal(config.Folder)
	if config.OnSdCard {
		folder = pocketbook.SdCard(config.Folder)
	}

	err = synchroniser.Sync(folder, config.AllowDeleteFiles)
	if err != nil {
		panic(err)
	}
}
