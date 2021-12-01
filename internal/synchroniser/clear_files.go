// nolint: wrapcheck
package synchroniser

import (
	"fmt"
	"log"
	"os"
	"pb-dropbox-downloader/internal/utils"
)

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
