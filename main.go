package main

import (
	"fmt"
	"io"
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

func mainInternal(w io.Writer) {
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

	fmt.Fprintln(w, account.Name.DisplayName)
	fmt.Fprintln(w, account.Email)

	dropboxClient := dropbox.NewClient(dropboxLibClient.Files)

	fs := osfs.New("")

	storage := datastorage.NewFileStorage(fs, pocketbook.Share(databaseFileName))

	synchroniser := synchroniser.NewSynchroniser(
		synchroniser.WithStorage(storage),
		synchroniser.WithFileSystem(fs),
		synchroniser.WithDropboxClient(dropboxClient),
		synchroniser.WithOutput(w),
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
