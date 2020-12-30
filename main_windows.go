package main

import (
	"log"
	"os"
	"path"
	"pb-dropbox-downloader/utils"
)

func main() {
	defer utils.PanicInterceptor(os.Exit, fatalExitCode)

	logfile, err := openLogFile(path.Join("./testdata", logFileName))
	if err != nil {
		panic(err)
	}

	defer logfile.Close()
	log.SetOutput(logfile)

	config, err := loadConfig(path.Join("./testdata", configFileName))
	if err != nil {
		panic(err)
	}

	synchroniser, err := createSynchroniser(config.AccessToken, path.Join("./testdata", databaseFileName))
	if err != nil {
		panic(err)
	}

	folder := path.Join("./testdata/internal", config.Folder)
	if config.OnSdCard {
		folder = path.Join("./testdata/sdcard", config.Folder)
	}

	err = synchroniser.Sync(folder, config.AllowDeleteFiles)
	if err != nil {
		panic(err)
	}
}
