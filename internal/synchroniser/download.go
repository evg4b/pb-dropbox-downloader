// nolint: wrapcheck
package synchroniser

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"pb-dropbox-downloader/internal/dropbox"
	"pb-dropbox-downloader/internal/utils"

	"github.com/hashicorp/go-multierror"

	"golang.org/x/sync/errgroup"
)

type dataChannel = chan dropbox.RemoteFile

type printInfo struct {
	name    string
	size    uint64
	success bool
}

func (s *DropboxSynchroniser) download(ctx context.Context, folder string, files []dropbox.RemoteFile) error {
	if len(files) == 0 {
		s.printf("No files to download")

		return nil
	}

	s.printf("Found %d files to download:\n", len(files))

	source := make(dataChannel)
	results := make(chan printInfo)

	tasksGroup, _ := errgroup.WithContext(ctx)

	for i := 0; i < calculateTheadsCount(s.maxParallelism, files); i++ {
		tasksGroup.Go(s.createDownloadThread(results, folder, source))
	}

	go s.printerThread(results)
	defer close(results)

	for _, file := range files {
		source <- file
	}

	close(source)

	return tasksGroup.Wait()
}

func substring(s string, num int) string {
	chars := []rune(s)

	return string(chars[:num])
}

func (s *DropboxSynchroniser) createDownloadThread(target chan printInfo, folder string, source dataChannel) func() error {
	return func() error {
		var result *multierror.Error

		for file := range source {
			fileName := filepath.Base(file.Path)
			if err := s.downloadFile(file, folder); err != nil {
				target <- printInfo{name: fileName}

				result = multierror.Append(result, err)

				continue
			}

			target <- printInfo{
				name:    fileName,
				size:    file.Size,
				success: true,
			}

			if err := s.storage.Add(file.Path, file.Hash); err != nil {
				result = multierror.Append(result, err)
			}
		}

		return result.ErrorOrNil()
	}
}

func (s *DropboxSynchroniser) downloadFile(file dropbox.RemoteFile, folder string) error {
	fileReader, err := s.dropbox.DownloadFile(file.Path)
	if err != nil {
		return err
	}

	defer fileReader.Close()

	data, err := ioutil.ReadAll(fileReader)
	if err != nil {
		return err
	}

	localFile, err := s.files.Create(utils.JoinPath(folder, file.Path))
	if err != nil {
		return err
	}

	defer localFile.Close()

	_, err = localFile.Write(data)

	return err
}
