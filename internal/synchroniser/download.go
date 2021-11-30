package synchroniser

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"pb-dropbox-downloader/internal/dropbox"
	"pb-dropbox-downloader/utils"
	"sync"

	"github.com/c2h5oh/datasize"
)

type dataChannel = chan dropbox.RemoteFile

func (db *DropboxSynchroniser) download(folder string, files []dropbox.RemoteFile) error {
	if len(files) == 0 {
		db.printf("no files to download")

		return nil
	}

	db.printf("founded %d files to download", len(files))

	source := make(dataChannel)
	wg := db.startThreads(folder, source)
	for _, file := range files {
		source <- file
	}

	close(source)

	wg.Wait()

	return nil
}

func (db *DropboxSynchroniser) startThreads(folder string, source dataChannel) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	wg.Add(db.maxParallelism)
	for i := 0; i < db.maxParallelism; i++ {
		go db.downloadThread(&wg, folder, source)
	}

	return &wg
}

func (db *DropboxSynchroniser) downloadThread(wg *sync.WaitGroup, folder string, source dataChannel) {
	defer wg.Done()

	for file := range source {
		err := db.downloadFile(file, folder)
		if err != nil {
			db.printf("%s .... [filed]", filepath.Base(file.Path))
			log.Println(err)

			continue
		}

		db.printf("%s (%s) .... [ok]", filepath.Base(file.Path), datasize.ByteSize(file.Size).HumanReadable())
		db.storage.Add(file.Path, file.Hash)
	}
}

func (db *DropboxSynchroniser) downloadFile(file dropbox.RemoteFile, folder string) error {
	fileReader, err := db.dropbox.DownloadFile(file.Path)
	if err != nil {
		return err
	}

	defer fileReader.Close()

	data, err := ioutil.ReadAll(fileReader)
	if err != nil {
		return err
	}

	localFile, err := db.files.Create(utils.JoinPath(folder, file.Path))
	if err != nil {
		return err
	}

	defer localFile.Close()

	_, err = localFile.Write(data)

	return err
}
