package main

import (
	"io"
	"log"
	"os"
	"pb-dropbox-downloader/internal/config"
	"pb-dropbox-downloader/internal/datastorage"
	"pb-dropbox-downloader/internal/dropbox"
	"pb-dropbox-downloader/internal/pocketbook"
	"pb-dropbox-downloader/internal/synchroniser"

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

func mainInternal(w io.Writer) error {
	fs := osfs.New("")

	const logfilePerm = 0755
	logfile, err := os.OpenFile(pocketbook.Share(logFileName), os.O_CREATE|os.O_APPEND, logfilePerm)
	if err != nil {
		return err
	}

	defer logfile.Close()
	log.SetOutput(logfile)

	syncConfig, err := config.LoadConfig(fs, pocketbook.ConfigPath(configFileName))
	if err != nil {
		return err
	}

	dropboxLibClient := dropboxLib.New(dropboxLib.NewConfig(syncConfig.AccessToken))
	dropboxClient := dropbox.NewClient(dropbox.WithGoDropbox(dropboxLibClient))
	storage := datastorage.NewFileStorage(fs, pocketbook.Share(databaseFileName))

	synchroniser := synchroniser.NewSynchroniser(
		synchroniser.WithStorage(storage),
		synchroniser.WithFileSystem(fs),
		synchroniser.WithDropboxClient(dropboxClient),
		synchroniser.WithOutput(w),
		synchroniser.WithMaxParallelism(parallelism),
	)

	folder := pocketbook.Internal(syncConfig.Folder)
	if syncConfig.OnSdCard {
		folder = pocketbook.SdCard(syncConfig.Folder)
	}

	return synchroniser.Sync(folder, syncConfig.AllowDeleteFiles)
}
