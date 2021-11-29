package synchroniser

import (
	"fmt"
	"os"
	"path/filepath"
	"pb-dropbox-downloader/infrastructure"
	"pb-dropbox-downloader/utils"
	"strings"
)

// nolint: cyclop
// Sync synchronies folder with application folder in drop box.
func (db *DropboxSynchroniser) Sync(folder string, remove bool) error {
	fmt.Fprintln(db.output, logo)
	fmt.Fprintf(db.output, "Account: %s\n", db.dropbox.AccountDisplayName())
	fmt.Fprintf(db.output, "Email: %s\n", db.dropbox.AccountEmail())

	normalizedFolder := filepath.ToSlash(folder)

	files, err := db.files.ReadDir(normalizedFolder)
	if err != nil {
		return err
	}

	err = db.refreshStorage(files)
	if err != nil {
		return err
	}

	remotesFiles, err := db.dropbox.GetFiles()
	if err != nil {
		return err
	}

	filesToDownload := []infrastructure.RemoteFile{}
	for _, remoteFile := range remotesFiles {
		if hash, ok := db.storage.Get(remoteFile.Path); ok {
			if strings.EqualFold(hash, remoteFile.Hash) {
				continue
			}
		}

		filesToDownload = append(filesToDownload, remoteFile)
	}

	err = db.download(normalizedFolder, filesToDownload)
	if err != nil {
		return err
	}

	err = db.storage.Commit()
	if err != nil {
		return err
	}

	if remove {
		filesToDelete := utils.FilterSliceBy(files, func(file os.FileInfo) bool {
			return !db.storage.KeyExists(file.Name())
		})
		err = db.delete(normalizedFolder, filesToDelete)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DropboxSynchroniser) refreshStorage(files []os.FileInfo) error {
	storageMap, err := db.storage.ToMap()
	if err != nil {
		return err
	}

	filteredMap := utils.FilterMapBy(storageMap, func(key, _ string) bool {
		return utils.FileSliceContins(files, key)
	})

	return db.storage.FromMap(filteredMap)
}
