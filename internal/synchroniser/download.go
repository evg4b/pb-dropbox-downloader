// nolint: wrapcheck
package synchroniser

import (
	"context"
	"io/ioutil"
	"log"
	"path/filepath"
	"pb-dropbox-downloader/internal/dropbox"
	"pb-dropbox-downloader/internal/utils"

	"github.com/c2h5oh/datasize"
	"github.com/hashicorp/go-multierror"

	"golang.org/x/sync/errgroup"
)

type dataChannel = chan dropbox.RemoteFile

func (s *DropboxSynchroniser) download(folder string, files []dropbox.RemoteFile) error {
	if len(files) == 0 {
		s.printf("no files to download")
	}

	s.printf("Found %d files to download:\n", len(files))

	source := make(dataChannel)

	tasksGroup, _ := errgroup.WithContext(context.Background())
	for i := 0; i < calculateTheadsCount(s.maxParallelism, files); i++ {
		tasksGroup.Go(s.createDownloadThread(folder, source))
	}

	for _, file := range files {
		source <- file
	}

	close(source)

	return tasksGroup.Wait()
}

func (s *DropboxSynchroniser) createDownloadThread(folder string, source dataChannel) func() error {
	return func() error {
		var result *multierror.Error

		for file := range source {
			err := s.downloadFile(file, folder)
			if err != nil {
				s.printf("%s .... [filed]", filepath.Base(file.Path))
				log.Println(err)

				result = multierror.Append(result, err)

				continue
			}

			s.printf("%s (%s) .... [ok]", filepath.Base(file.Path), datasize.ByteSize(file.Size).HumanReadable())
			err = s.storage.Add(file.Path, file.Hash)
			if err != nil {
				s.printf("%s .... [filed]", filepath.Base(file.Path))
				log.Println(err)

				result = multierror.Append(result, err)

				continue
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
