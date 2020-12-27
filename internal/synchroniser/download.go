package synchroniser

import (
	"log"
	"path"
	"pb-dropbox-downloader/infrastructure"
	"sync"
)

type dataChannel = chan infrastructure.RemoteFile

func (db *DropboxSynchroniser) download(folder string, files []infrastructure.RemoteFile) error {
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
		go db.downloadThread(&wg, folder, source, i)
	}

	return &wg
}

func (db *DropboxSynchroniser) downloadThread(wg *sync.WaitGroup, folder string, source dataChannel, id int) {
	defer wg.Done()

	for file := range source {
		fileReader, err := db.dropbox.DownloadFile(file.Path)
		if err != nil {
			db.printf("%s .... [filed]", file.Path)
			log.Println(err)

			continue
		}

		err = db.files.CopyDataToFile(path.Join(folder, file.Path), fileReader)
		if err != nil {
			db.printf("%s .... [filed]", file.Path)
			log.Println(err)

			continue
		}

		db.printf("%s .... [ok]", file.Path)
	}
}