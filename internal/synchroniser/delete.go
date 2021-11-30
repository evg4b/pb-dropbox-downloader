package synchroniser

import (
	"fmt"
	"log"
	"pb-dropbox-downloader/internal/utils"
)

func (db *DropboxSynchroniser) delete(folder string, files []string) error {
	for _, file := range files {
		err := db.files.Remove(utils.JoinPath(folder, file))
		if err != nil {
			fmt.Fprintf(db.output, "%s .... [filed]\n", file)
			log.Println(err)

			return err
		}

		fmt.Fprintf(db.output, "%s .... [ok]\n", file)
		err = db.storage.Remove(file)
		if err != nil {
			fmt.Fprintf(db.output, "%s .... [filed]\n", file)
			log.Println(err)

			return err
		}
	}

	return nil
}
