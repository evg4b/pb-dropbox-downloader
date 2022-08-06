package app

import (
	"context"
	"fmt"
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

var (
	parallelism      = 3
	logFileName      = "pb-dropbox-downloader.log"
	databaseFileName = "pb-dropbox-downloader.bin"
	configFileName   = "pb-dropbox-downloader-config.json"
	version          = "X.X.X"
)

const perm = 0755

func Run(ctx context.Context, w io.Writer) error {
	fs := osfs.New("")

	if err := fs.MkdirAll(pocketbook.ConfigPath(), perm); err != nil {
		return fmt.Errorf("failed to create config dir: %w", err)
	}

	if err := fs.MkdirAll(pocketbook.Share(), perm); err != nil {
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

	dropboxClient := dropbox.NewClient(dropbox.WithGoDropbox(
		dropboxLib.New(dropboxLib.NewConfig(syncConfig.AccessToken)),
	))

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
		synchroniser.WithVersion(version),
	)

	folder := pocketbook.Internal(syncConfig.Folder)
	if syncConfig.OnSdCard {
		folder = pocketbook.SdCard(syncConfig.Folder)
	}

	err = synchroniser.Sync(ctx, folder, syncConfig.AllowDeleteFiles)
	if err != nil {
		return fmt.Errorf("synchronization finished with error: %w", err)
	}

	fmt.Fprintln(w)

	err = pocketbook.RefreshScanner(ctx)
	if err != nil {
		return fmt.Errorf("filed to update book list: %w", err)
	}

	return nil
}
