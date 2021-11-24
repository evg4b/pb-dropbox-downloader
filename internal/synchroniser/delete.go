package synchroniser

import (
	"fmt"
	"log"
	"pb-dropbox-downloader/utils"
)

func (db *DropboxSynchroniser) delete(folder string, files []string) error {
	for _, file := range files {
		err := db.files.DeleteFile(utils.JoinPath(folder, file))
		if err != nil {
			fmt.Fprintf(db.output, "%s .... [filed]\n", file)
			log.Println(err)

			return err
		}

		fmt.Fprintf(db.output, "%s .... [ok]\n", file)
		db.storage.Remove(file)
	}

	return nil
}
