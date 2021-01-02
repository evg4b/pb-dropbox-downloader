package synchroniser

import (
	"fmt"
	"log"
	"path/filepath"
)

func (db *DropboxSynchroniser) delete(folder string, files []string) error {
	for _, file := range files {
		err := db.files.DeleteFile(filepath.Join(folder, file))
		if err != nil {
			fmt.Fprintln(db.output, fmt.Sprintf("%s .... [filed]", file))
			log.Println(err)

			return err
		}

		fmt.Fprintln(db.output, fmt.Sprintf("%s .... [ok]", file))
		db.storage.Remove(file)
	}

	return nil
}
