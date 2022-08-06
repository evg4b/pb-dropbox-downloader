// nolint: wrapcheck
package synchroniser

import (
	"context"
	"os"
	"path/filepath"
	"pb-dropbox-downloader/internal/dropbox"
	"strings"
)

// Sync synchronies folder with application folder in drop box.
func (s *DropboxSynchroniser) Sync(ctx context.Context, folder string, remove bool) error {
	s.infoHeader()

	if err := s.files.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}

	normalizedFolder := filepath.ToSlash(folder)
	files, err := s.getLocalFiles(normalizedFolder)
	if err != nil {
		return err
	}

	filesToDownload, err := s.getFilesToDownload()
	if err != nil {
		return err
	}

	err = s.download(ctx, normalizedFolder, filesToDownload)
	if err != nil {
		return err
	}

	err = s.storage.Commit()
	if err != nil {
		return err
	}

	if remove {
		err = s.deleteFiles(normalizedFolder, files)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DropboxSynchroniser) getLocalFiles(folder string) ([]os.FileInfo, error) {
	files, err := s.files.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	err = s.refreshStorage(files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (s *DropboxSynchroniser) getFilesToDownload() ([]dropbox.RemoteFile, error) {
	remoteFiles, err := s.dropbox.GetFiles()
	if err != nil {
		return nil, err
	}

	filesToDownload := []dropbox.RemoteFile{}
	for _, remoteFile := range remoteFiles {
		if hash, err := s.storage.Get(remoteFile.Path); err == nil {
			if strings.EqualFold(hash, remoteFile.Hash) {
				continue
			}
		}

		filesToDownload = append(filesToDownload, remoteFile)
	}

	return filesToDownload, nil
}

func (s *DropboxSynchroniser) refreshStorage(files []os.FileInfo) error {
	storageFiles, err := s.storage.ToMap()
	if err != nil {
		return err
	}

	filteredMap := filterMapBy(storageFiles, func(key, _ string) bool {
		return fileSliceContins(files, key)
	})

	s.storage.FromMap(filteredMap)

	return s.storage.Commit()
}
