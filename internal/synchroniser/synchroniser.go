package synchroniser

import (
	"pb-dropbox-downloader/infrastructure"
	"pb-dropbox-downloader/internal"
	"pb-dropbox-downloader/utils"
	"strings"
)

type DropboxSynchroniser struct {
	storage internal.DataStorage
	files   infrastructure.FileSystem
	dropbox infrastructure.Dropbox
}

// NewSynchroniser creates and initialize new instance of DropboxSynchroniser create
func NewSynchroniser(
	storage internal.DataStorage,
	files infrastructure.FileSystem,
	dropbox infrastructure.Dropbox,
) *DropboxSynchroniser {
	return &DropboxSynchroniser{
		storage: storage,
		files:   files,
		dropbox: dropbox,
	}
}

func (db *DropboxSynchroniser) Sync(folder string, remove bool) error {
	files := db.files.GetFilesInFolder(folder)
	err := db.refreshStorage(files)
	if err != nil {
		return err
	}

	remotesFiles := db.dropbox.GetFiles()
	filesToDownload := []infrastructure.RemoteFile{}
	for _, remoteFile := range remotesFiles {
		if hash, ok := db.storage.Get(remoteFile.Path); ok {
			if strings.EqualFold(hash, remoteFile.Hash) {
				continue
			}
		}

		filesToDownload = append(filesToDownload, remoteFile)
	}

	db.download(remotesFiles)

	return nil
}

func (db *DropboxSynchroniser) refreshStorage(files []string) error {
	storageMap, err := db.storage.ToMap()
	if err != nil {
		return err
	}

	filteredMap := utils.FilterMapBy(storageMap, func(key, _ string) bool {
		return utils.Contins(files, key)
	})

	return db.storage.FromMap(filteredMap)
}
