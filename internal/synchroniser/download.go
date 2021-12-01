// nolint: wrapcheck
package synchroniser

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"pb-dropbox-downloader/internal/dropbox"
	"pb-dropbox-downloader/internal/utils"
	"sync"

	"github.com/c2h5oh/datasize"
)

type dataChannel = chan dropbox.RemoteFile

func (ds *DropboxSynchroniser) download(folder string, files []dropbox.RemoteFile) error {
	if len(files) == 0 {
		ds.printf("no files to download")

		return nil
	}

	ds.printf("founded %d files to download", len(files))

	source := make(dataChannel)
	wg := ds.startThreads(folder, source)
	for _, file := range files {
		source <- file
	}

	close(source)

	wg.Wait()

	return nil
}

func (ds *DropboxSynchroniser) startThreads(folder string, source dataChannel) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	wg.Add(ds.maxParallelism)
	for i := 0; i < ds.maxParallelism; i++ {
		go ds.downloadThread(&wg, folder, source)
	}

	return &wg
}

func (ds *DropboxSynchroniser) downloadThread(wg *sync.WaitGroup, folder string, source dataChannel) {
	defer wg.Done()

	for file := range source {
		err := ds.downloadFile(file, folder)
		if err != nil {
			ds.printf("%s .... [filed]", filepath.Base(file.Path))
			log.Println(err)

			continue
		}

		ds.printf("%s (%s) .... [ok]", filepath.Base(file.Path), datasize.ByteSize(file.Size).HumanReadable())
		err = ds.storage.Add(file.Path, file.Hash)
		if err != nil {
			ds.printf("%s .... [filed]", filepath.Base(file.Path))
			log.Println(err)

			continue
		}
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
