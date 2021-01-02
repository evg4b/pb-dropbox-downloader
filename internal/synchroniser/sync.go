package synchroniser

import (
	"path/filepath"
	"pb-dropbox-downloader/infrastructure"
	"pb-dropbox-downloader/utils"
	"strings"
)

// Sync synchronies folder with application folder in drop box
func (db *DropboxSynchroniser) Sync(folder string, remove bool) error {
	normalizedFolder := filepath.FromSlash(folder)
	files := db.files.GetFilesInFolder(normalizedFolder)
	err := db.refreshStorage(files)
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
		filesToDelete := utils.FilterSliceBy(files, func(key string) bool {
			return !db.storage.KeyExists(key)
		})
		err = db.delete(normalizedFolder, filesToDelete)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DropboxSynchroniser) refreshStorage(files []string) error {
	storageMap, err := db.storage.ToMap()
	if err != nil {
		return err
	}

	filteredMap := utils.FilterMapBy(storageMap, func(key, _ string) bool {
		return utils.SliceContins(files, key)
	})

	return db.storage.FromMap(filteredMap)
}
