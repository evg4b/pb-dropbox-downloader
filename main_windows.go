package main

import (
	"log"
	"os"
	"pb-dropbox-downloader/utils"
)

func main() {
	defer utils.PanicInterceptor(os.Exit, fatalExitCode)

	logfile, err := openLogFile(utils.JoinPath("./testdata", logFileName))
	if err != nil {
		panic(err)
	}

	defer logfile.Close()
	log.SetOutput(logfile)

	config, err := loadConfig(utils.JoinPath("./testdata", configFileName))
	if err != nil {
		panic(err)
	}

	synchroniser, err := createSynchroniser(config.AccessToken, utils.JoinPath("./testdata", databaseFileName))
	if err != nil {
		panic(err)
	}

	folder := utils.JoinPath("./testdata/internal", config.Folder)
	if config.OnSdCard {
		folder = utils.JoinPath("./testdata/sdcard", config.Folder)
	}

	err = synchroniser.Sync(folder, config.AllowDeleteFiles)
	if err != nil {
		panic(err)
	}
}
