package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"pb-dropbox-downloader/internal/config"
	"pb-dropbox-downloader/internal/datastorage"
	"pb-dropbox-downloader/internal/dropbox"
	"pb-dropbox-downloader/internal/pocketbook"
	"pb-dropbox-downloader/internal/synchroniser"

	"github.com/go-git/go-billy/v5/osfs"
	dropboxLib "github.com/tj/go-dropbox"
)

const (
	perm             = 0755
	fatalExitCode    = 500
	parallelism      = 3
	logFileName      = "pb-dropbox-downloader.log"
	databaseFileName = "pb-dropbox-downloader.bin"
	configFileName   = "pb-dropbox-downloader-config.json"
)

func mainInternal(w io.Writer) error {
	fs := osfs.New("")

	if err := os.MkdirAll(pocketbook.ConfigPath(), perm); err != nil {
		return fmt.Errorf("failed to create config dir: %w", err)
	}

	if err := os.MkdirAll(pocketbook.ConfigPath(), perm); err != nil {
		return fmt.Errorf("failed to create share dir: %w", err)
	}

	logfile, err := os.OpenFile(pocketbook.Share(logFileName), os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	defer logfile.Close()
	log.SetOutput(logfile)

	syncConfig, err := config.LoadConfig(fs, pocketbook.ConfigPath(configFileName))
	if err != nil {
		return fmt.Errorf("failed loaded configuration: %w", err)
	}

	dropboxLibClient := dropboxLib.New(dropboxLib.NewConfig(syncConfig.AccessToken))
	dropboxClient := dropbox.NewClient(dropbox.WithGoDropbox(dropboxLibClient))
	storage := datastorage.NewFileStorage(
		datastorage.WithFilesystem(fs),
		datastorage.WithConfigPath(pocketbook.Share(databaseFileName)),
	)

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

	err = synchroniser.Sync(folder, syncConfig.AllowDeleteFiles)
	if err != nil {
		return fmt.Errorf("synchronization finished with error: %w", err)
	}

	cmd := exec.Command("/bin/killall", "scanner.app")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}
