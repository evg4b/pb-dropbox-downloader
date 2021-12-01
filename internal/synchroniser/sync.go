// nolint: wrapcheck
package synchroniser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"pb-dropbox-downloader/internal/dropbox"
	"pb-dropbox-downloader/internal/utils"
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
	remotesFiles, err := ds.dropbox.GetFiles()
	if err != nil {
		return nil, err
	}

	filesToDownload := []dropbox.RemoteFile{}
	for _, remoteFile := range remotesFiles {
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
	storageMap, err := ds.storage.ToMap()
	if err != nil {
		return err
	}

	filteredMap := filterMapBy(storageMap, func(key, _ string) bool {
		return fileSliceContins(files, key)
	})

	ds.storage.FromMap(filteredMap)

	return ds.storage.Commit()
}

func (ds *DropboxSynchroniser) deleteFiles(folder string, files []os.FileInfo) error {
	for _, file := range files {
		fileName := file.Name()
		exist, err := ds.storage.KeyExists(fileName)
		if err != nil {
			return err
		}

		if !exist {
			err := ds.files.Remove(utils.JoinPath(folder, fileName))
			if err != nil {
				fmt.Fprintf(ds.output, "%s .... [filed]\n", file)
				log.Println(err)

				return err
			}

			fmt.Fprintf(ds.output, "%s .... [ok]\n", file)
			err = ds.storage.Remove(fileName)
			if err != nil {
				fmt.Fprintf(ds.output, "%s .... [filed]\n", file)
				log.Println(err)

				return err
			}
		}
	}

	return nil
}
