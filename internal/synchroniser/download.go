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

func (ds *DropboxSynchroniser) download(folder string, files []dropbox.RemoteFile) error {
	if len(files) == 0 {
		ds.printf("no files to download")
	}

	ds.printf("founded %d files to download", len(files))

	source := make(dataChannel)

	tasksGroup, _ := errgroup.WithContext(context.Background())
	for i := 0; i < calculateTheadsCount(ds.maxParallelism, files); i++ {
		tasksGroup.Go(ds.createDownloadThread(folder, source))
	}

	for _, file := range files {
		source <- file
	}

	close(source)

	return tasksGroup.Wait()
}

func (ds *DropboxSynchroniser) createDownloadThread(folder string, source dataChannel) func() error {
	return func() error {
		var result *multierror.Error

		for file := range source {
			err := ds.downloadFile(file, folder)
			if err != nil {
				ds.printf("%s .... [filed]", filepath.Base(file.Path))
				log.Println(err)

				result = multierror.Append(result, err)

				continue
			}

			ds.printf("%s (%s) .... [ok]", filepath.Base(file.Path), datasize.ByteSize(file.Size).HumanReadable())
			err = ds.storage.Add(file.Path, file.Hash)
			if err != nil {
				ds.printf("%s .... [filed]", filepath.Base(file.Path))
				log.Println(err)

				result = multierror.Append(result, err)

				continue
			}
		}

		return result.ErrorOrNil()
	}
}

func (ds *DropboxSynchroniser) downloadFile(file dropbox.RemoteFile, folder string) error {
	fileReader, err := ds.dropbox.DownloadFile(file.Path)
	if err != nil {
		return err
	}

	defer fileReader.Close()

	data, err := ioutil.ReadAll(fileReader)
	if err != nil {
		return err
	}

	localFile, err := ds.files.Create(utils.JoinPath(folder, file.Path))
	if err != nil {
		return err
	}

	defer localFile.Close()

	_, err = localFile.Write(data)

	return err
}
