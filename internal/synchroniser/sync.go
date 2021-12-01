// nolint: wrapcheck
package synchroniser

import (
	"os"
	"path/filepath"
	"pb-dropbox-downloader/internal/dropbox"
	"strings"
)

// Sync synchronies folder with application folder in drop box.
func (ds *DropboxSynchroniser) Sync(folder string, remove bool) error {
	ds.infoHeader()

	normalizedFolder := filepath.ToSlash(folder)
	files, err := ds.getLocalFiles(normalizedFolder)
	if err != nil {
		return err
	}

	filesToDownload, err := ds.getFilesToDownload()
	if err != nil {
		return err
	}

	err = ds.download(normalizedFolder, filesToDownload)
	if err != nil {
		return err
	}

	err = ds.storage.Commit()
	if err != nil {
		return err
	}

	if remove {
		err = ds.deleteFiles(normalizedFolder, files)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ds *DropboxSynchroniser) getLocalFiles(folder string) ([]os.FileInfo, error) {
	files, err := ds.files.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	err = ds.refreshStorage(files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (ds *DropboxSynchroniser) getFilesToDownload() ([]dropbox.RemoteFile, error) {
	remoteFiles, err := ds.dropbox.GetFiles()
	if err != nil {
		return nil, err
	}

	filesToDownload := []dropbox.RemoteFile{}
	for _, remoteFile := range remoteFiles {
		if hash, err := ds.storage.Get(remoteFile.Path); err == nil {
			if strings.EqualFold(hash, remoteFile.Hash) {
				continue
			}
		}

		filesToDownload = append(filesToDownload, remoteFile)
	}

	return filesToDownload, nil
}

func (ds *DropboxSynchroniser) refreshStorage(files []os.FileInfo) error {
	storageFiles, err := ds.storage.ToMap()
	if err != nil {
		return err
	}

	filteredMap := filterMapBy(storageFiles, func(key, _ string) bool {
		return fileSliceContins(files, key)
	})

	ds.storage.FromMap(filteredMap)

	return ds.storage.Commit()
}
